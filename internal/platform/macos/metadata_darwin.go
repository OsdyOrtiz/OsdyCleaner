//go:build darwin

// Package macos observes one local filesystem path without following links.
package macos

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"
	"syscall"

	"github.com/osdy/OsdyCleaner/internal/core"
	"golang.org/x/sys/unix"
)

var (
	ErrMetadataMissing            = errors.New("metadata path is missing")
	ErrMetadataSymlink            = errors.New("metadata path is a symbolic link")
	ErrMetadataNotDirectory       = errors.New("metadata path is not a directory")
	ErrMetadataDeviceMismatch     = errors.New("metadata path is on another device")
	ErrMetadataNonLocalFilesystem = errors.New("metadata path is not on a local filesystem")
	ErrMetadataInaccessible       = errors.New("metadata is inaccessible")
	ErrMetadataInvalid            = errors.New("metadata is invalid")
)

const localFilesystemFlag = 0x00001000

// MetadataPathError supplies stable path and operation context without exposing raw OS text as a fact.
type MetadataPathError struct {
	Operation, Path string
	Err             error
}

func (e *MetadataPathError) Error() string {
	return fmt.Sprintf("metadata %s %q: %v", e.Operation, e.Path, e.Err)
}
func (e *MetadataPathError) Unwrap() error { return e.Err }

type Metadata struct {
	mode               fs.FileMode
	device             uint64
	identity           core.OptionalFilesystemIdentity
	linkCount, logical uint64
	allocation         core.Allocation
}

func (m Metadata) Mode() fs.FileMode                         { return m.mode }
func (m Metadata) Device() uint64                            { return m.device }
func (m Metadata) Identity() core.OptionalFilesystemIdentity { return m.identity }
func (m Metadata) LinkCount() uint64                         { return m.linkCount }
func (m Metadata) LogicalSize() uint64                       { return m.logical }
func (m Metadata) Allocation() core.Allocation               { return m.allocation }

type componentOps struct {
	openRoot func(string, int, uint32) (int, error)
	openat   func(int, string, int, uint32) (int, error)
	fstat    func(int, *syscall.Stat_t) error
	fstatfs  func(int, *syscall.Statfs_t) error
	close    func(int) error
}
type MetadataAdapter struct {
	device uint64
	ops    componentOps
}

func NewMetadataAdapter(device uint64) (MetadataAdapter, error) {
	if device == 0 {
		return MetadataAdapter{}, fmt.Errorf("metadata adapter requires a device")
	}
	return newComponentMetadataAdapter(device, componentOps{openRoot: unix.Open, openat: openat, fstat: syscall.Fstat, fstatfs: syscall.Fstatfs, close: syscall.Close}), nil
}
func newComponentMetadataAdapter(device uint64, ops componentOps) MetadataAdapter {
	return MetadataAdapter{device, ops}
}

// Inspect resolves clean absolute components from a trusted root; only final-descriptor facts escape.
func (a MetadataAdapter) Inspect(path string, requireDirectory bool) (result Metadata, err error) {
	parts, err := parseCleanAbsolutePath(path)
	if err != nil {
		return Metadata{}, a.pathError("path", path, ErrMetadataInaccessible)
	}
	base := syscall.O_RDONLY | syscall.O_NOFOLLOW | syscall.O_CLOEXEC | syscall.O_NONBLOCK
	fd, err := a.ops.openRoot("/", base|syscall.O_DIRECTORY, 0)
	if err != nil {
		return Metadata{}, a.pathError("open", path, classify(err))
	}
	fds := []int{fd}
	defer func() {
		for i := len(fds) - 1; i >= 0; i-- {
			if closeErr := a.ops.close(fds[i]); err == nil && closeErr != nil {
				result = Metadata{}
				err = a.pathError("close", path, ErrMetadataInaccessible)
			}
		}
	}()
	for i, part := range parts {
		flags := base
		if i < len(parts)-1 || requireDirectory {
			flags |= syscall.O_DIRECTORY
		}
		fd, err = a.ops.openat(fd, part, flags, 0)
		if err != nil {
			return Metadata{}, a.pathError("openat", path, classify(err))
		}
		fds = append(fds, fd)
	}
	var stat syscall.Stat_t
	if err := a.ops.fstat(fd, &stat); err != nil {
		return Metadata{}, a.pathError("fstat", path, ErrMetadataInaccessible)
	}
	if !validStat(&stat) {
		return Metadata{}, a.pathError("fstat", path, ErrMetadataInvalid)
	}
	if uint64(stat.Dev) != a.device {
		return Metadata{}, a.pathError("fstat", path, ErrMetadataDeviceMismatch)
	}
	if requireDirectory && stat.Mode&syscall.S_IFMT != syscall.S_IFDIR {
		return Metadata{}, a.pathError("fstat", path, ErrMetadataNotDirectory)
	}
	var volume syscall.Statfs_t
	if err := a.ops.fstatfs(fd, &volume); err != nil {
		return Metadata{}, a.pathError("fstatfs", path, ErrMetadataInaccessible)
	}
	if volume.Flags&localFilesystemFlag == 0 {
		return Metadata{}, a.pathError("fstatfs", path, ErrMetadataNonLocalFilesystem)
	}
	identity, err := core.NewFilesystemIdentity(uint64(stat.Dev), uint64(stat.Ino))
	if err != nil {
		return Metadata{}, a.pathError("fstat", path, ErrMetadataInvalid)
	}
	result = Metadata{mode: fileMode(stat.Mode), device: uint64(stat.Dev), identity: core.KnownFilesystemIdentity(identity), linkCount: uint64(stat.Nlink), allocation: core.UnknownAllocation()}
	if stat.Mode&syscall.S_IFMT == syscall.S_IFREG {
		result.logical = uint64(stat.Size)
		if uint64(stat.Blocks) > ^uint64(0)/512 {
			return Metadata{}, a.pathError("fstat", path, ErrMetadataInvalid)
		}
		result.allocation = core.KnownAllocation(uint64(stat.Blocks) * 512)
	}
	return result, nil
}

func fileMode(mode uint16) fs.FileMode {
	result := fs.FileMode(mode & 0o777)
	switch mode & syscall.S_IFMT {
	case syscall.S_IFDIR:
		result |= fs.ModeDir
	case syscall.S_IFIFO:
		result |= fs.ModeNamedPipe
	case syscall.S_IFLNK:
		result |= fs.ModeSymlink
	case syscall.S_IFSOCK:
		result |= fs.ModeSocket
	case syscall.S_IFCHR:
		result |= fs.ModeCharDevice
	case syscall.S_IFBLK:
		result |= fs.ModeDevice
	}
	return result
}

func parseCleanAbsolutePath(path string) ([]string, error) {
	if path == "/" || len(path) < 2 || path[0] != '/' || path[len(path)-1] == '/' {
		return nil, ErrMetadataInaccessible
	}
	parts := strings.Split(path[1:], "/")
	for _, part := range parts {
		if part == "" || part == "." || part == ".." {
			return nil, ErrMetadataInaccessible
		}
	}
	return parts, nil
}

func openat(dirfd int, path string, flags int, mode uint32) (int, error) {
	return unix.Openat(dirfd, path, flags, mode)
}
func (a MetadataAdapter) pathError(operation, path string, err error) error {
	return &MetadataPathError{Operation: operation, Path: path, Err: err}
}
func classify(err error) error {
	if errors.Is(err, fs.ErrNotExist) {
		return ErrMetadataMissing
	}
	if errors.Is(err, syscall.ELOOP) || errors.Is(err, syscall.EMLINK) {
		return ErrMetadataSymlink
	}
	if errors.Is(err, syscall.ENOTDIR) {
		return ErrMetadataNotDirectory
	}
	return ErrMetadataInaccessible
}
func validStat(stat *syscall.Stat_t) bool {
	return stat.Dev > 0 && stat.Ino > 0 && stat.Nlink > 0 && stat.Size >= 0 && stat.Blocks >= 0
}

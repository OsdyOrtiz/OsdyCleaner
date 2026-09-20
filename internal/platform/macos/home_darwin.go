//go:build darwin

package macos

import (
	"errors"
	"fmt"
	"syscall"

	"github.com/osdy/OsdyCleaner/internal/scan"
	"golang.org/x/sys/unix"
)

// NewHomeInspector validates the process home through no-follow descriptors.
func NewHomeInspector() (scan.HomeInspector, error) { return homeInspector{}, nil }

type homeInspector struct{}

func (homeInspector) InspectHome(home string) error {
	components, err := scan.NewAbsoluteComponents(home)
	if err != nil || components.ValidateScanRoot() != nil {
		return fmt.Errorf("inspect home: %w", scan.ErrInvalidMetadata)
	}
	fd, err := unix.Open("/", walkerFlags, 0)
	if err != nil {
		return homeError(err)
	}
	defer func() { _ = syscall.Close(fd) }()
	for _, component := range components.Components() {
		var entry unix.Stat_t
		if statErr := unix.Fstatat(fd, component, &entry, unix.AT_SYMLINK_NOFOLLOW); statErr != nil {
			return homeError(statErr)
		}
		if entry.Mode&syscall.S_IFMT == syscall.S_IFLNK {
			return fmt.Errorf("inspect home: %w", scan.ErrSymlink)
		}
		next, openErr := unix.Openat(fd, component, walkerFlags, 0)
		if openErr != nil {
			return homeError(openErr)
		}
		if closeErr := syscall.Close(fd); closeErr != nil {
			_ = syscall.Close(next)
			return fmt.Errorf("inspect home: %w", scan.ErrInaccessible)
		}
		fd = next
	}
	var stat syscall.Stat_t
	if err := syscall.Fstat(fd, &stat); err != nil || stat.Mode&syscall.S_IFMT != syscall.S_IFDIR || stat.Dev == 0 || stat.Ino == 0 {
		return fmt.Errorf("inspect home: %w", scan.ErrInaccessible)
	}
	var volume syscall.Statfs_t
	if err := syscall.Fstatfs(fd, &volume); err != nil || volume.Flags&localFilesystemFlag == 0 {
		return fmt.Errorf("inspect home: %w", scan.ErrNonLocal)
	}
	return nil
}

func homeError(err error) error {
	if errors.Is(err, syscall.ELOOP) {
		return fmt.Errorf("inspect home: %w", scan.ErrSymlink)
	}
	return fmt.Errorf("inspect home: %w", scan.ErrInaccessible)
}

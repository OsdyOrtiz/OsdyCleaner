//go:build !darwin

package macos

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/osdy/OsdyCleaner/internal/core"
)

var (
	ErrMetadataMissing            = errors.New("metadata path is missing")
	ErrMetadataSymlink            = errors.New("metadata path is a symbolic link")
	ErrMetadataNotDirectory       = errors.New("metadata path is not a directory")
	ErrMetadataDeviceMismatch     = errors.New("metadata path is on another device")
	ErrMetadataNonLocalFilesystem = errors.New("metadata path is not on a local filesystem")
	ErrMetadataInaccessible       = errors.New("metadata is inaccessible")
	ErrMetadataInvalid            = errors.New("metadata is invalid")
	ErrUnsupportedEnvironment     = errors.New("unsupported environment")
)

// UnsupportedEnvironmentError identifies a platform where macOS metadata is unavailable.
// MetadataPathError supplies stable path and operation context without exposing raw OS text as a fact.
type MetadataPathError struct {
	Operation, Path string
	Err             error
}

func (e *MetadataPathError) Error() string {
	return fmt.Sprintf("metadata %s %q: %v", e.Operation, e.Path, e.Err)
}
func (e *MetadataPathError) Unwrap() error { return e.Err }

type UnsupportedEnvironmentError struct{}

func (UnsupportedEnvironmentError) Error() string {
	return "macOS metadata observation is unsupported on this platform"
}
func (UnsupportedEnvironmentError) Unwrap() error { return ErrUnsupportedEnvironment }

// Metadata mirrors the Darwin observation API without inventing unavailable facts.
type Metadata struct {
	mode       fs.FileMode
	device     uint64
	identity   core.OptionalFilesystemIdentity
	linkCount  uint64
	logical    uint64
	allocation core.Allocation
}

func (m Metadata) Mode() fs.FileMode                         { return m.mode }
func (m Metadata) Device() uint64                            { return m.device }
func (m Metadata) Identity() core.OptionalFilesystemIdentity { return m.identity }
func (m Metadata) LinkCount() uint64                         { return m.linkCount }
func (m Metadata) LogicalSize() uint64                       { return m.logical }
func (m Metadata) Allocation() core.Allocation               { return m.allocation }

type MetadataAdapter struct{}

// NewMetadataAdapter fails before any home resolution or filesystem observation.
func NewMetadataAdapter(uint64) (MetadataAdapter, error) {
	return MetadataAdapter{}, UnsupportedEnvironmentError{}
}

func (MetadataAdapter) Inspect(string, bool) (Metadata, error) {
	return Metadata{}, UnsupportedEnvironmentError{}
}

//go:build !darwin

package macos

import (
	"errors"
	"testing"
)

func TestUnsupportedEnvironmentErrors(t *testing.T) {
	adapter, err := NewMetadataAdapter(1)
	if !errors.Is(err, ErrUnsupportedEnvironment) {
		t.Fatalf("NewMetadataAdapter() error = %v, want unsupported sentinel", err)
	}
	var unsupported UnsupportedEnvironmentError
	if !errors.As(err, &unsupported) {
		t.Fatalf("NewMetadataAdapter() error = %T, want UnsupportedEnvironmentError", err)
	}
	_, err = adapter.Inspect("/no-filesystem-access", false)
	if !errors.Is(err, ErrUnsupportedEnvironment) || !errors.As(err, &unsupported) {
		t.Fatalf("Inspect() error = %v, want typed unsupported error", err)
	}
}

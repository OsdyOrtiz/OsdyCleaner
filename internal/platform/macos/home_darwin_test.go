//go:build darwin

package macos

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/scan"
)

func TestHomeInspectorFixtureBoundaries(t *testing.T) {
	home := t.TempDir()
	link := filepath.Join(t.TempDir(), "home")
	if err := os.Symlink(home, link); err != nil {
		t.Fatal(err)
	}
	inspector, err := NewHomeInspector()
	if err != nil {
		t.Fatal(err)
	}
	realHome, err := filepath.EvalSymlinks(home)
	if err != nil {
		t.Fatal(err)
	}
	if err := inspector.InspectHome(realHome); err != nil {
		t.Fatalf("InspectHome(valid fixture) = %v", err)
	}
	if err := inspector.InspectHome(link); !errors.Is(err, scan.ErrSymlink) {
		t.Fatalf("InspectHome(symlink) = %v, want symlink", err)
	}
}

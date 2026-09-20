package cli

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/scan"
)

type poisonHome struct{ called bool }

func (h *poisonHome) Home() (string, error) {
	h.called = true
	return "", errors.New("home must not be resolved")
}

func TestProductionUnsupportedStopsBeforeHomeResolution(t *testing.T) {
	home := &poisonHome{}
	dependencies := newProductionDependencies(home, nil, nil, errors.Join(scan.ErrUnsupported, errors.New("unsupported")), nil, nil, nil)
	if _, err := dependencies.Scan(context.Background()); !errors.Is(err, scan.ErrUnsupported) || home.called {
		t.Fatalf("err=%v home_called=%t", err, home.called)
	}
}

type nonTerminalFile struct{}

func (nonTerminalFile) Read([]byte) (int, error)  { return 0, io.EOF }
func (nonTerminalFile) Write([]byte) (int, error) { return 0, nil }

func TestProductionTerminalDetectionRequiresBothDescriptors(t *testing.T) {
	if terminal(nonTerminalFile{}, nonTerminalFile{}) {
		t.Fatal("non-file streams accepted as terminals")
	}
}

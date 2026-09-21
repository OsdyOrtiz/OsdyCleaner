package cli

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/core"
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

type productionInspector struct{}

func (productionInspector) InspectHome(string) error { return nil }

type productionWalker struct{ acquisitions int }

func (w *productionWalker) AcquireRoot(context.Context, scan.AbsoluteComponents, scan.AbsoluteComponents, scan.WalkLimits) (scan.TrustedRoot, error) {
	w.acquisitions++
	facts, err := scan.NewRootFacts(scan.Identity{Device: 1, Inode: uint64(w.acquisitions)}, true)
	if err != nil {
		return nil, err
	}
	return scan.NewTrustedRoot(facts, func() error { return nil })
}

func (w *productionWalker) WalkFacts(context.Context, scan.TrustedRoot, func(scan.DirectoryFact) error) error {
	return nil
}

func TestProductionSnapshotForwardsProgressObserver(t *testing.T) {
	walker := &productionWalker{}
	var events []scan.ProgressEvent
	snapshot, err := productionSnapshot(context.Background(), func(event scan.ProgressEvent) {
		events = append(events, event)
	}, fixedHome{home: "/fixture/home"}, productionInspector{}, walker, nil)
	if err != nil || snapshot.Outcome() != core.OutcomeComplete || walker.acquisitions != 5 || len(events) == 0 {
		t.Fatalf("err=%v outcome=%q acquisitions=%d events=%d", err, snapshot.Outcome(), walker.acquisitions, len(events))
	}
}

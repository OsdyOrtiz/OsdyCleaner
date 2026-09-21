package cli

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/osdy/OsdyCleaner/internal/core"
	"github.com/osdy/OsdyCleaner/internal/platform/macos"
	"github.com/osdy/OsdyCleaner/internal/report"
	"github.com/osdy/OsdyCleaner/internal/scan"
	"github.com/osdy/OsdyCleaner/internal/tui"
	"golang.org/x/term"
)

const productionRuleSetVersion = "builtins-v1"

type fixedHome struct {
	home string
	err  error
}

func (h fixedHome) Home() (string, error) { return h.home, h.err }

// ProductionDependencies builds the fixed read-only application outside Cobra.
func ProductionDependencies(input, output, diagnostics *os.File) Dependencies {
	inspector, err := macos.NewHomeInspector()
	if err != nil {
		return newProductionDependencies(nil, nil, nil, err, input, output, diagnostics)
	}
	home, homeErr := os.UserHomeDir()
	return newProductionDependencies(fixedHome{home: home, err: homeErr}, inspector, macos.NewDescriptorWalker(), nil, input, output, diagnostics)
}

func newProductionDependencies(home scan.HomeResolver, inspector scan.HomeInspector, walker scan.DescriptorWalker, initialErr error, input, output, diagnostics *os.File) Dependencies {
	scanSnapshot := func(ctx context.Context, observer scan.ProgressObserver) (core.Snapshot, error) {
		return productionSnapshot(ctx, observer, home, inspector, walker, initialErr)
	}
	return Dependencies{
		Scan: func(ctx context.Context) (core.Snapshot, error) {
			return scanSnapshot(ctx, nil)
		},
		InteractiveScan: func(ctx context.Context) (core.Snapshot, error) {
			return tui.RunScan(ctx, scanSnapshot, input, output)
		},
		RenderText: report.RenderText,
		RenderJSON: report.RenderJSON,
		Input:      input, Output: output, Error: diagnostics, IsTerminal: terminal,
	}
}

// productionSnapshot creates the one finalized schema-v1 snapshot for either CLI mode.
func productionSnapshot(ctx context.Context, observer scan.ProgressObserver, home scan.HomeResolver, inspector scan.HomeInspector, walker scan.DescriptorWalker, initialErr error) (core.Snapshot, error) {
	if initialErr != nil {
		return core.Snapshot{}, initialErr
	}
	if home == nil || inspector == nil || walker == nil {
		return core.Snapshot{}, errors.New("production scan dependencies are unavailable")
	}
	definitions, err := scan.ResolveBuiltins(home, inspector)
	if err != nil {
		return core.Snapshot{}, err
	}
	limits := scan.DefaultWalkLimits()
	scanner := scan.NewScanner(home, definitions, walker, limits)
	if _, err := scanner.ScanWithProgress(ctx, observer); err != nil {
		return core.Snapshot{}, err
	}
	facts, err := scanner.FinalizedFileFacts()
	if err != nil {
		return core.Snapshot{}, err
	}
	policy, err := core.NewScanPolicy(true, true, uint64(limits.MaxEntries), uint64(limits.MaxPathBytes))
	if err != nil {
		return core.Snapshot{}, err
	}
	return core.NewSnapshot(productionRuleSetVersion, policy, facts.Roots(), facts.Findings(), facts.Warnings())
}

func terminal(input io.Reader, output io.Writer) bool {
	in, inputOK := input.(interface{ Fd() uintptr })
	out, outputOK := output.(interface{ Fd() uintptr })
	return inputOK && outputOK && term.IsTerminal(int(in.Fd())) && term.IsTerminal(int(out.Fd()))
}

package tui

import (
	"context"
	"fmt"
	"io"
	"os"

	tea "charm.land/bubbletea/v2"
	"golang.org/x/term"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// Run owns terminal I/O only; the model remains a pure snapshot viewer.
func Run(snapshot core.Snapshot, input, output *os.File) error {
	if input == nil || output == nil {
		return fmt.Errorf("read-only viewer requires terminal input and output")
	}
	return run(snapshot, input, output, func(fd uintptr) bool { return term.IsTerminal(int(fd)) })
}

// RunScan owns an interactive scan session and waits for scanner shutdown.
func RunScan(parent context.Context, fn ScanFunc, input, output *os.File) (core.Snapshot, error) {
	if input == nil || output == nil {
		return core.Snapshot{}, fmt.Errorf("read-only viewer requires terminal input and output")
	}
	return runScan(parent, fn, input, output, func(fd uintptr) bool { return term.IsTerminal(int(fd)) })
}

type program interface{ Run() (tea.Model, error) }
type programOptions struct{ alternateScreen bool }

var startProgram = func(model Model, input io.Reader, output io.Writer, _ programOptions) program {
	return tea.NewProgram(model, tea.WithInput(input), tea.WithOutput(output))
}

func configureProgramModel(model Model, options programOptions) Model {
	model.alternateScreen = options.alternateScreen
	return model
}

func run(snapshot core.Snapshot, input io.Reader, output io.Writer, terminal func(uintptr) bool) error {
	if err := validateTerminal(input, output, terminal); err != nil {
		return err
	}
	options := programOptions{alternateScreen: true}
	_, err := startProgram(configureProgramModel(NewModel(snapshot), options), input, output, options).Run()
	return err
}

func runScan(parent context.Context, fn ScanFunc, input io.Reader, output io.Writer, terminal func(uintptr) bool) (snapshot core.Snapshot, err error) {
	if parent == nil || fn == nil {
		return core.Snapshot{}, fmt.Errorf("interactive scan requires context and scan function")
	}
	if err := validateTerminal(input, output, terminal); err != nil {
		return core.Snapshot{}, err
	}
	model := newScanModel(parent, fn)
	defer func() {
		model.session.cancel()
		finalSnapshot, scanErr := model.session.waitShutdown()
		snapshot = finalSnapshot
		if err == nil {
			err = scanErr
		}
	}()
	options := programOptions{alternateScreen: true}
	_, err = startProgram(configureProgramModel(model, options), input, output, options).Run()
	return core.Snapshot{}, err
}

func validateTerminal(input io.Reader, output io.Writer, terminal func(uintptr) bool) error {
	in, inputOK := input.(interface{ Fd() uintptr })
	out, outputOK := output.(interface{ Fd() uintptr })
	if !inputOK || !outputOK || !terminal(in.Fd()) || !terminal(out.Fd()) {
		return fmt.Errorf("read-only viewer requires terminal input and output")
	}
	return nil
}

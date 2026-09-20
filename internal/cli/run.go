package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/osdy/OsdyCleaner/internal/core"
	"github.com/osdy/OsdyCleaner/internal/scan"
)

type exitError struct {
	code int
	err  error
}

func (e *exitError) Error() string { return e.err.Error() }
func (e *exitError) Unwrap() error { return e.err }

func inputError(format string, args ...any) error {
	return &exitError{code: ExitInput, err: fmt.Errorf(format, args...)}
}
func failure(err error) error {
	return &exitError{code: ExitFailure, err: err}
}

func run(ctx context.Context, d Dependencies, format string) error {
	if d.Scan == nil || d.RenderText == nil || d.RenderJSON == nil || d.View == nil || d.Output == nil {
		return failure(errors.New("read-only scan dependencies are unavailable"))
	}
	if format == "" {
		if d.IsTerminal != nil && d.IsTerminal(d.Input, d.Output) {
			format = "tui"
		} else {
			format = "text"
		}
	}
	switch format {
	case "text", "json":
	case "tui":
		if d.IsTerminal == nil || !d.IsTerminal(d.Input, d.Output) {
			return inputError("read-only viewer requires terminal input and output")
		}
	default:
		return inputError("unsupported format %q", format)
	}
	var snapshot core.Snapshot
	snapshot, err := d.Scan(ctx)
	if err != nil {
		if errors.Is(err, scan.ErrUnsupported) {
			return &exitError{code: ExitUnsupported, err: err}
		}
		return failure(err)
	}
	if format == "tui" {
		if err := d.View(snapshot); err != nil {
			return failure(err)
		}
		return outcomeExit(snapshot)
	}
	var bytes []byte
	if format == "json" {
		bytes, err = d.RenderJSON(snapshot)
	} else {
		bytes, err = d.RenderText(snapshot)
	}
	if err != nil {
		return failure(err)
	}
	if _, err := d.Output.Write(bytes); err != nil {
		return failure(err)
	}
	return outcomeExit(snapshot)
}

func outcomeExit(snapshot core.Snapshot) error {
	switch snapshot.Outcome() {
	case core.OutcomeComplete:
		return nil
	case core.OutcomePartial:
		return &exitError{code: ExitPartial, err: errors.New("scan completed with partial visibility")}
	case core.OutcomeCancelled:
		return &exitError{code: ExitCancelled, err: errors.New("scan cancelled")}
	default:
		return failure(errors.New("invalid scan outcome"))
	}
}

func exitCode(err error) int {
	var exit *exitError
	if errors.As(err, &exit) {
		return exit.code
	}
	return ExitInput
}

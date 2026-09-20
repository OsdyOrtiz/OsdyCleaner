// Package cli adapts the fixed read-only scan application to Cobra.
package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/osdy/OsdyCleaner/internal/core"
	"github.com/spf13/cobra"
)

const (
	ExitComplete    = 0
	ExitFailure     = 1
	ExitInput       = 2
	ExitPartial     = 3
	ExitUnsupported = 4
	ExitCancelled   = 130
)

// Dependencies are application boundaries; production construction belongs to main.
type Dependencies struct {
	Scan       func(context.Context) (core.Snapshot, error)
	RenderText func(core.Snapshot) ([]byte, error)
	RenderJSON func(core.Snapshot) ([]byte, error)
	View       func(core.Snapshot) error
	Input      io.Reader
	Output     io.Writer
	Error      io.Writer
	IsTerminal func(io.Reader, io.Writer) bool
}

// NewCommand constructs the grammar only. Its RunE delegates to the application.
func NewCommand(ctx context.Context, dependencies Dependencies) *cobra.Command {
	var format string
	command := &cobra.Command{
		Use:           "osdy",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	scan := &cobra.Command{
		Use:   "scan",
		Short: "scan fixed read-only cache areas",
		Args:  cobra.NoArgs,
		RunE: func(*cobra.Command, []string) error {
			return run(ctx, dependencies, format)
		},
	}
	scan.Flags().StringVar(&format, "format", "", "output format: text, json, or tui")
	command.AddCommand(scan)
	return command
}

// Execute runs a command with stable diagnostics and returns its process exit class.
func Execute(ctx context.Context, dependencies Dependencies, args []string) int {
	command := NewCommand(ctx, dependencies)
	command.SetArgs(args)
	if err := command.ExecuteContext(ctx); err != nil {
		code := exitCode(err)
		if dependencies.Error != nil && code != ExitPartial && code != ExitCancelled {
			fmt.Fprintf(dependencies.Error, "osdy: %v\n", err)
		}
		return code
	}
	return ExitComplete
}

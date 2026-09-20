package main

import (
	"context"
	"os"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/cli"
)

func TestRunStopsSignalsAndPassesContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	stopped := false
	code := run([]string{"scan"}, cli.Dependencies{}, func(context.Context, ...os.Signal) (context.Context, context.CancelFunc) {
		return ctx, func() { stopped = true }
	}, func(got context.Context, _ cli.Dependencies, args []string) int {
		if got.Err() != context.Canceled || len(args) != 1 || args[0] != "scan" {
			t.Fatalf("context=%v args=%v", got.Err(), args)
		}
		return cli.ExitCancelled
	})
	if code != cli.ExitCancelled || !stopped {
		t.Fatalf("code=%d stopped=%t", code, stopped)
	}
}

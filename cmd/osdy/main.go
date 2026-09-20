package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/osdy/OsdyCleaner/internal/cli"
)

type notifyContext func(context.Context, ...os.Signal) (context.Context, context.CancelFunc)
type execute func(context.Context, cli.Dependencies, []string) int

func main() {
	dependencies := cli.ProductionDependencies(os.Stdin, os.Stdout, os.Stderr)
	os.Exit(run(os.Args[1:], dependencies, signal.NotifyContext, cli.Execute))
}

func run(args []string, dependencies cli.Dependencies, notify notifyContext, execute execute) int {
	ctx, stop := notify(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return execute(ctx, dependencies, args)
}

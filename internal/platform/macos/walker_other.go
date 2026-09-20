//go:build !darwin

package macos

import (
	"context"

	"github.com/osdy/OsdyCleaner/internal/scan"
)

func NewDescriptorWalker() scan.DescriptorWalker { return unsupportedWalker{} }

type unsupportedWalker struct{}

func (unsupportedWalker) WalkFacts(context.Context, scan.TrustedRoot, func(scan.DirectoryFact) error) error {
	return unsupportedWalkerError{}
}

func (unsupportedWalker) AcquireRoot(context.Context, scan.AbsoluteComponents, scan.AbsoluteComponents, scan.WalkLimits) (scan.TrustedRoot, error) {
	return nil, unsupportedWalkerError{}
}

type unsupportedWalkerError struct{}

func (unsupportedWalkerError) Error() string { return "macOS descriptor walker is unsupported" }
func (unsupportedWalkerError) Unwrap() error { return scan.ErrUnsupported }

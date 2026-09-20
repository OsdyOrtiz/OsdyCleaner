//go:build !darwin

package macos

import "github.com/osdy/OsdyCleaner/internal/scan"

func NewHomeInspector() (scan.HomeInspector, error) { return nil, unsupportedWalkerError{} }

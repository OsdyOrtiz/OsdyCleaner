package scan

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// HomeResolver supplies the current home without selecting scan roots.
type HomeResolver interface {
	Home() (string, error)
}

// HomeInspector validates a home boundary without exposing filesystem APIs here.
type HomeInspector interface {
	InspectHome(home string) error
}

// BuiltinDefinition is one fixed, read-only built-in root fact.
type BuiltinDefinition struct {
	AreaID        core.AreaID
	DisplayName   string
	DisplayPath   string
	RelativePath  string
	RootPath      string
	FindingReason core.FindingReason
}

var builtinDefinitions = []BuiltinDefinition{
	{core.AreaNPMCache, "npm cache", "~/.npm", ".npm", "", core.FindingBuiltInNPMCacheEntry},
	{core.AreaHomebrewCache, "Homebrew cache", "~/Library/Caches/Homebrew", "Library/Caches/Homebrew", "", core.FindingBuiltInHomebrewCacheEntry},
	{core.AreaGradleCaches, "Gradle caches", "~/.gradle/caches", ".gradle/caches", "", core.FindingBuiltInGradleCacheEntry},
	{core.AreaXcodeDerivedData, "Xcode DerivedData", "~/Library/Developer/Xcode/DerivedData", "Library/Developer/Xcode/DerivedData", "", core.FindingBuiltInXcodeDerivedDataEntry},
	{core.AreaCoreSimulator, "CoreSimulator", "~/Library/Developer/CoreSimulator", "Library/Developer/CoreSimulator", "", core.FindingBuiltInCoreSimulatorEntry},
}

// ResolveBuiltins returns copied canonical definitions rooted beneath an inspected home.
func ResolveBuiltins(resolver HomeResolver, inspector HomeInspector) ([]BuiltinDefinition, error) {
	home, err := resolver.Home()
	if err != nil {
		return nil, fmt.Errorf("resolve home: %w", err)
	}
	if home == "" || home != strings.TrimSpace(home) || !filepath.IsAbs(home) || home != filepath.Clean(home) {
		return nil, fmt.Errorf("home must be a non-empty absolute path")
	}
	if err := inspector.InspectHome(home); err != nil {
		return nil, fmt.Errorf("inspect home: %w", err)
	}
	definitions := make([]BuiltinDefinition, len(builtinDefinitions))
	copy(definitions, builtinDefinitions)
	for i := range definitions {
		definitions[i].RootPath = filepath.Join(home, definitions[i].RelativePath)
	}
	return definitions, nil
}

package scan

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/core"
)

type fixtureHomeResolver struct {
	home string
	err  error
}

func (r fixtureHomeResolver) Home() (string, error) { return r.home, r.err }

type recordingHomeInspector struct {
	err   error
	calls int
}

func (i *recordingHomeInspector) InspectHome(string) error {
	i.calls++
	return i.err
}

func TestResolveBuiltins(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	inspector := &recordingHomeInspector{}
	got, err := ResolveBuiltins(fixtureHomeResolver{home: home}, inspector)
	if err != nil {
		t.Fatalf("ResolveBuiltins() error = %v", err)
	}
	want := []BuiltinDefinition{
		{AreaID: core.AreaNPMCache, DisplayName: "npm cache", DisplayPath: "~/.npm", RelativePath: ".npm", RootPath: filepath.Join(home, ".npm"), FindingReason: core.FindingBuiltInNPMCacheEntry},
		{AreaID: core.AreaHomebrewCache, DisplayName: "Homebrew cache", DisplayPath: "~/Library/Caches/Homebrew", RelativePath: "Library/Caches/Homebrew", RootPath: filepath.Join(home, "Library/Caches/Homebrew"), FindingReason: core.FindingBuiltInHomebrewCacheEntry},
		{AreaID: core.AreaGradleCaches, DisplayName: "Gradle caches", DisplayPath: "~/.gradle/caches", RelativePath: ".gradle/caches", RootPath: filepath.Join(home, ".gradle/caches"), FindingReason: core.FindingBuiltInGradleCacheEntry},
		{AreaID: core.AreaXcodeDerivedData, DisplayName: "Xcode DerivedData", DisplayPath: "~/Library/Developer/Xcode/DerivedData", RelativePath: "Library/Developer/Xcode/DerivedData", RootPath: filepath.Join(home, "Library/Developer/Xcode/DerivedData"), FindingReason: core.FindingBuiltInXcodeDerivedDataEntry},
		{AreaID: core.AreaCoreSimulator, DisplayName: "CoreSimulator", DisplayPath: "~/Library/Developer/CoreSimulator", RelativePath: "Library/Developer/CoreSimulator", RootPath: filepath.Join(home, "Library/Developer/CoreSimulator"), FindingReason: core.FindingBuiltInCoreSimulatorEntry},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("definitions = %#v; want %#v", got, want)
	}
	if inspector.calls != 1 {
		t.Fatalf("inspector calls = %d; want 1", inspector.calls)
	}
}

func TestResolveBuiltinsHomeValidation(t *testing.T) {
	for _, tc := range []struct {
		name string
		home string
	}{
		{name: "empty"},
		{name: "whitespace", home: " /poison"},
		{name: "relative", home: "relative/home"},
		{name: "invalid absolute whitespace", home: "/poison "},
		{name: "dot component", home: "/fixture/./home"},
		{name: "parent component", home: "/fixture/../home"},
		{name: "redundant separator", home: "/fixture//home"},
		{name: "non-root trailing separator", home: "/fixture/home/"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			inspector := &recordingHomeInspector{err: errors.New("must not inspect invalid home")}
			_, err := ResolveBuiltins(fixtureHomeResolver{home: tc.home}, inspector)
			if err == nil {
				t.Fatal("ResolveBuiltins() succeeded for invalid home")
			}
			if inspector.calls != 0 {
				t.Fatalf("invalid home invoked inspector %d times", inspector.calls)
			}
		})
	}
}

func TestResolveBuiltinsCleanRootHome(t *testing.T) {
	inspector := &recordingHomeInspector{}
	definitions, err := ResolveBuiltins(fixtureHomeResolver{home: "/"}, inspector)
	if err != nil {
		t.Fatalf("ResolveBuiltins(/) error = %v", err)
	}
	if inspector.calls != 1 {
		t.Fatalf("root home inspector calls = %d; want 1", inspector.calls)
	}
	if len(definitions) != 5 || definitions[0].RootPath != "/.npm" {
		t.Fatalf("root home definitions = %#v", definitions)
	}
}

func TestResolveBuiltinsPoisonCases(t *testing.T) {
	resolverErr := errors.New("resolver failed")
	resolverInspector := &recordingHomeInspector{}
	if _, err := ResolveBuiltins(fixtureHomeResolver{err: resolverErr}, resolverInspector); !errors.Is(err, resolverErr) {
		t.Fatalf("resolver error = %v; want wrapped %v", err, resolverErr)
	}
	if resolverInspector.calls != 0 {
		t.Fatalf("resolver failure invoked inspector %d times", resolverInspector.calls)
	}

	poison := &recordingHomeInspector{err: errors.New("poison home is not inspectable")}
	if _, err := ResolveBuiltins(fixtureHomeResolver{home: "/poison"}, poison); err == nil {
		t.Fatal("poison inspectability failure was accepted")
	}
	if poison.calls != 1 {
		t.Fatalf("poison inspector calls = %d; want 1", poison.calls)
	}

	first, err := ResolveBuiltins(fixtureHomeResolver{home: "/fixture/home"}, &recordingHomeInspector{})
	if err != nil {
		t.Fatal(err)
	}
	first[0].DisplayName = "mutated"
	second, err := ResolveBuiltins(fixtureHomeResolver{home: "/fixture/home"}, &recordingHomeInspector{})
	if err != nil {
		t.Fatal(err)
	}
	if second[0].DisplayName != "npm cache" || len(second) != 5 {
		t.Fatalf("returned table leaked mutation or extra roots: %#v", second)
	}
	if first[0].RootPath != "/fixture/home/.npm" || second[1].RootPath != "/fixture/home/Library/Caches/Homebrew" {
		t.Fatalf("lexical root paths = %q, %q", first[0].RootPath, second[1].RootPath)
	}
}

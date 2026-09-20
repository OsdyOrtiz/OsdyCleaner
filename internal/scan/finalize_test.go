package scan

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/core"
)

func TestFinalizeCanonicalOrdering(t *testing.T) {
	identityA := knownIdentity(t, 1, 10)
	identityB := knownIdentity(t, 1, 20)
	roots := []rawRoot{
		{area: core.AreaHomebrewCache, displayName: "Homebrew cache", displayPath: "~/Library/Caches/Homebrew", observations: []rawObservation{{displayPath: "~/Library/Caches/Homebrew/zeta", reason: core.FindingBuiltInHomebrewCacheEntry, logicalBytes: 7, allocation: core.KnownAllocation(4), identity: identityB}, {displayPath: "~/Library/Caches/Homebrew/alpha", reason: core.FindingBuiltInHomebrewCacheEntry, logicalBytes: 3, allocation: core.KnownAllocation(2), identity: identityA}}},
		{area: core.AreaNPMCache, displayName: "npm cache", displayPath: "~/.npm", observations: []rawObservation{{displayPath: "~/.npm/zeta", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 11, allocation: core.KnownAllocation(8), identity: knownIdentity(t, 2, 20)}, {displayPath: "~/.npm/alpha", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 5, allocation: core.KnownAllocation(4), identity: knownIdentity(t, 2, 10)}}},
	}

	first := mustFinalize(t, roots)
	reordered := []rawRoot{roots[1], roots[0]}
	reordered[0].observations = []rawObservation{roots[1].observations[1], roots[1].observations[0]}
	reordered[1].observations = []rawObservation{roots[0].observations[1], roots[0].observations[0]}
	second := mustFinalize(t, reordered)

	want := []string{"npm-cache|~/.npm/alpha|built_in_npm_cache_entry", "npm-cache|~/.npm/zeta|built_in_npm_cache_entry", "homebrew-cache|~/Library/Caches/Homebrew/alpha|built_in_homebrew_cache_entry", "homebrew-cache|~/Library/Caches/Homebrew/zeta|built_in_homebrew_cache_entry"}
	if got := findingKeys(first.findings); !reflect.DeepEqual(got, want) {
		t.Fatalf("finding order = %v; want %v", got, want)
	}
	if !reflect.DeepEqual(factsSignature(first), factsSignature(second)) {
		t.Fatalf("reordered discovery changed facts:\nfirst  %v\nsecond %v", factsSignature(first), factsSignature(second))
	}
}

func TestFinalizeRejectsInvalidRawPath(t *testing.T) {
	_, err := finalizeRawRoots([]rawRoot{{area: core.AreaNPMCache, displayName: "npm cache", displayPath: "~/.npm", observations: []rawObservation{{displayPath: "~/.npm/a/../b", reason: core.FindingBuiltInNPMCacheEntry, identity: core.UnknownFilesystemIdentity(), allocation: core.UnknownAllocation()}}}})
	if err == nil {
		t.Fatal("invalid raw path was silently omitted from warnings")
	}
}

func TestFinalizeTiedInputsDeterministic(t *testing.T) {
	shared := knownIdentity(t, 8, 8)
	a := rawRoot{area: core.AreaNPMCache, displayName: "a", displayPath: "~/.npm/a", observations: []rawObservation{{displayPath: "~/.npm/a/tie", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 9, allocation: core.KnownAllocation(9), identity: shared}, {displayPath: "~/.npm/a/tie", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 3, allocation: core.KnownAllocation(3), identity: shared}}}
	b := rawRoot{area: core.AreaNPMCache, displayName: "b", displayPath: "~/.npm/b", observations: []rawObservation{{displayPath: "~/.npm/b/tie", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 7, allocation: core.KnownAllocation(7), identity: shared}}}
	first := mustFinalize(t, []rawRoot{b, a})
	a.observations[0], a.observations[1] = a.observations[1], a.observations[0]
	second := mustFinalize(t, []rawRoot{a, b})
	if !reflect.DeepEqual(factsSignature(first), factsSignature(second)) {
		t.Fatalf("tied input permutation changed attribution:\nfirst  %v\nsecond %v", factsSignature(first), factsSignature(second))
	}
}

func TestFinalizeRepeatedIdentity(t *testing.T) {
	shared := knownIdentity(t, 9, 99)
	facts := mustFinalize(t, []rawRoot{
		{area: core.AreaHomebrewCache, displayName: "Homebrew cache", displayPath: "~/Library/Caches/Homebrew", observations: []rawObservation{{displayPath: "~/Library/Caches/Homebrew/archive", reason: core.FindingBuiltInHomebrewCacheEntry, logicalBytes: 12, allocation: core.KnownAllocation(8), identity: shared, linkCount: 3}}},
		{area: core.AreaNPMCache, displayName: "npm cache", displayPath: "~/.npm", observations: []rawObservation{{displayPath: "~/.npm/z-link", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 12, allocation: core.KnownAllocation(8), identity: shared, linkCount: 3}, {displayPath: "~/.npm/a-link", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 12, allocation: core.KnownAllocation(8), identity: shared, linkCount: 3}}},
	})

	canonical := findingByPath(t, facts.findings, "~/.npm/a-link")
	if canonical.Estimate().Logical().KnownBytes() != 12 || canonical.Estimate().Allocated().KnownBytes() != 8 {
		t.Fatalf("canonical identity attribution = logical %d allocated %d; want 12 and 8", canonical.Estimate().Logical().KnownBytes(), canonical.Estimate().Allocated().KnownBytes())
	}
	if got := canonical.Accounting(); got.AttributedObjects() != 1 || got.AlreadyAccountedPaths() != 0 {
		t.Fatalf("canonical accounting = %+v; want one attributed object", got)
	}
	for _, displayPath := range []string{"~/.npm/z-link", "~/Library/Caches/Homebrew/archive"} {
		finding := findingByPath(t, facts.findings, displayPath)
		if finding.Estimate().Logical().KnownBytes() != 0 || finding.Estimate().Allocated().KnownBytes() != 0 {
			t.Errorf("later identity path %q contributed bytes: %+v", displayPath, finding.Estimate())
		}
		if got := finding.Accounting(); got.AttributedObjects() != 0 || got.AlreadyAccountedPaths() != 1 || got.HardLinkObservations() != 1 {
			t.Errorf("later identity accounting for %q = %+v", displayPath, got)
		}
		assertRelatedWarning(t, facts.warnings, displayPath, core.WarningIdentityAlreadyAccounted, "~/.npm/a-link")
	}
}

func TestFinalizeFindingGroups(t *testing.T) {
	facts := mustFinalize(t, []rawRoot{
		{area: core.AreaHomebrewCache, displayName: "Homebrew cache", displayPath: "~/Library/Caches/Homebrew"},
		{area: core.AreaNPMCache, displayName: "npm cache", displayPath: "~/.npm", observations: []rawObservation{{displayPath: "~/.npm/pkg/nested/b", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 7, allocation: core.KnownAllocation(6), identity: knownIdentity(t, 1, 3)}, {displayPath: "~/.npm/direct.txt", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 2, allocation: core.KnownAllocation(2), identity: knownIdentity(t, 1, 1)}, {displayPath: "~/.npm/pkg/a", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 5, allocation: core.KnownAllocation(4), identity: knownIdentity(t, 1, 2)}}},
	})

	if got, want := findingKeys(facts.findings), []string{"npm-cache|~/.npm/direct.txt|built_in_npm_cache_entry", "npm-cache|~/.npm/pkg|built_in_npm_cache_entry"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("immediate-child groups = %v; want %v", got, want)
	}
	group := findingByPath(t, facts.findings, "~/.npm/pkg")
	if group.Estimate().Logical().KnownBytes() != 12 || group.Accounting().MeasuredObjects() != 2 {
		t.Fatalf("nested group = estimate %+v accounting %+v; want logical 12 across two objects", group.Estimate(), group.Accounting())
	}
	if len(facts.roots) != 2 || facts.roots[0].AreaID() != core.AreaNPMCache || facts.roots[0].FindingCount() != 2 {
		t.Fatalf("canonical npm root facts = %+v; want two findings", facts.roots)
	}
	empty := facts.roots[1]
	if empty.AreaID() != core.AreaHomebrewCache || empty.FindingCount() != 0 || empty.Estimate().Logical().KnownBytes() != 0 || empty.Estimate().Allocated().KnownBytes() != 0 {
		t.Fatalf("empty root facts = %+v; want complete zero with no findings", empty)
	}
}

func TestFinalizeUnknownMetadata(t *testing.T) {
	facts := mustFinalize(t, []rawRoot{{area: core.AreaNPMCache, displayName: "npm cache", displayPath: "~/.npm", observations: []rawObservation{{displayPath: "~/.npm/uncertain", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 9, allocation: core.UnknownAllocation(), identity: core.UnknownFilesystemIdentity()}}}})

	finding := findingByPath(t, facts.findings, "~/.npm/uncertain")
	if finding.Estimate().Logical().KnownBytes() != 9 {
		t.Fatalf("unknown identity suppressed observed logical bytes: %+v", finding.Estimate())
	}
	allocated := finding.Estimate().Allocated()
	if allocated.KnownBytes() != 0 || allocated.Completeness() != core.CompletenessUnknown {
		t.Fatalf("unknown allocation = %+v; want zero known bytes and unknown completeness", allocated)
	}
	if finding.Estimate().DeduplicationComplete() {
		t.Fatal("unknown identity claimed complete deduplication")
	}
	if got := finding.Accounting(); got.MeasuredObjects() != 1 || got.AttributedObjects() != 1 || got.UnknownIdentityObjects() != 1 {
		t.Fatalf("unknown identity accounting = %+v; want measured and attributed uncertainty", got)
	}
	assertWarning(t, facts.warnings, "~/.npm/uncertain", core.WarningIdentityUnavailable)
	assertWarning(t, facts.warnings, "~/.npm/uncertain", core.WarningAllocationUnavailable)
}

func TestFinalizeTriangulatesTimingUncertaintyAndCaveats(t *testing.T) {
	shared := knownIdentity(t, 3, 30)
	cases := []struct {
		name  string
		roots []rawRoot
	}{
		{
			name: "tied raw inputs across roots",
			roots: []rawRoot{
				{area: core.AreaCoreSimulator, displayName: "CoreSimulator", displayPath: "~/Library/Developer/CoreSimulator", observations: []rawObservation{{displayPath: "~/Library/Developer/CoreSimulator/device", reason: core.FindingBuiltInCoreSimulatorEntry, logicalBytes: 12, allocation: core.KnownAllocation(8), identity: shared}}},
				{area: core.AreaNPMCache, displayName: "npm cache", displayPath: "~/.npm", observations: []rawObservation{{displayPath: "~/.npm/package", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 12, allocation: core.KnownAllocation(8), identity: shared}, {displayPath: "~/.npm/unknown", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 13, allocation: core.UnknownAllocation(), identity: core.UnknownFilesystemIdentity()}}},
			},
		},
		{
			name: "non-tied raw inputs across roots",
			roots: []rawRoot{
				{area: core.AreaCoreSimulator, displayName: "CoreSimulator", displayPath: "~/Library/Developer/CoreSimulator", observations: []rawObservation{{displayPath: "~/Library/Developer/CoreSimulator/device", reason: core.FindingBuiltInCoreSimulatorEntry, logicalBytes: 99, allocation: core.KnownAllocation(64), identity: shared}}},
				{area: core.AreaNPMCache, displayName: "npm cache", displayPath: "~/.npm", observations: []rawObservation{{displayPath: "~/.npm/package", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 12, allocation: core.KnownAllocation(8), identity: shared}, {displayPath: "~/.npm/unknown", reason: core.FindingBuiltInNPMCacheEntry, logicalBytes: 13, allocation: core.UnknownAllocation(), identity: core.UnknownFilesystemIdentity()}}},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			first := mustFinalize(t, tc.roots)
			permuted := []rawRoot{tc.roots[1], tc.roots[0]}
			permuted[0].observations = append([]rawObservation(nil), tc.roots[1].observations...)
			permuted[1].observations = append([]rawObservation(nil), tc.roots[0].observations...)
			second := mustFinalize(t, permuted)
			if !reflect.DeepEqual(factsSignature(first), factsSignature(second)) {
				t.Fatal("worker/discovery timing chose different facts")
			}
			canonical := findingByPath(t, first.findings, "~/.npm/package")
			if canonical.Estimate().Logical().KnownBytes() != 12 || canonical.Accounting().AttributedObjects() != 1 {
				t.Fatalf("canonical npm attribution = %+v %+v", canonical.Estimate(), canonical.Accounting())
			}
			later := findingByPath(t, first.findings, "~/Library/Developer/CoreSimulator/device")
			if later.Estimate().Logical().KnownBytes() != 0 || later.Estimate().Allocated().KnownBytes() != 0 || later.Accounting().AlreadyAccountedPaths() != 1 {
				t.Fatalf("later accounting attribution = %+v %+v", later.Estimate(), later.Accounting())
			}
			unknown := findingByPath(t, first.findings, "~/.npm/unknown")
			if unknown.Estimate().Logical().KnownBytes() != 13 || unknown.Estimate().Allocated().KnownBytes() != 0 || unknown.Estimate().Allocated().Completeness() != core.CompletenessUnknown || unknown.Estimate().DeduplicationComplete() {
				t.Fatalf("unknown metadata estimate = %+v", unknown.Estimate())
			}
			assertStableWarning(t, first.warnings, "~/.npm", core.WarningAPFSEstimateUncertain)
			assertStableWarning(t, first.warnings, "~/Library/Developer/CoreSimulator", core.WarningAPFSEstimateUncertain)
			assertStableWarning(t, first.warnings, "~/Library/Developer/CoreSimulator", core.WarningCoreSimulatorSemantics)
			for _, warning := range first.warnings {
				if strings.Contains(strings.ToLower(warning.Message()), "duplicate") {
					t.Fatalf("accounting warning used duplicate-file language: %q", warning.Message())
				}
			}
		})
	}
}

func assertStableWarning(t *testing.T, warnings []core.Warning, displayPath string, code core.WarningCode) {
	t.Helper()
	for _, warning := range warnings {
		if warning.DisplayPath() == displayPath && warning.Code() == code {
			if warning.Message() != code.Message() {
				t.Fatalf("warning %s message = %q; want %q", code, warning.Message(), code.Message())
			}
			return
		}
	}
	t.Fatalf("stable warning %s at %q not found", code, displayPath)
}

func knownIdentity(t *testing.T, device, inode uint64) core.OptionalFilesystemIdentity {
	t.Helper()
	identity, err := core.NewFilesystemIdentity(device, inode)
	if err != nil {
		t.Fatal(err)
	}
	return core.KnownFilesystemIdentity(identity)
}

func mustFinalize(t *testing.T, roots []rawRoot) finalizedFacts {
	t.Helper()
	facts, err := finalizeRawRoots(roots)
	if err != nil {
		t.Fatalf("finalizeRawRoots: %v", err)
	}
	return facts
}

func findingKeys(findings []core.Finding) []string {
	keys := make([]string, len(findings))
	for i, finding := range findings {
		keys[i] = fmt.Sprintf("%s|%s|%s", finding.AreaID(), finding.DisplayPath(), finding.Reason())
	}
	return keys
}

func findingByPath(t *testing.T, findings []core.Finding, displayPath string) core.Finding {
	t.Helper()
	for _, finding := range findings {
		if finding.DisplayPath() == displayPath {
			return finding
		}
	}
	t.Fatalf("finding %q not found in %v", displayPath, findingKeys(findings))
	return core.Finding{}
}

func factsSignature(facts finalizedFacts) any {
	return []any{facts.roots, facts.findings, facts.warnings}
}

func assertRelatedWarning(t *testing.T, warnings []core.Warning, displayPath string, code core.WarningCode, related string) {
	t.Helper()
	for _, warning := range warnings {
		gotRelated, ok := warning.RelatedPath()
		if warning.DisplayPath() == displayPath && warning.Code() == code && ok && gotRelated == related {
			return
		}
	}
	t.Errorf("warning %s at %q related to %q not found", code, displayPath, related)
}

func assertWarning(t *testing.T, warnings []core.Warning, displayPath string, code core.WarningCode) {
	t.Helper()
	for _, warning := range warnings {
		if warning.DisplayPath() == displayPath && warning.Code() == code {
			return
		}
	}
	t.Errorf("warning %s at %q not found", code, displayPath)
}

package core

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewEstimate(t *testing.T) {
	logical := mustValueEstimate(t, 12, BasisRegularFileStatSize, CompletenessComplete)
	allocated := mustValueEstimate(t, 8, BasisStatBlocks512, CompletenessComplete)
	estimate, err := NewEstimate(logical, allocated, true)
	if err != nil {
		t.Fatalf("NewEstimate: %v", err)
	}
	if estimate.Logical() != logical || estimate.Allocated() != allocated || !estimate.DeduplicationComplete() {
		t.Fatal("estimate accessors changed values")
	}
	for _, tt := range []struct {
		name         string
		bytes        uint64
		basis        EstimateBasis
		completeness EstimateCompleteness
		valid        bool
	}{
		{"logical complete", 1, BasisRegularFileStatSize, CompletenessComplete, true},
		{"logical incomplete", 1, BasisRegularFileStatSize, CompletenessIncomplete, true},
		{"allocation unknown", 0, BasisStatBlocks512, CompletenessUnknown, true},
		{"invalid basis", 1, "other", CompletenessComplete, false},
		{"invalid completeness", 1, BasisStatBlocks512, "other", false},
		{"unknown with bytes", 1, BasisStatBlocks512, CompletenessUnknown, false},
		{"unknown logical", 0, BasisRegularFileStatSize, CompletenessUnknown, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewValueEstimate(tt.bytes, tt.basis, tt.completeness)
			if (err == nil) != tt.valid {
				t.Fatalf("NewValueEstimate error = %v; valid=%t", err, tt.valid)
			}
		})
	}
	incomplete := mustValueEstimate(t, 1, BasisRegularFileStatSize, CompletenessIncomplete)
	if _, err := NewEstimate(incomplete, allocated, true); err == nil {
		t.Fatal("incomplete logical estimate accepted complete allocation")
	}
	for _, counts := range [][5]uint64{{2, 1, 1, 1, 0}, {1, 1, 0, 0, 1}, {0, 0, 0, 0, 0}} {
		if _, err := NewAccounting(counts[0], counts[1], counts[2], counts[3], counts[4]); err != nil {
			t.Errorf("valid accounting %v rejected: %v", counts, err)
		}
	}
	for _, counts := range [][5]uint64{{1, 0, 0, 0, 0}, {1, 2, 0, 0, 0}, {1, 1, 0, 2, 0}, {1, 0, 1, 0, 1}} {
		if _, err := NewAccounting(counts[0], counts[1], counts[2], counts[3], counts[4]); err == nil {
			t.Errorf("inconsistent accounting %v accepted", counts)
		}
	}
}
func TestNewFinding(t *testing.T) {
	estimate := completeEstimate(t)
	accounting, _ := NewAccounting(2, 1, 1, 1, 0)
	input := []WarningCode{WarningSymlinkSkipped, WarningAPFSEstimateUncertain, WarningSymlinkSkipped}
	finding, err := NewFinding(AreaNPMCache, "~/.npm/pkg", FindingBuiltInNPMCacheEntry, estimate, accounting, input)
	if err != nil {
		t.Fatalf("NewFinding: %v", err)
	}
	input[0] = WarningRootInaccessible
	if finding.AreaID() != AreaNPMCache || finding.DisplayPath() != "~/.npm/pkg" ||
		finding.Reason() != FindingBuiltInNPMCacheEntry || finding.Risk() != RiskManualReview ||
		finding.Estimate() != estimate || finding.Accounting() != accounting {
		t.Fatal("finding accessors changed facts or risk")
	}
	assertWarningCodes(t, finding.WarningCodes(), []WarningCode{WarningAPFSEstimateUncertain, WarningSymlinkSkipped})
	for _, tt := range []struct {
		area   AreaID
		path   string
		reason FindingReason
		codes  []WarningCode
	}{
		{"unknown", "~/.npm/pkg", FindingBuiltInNPMCacheEntry, nil},
		{AreaNPMCache, "/tmp/pkg", FindingBuiltInNPMCacheEntry, nil},
		{AreaNPMCache, "~/.npm/../tmp", FindingBuiltInNPMCacheEntry, nil},
		{AreaNPMCache, "~/.npm/pkg", FindingBuiltInGradleCacheEntry, nil},
		{AreaNPMCache, "~/.npm/pkg", FindingBuiltInNPMCacheEntry, []WarningCode{"unknown"}},
	} {
		if _, err := NewFinding(tt.area, tt.path, tt.reason, estimate, accounting, tt.codes); err == nil {
			t.Errorf("invalid finding accepted: %+v", tt)
		}
	}
	if _, err := NewFinding(AreaNPMCache, "~/.npm/pkg", FindingBuiltInNPMCacheEntry, Estimate{}, accounting, nil); err == nil {
		t.Fatal("finding accepted invalid estimate")
	}
}
func TestNewRootObservation(t *testing.T) {
	estimate := completeEstimate(t)
	statuses := []RootStatus{RootScanned, RootMissing, RootInaccessible, RootSkipped, RootBoundaryLimited, RootPartial, RootCancelled}
	reasons := []RootReasonCode{ReasonCompleted, ReasonNotFound, ReasonRootInspectionFailed, ReasonRootSymlink, ReasonRootDeviceBoundary, ReasonRootNotDirectory, ReasonCancelledBeforeStart, ReasonDeviceBoundary, ReasonEntryVisibilityGap, ReasonCancelledDuringScan}
	valid := map[[2]string]bool{
		{string(RootScanned), string(ReasonCompleted)}: true, {string(RootMissing), string(ReasonNotFound)}: true,
		{string(RootInaccessible), string(ReasonRootInspectionFailed)}: true, {string(RootSkipped), string(ReasonRootSymlink)}: true,
		{string(RootSkipped), string(ReasonRootDeviceBoundary)}: true, {string(RootSkipped), string(ReasonRootNotDirectory)}: true,
		{string(RootSkipped), string(ReasonCancelledBeforeStart)}: true, {string(RootBoundaryLimited), string(ReasonDeviceBoundary)}: true,
		{string(RootPartial), string(ReasonEntryVisibilityGap)}: true, {string(RootCancelled), string(ReasonCancelledDuringScan)}: true,
	}
	for _, status := range statuses {
		for _, reason := range reasons {
			_, err := NewRootObservation(AreaNPMCache, "npm cache", "~/.npm", status, reason, estimate, 3, nil)
			if (err == nil) != valid[[2]string{string(status), string(reason)}] {
				t.Errorf("status/reason %s/%s error = %v", status, reason, err)
			}
		}
	}
	for _, tt := range []struct {
		area       AreaID
		name, path string
	}{{"bad", "npm", "~/.npm"}, {AreaNPMCache, "", "~/.npm"}, {AreaNPMCache, "npm", "~/../tmp"}} {
		if _, err := NewRootObservation(tt.area, tt.name, tt.path, RootScanned, ReasonCompleted, estimate, 0, nil); err == nil {
			t.Errorf("invalid root accepted: %+v", tt)
		}
	}
}
func TestWarningCopyBehavior(t *testing.T) {
	related := "~/.npm/original"
	warning, err := NewWarning(AreaNPMCache, "~/.npm/pkg", WarningEntryInaccessible, &related)
	if err != nil {
		t.Fatal(err)
	}
	related = "~/.npm/mutated"
	gotRelated, ok := warning.RelatedPath()
	if !ok || gotRelated != "~/.npm/original" || warning.Message() != WarningEntryInaccessible.Message() || strings.Contains(warning.Message(), "permission denied") {
		t.Fatal("warning did not preserve copied, stable code-owned facts")
	}
	withoutRelated, err := NewWarning(AreaNPMCache, "~/.npm", WarningAPFSEstimateUncertain, nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := withoutRelated.RelatedPath(); ok {
		t.Fatal("absent related path reported present")
	}
	for _, tt := range []struct {
		area    AreaID
		path    string
		code    WarningCode
		related *string
	}{
		{"bad", "~/.npm", WarningEntryChanged, nil}, {AreaNPMCache, "/tmp", WarningEntryChanged, nil},
		{AreaNPMCache, "~/.npm", "bad", nil}, {AreaNPMCache, "~/.npm", WarningEntryChanged, ptr("/tmp")},
	} {
		if _, err := NewWarning(tt.area, tt.path, tt.code, tt.related); err == nil {
			t.Errorf("invalid warning accepted: %+v", tt)
		}
	}
	codes := []WarningCode{WarningSymlinkSkipped, WarningAPFSEstimateUncertain, WarningSymlinkSkipped}
	accounting, _ := NewAccounting(1, 1, 0, 0, 0)
	finding, _ := NewFinding(AreaNPMCache, "~/.npm/pkg", FindingBuiltInNPMCacheEntry, completeEstimate(t), accounting, codes)
	root, _ := NewRootObservation(AreaNPMCache, "npm", "~/.npm", RootScanned, ReasonCompleted, completeEstimate(t), 1, codes)
	codes[0] = WarningRootInaccessible
	returned := finding.WarningCodes()
	returned[0] = WarningRootInaccessible
	assertWarningCodes(t, finding.WarningCodes(), []WarningCode{WarningAPFSEstimateUncertain, WarningSymlinkSkipped})
	assertWarningCodes(t, root.WarningCodes(), []WarningCode{WarningAPFSEstimateUncertain, WarningSymlinkSkipped})
}
func TestNewScanPolicy(t *testing.T) {
	for _, tt := range []struct {
		name                     string
		noLinks, sameDevice      bool
		maxEntries, maxPathBytes uint64
		valid                    bool
	}{
		{"fixed safe policy", true, true, 250000, 64 << 20, true},
		{"allows links", false, true, 1, 1, false},
		{"allows devices", true, false, 1, 1, false},
		{"zero entries", true, true, 0, 1, false},
		{"zero path budget", true, true, 1, 0, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			policy, err := NewScanPolicy(tt.noLinks, tt.sameDevice, tt.maxEntries, tt.maxPathBytes)
			if (err == nil) != tt.valid {
				t.Fatalf("NewScanPolicy error = %v; valid=%t", err, tt.valid)
			}
			if tt.valid && (!policy.NoSymlinkFollowing() || !policy.SameDeviceOnly() || policy.MaxEntriesPerRoot() != tt.maxEntries || policy.MaxPathBytesPerRoot() != tt.maxPathBytes) {
				t.Fatal("policy accessors changed fixed facts")
			}
		})
	}
}
func TestNewSnapshot(t *testing.T) {
	policy := mustPolicy(t)
	roots := canonicalRoots(t, RootScanned)
	roots[1], roots[4] = roots[4], roots[1]
	findings := []Finding{
		mustFinding(t, AreaCoreSimulator, "~/Library/Developer/CoreSimulator/z", 5, 0, CompletenessUnknown),
		mustFinding(t, AreaNPMCache, "~/.npm/b", 7, 4, CompletenessComplete),
		mustFinding(t, AreaNPMCache, "~/.npm/a", 3, 2, CompletenessComplete),
	}
	findings = append(findings, findings[1])
	warnings := []Warning{
		mustWarning(t, AreaCoreSimulator, "~/Library/Developer/CoreSimulator/z", WarningCoreSimulatorSemantics),
		mustWarning(t, AreaNPMCache, "~/.npm/b", WarningEntryChanged),
		mustWarning(t, AreaNPMCache, "~/.npm/a", WarningAPFSEstimateUncertain),
	}
	warnings = append(warnings, warnings[1])
	snapshot, err := NewSnapshot("1", policy, roots, findings, warnings)
	if err != nil {
		t.Fatalf("NewSnapshot: %v", err)
	}
	if snapshot.SchemaVersion() != 1 || snapshot.RuleSetVersion() != "1" || snapshot.ScanPolicy() != policy || snapshot.Outcome() != OutcomeComplete || !snapshot.Complete() {
		t.Fatal("snapshot fixed or derived facts changed")
	}
	if got := snapshot.Roots(); len(got) != 5 || got[0].AreaID() != AreaNPMCache || got[4].AreaID() != AreaCoreSimulator {
		t.Fatalf("roots not canonical: %v", rootAreas(got))
	}
	if got := snapshot.Findings(); len(got) != 3 || got[0].DisplayPath() != "~/.npm/a" || got[2].AreaID() != AreaCoreSimulator {
		t.Fatalf("findings not canonical: %v", findingPaths(got))
	}
	if got := snapshot.Warnings(); len(got) != 3 || got[0].DisplayPath() != "~/.npm/a" || got[2].AreaID() != AreaCoreSimulator {
		t.Fatal("warnings not canonical or deduplicated")
	}
	if got := snapshot.Estimate(); got.Logical().KnownBytes() != 15 || got.Allocated().KnownBytes() != 6 || got.Allocated().Completeness() != CompletenessIncomplete {
		t.Fatalf("aggregate estimate = %#v", got)
	}
	roots[0] = roots[1]
	findings[0] = findings[1]
	warnings[0] = warnings[1]
	gotFindings, gotWarnings := snapshot.Findings(), snapshot.Warnings()
	gotFindings[0] = gotFindings[1]
	gotWarnings[0] = gotWarnings[1]
	if snapshot.Roots()[4].AreaID() != AreaCoreSimulator || snapshot.Findings()[0].DisplayPath() != "~/.npm/a" || snapshot.Warnings()[0].DisplayPath() != "~/.npm/a" {
		t.Fatal("snapshot input or accessor slices exposed mutation")
	}

	badRoots := canonicalRoots(t, RootScanned)
	badRoots[4] = badRoots[0]
	for _, tt := range []struct {
		name     string
		version  string
		roots    []RootObservation
		findings []Finding
	}{
		{"empty version", "", canonicalRoots(t, RootScanned), nil},
		{"missing root", "1", canonicalRoots(t, RootScanned)[:4], nil},
		{"duplicate root", "1", badRoots, nil},
		{"finding outside root", "1", canonicalRoots(t, RootScanned), []Finding{mustFinding(t, AreaNPMCache, "~/.gradle/caches/x", 1, 1, CompletenessComplete)}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewSnapshot(tt.version, policy, tt.roots, tt.findings, nil); err == nil {
				t.Fatal("invalid snapshot accepted")
			}
		})
	}
}

func TestSnapshotAccessors(t *testing.T) {
	roots := canonicalRoots(t, RootMissing)
	snapshot, err := NewSnapshot("1", mustPolicy(t), roots, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if snapshot.Outcome() != OutcomeComplete || !snapshot.Complete() || snapshot.Estimate().Logical().KnownBytes() != 0 || snapshot.Estimate().Allocated().KnownBytes() != 0 || len(snapshot.Findings()) != 0 {
		t.Fatal("all-missing snapshot was not complete zero")
	}
	returned := snapshot.Roots()
	returned[0] = returned[1]
	if snapshot.Roots()[0].AreaID() != AreaNPMCache {
		t.Fatal("root accessor exposed snapshot mutation")
	}
	if got := reflect.TypeOf(snapshot); got.NumField() != 9 {
		t.Fatalf("Snapshot retains unexpected metadata fields: %d", got.NumField())
	}
}

func TestSnapshotOutcomes(t *testing.T) {
	for _, tt := range []struct {
		name              string
		statuses          []RootStatus
		want              Outcome
		complete          bool
		allocation        EstimateCompleteness
		unknownAllocation bool
	}{
		{"scanned and missing", []RootStatus{RootScanned, RootMissing}, OutcomeComplete, true, CompletenessComplete, false},
		{"inaccessible", []RootStatus{RootInaccessible}, OutcomePartial, false, CompletenessIncomplete, false},
		{"skipped", []RootStatus{RootSkipped}, OutcomePartial, false, CompletenessIncomplete, false},
		{"boundary limited", []RootStatus{RootBoundaryLimited}, OutcomePartial, false, CompletenessIncomplete, false},
		{"partial with unknown allocation", []RootStatus{RootPartial}, OutcomePartial, false, CompletenessIncomplete, true},
		{"cancelled with unknown allocation outranks partial", []RootStatus{RootPartial, RootCancelled}, OutcomeCancelled, false, CompletenessIncomplete, true},
		{"safety skip", []RootStatus{RootSkipped}, OutcomePartial, false, CompletenessIncomplete, false},
		{"cancelled active outranks gaps", []RootStatus{RootInaccessible, RootCancelled}, OutcomeCancelled, false, CompletenessIncomplete, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			roots := canonicalRoots(t, RootScanned)
			for i, status := range tt.statuses {
				roots[i] = canonicalRoot(t, roots[i].AreaID(), status)
			}
			var findings []Finding
			if tt.unknownAllocation {
				findings = []Finding{mustFinding(t, AreaNPMCache, "~/.npm/unknown-allocation", 1, 0, CompletenessUnknown)}
			}
			snapshot, err := NewSnapshot("1", mustPolicy(t), roots, findings, nil)
			if err != nil {
				t.Fatal(err)
			}
			if snapshot.Outcome() != tt.want || snapshot.Complete() != tt.complete || snapshot.Estimate().Logical().Completeness() != tt.allocation || snapshot.Estimate().Allocated().Completeness() != tt.allocation {
				t.Fatalf("outcome=%s complete=%t estimate=%v", snapshot.Outcome(), snapshot.Complete(), snapshot.Estimate())
			}
		})
	}
	roots := canonicalRoots(t, RootScanned)
	roots[2] = canonicalRoot(t, AreaGradleCaches, RootSkipped)
	roots[2].reason = ReasonCancelledBeforeStart
	snapshot, err := NewSnapshot("1", mustPolicy(t), roots, nil, nil)
	if err != nil || snapshot.Outcome() != OutcomeCancelled {
		t.Fatalf("cancelled-before-start outcome = %s, %v", snapshot.Outcome(), err)
	}
}

func mustPolicy(t *testing.T) ScanPolicy {
	t.Helper()
	policy, err := NewScanPolicy(true, true, 250000, 64<<20)
	if err != nil {
		t.Fatal(err)
	}
	return policy
}

func canonicalRoots(t *testing.T, status RootStatus) []RootObservation {
	t.Helper()
	areas := []AreaID{AreaNPMCache, AreaHomebrewCache, AreaGradleCaches, AreaXcodeDerivedData, AreaCoreSimulator}
	roots := make([]RootObservation, len(areas))
	for i, area := range areas {
		roots[i] = canonicalRoot(t, area, status)
	}
	return roots
}

func canonicalRoot(t *testing.T, area AreaID, status RootStatus) RootObservation {
	t.Helper()
	names := []string{"npm cache", "Homebrew cache", "Gradle caches", "Xcode DerivedData", "CoreSimulator"}
	paths := []string{"~/.npm", "~/Library/Caches/Homebrew", "~/.gradle/caches", "~/Library/Developer/Xcode/DerivedData", "~/Library/Developer/CoreSimulator"}
	reasons := map[RootStatus]RootReasonCode{RootScanned: ReasonCompleted, RootMissing: ReasonNotFound, RootInaccessible: ReasonRootInspectionFailed, RootSkipped: ReasonRootSymlink, RootBoundaryLimited: ReasonDeviceBoundary, RootPartial: ReasonEntryVisibilityGap, RootCancelled: ReasonCancelledDuringScan}
	rank := AreaRank(area)
	root, err := NewRootObservation(area, names[rank], paths[rank], status, reasons[status], completeEstimate(t), 0, []WarningCode{WarningSymlinkSkipped, WarningSymlinkSkipped})
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func mustFinding(t *testing.T, area AreaID, displayPath string, logical, allocated uint64, completeness EstimateCompleteness) Finding {
	t.Helper()
	estimate, _ := NewEstimate(mustValueEstimate(t, logical, BasisRegularFileStatSize, CompletenessComplete), mustValueEstimate(t, allocated, BasisStatBlocks512, completeness), true)
	accounting, _ := NewAccounting(1, 1, 0, 0, 0)
	finding, err := NewFinding(area, displayPath, mustFindingReason(area), estimate, accounting, []WarningCode{WarningSymlinkSkipped, WarningSymlinkSkipped})
	if err != nil {
		t.Fatal(err)
	}
	return finding
}

func mustFindingReason(area AreaID) FindingReason { reason, _ := findingReason(area); return reason }
func mustWarning(t *testing.T, area AreaID, displayPath string, code WarningCode) Warning {
	t.Helper()
	w, err := NewWarning(area, displayPath, code, nil)
	if err != nil {
		t.Fatal(err)
	}
	return w
}
func rootAreas(roots []RootObservation) []AreaID {
	out := make([]AreaID, len(roots))
	for i := range roots {
		out[i] = roots[i].AreaID()
	}
	return out
}
func findingPaths(findings []Finding) []string {
	out := make([]string, len(findings))
	for i := range findings {
		out[i] = findings[i].DisplayPath()
	}
	return out
}

func mustValueEstimate(t *testing.T, bytes uint64, basis EstimateBasis, completeness EstimateCompleteness) ValueEstimate {
	t.Helper()
	value, err := NewValueEstimate(bytes, basis, completeness)
	if err != nil {
		t.Fatal(err)
	}
	return value
}

func completeEstimate(t *testing.T) Estimate {
	t.Helper()
	estimate, err := NewEstimate(mustValueEstimate(t, 10, BasisRegularFileStatSize, CompletenessComplete), mustValueEstimate(t, 8, BasisStatBlocks512, CompletenessComplete), true)
	if err != nil {
		t.Fatal(err)
	}
	return estimate
}
func assertWarningCodes(t *testing.T, got, want []WarningCode) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("warning codes = %v; want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("warning codes = %v; want %v", got, want)
		}
	}
}
func ptr(value string) *string { return &value }

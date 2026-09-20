package report

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	stdstrings "strings"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/core"
)

var updateGoldens = flag.Bool("update", false, "update report golden files")

func TestJSONReportGoldens(t *testing.T) {
	for _, tt := range []struct {
		name    string
		status  core.RootStatus
		reason  core.RootReasonCode
		unknown bool
	}{
		{"complete", core.RootScanned, core.ReasonCompleted, false},
		{"partial", core.RootPartial, core.ReasonEntryVisibilityGap, true},
		{"cancelled", core.RootCancelled, core.ReasonCancelledDuringScan, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderJSON(reportSnapshot(t, tt.status, tt.reason, tt.unknown))
			if err != nil {
				t.Fatal(err)
			}
			path := filepath.Join("testdata", tt.name+"_v1.json.golden")
			if *updateGoldens {
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(path, got, 0o644); err != nil {
					t.Fatal(err)
				}
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("golden mismatch (-want +got):\n-%s\n+%s", want, got)
			}
			if bytes.Count(got, []byte{'\n'}) != 1 || got[len(got)-1] != '\n' {
				t.Fatalf("golden must have exactly one trailing newline: %q", got)
			}
		})
	}
}

func TestJSONReport(t *testing.T) {
	for _, tt := range []struct {
		name    string
		status  core.RootStatus
		reason  core.RootReasonCode
		unknown bool
		outcome string
	}{
		{"complete", core.RootScanned, core.ReasonCompleted, false, "complete"},
		{"partial", core.RootPartial, core.ReasonEntryVisibilityGap, true, "partial"},
		{"cancelled", core.RootCancelled, core.ReasonCancelledDuringScan, true, "cancelled"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			snapshot := reportSnapshot(t, tt.status, tt.reason, tt.unknown)
			got, err := RenderJSON(snapshot)
			if err != nil {
				t.Fatal(err)
			}
			if got[len(got)-1] != '\n' || bytes.Count(got, []byte{'\n'}) != 1 {
				t.Fatalf("report must have exactly one trailing newline: %q", got)
			}
			var document map[string]any
			if err := json.Unmarshal(got, &document); err != nil {
				t.Fatalf("report must parse: %v", err)
			}
			if document["schema_version"] != float64(1) || document["outcome"] != tt.outcome || document["complete"] != (tt.outcome == "complete") {
				t.Fatalf("unexpected shared facts: %#v", document)
			}
			roots := document["roots"].([]any)
			if len(roots) != 5 || roots[0].(map[string]any)["area_id"] != "npm-cache" || roots[4].(map[string]any)["area_id"] != "core-simulator" {
				t.Fatalf("canonical roots missing: %#v", roots)
			}
			finding := document["findings"].([]any)[0].(map[string]any)
			if finding["risk"] != "manual_review" || finding["estimate"].(map[string]any)["allocated"].(map[string]any)["known_bytes"] == nil {
				t.Fatalf("finding typed facts missing: %#v", finding)
			}
			if finding["estimate"].(map[string]any)["allocated"].(map[string]any)["completeness"] != map[bool]string{true: "unknown", false: "complete"}[tt.unknown] {
				t.Fatalf("allocation completeness missing: %#v", finding)
			}
			warning := document["warnings"].([]any)[0].(map[string]any)
			if warning["related_path"] != nil || warning["message"] == "" {
				t.Fatalf("warning null/message facts missing: %#v", warning)
			}
			if len(document["caveats"].([]any)) < 2 || !bytes.Contains(got, []byte("APFS")) || !bytes.Contains(got, []byte("CoreSimulator")) {
				t.Fatalf("honest caveats missing: %s", got)
			}
		})
	}
}

func TestReportDeterminism(t *testing.T) {
	snapshot := reportSnapshot(t, core.RootPartial, core.ReasonEntryVisibilityGap, true)
	first, err := RenderJSON(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	for range 10 {
		got, err := RenderJSON(snapshot)
		if err != nil || !bytes.Equal(first, got) {
			t.Fatalf("report changed: %v\n%s\n%s", err, first, got)
		}
	}
	if !bytes.HasPrefix(first, []byte(`{"schema_version":1,"rule_set_version":`)) ||
		!bytes.Contains(first, []byte(`"policy":{"no_symlink_following":true,"same_device_only":true`)) ||
		!bytes.Contains(first, []byte(`"roots":[{"area_id":"npm-cache","display_name":"npm cache"`)) {
		t.Fatalf("field order changed: %s", first)
	}
	for _, forbidden := range [][]byte{[]byte("/Users/"), []byte(`"host"`), []byte(`"user"`), []byte(`"time"`), []byte(`"random"`)} {
		if bytes.Contains(first, forbidden) {
			t.Fatalf("report contains forbidden runtime data %q", forbidden)
		}
	}
	empty, err := core.NewSnapshot("rules-v1", snapshot.ScanPolicy(), snapshot.Roots(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	emptyJSON, err := RenderJSON(empty)
	if err != nil || !bytes.Contains(emptyJSON, []byte(`"findings":[],"warnings":[]`)) {
		t.Fatalf("empty collections must be arrays: %v %s", err, emptyJSON)
	}
}

func TestTextReportGoldens(t *testing.T) {
	for _, tt := range []struct {
		name    string
		status  core.RootStatus
		reason  core.RootReasonCode
		unknown bool
	}{
		{"complete", core.RootScanned, core.ReasonCompleted, false},
		{"partial", core.RootPartial, core.ReasonEntryVisibilityGap, true},
		{"cancelled", core.RootCancelled, core.ReasonCancelledDuringScan, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderText(reportSnapshot(t, tt.status, tt.reason, tt.unknown))
			if err != nil {
				t.Fatal(err)
			}
			path := filepath.Join("testdata", tt.name+"_v1.txt.golden")
			if *updateGoldens {
				if err := os.WriteFile(path, got, 0o644); err != nil {
					t.Fatal(err)
				}
			}
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("golden mismatch (-want +got):\n-%s\n+%s", want, got)
			}
			if got[len(got)-1] != '\n' || bytes.HasSuffix(got[:len(got)-1], []byte{'\n'}) {
				t.Fatalf("report needs one trailing newline: %q", got)
			}
		})
	}
}

func TestTextReport(t *testing.T) {
	for _, tt := range []struct {
		name, outcome string
		status        core.RootStatus
		reason        core.RootReasonCode
		unknown       bool
	}{
		{"complete", "complete", core.RootScanned, core.ReasonCompleted, false},
		{"partial", "partial", core.RootPartial, core.ReasonEntryVisibilityGap, true},
		{"cancelled", "cancelled", core.RootCancelled, core.ReasonCancelledDuringScan, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderText(reportSnapshot(t, tt.status, tt.reason, tt.unknown))
			if err != nil {
				t.Fatal(err)
			}
			for _, want := range []string{"Schema: 1", "Outcome: " + tt.outcome, "Manual review required", "regular_file_stat_size", "stat_blocks_512", "APFS clones, snapshots, compression, and hard links", "CoreSimulator data requires product-aware manual review"} {
				if !stdstrings.Contains(string(got), want) {
					t.Errorf("missing %q in %s", want, got)
				}
			}
			for _, root := range reportSnapshot(t, tt.status, tt.reason, tt.unknown).Roots() {
				if !stdstrings.Contains(string(got), root.DisplayPath()) || !stdstrings.Contains(string(got), string(root.Status())) {
					t.Errorf("root omitted: %s", root.AreaID())
				}
			}
			if tt.outcome != "complete" && stdstrings.Index(string(got), "Outcome:") > stdstrings.Index(string(got), "Roots:") {
				t.Fatal("incomplete state is not prominent")
			}
			if !stdstrings.Contains(string(got), "not guaranteed reclaimable space") {
				t.Error("report must deny a reclaim guarantee")
			}
			for _, forbidden := range []string{"safe to delete", "deletion", "cleanup-eligible"} {
				if stdstrings.Contains(stdstrings.ToLower(string(got)), forbidden) {
					t.Errorf("unsafe promise %q", forbidden)
				}
			}
		})
	}
}

func TestReportEquivalence(t *testing.T) {
	snapshot := variedEquivalenceSnapshot(t)
	text, err := RenderText(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	jsonBytes, err := RenderJSON(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	var document map[string]any
	if err := json.Unmarshal(jsonBytes, &document); err != nil {
		t.Fatal(err)
	}
	sections := splitTextReport(t, text)
	policy := document["policy"].(map[string]any)
	requireTextLines(t, sections.header, []string{
		fmt.Sprintf("Schema: %.0f", document["schema_version"]),
		fmt.Sprintf("Rule set: %s", document["rule_set_version"]),
		fmt.Sprintf("Outcome: %s (complete: %t)", document["outcome"], document["complete"]),
		fmt.Sprintf("Policy: no_symlink_following=%t same_device_only=%t max_entries_per_root=%.0f max_path_bytes_per_root=%.0f", policy["no_symlink_following"], policy["same_device_only"], policy["max_entries_per_root"], policy["max_path_bytes_per_root"]),
		textEstimate("Estimate", document["estimate"].(map[string]any)),
	})
	roots, findings, warnings, caveats := document["roots"].([]any), document["findings"].([]any), document["warnings"].([]any), document["caveats"].([]any)
	requireSection(t, sections, "Roots:", len(roots)*2)
	for i, value := range roots {
		root := value.(map[string]any)
		requireTextLines(t, sections.records["Roots:"][i*2:i*2+2], []string{
			fmt.Sprintf("- %s | %s | %s | status=%s reason=%s findings=%.0f warnings=%v", root["area_id"], root["display_name"], root["display_path"], root["status"], root["reason"], root["finding_count"], root["warning_codes"]),
			textEstimate("  estimate", root["estimate"].(map[string]any)),
		})
	}
	requireSection(t, sections, "Findings:", len(findings)*2)
	for i, value := range findings {
		finding := value.(map[string]any)
		accounting := finding["accounting"].(map[string]any)
		requireTextLines(t, sections.records["Findings:"][i*2:i*2+2], []string{
			fmt.Sprintf("- %s | %s | reason=%s risk=%s warnings=%v accounting=%.0f/%.0f/%.0f/%.0f/%.0f", finding["area_id"], finding["display_path"], finding["reason"], finding["risk"], finding["warning_codes"], accounting["measured_objects"], accounting["attributed_objects"], accounting["already_accounted_paths"], accounting["hard_link_observations"], accounting["unknown_identity_objects"]),
			textEstimate("  estimate", finding["estimate"].(map[string]any)),
		})
	}
	requireSection(t, sections, "Warnings:", len(warnings))
	for i, value := range warnings {
		warning := value.(map[string]any)
		related := "none"
		if warning["related_path"] != nil {
			related = warning["related_path"].(string)
		}
		requireTextLines(t, sections.records["Warnings:"][i:i+1], []string{fmt.Sprintf("- %s | %s | code=%s related_path=%s | %s", warning["area_id"], warning["display_path"], warning["code"], related, warning["message"])})
	}
	requireSection(t, sections, "Caveats:", len(caveats)+1)
	requireTextLines(t, sections.records["Caveats:"], []string{
		"- Manual review required. Estimate—not guaranteed reclaimable space.",
		"- " + caveats[0].(string),
		"- " + stdstrings.TrimSuffix(caveats[1].(string), ".") + "; no device lifecycle or cleanup semantics are inferred.",
	})
}

type textReportSections struct {
	header, names []string
	records       map[string][]string
}

func splitTextReport(t *testing.T, text []byte) textReportSections {
	t.Helper()
	markers := []string{"Roots:", "Findings:", "Warnings:", "Caveats:"}
	out := textReportSections{records: make(map[string][]string)}
	current := -1
	for _, line := range stdstrings.Split(stdstrings.TrimSuffix(string(text), "\n"), "\n") {
		if current+1 < len(markers) && line == markers[current+1] {
			current++
			out.names = append(out.names, line)
			continue
		}
		if current < 0 {
			out.header = append(out.header, line)
			continue
		}
		out.records[markers[current]] = append(out.records[markers[current]], line)
	}
	if stdstrings.Join(out.names, "|") != stdstrings.Join(markers, "|") {
		t.Fatalf("sections out of order: %q", out.names)
	}
	return out
}

func requireSection(t *testing.T, sections textReportSections, name string, count int) {
	t.Helper()
	if got := len(sections.records[name]); got != count {
		t.Fatalf("%s record count = %d, want %d: %q", name, got, count, sections.records[name])
	}
}

func requireTextLines(t *testing.T, got, want []string) {
	t.Helper()
	if stdstrings.Join(got, "\n") != stdstrings.Join(want, "\n") {
		t.Errorf("text record mismatch (-want +got):\n-%s\n+%s", stdstrings.Join(want, "\n"), stdstrings.Join(got, "\n"))
	}
}

func textEstimate(label string, estimate map[string]any) string {
	logical, allocated := estimate["logical"].(map[string]any), estimate["allocated"].(map[string]any)
	return fmt.Sprintf("%s: logical=%.0f basis=%s completeness=%s; allocated=%.0f basis=%s completeness=%s; deduplication_complete=%t", label, logical["known_bytes"], logical["basis"], logical["completeness"], allocated["known_bytes"], allocated["basis"], allocated["completeness"], estimate["deduplication_complete"])
}

func variedEquivalenceSnapshot(t *testing.T) core.Snapshot {
	t.Helper()
	estimate := func(logicalBytes, allocatedBytes uint64, logicalState, allocatedState core.EstimateCompleteness, dedup bool) core.Estimate {
		logical, err := core.NewValueEstimate(logicalBytes, core.BasisRegularFileStatSize, logicalState)
		if err != nil {
			t.Fatal(err)
		}
		allocated, err := core.NewValueEstimate(allocatedBytes, core.BasisStatBlocks512, allocatedState)
		if err != nil {
			t.Fatal(err)
		}
		value, err := core.NewEstimate(logical, allocated, dedup)
		if err != nil {
			t.Fatal(err)
		}
		return value
	}
	policy, err := core.NewScanPolicy(true, true, 37, 4097)
	if err != nil {
		t.Fatal(err)
	}
	areas := []core.AreaID{core.AreaNPMCache, core.AreaHomebrewCache, core.AreaGradleCaches, core.AreaXcodeDerivedData, core.AreaCoreSimulator}
	names := []string{"npm cache", "Homebrew cache", "Gradle caches", "Xcode DerivedData", "CoreSimulator"}
	paths := []string{"~/.npm", "~/Library/Caches/Homebrew", "~/.gradle/caches", "~/Library/Developer/Xcode/DerivedData", "~/Library/Developer/CoreSimulator"}
	reasons := []core.FindingReason{core.FindingBuiltInNPMCacheEntry, core.FindingBuiltInHomebrewCacheEntry, core.FindingBuiltInGradleCacheEntry, core.FindingBuiltInXcodeDerivedDataEntry, core.FindingBuiltInCoreSimulatorEntry}
	codes := [][]core.WarningCode{{core.WarningAPFSEstimateUncertain, core.WarningEntryChanged}, {core.WarningIdentityUnavailable}, {core.WarningAllocationUnavailable, core.WarningHardLinkObserved}, {core.WarningPathBudgetReached}, {core.WarningCoreSimulatorSemantics}}
	roots := make([]core.RootObservation, len(areas))
	findings := make([]core.Finding, len(areas))
	warnings := make([]core.Warning, len(areas))
	for i := range areas {
		rootEstimate := estimate(uint64(101+i), map[bool]uint64{true: uint64(201 + i), false: 0}[i != 2], map[bool]core.EstimateCompleteness{true: core.CompletenessIncomplete, false: core.CompletenessComplete}[i%2 == 1], map[bool]core.EstimateCompleteness{true: core.CompletenessUnknown, false: core.CompletenessIncomplete}[i == 2], i%2 == 0)
		roots[i], err = core.NewRootObservation(areas[i], names[i], paths[i], core.RootScanned, core.ReasonCompleted, rootEstimate, 1, codes[i])
		if err != nil {
			t.Fatal(err)
		}
		findingEstimate := estimate(uint64(11+i), map[bool]uint64{true: uint64(21 + i), false: 0}[i != 2], map[bool]core.EstimateCompleteness{true: core.CompletenessIncomplete, false: core.CompletenessComplete}[i%2 == 1], map[bool]core.EstimateCompleteness{true: core.CompletenessUnknown, false: core.CompletenessIncomplete}[i == 2], i%2 == 1)
		accounting, err := core.NewAccounting(uint64(100+i), 90, uint64(10+i), uint64(20+i), uint64(30+i))
		if err != nil {
			t.Fatal(err)
		}
		findings[i], err = core.NewFinding(areas[i], paths[i]+"/entry-"+fmt.Sprint(i), reasons[i], findingEstimate, accounting, codes[i])
		if err != nil {
			t.Fatal(err)
		}
		var related *string
		if i == 2 {
			value := paths[i] + "/related-entry"
			related = &value
		}
		warnings[i], err = core.NewWarning(areas[i], paths[i]+"/warning-"+fmt.Sprint(i), codes[i][0], related)
		if err != nil {
			t.Fatal(err)
		}
	}
	snapshot, err := core.NewSnapshot("rules-varied-v1", policy, roots, findings, warnings)
	if err != nil {
		t.Fatal(err)
	}
	return snapshot
}

func TestTextReportDeterminism(t *testing.T) {
	first, err := RenderText(reportSnapshot(t, core.RootPartial, core.ReasonEntryVisibilityGap, true))
	if err != nil {
		t.Fatal(err)
	}
	for range 10 {
		got, err := RenderText(reportSnapshot(t, core.RootPartial, core.ReasonEntryVisibilityGap, true))
		if err != nil || !bytes.Equal(first, got) {
			t.Fatalf("text changed: %v", err)
		}
	}
}

func reportSnapshot(t *testing.T, status core.RootStatus, reason core.RootReasonCode, unknown bool) core.Snapshot {
	t.Helper()
	policy, _ := core.NewScanPolicy(true, true, 10, 100)
	logical, _ := core.NewValueEstimate(12, core.BasisRegularFileStatSize, map[bool]core.EstimateCompleteness{true: core.CompletenessIncomplete, false: core.CompletenessComplete}[unknown])
	allocated, _ := core.NewValueEstimate(map[bool]uint64{true: 0, false: 8}[unknown], core.BasisStatBlocks512, map[bool]core.EstimateCompleteness{true: core.CompletenessUnknown, false: core.CompletenessComplete}[unknown])
	estimate, _ := core.NewEstimate(logical, allocated, !unknown)
	areas := []core.AreaID{core.AreaNPMCache, core.AreaHomebrewCache, core.AreaGradleCaches, core.AreaXcodeDerivedData, core.AreaCoreSimulator}
	names := []string{"npm cache", "Homebrew cache", "Gradle caches", "Xcode DerivedData", "CoreSimulator"}
	paths := []string{"~/.npm", "~/Library/Caches/Homebrew", "~/.gradle/caches", "~/Library/Developer/Xcode/DerivedData", "~/Library/Developer/CoreSimulator"}
	roots := make([]core.RootObservation, len(areas))
	for i := range areas {
		rootStatus, rootReason := core.RootScanned, core.ReasonCompleted
		if i == 0 {
			rootStatus, rootReason = status, reason
		}
		roots[i], _ = core.NewRootObservation(areas[i], names[i], paths[i], rootStatus, rootReason, estimate, 1, []core.WarningCode{core.WarningAPFSEstimateUncertain})
	}
	accounting, _ := core.NewAccounting(1, 1, 0, 0, 0)
	finding, _ := core.NewFinding(core.AreaNPMCache, "~/.npm/cache", core.FindingBuiltInNPMCacheEntry, estimate, accounting, []core.WarningCode{core.WarningAPFSEstimateUncertain})
	warning, _ := core.NewWarning(core.AreaNPMCache, "~/.npm/cache", core.WarningAPFSEstimateUncertain, nil)
	snapshot, err := core.NewSnapshot("rules-v1", policy, roots, []core.Finding{finding}, []core.Warning{warning})
	if err != nil {
		t.Fatal(err)
	}
	return snapshot
}

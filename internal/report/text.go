package report

import (
	"fmt"
	stdstrings "strings"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// RenderText projects a schema-v1 snapshot without rescanning or recomputing it.
func RenderText(snapshot core.Snapshot) ([]byte, error) {
	if snapshot.SchemaVersion() != 1 {
		return nil, fmt.Errorf("unsupported snapshot schema %d", snapshot.SchemaVersion())
	}
	var out stdstrings.Builder
	policy, estimate := snapshot.ScanPolicy(), snapshot.Estimate()
	fmt.Fprintf(&out, "Schema: %d\nRule set: %s\nOutcome: %s (complete: %t)\n", snapshot.SchemaVersion(), snapshot.RuleSetVersion(), snapshot.Outcome(), snapshot.Complete())
	if !snapshot.Complete() {
		out.WriteString("INCOMPLETE: available facts may not cover every root.\n")
	}
	fmt.Fprintf(&out, "Policy: no_symlink_following=%t same_device_only=%t max_entries_per_root=%d max_path_bytes_per_root=%d\n", policy.NoSymlinkFollowing(), policy.SameDeviceOnly(), policy.MaxEntriesPerRoot(), policy.MaxPathBytesPerRoot())
	writeEstimate(&out, "Estimate", estimate)
	out.WriteString("Roots:\n")
	for _, root := range snapshot.Roots() {
		fmt.Fprintf(&out, "- %s | %s | %s | status=%s reason=%s findings=%d warnings=%v\n", root.AreaID(), root.DisplayName(), root.DisplayPath(), root.Status(), root.Reason(), root.FindingCount(), root.WarningCodes())
		writeEstimate(&out, "  estimate", root.Estimate())
	}
	out.WriteString("Findings:\n")
	for _, finding := range snapshot.Findings() {
		accounting := finding.Accounting()
		fmt.Fprintf(&out, "- %s | %s | reason=%s risk=%s warnings=%v accounting=%d/%d/%d/%d/%d\n", finding.AreaID(), finding.DisplayPath(), finding.Reason(), finding.Risk(), finding.WarningCodes(), accounting.MeasuredObjects(), accounting.AttributedObjects(), accounting.AlreadyAccountedPaths(), accounting.HardLinkObservations(), accounting.UnknownIdentityObjects())
		writeEstimate(&out, "  estimate", finding.Estimate())
	}
	out.WriteString("Warnings:\n")
	for _, warning := range snapshot.Warnings() {
		related := "none"
		if path, ok := warning.RelatedPath(); ok {
			related = path
		}
		fmt.Fprintf(&out, "- %s | %s | code=%s related_path=%s | %s\n", warning.AreaID(), warning.DisplayPath(), warning.Code(), related, warning.Message())
	}
	out.WriteString("Caveats:\n- Manual review required. Estimate—not guaranteed reclaimable space.\n- Logical and allocated sizes are estimates; APFS clones, snapshots, compression, and hard links can make reclaimed space lower or different.\n- CoreSimulator data requires product-aware manual review; no device lifecycle or cleanup semantics are inferred.\n")
	return []byte(out.String()), nil
}

func writeEstimate(out *stdstrings.Builder, label string, estimate core.Estimate) {
	logical, allocated := estimate.Logical(), estimate.Allocated()
	fmt.Fprintf(out, "%s: logical=%d basis=%s completeness=%s; allocated=%d basis=%s completeness=%s; deduplication_complete=%t\n", label, logical.KnownBytes(), logical.Basis(), logical.Completeness(), allocated.KnownBytes(), allocated.Basis(), allocated.Completeness(), estimate.DeduplicationComplete())
}

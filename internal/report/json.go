// Package report projects finalized scan snapshots into stable presentation formats.
package report

import (
	"encoding/json"
	"fmt"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// RenderJSON projects a schema-v1 snapshot without rescanning or recomputing it.
func RenderJSON(snapshot core.Snapshot) ([]byte, error) {
	if snapshot.SchemaVersion() != 1 {
		return nil, fmt.Errorf("unsupported snapshot schema %d", snapshot.SchemaVersion())
	}
	policy := snapshot.ScanPolicy()
	document := jsonReport{
		SchemaVersion: snapshot.SchemaVersion(), RuleSetVersion: snapshot.RuleSetVersion(), Outcome: snapshot.Outcome(), Complete: snapshot.Complete(),
		Policy:   policyDTO{policy.NoSymlinkFollowing(), policy.SameDeviceOnly(), policy.MaxEntriesPerRoot(), policy.MaxPathBytesPerRoot()},
		Estimate: estimateDTOFrom(snapshot.Estimate()), Roots: []rootDTO{}, Findings: []findingDTO{}, Warnings: []warningDTO{}, Caveats: []string{
			"Logical and allocated sizes are estimates; APFS clones, snapshots, compression, and hard links can make reclaimed space lower or different.",
			"CoreSimulator data requires product-aware manual review.",
		},
	}
	for _, root := range snapshot.Roots() {
		document.Roots = append(document.Roots, rootDTO{string(root.AreaID()), root.DisplayName(), root.DisplayPath(), string(root.Status()), string(root.Reason()), estimateDTOFrom(root.Estimate()), root.FindingCount(), strings(root.WarningCodes())})
	}
	for _, finding := range snapshot.Findings() {
		accounting := finding.Accounting()
		document.Findings = append(document.Findings, findingDTO{string(finding.AreaID()), finding.DisplayPath(), string(finding.Reason()), string(finding.Risk()), estimateDTOFrom(finding.Estimate()), accountingDTO{accounting.MeasuredObjects(), accounting.AttributedObjects(), accounting.AlreadyAccountedPaths(), accounting.HardLinkObservations(), accounting.UnknownIdentityObjects()}, strings(finding.WarningCodes())})
	}
	for _, warning := range snapshot.Warnings() {
		var related *string
		if path, ok := warning.RelatedPath(); ok {
			related = &path
		}
		document.Warnings = append(document.Warnings, warningDTO{string(warning.AreaID()), warning.DisplayPath(), string(warning.Code()), warning.Message(), related})
	}
	bytes, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	return append(bytes, '\n'), nil
}

type jsonReport struct {
	SchemaVersion  int          `json:"schema_version"`
	RuleSetVersion string       `json:"rule_set_version"`
	Outcome        core.Outcome `json:"outcome"`
	Complete       bool         `json:"complete"`
	Policy         policyDTO    `json:"policy"`
	Estimate       estimateDTO  `json:"estimate"`
	Roots          []rootDTO    `json:"roots"`
	Findings       []findingDTO `json:"findings"`
	Warnings       []warningDTO `json:"warnings"`
	Caveats        []string     `json:"caveats"`
}
type policyDTO struct {
	NoSymlinkFollowing  bool   `json:"no_symlink_following"`
	SameDeviceOnly      bool   `json:"same_device_only"`
	MaxEntriesPerRoot   uint64 `json:"max_entries_per_root"`
	MaxPathBytesPerRoot uint64 `json:"max_path_bytes_per_root"`
}
type estimateDTO struct {
	Logical               valueEstimateDTO `json:"logical"`
	Allocated             valueEstimateDTO `json:"allocated"`
	DeduplicationComplete bool             `json:"deduplication_complete"`
}
type valueEstimateDTO struct {
	KnownBytes   uint64                    `json:"known_bytes"`
	Basis        core.EstimateBasis        `json:"basis"`
	Completeness core.EstimateCompleteness `json:"completeness"`
}
type rootDTO struct {
	AreaID       string      `json:"area_id"`
	DisplayName  string      `json:"display_name"`
	DisplayPath  string      `json:"display_path"`
	Status       string      `json:"status"`
	Reason       string      `json:"reason"`
	Estimate     estimateDTO `json:"estimate"`
	FindingCount uint64      `json:"finding_count"`
	WarningCodes []string    `json:"warning_codes"`
}
type findingDTO struct {
	AreaID       string        `json:"area_id"`
	DisplayPath  string        `json:"display_path"`
	Reason       string        `json:"reason"`
	Risk         string        `json:"risk"`
	Estimate     estimateDTO   `json:"estimate"`
	Accounting   accountingDTO `json:"accounting"`
	WarningCodes []string      `json:"warning_codes"`
}
type accountingDTO struct {
	MeasuredObjects        uint64 `json:"measured_objects"`
	AttributedObjects      uint64 `json:"attributed_objects"`
	AlreadyAccountedPaths  uint64 `json:"already_accounted_paths"`
	HardLinkObservations   uint64 `json:"hard_link_observations"`
	UnknownIdentityObjects uint64 `json:"unknown_identity_objects"`
}
type warningDTO struct {
	AreaID      string  `json:"area_id"`
	DisplayPath string  `json:"display_path"`
	Code        string  `json:"code"`
	Message     string  `json:"message"`
	RelatedPath *string `json:"related_path"`
}

func estimateDTOFrom(estimate core.Estimate) estimateDTO {
	return estimateDTO{valueEstimateDTO{estimate.Logical().KnownBytes(), estimate.Logical().Basis(), estimate.Logical().Completeness()}, valueEstimateDTO{estimate.Allocated().KnownBytes(), estimate.Allocated().Basis(), estimate.Allocated().Completeness()}, estimate.DeduplicationComplete()}
}
func strings(codes []core.WarningCode) []string {
	out := make([]string, len(codes))
	for i, code := range codes {
		out[i] = string(code)
	}
	return out
}

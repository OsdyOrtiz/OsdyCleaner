package core

import (
	"fmt"
	"path"
	"sort"
	"strings"
)

const snapshotSchemaVersion = 1

type ScanPolicy struct {
	noSymlinkFollowing, sameDeviceOnly     bool
	maxEntriesPerRoot, maxPathBytesPerRoot uint64
}

func NewScanPolicy(noLinks, sameDevice bool, maxEntries, maxPathBytes uint64) (ScanPolicy, error) {
	if !noLinks || !sameDevice || maxEntries == 0 || maxPathBytes == 0 {
		return ScanPolicy{}, fmt.Errorf("invalid scan policy")
	}
	return ScanPolicy{noLinks, sameDevice, maxEntries, maxPathBytes}, nil
}
func (p ScanPolicy) NoSymlinkFollowing() bool    { return p.noSymlinkFollowing }
func (p ScanPolicy) SameDeviceOnly() bool        { return p.sameDeviceOnly }
func (p ScanPolicy) MaxEntriesPerRoot() uint64   { return p.maxEntriesPerRoot }
func (p ScanPolicy) MaxPathBytesPerRoot() uint64 { return p.maxPathBytesPerRoot }

type Snapshot struct {
	schemaVersion  int
	ruleSetVersion string
	outcome        Outcome
	complete       bool
	policy         ScanPolicy
	estimate       Estimate
	roots          []RootObservation
	findings       []Finding
	warnings       []Warning
}

func NewSnapshot(ruleVersion string, policy ScanPolicy, roots []RootObservation, findings []Finding, warnings []Warning) (Snapshot, error) {
	if strings.TrimSpace(ruleVersion) == "" || !validPolicy(policy) || len(roots) != 5 {
		return Snapshot{}, fmt.Errorf("invalid snapshot")
	}
	canonicalRoots := append([]RootObservation(nil), roots...)
	sort.Slice(canonicalRoots, func(i, j int) bool { return AreaRank(canonicalRoots[i].area) < AreaRank(canonicalRoots[j].area) })
	for rank, root := range canonicalRoots {
		if AreaRank(root.area) != rank || !canonicalRootFact(root) {
			return Snapshot{}, fmt.Errorf("invalid snapshot roots")
		}
	}
	canonicalFindings := append([]Finding(nil), findings...)
	for _, finding := range canonicalFindings {
		if !contained(finding.area, finding.displayPath) {
			return Snapshot{}, fmt.Errorf("invalid snapshot finding")
		}
	}
	sort.Slice(canonicalFindings, func(i, j int) bool {
		a, b := canonicalFindings[i], canonicalFindings[j]
		return AreaRank(a.area) < AreaRank(b.area) || (a.area == b.area && (a.displayPath < b.displayPath || (a.displayPath == b.displayPath && a.reason < b.reason)))
	})
	canonicalFindings = deduplicateFindings(canonicalFindings)
	canonicalWarnings := append([]Warning(nil), warnings...)
	for _, warning := range canonicalWarnings {
		if !contained(warning.area, warning.displayPath) {
			return Snapshot{}, fmt.Errorf("invalid snapshot warning")
		}
	}
	sort.Slice(canonicalWarnings, func(i, j int) bool { return warningLess(canonicalWarnings[i], canonicalWarnings[j]) })
	canonicalWarnings = deduplicateWarnings(canonicalWarnings)
	outcome := deriveOutcome(canonicalRoots)
	estimate, err := aggregateEstimate(canonicalFindings, outcome == OutcomeComplete)
	if err != nil {
		return Snapshot{}, err
	}
	return Snapshot{snapshotSchemaVersion, ruleVersion, outcome, outcome == OutcomeComplete, policy, estimate, canonicalRoots, canonicalFindings, canonicalWarnings}, nil
}
func (s Snapshot) SchemaVersion() int       { return s.schemaVersion }
func (s Snapshot) RuleSetVersion() string   { return s.ruleSetVersion }
func (s Snapshot) Outcome() Outcome         { return s.outcome }
func (s Snapshot) Complete() bool           { return s.complete }
func (s Snapshot) ScanPolicy() ScanPolicy   { return s.policy }
func (s Snapshot) Estimate() Estimate       { return s.estimate }
func (s Snapshot) Roots() []RootObservation { return append([]RootObservation(nil), s.roots...) }
func (s Snapshot) Findings() []Finding      { return append([]Finding(nil), s.findings...) }
func (s Snapshot) Warnings() []Warning      { return append([]Warning(nil), s.warnings...) }

func validPolicy(p ScanPolicy) bool {
	_, err := NewScanPolicy(p.noSymlinkFollowing, p.sameDeviceOnly, p.maxEntriesPerRoot, p.maxPathBytesPerRoot)
	return err == nil
}
func canonicalRootFact(root RootObservation) bool {
	names := [...]string{"npm cache", "Homebrew cache", "Gradle caches", "Xcode DerivedData", "CoreSimulator"}
	paths := [...]string{"~/.npm", "~/Library/Caches/Homebrew", "~/.gradle/caches", "~/Library/Developer/Xcode/DerivedData", "~/Library/Developer/CoreSimulator"}
	rank := AreaRank(root.area)
	return rank >= 0 && root.displayName == names[rank] && root.displayPath == paths[rank] && validStatusReason(root.status, root.reason) && validEstimate(root.estimate)
}
func rootPath(area AreaID) string {
	paths := [...]string{"~/.npm", "~/Library/Caches/Homebrew", "~/.gradle/caches", "~/Library/Developer/Xcode/DerivedData", "~/Library/Developer/CoreSimulator"}
	if rank := AreaRank(area); rank >= 0 {
		return paths[rank]
	}
	return ""
}
func contained(area AreaID, displayPath string) bool {
	root := rootPath(area)
	return displayPath == root || strings.HasPrefix(displayPath, root+"/")
}
func deriveOutcome(roots []RootObservation) Outcome {
	outcome := OutcomeComplete
	for _, root := range roots {
		if root.status == RootCancelled || (root.status == RootSkipped && root.reason == ReasonCancelledBeforeStart) {
			return OutcomeCancelled
		}
		if root.status != RootScanned && root.status != RootMissing {
			outcome = OutcomePartial
		}
	}
	return outcome
}
func aggregateEstimate(findings []Finding, complete bool) (Estimate, error) {
	var logical, allocated uint64
	logicalState, allocatedState := CompletenessComplete, CompletenessComplete
	dedup, anyAllocation := true, false
	for _, finding := range findings {
		e := finding.estimate
		if ^uint64(0)-logical < e.logical.knownBytes || ^uint64(0)-allocated < e.allocated.knownBytes {
			return Estimate{}, fmt.Errorf("snapshot estimate overflow")
		}
		logical += e.logical.knownBytes
		allocated += e.allocated.knownBytes
		if e.logical.completeness != CompletenessComplete {
			logicalState = CompletenessIncomplete
		}
		if e.allocated.completeness == CompletenessUnknown {
			allocatedState = CompletenessUnknown
		} else {
			anyAllocation = true
			if e.allocated.completeness == CompletenessIncomplete {
				allocatedState = CompletenessIncomplete
			}
		}
		dedup = dedup && e.deduplicationComplete
	}
	if allocatedState == CompletenessUnknown && anyAllocation {
		allocatedState = CompletenessIncomplete
	}
	if !complete {
		logicalState = CompletenessIncomplete
		allocatedState = CompletenessIncomplete
	}
	logicalValue, _ := NewValueEstimate(logical, BasisRegularFileStatSize, logicalState)
	allocatedValue, _ := NewValueEstimate(allocated, BasisStatBlocks512, allocatedState)
	return NewEstimate(logicalValue, allocatedValue, dedup)
}
func deduplicateFindings(in []Finding) []Finding {
	out := in[:0]
	for _, finding := range in {
		if len(out) == 0 || finding.area != out[len(out)-1].area || finding.displayPath != out[len(out)-1].displayPath || finding.reason != out[len(out)-1].reason {
			out = append(out, finding)
		}
	}
	return out
}
func deduplicateWarnings(in []Warning) []Warning {
	out := in[:0]
	for _, warning := range in {
		if len(out) == 0 || warningLess(out[len(out)-1], warning) || warningLess(warning, out[len(out)-1]) {
			out = append(out, warning)
		}
	}
	return out
}
func warningLess(a, b Warning) bool {
	if AreaRank(a.area) != AreaRank(b.area) {
		return AreaRank(a.area) < AreaRank(b.area)
	}
	if a.displayPath != b.displayPath {
		return a.displayPath < b.displayPath
	}
	if a.code != b.code {
		return a.code < b.code
	}
	ar, _ := a.RelatedPath()
	br, _ := b.RelatedPath()
	return ar < br
}

type FindingReason string

const (
	FindingBuiltInNPMCacheEntry, FindingBuiltInHomebrewCacheEntry       FindingReason = "built_in_npm_cache_entry", "built_in_homebrew_cache_entry"
	FindingBuiltInGradleCacheEntry, FindingBuiltInXcodeDerivedDataEntry FindingReason = "built_in_gradle_cache_entry", "built_in_xcode_derived_data_entry"
	FindingBuiltInCoreSimulatorEntry                                    FindingReason = "built_in_core_simulator_entry"
)

type Finding struct {
	area         AreaID
	displayPath  string
	reason       FindingReason
	estimate     Estimate
	accounting   Accounting
	warningCodes []WarningCode
}

func NewFinding(area AreaID, displayPath string, reason FindingReason, estimate Estimate, accounting Accounting, codes []WarningCode) (Finding, error) {
	wantReason, validArea := findingReason(area)
	canonical, err := canonicalWarningCodes(codes)
	if !validArea || !validDisplayPath(displayPath) || reason != wantReason || !validEstimate(estimate) || err != nil {
		return Finding{}, fmt.Errorf("invalid finding")
	}
	return Finding{area, displayPath, reason, estimate, accounting, canonical}, nil
}
func (f Finding) AreaID() AreaID              { return f.area }
func (f Finding) DisplayPath() string         { return f.displayPath }
func (f Finding) Reason() FindingReason       { return f.reason }
func (Finding) Risk() Risk                    { return RiskManualReview }
func (f Finding) Estimate() Estimate          { return f.estimate }
func (f Finding) Accounting() Accounting      { return f.accounting }
func (f Finding) WarningCodes() []WarningCode { return append([]WarningCode(nil), f.warningCodes...) }

type RootObservation struct {
	area                     AreaID
	displayName, displayPath string
	status                   RootStatus
	reason                   RootReasonCode
	estimate                 Estimate
	findingCount             uint64
	warningCodes             []WarningCode
}

func NewRootObservation(area AreaID, name, displayPath string, status RootStatus, reason RootReasonCode, estimate Estimate, findings uint64, codes []WarningCode) (RootObservation, error) {
	canonical, err := canonicalWarningCodes(codes)
	if AreaRank(area) < 0 || strings.TrimSpace(name) == "" || !validDisplayPath(displayPath) ||
		!validStatusReason(status, reason) || !validEstimate(estimate) || err != nil {
		return RootObservation{}, fmt.Errorf("invalid root observation")
	}
	return RootObservation{area, name, displayPath, status, reason, estimate, findings, canonical}, nil
}
func (r RootObservation) AreaID() AreaID         { return r.area }
func (r RootObservation) DisplayName() string    { return r.displayName }
func (r RootObservation) DisplayPath() string    { return r.displayPath }
func (r RootObservation) Status() RootStatus     { return r.status }
func (r RootObservation) Reason() RootReasonCode { return r.reason }
func (r RootObservation) Estimate() Estimate     { return r.estimate }
func (r RootObservation) FindingCount() uint64   { return r.findingCount }
func (r RootObservation) WarningCodes() []WarningCode {
	return append([]WarningCode(nil), r.warningCodes...)
}

type Warning struct {
	area        AreaID
	displayPath string
	code        WarningCode
	relatedPath *string
}

func NewWarning(area AreaID, displayPath string, code WarningCode, relatedPath *string) (Warning, error) {
	if AreaRank(area) < 0 || !validDisplayPath(displayPath) || code.Message() == "" ||
		(relatedPath != nil && !validDisplayPath(*relatedPath)) {
		return Warning{}, fmt.Errorf("invalid warning")
	}
	var related *string
	if relatedPath != nil {
		value := *relatedPath
		related = &value
	}
	return Warning{area, displayPath, code, related}, nil
}
func (w Warning) AreaID() AreaID      { return w.area }
func (w Warning) DisplayPath() string { return w.displayPath }
func (w Warning) Code() WarningCode   { return w.code }
func (w Warning) RelatedPath() (string, bool) {
	if w.relatedPath == nil {
		return "", false
	}
	return *w.relatedPath, true
}
func (w Warning) Message() string { return w.code.Message() }

func validEstimate(e Estimate) bool {
	_, err := NewEstimate(e.logical, e.allocated, e.deduplicationComplete)
	return err == nil
}
func validDisplayPath(value string) bool {
	return value == "~" || (strings.HasPrefix(value, "~/") && !strings.Contains(value, "\\") && path.Clean(value) == value)
}
func findingReason(area AreaID) (FindingReason, bool) {
	switch area {
	case AreaNPMCache:
		return FindingBuiltInNPMCacheEntry, true
	case AreaHomebrewCache:
		return FindingBuiltInHomebrewCacheEntry, true
	case AreaGradleCaches:
		return FindingBuiltInGradleCacheEntry, true
	case AreaXcodeDerivedData:
		return FindingBuiltInXcodeDerivedDataEntry, true
	case AreaCoreSimulator:
		return FindingBuiltInCoreSimulatorEntry, true
	default:
		return "", false
	}
}
func validStatusReason(status RootStatus, reason RootReasonCode) bool {
	switch status {
	case RootScanned:
		return reason == ReasonCompleted
	case RootMissing:
		return reason == ReasonNotFound
	case RootInaccessible:
		return reason == ReasonRootInspectionFailed
	case RootSkipped:
		return reason == ReasonRootSymlink || reason == ReasonRootDeviceBoundary || reason == ReasonRootNotDirectory || reason == ReasonCancelledBeforeStart
	case RootBoundaryLimited:
		return reason == ReasonDeviceBoundary
	case RootPartial:
		return reason == ReasonEntryVisibilityGap
	case RootCancelled:
		return reason == ReasonCancelledDuringScan
	default:
		return false
	}
}
func canonicalWarningCodes(input []WarningCode) ([]WarningCode, error) {
	seen := make(map[WarningCode]bool, len(input))
	out := make([]WarningCode, 0, len(input))
	for _, code := range input {
		if code.Message() == "" {
			return nil, fmt.Errorf("invalid warning code")
		}
		if !seen[code] {
			seen[code] = true
			out = append(out, code)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out, nil
}

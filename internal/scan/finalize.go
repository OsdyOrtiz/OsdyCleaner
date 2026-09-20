package scan

import (
	"cmp"
	"sort"
	"strings"

	"github.com/osdy/OsdyCleaner/internal/core"
)

type rawObservation struct {
	displayPath  string
	reason       core.FindingReason
	logicalBytes uint64
	allocation   core.Allocation
	identity     core.OptionalFilesystemIdentity
	linkCount    uint64
	nonRegular   bool
}
type rawRoot struct {
	area                     core.AreaID
	displayName, displayPath string
	observations             []rawObservation
}
type finalizedFacts struct {
	roots    []core.RootObservation
	findings []core.Finding
	warnings []core.Warning
}

// FinalizedFileFacts is the immutable-by-convention scanner result boundary.
// Its accessors return copied slices so presentation consumers cannot mutate
// scanner-owned canonical findings, warnings, or root observations.
type FinalizedFileFacts struct{ facts finalizedFacts }

func (f FinalizedFileFacts) Roots() []core.RootObservation {
	return append([]core.RootObservation(nil), f.facts.roots...)
}
func (f FinalizedFileFacts) Findings() []core.Finding {
	return append([]core.Finding(nil), f.facts.findings...)
}
func (f FinalizedFileFacts) Warnings() []core.Warning {
	return append([]core.Warning(nil), f.facts.warnings...)
}

// finalizeDirectoryObservation keeps scanner callback state on the same immutable
// root-observation construction boundary as file finalization.
func finalizeDirectoryObservation(definition BuiltinDefinition, status core.RootStatus, reason core.RootReasonCode, codes []core.WarningCode) core.RootObservation {
	return rootObservation(definition, status, reason, codes)
}

type findingAccumulator struct {
	area                                                                          core.AreaID
	path                                                                          string
	reason                                                                        core.FindingReason
	logical, allocated, measured, attributed, already, hardLinks, unknownIdentity uint64
	allocationKnown, allocationUnknown, deduplicationComplete                     bool
	codes                                                                         []core.WarningCode
}

func finalizeRawRoots(input []rawRoot) (finalizedFacts, error) {
	roots := append([]rawRoot(nil), input...)
	for i := range roots {
		roots[i].observations = append([]rawObservation(nil), roots[i].observations...)
		sort.Slice(roots[i].observations, func(a, b int) bool { return observationCompare(roots[i].observations[a], roots[i].observations[b]) < 0 })
	}
	sort.Slice(roots, func(i, j int) bool { return rootCompare(roots[i], roots[j]) < 0 })
	seen := make(map[core.FilesystemIdentity]string)
	facts := finalizedFacts{}
	for _, root := range roots {
		observations := root.observations
		groups := make([]findingAccumulator, 0)
		var rootCodes []core.WarningCode
		for _, observation := range observations {
			// This constructor is the core-owned invariant for area and normalized display path.
			if _, err := core.NewWarning(root.area, observation.displayPath, core.WarningAPFSEstimateUncertain, nil); err != nil {
				return finalizedFacts{}, err
			}
			if observation.nonRegular {
				continue
			}
			groupPath := immediateChild(root.displayPath, observation.displayPath)
			if len(groups) == 0 || groups[len(groups)-1].path != groupPath || groups[len(groups)-1].reason != observation.reason {
				groups = append(groups, findingAccumulator{area: root.area, path: groupPath, reason: observation.reason, deduplicationComplete: true})
			}
			group := &groups[len(groups)-1]
			group.measured++
			codes := make([]core.WarningCode, 0, 3)
			attributed := true
			if identity, known := observation.identity.Value(); known {
				if canonical, exists := seen[identity]; exists {
					attributed = false
					group.already++
					group.codes = append(group.codes, core.WarningIdentityAlreadyAccounted)
					rootCodes = append(rootCodes, core.WarningIdentityAlreadyAccounted)
					warning, err := core.NewWarning(root.area, observation.displayPath, core.WarningIdentityAlreadyAccounted, &canonical)
					if err != nil {
						return finalizedFacts{}, err
					}
					facts.warnings = append(facts.warnings, warning)
				} else {
					seen[identity] = observation.displayPath
				}
			} else {
				group.unknownIdentity++
				group.deduplicationComplete = false
				codes = append(codes, core.WarningIdentityUnavailable)
			}
			if observation.linkCount > 1 {
				group.hardLinks++
				codes = append(codes, core.WarningHardLinkObserved)
			}
			if attributed {
				group.attributed++
				group.logical += observation.logicalBytes
				if bytes, known := observation.allocation.Bytes(); known {
					group.allocated += bytes
					group.allocationKnown = true
				} else {
					group.allocationUnknown = true
					codes = append(codes, core.WarningAllocationUnavailable)
				}
			}
			for _, code := range codes {
				warning, err := core.NewWarning(root.area, observation.displayPath, code, nil)
				if err != nil {
					return finalizedFacts{}, err
				}
				facts.warnings = append(facts.warnings, warning)
				group.codes = append(group.codes, code)
				rootCodes = append(rootCodes, code)
			}
		}
		start := len(facts.findings)
		for _, group := range groups {
			finding, err := group.finding()
			if err != nil {
				return finalizedFacts{}, err
			}
			facts.findings = append(facts.findings, finding)
		}
		estimate, err := estimateFindings(facts.findings[start:])
		if err != nil {
			return finalizedFacts{}, err
		}
		if len(groups) > 0 {
			caveats := []core.WarningCode{core.WarningAPFSEstimateUncertain}
			if root.area == core.AreaCoreSimulator {
				caveats = append(caveats, core.WarningCoreSimulatorSemantics)
			}
			for _, code := range caveats {
				warning, err := core.NewWarning(root.area, root.displayPath, code, nil)
				if err != nil {
					return finalizedFacts{}, err
				}
				facts.warnings = append(facts.warnings, warning)
				rootCodes = append(rootCodes, code)
			}
		}
		rootFact, err := core.NewRootObservation(root.area, root.displayName, root.displayPath, core.RootScanned, core.ReasonCompleted, estimate, uint64(len(groups)), rootCodes)
		if err != nil {
			return finalizedFacts{}, err
		}
		facts.roots = append(facts.roots, rootFact)
	}
	sort.Slice(facts.warnings, func(i, j int) bool { return warningKey(facts.warnings[i]) < warningKey(facts.warnings[j]) })
	facts.warnings = deduplicateWarnings(facts.warnings)
	return facts, nil
}

func (group findingAccumulator) finding() (core.Finding, error) {
	// Fixed bases and allocationState's exhaustive enum make these validating errors unreachable.
	logical, _ := core.NewValueEstimate(group.logical, core.BasisRegularFileStatSize, core.CompletenessComplete)
	allocated, _ := core.NewValueEstimate(group.allocated, core.BasisStatBlocks512, allocationState(group.allocationKnown, group.allocationUnknown))
	estimate, err := core.NewEstimate(logical, allocated, group.deduplicationComplete)
	if err != nil {
		return core.Finding{}, err
	}
	accounting, err := core.NewAccounting(group.measured, group.attributed, group.already, group.hardLinks, group.unknownIdentity)
	if err != nil {
		return core.Finding{}, err
	}
	return core.NewFinding(group.area, group.path, group.reason, estimate, accounting, group.codes)
}

func estimateFindings(findings []core.Finding) (core.Estimate, error) {
	var logical, allocated uint64
	known, unknown, dedup := false, false, true
	for _, finding := range findings {
		estimate := finding.Estimate()
		logical += estimate.Logical().KnownBytes()
		allocated += estimate.Allocated().KnownBytes()
		known = known || estimate.Allocated().Completeness() != core.CompletenessUnknown
		unknown = unknown || estimate.Allocated().Completeness() != core.CompletenessComplete
		dedup = dedup && estimate.DeduplicationComplete()
	}
	// Fixed bases and allocationState's exhaustive enum make these validating errors unreachable.
	logicalValue, _ := core.NewValueEstimate(logical, core.BasisRegularFileStatSize, core.CompletenessComplete)
	allocatedValue, _ := core.NewValueEstimate(allocated, core.BasisStatBlocks512, allocationState(known, unknown))
	return core.NewEstimate(logicalValue, allocatedValue, dedup)
}

func rootCompare(a, b rawRoot) int {
	if order := cmp.Or(cmp.Compare(core.AreaRank(a.area), core.AreaRank(b.area)), cmp.Compare(a.displayPath, b.displayPath), cmp.Compare(a.displayName, b.displayName)); order != 0 {
		return order
	}
	for i := 0; i < len(a.observations) && i < len(b.observations); i++ {
		if order := observationCompare(a.observations[i], b.observations[i]); order != 0 {
			return order
		}
	}
	return cmp.Compare(len(a.observations), len(b.observations))
}
func observationCompare(a, b rawObservation) int {
	aa, ak := a.allocation.Bytes()
	ba, bk := b.allocation.Bytes()
	ai, aik := a.identity.Value()
	bi, bik := b.identity.Value()
	return cmp.Or(cmp.Compare(a.displayPath, b.displayPath), cmp.Compare(boolRank(a.nonRegular), boolRank(b.nonRegular)), cmp.Compare(a.reason, b.reason), cmp.Compare(a.logicalBytes, b.logicalBytes), cmp.Compare(boolRank(ak), boolRank(bk)), cmp.Compare(aa, ba), cmp.Compare(boolRank(aik), boolRank(bik)), cmp.Compare(ai.Device(), bi.Device()), cmp.Compare(ai.Inode(), bi.Inode()), cmp.Compare(a.linkCount, b.linkCount))
}
func boolRank(value bool) int {
	if value {
		return 1
	}
	return 0
}

func allocationState(known, unknown bool) core.EstimateCompleteness {
	if unknown && known {
		return core.CompletenessIncomplete
	}
	if unknown {
		return core.CompletenessUnknown
	}
	return core.CompletenessComplete
}
func immediateChild(root, displayPath string) string {
	remainder := strings.TrimPrefix(displayPath, root+"/")
	return root + "/" + strings.SplitN(remainder, "/", 2)[0]
}
func warningKey(w core.Warning) string {
	related, _ := w.RelatedPath()
	return string(rune('0'+core.AreaRank(w.AreaID()))) + "|" + w.DisplayPath() + "|" + string(w.Code()) + "|" + related
}
func deduplicateWarnings(input []core.Warning) []core.Warning {
	out := input[:0]
	for _, warning := range input {
		if len(out) == 0 || warningKey(out[len(out)-1]) != warningKey(warning) {
			out = append(out, warning)
		}
	}
	return out
}

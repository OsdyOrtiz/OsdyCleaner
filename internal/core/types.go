package core

import "fmt"

// AreaID identifies one built-in scan area.
type AreaID string

const (
	AreaNPMCache, AreaHomebrewCache        AreaID = "npm-cache", "homebrew-cache"
	AreaGradleCaches, AreaXcodeDerivedData AreaID = "gradle-caches", "xcode-derived-data"
	AreaCoreSimulator                      AreaID = "core-simulator"
)

func ParseAreaID(value string) (AreaID, error) {
	id := AreaID(value)
	if AreaRank(id) < 0 {
		return "", invalidValue("area ID", value)
	}
	return id, nil
}

// AreaRank returns the canonical area order, or -1 for an unknown ID.
func AreaRank(id AreaID) int {
	switch id {
	case AreaNPMCache:
		return 0
	case AreaHomebrewCache:
		return 1
	case AreaGradleCaches:
		return 2
	case AreaXcodeDerivedData:
		return 3
	case AreaCoreSimulator:
		return 4
	default:
		return -1
	}
}

type RootStatus string

const (
	RootScanned, RootMissing         RootStatus = "scanned", "missing"
	RootInaccessible, RootSkipped    RootStatus = "inaccessible", "skipped"
	RootBoundaryLimited, RootPartial RootStatus = "boundary_limited", "partial"
	RootCancelled                    RootStatus = "cancelled"
)

func ParseRootStatus(value string) (RootStatus, error) {
	status := RootStatus(value)
	switch status {
	case RootScanned, RootMissing, RootInaccessible, RootSkipped,
		RootBoundaryLimited, RootPartial, RootCancelled:
		return status, nil
	default:
		return "", invalidValue("root status", value)
	}
}

type RootReasonCode string

const (
	ReasonCompleted, ReasonNotFound                     RootReasonCode = "completed", "not_found"
	ReasonRootInspectionFailed, ReasonRootSymlink       RootReasonCode = "root_inspection_failed", "root_symlink"
	ReasonRootDeviceBoundary, ReasonRootNotDirectory    RootReasonCode = "root_device_boundary", "root_not_directory"
	ReasonCancelledBeforeStart, ReasonDeviceBoundary    RootReasonCode = "cancelled_before_start", "device_boundary"
	ReasonEntryVisibilityGap, ReasonCancelledDuringScan RootReasonCode = "entry_visibility_gap", "cancelled_during_scan"
)

func ParseRootReasonCode(value string) (RootReasonCode, error) {
	reason := RootReasonCode(value)
	switch reason {
	case ReasonCompleted, ReasonNotFound, ReasonRootInspectionFailed,
		ReasonRootSymlink, ReasonRootDeviceBoundary, ReasonRootNotDirectory,
		ReasonCancelledBeforeStart, ReasonDeviceBoundary,
		ReasonEntryVisibilityGap, ReasonCancelledDuringScan:
		return reason, nil
	default:
		return "", invalidValue("root reason code", value)
	}
}

type Outcome string

const (
	OutcomeComplete, OutcomePartial Outcome = "complete", "partial"
	OutcomeCancelled                Outcome = "cancelled"
)

type Risk string

const RiskManualReview Risk = "manual_review"

type WarningCode string

const (
	WarningRootInaccessible, WarningRootSymlink               WarningCode = "root_inaccessible", "root_symlink"
	WarningRootDeviceBoundary, WarningRootNotDirectory        WarningCode = "root_device_boundary", "root_not_directory"
	WarningEntryInaccessible, WarningEntryChanged             WarningCode = "entry_inaccessible", "entry_changed"
	WarningSymlinkSkipped, WarningDeviceBoundary              WarningCode = "symlink_skipped", "device_boundary"
	WarningUnsupportedEntryType, WarningIdentityUnavailable   WarningCode = "unsupported_entry_type", "identity_unavailable"
	WarningAllocationUnavailable, WarningHardLinkObserved     WarningCode = "allocation_unavailable", "hard_link_observed"
	WarningIdentityAlreadyAccounted, WarningEntryLimitReached WarningCode = "identity_already_accounted", "entry_limit_reached"
	WarningPathBudgetReached, WarningCancelledDuringRoot      WarningCode = "path_budget_reached", "cancelled_during_root"
	WarningCancelledBeforeRoot, WarningAPFSEstimateUncertain  WarningCode = "cancelled_before_root", "apfs_estimate_uncertain"
	WarningCoreSimulatorSemantics                             WarningCode = "core_simulator_semantics"
)

var warningMessages = map[WarningCode]string{
	WarningRootInaccessible:         "Built-in root could not be inspected.",
	WarningRootSymlink:              "Built-in root is a symbolic link and was not followed.",
	WarningRootDeviceBoundary:       "Built-in root is on another device and was not scanned.",
	WarningRootNotDirectory:         "Built-in root is not a directory and was not scanned.",
	WarningEntryInaccessible:        "Entry could not be inspected.",
	WarningEntryChanged:             "Entry changed during inspection and was not measured.",
	WarningSymlinkSkipped:           "Symbolic link was skipped and its target was not measured.",
	WarningDeviceBoundary:           "Entry is on another device and was not scanned.",
	WarningUnsupportedEntryType:     "Unsupported entry type was not measured.",
	WarningIdentityUnavailable:      "Filesystem identity was unavailable; deduplication is incomplete.",
	WarningAllocationUnavailable:    "Allocated size was unavailable and remains unknown.",
	WarningHardLinkObserved:         "Multiple paths share one observed filesystem identity.",
	WarningIdentityAlreadyAccounted: "Filesystem identity was already included in totals.",
	WarningEntryLimitReached:        "Entry limit was reached before the root was fully inspected.",
	WarningPathBudgetReached:        "Path budget was reached before the root was fully inspected.",
	WarningCancelledDuringRoot:      "Scan was cancelled while inspecting this root.",
	WarningCancelledBeforeRoot:      "Scan was cancelled before this root was inspected.",
	WarningAPFSEstimateUncertain:    "APFS allocation is an estimate and not guaranteed reclaimable space.",
	WarningCoreSimulatorSemantics:   "CoreSimulator data requires product-aware manual review.",
}

func ParseWarningCode(value string) (WarningCode, error) {
	code := WarningCode(value)
	if _, known := warningMessages[code]; !known {
		return "", invalidValue("warning code", value)
	}
	return code, nil
}
func (code WarningCode) Message() string    { return warningMessages[code] }
func invalidValue(kind, value string) error { return fmt.Errorf("invalid %s %q", kind, value) }

type FilesystemIdentity struct {
	device uint64
	inode  uint64
}

func NewFilesystemIdentity(device, inode uint64) (FilesystemIdentity, error) {
	if device == 0 || inode == 0 {
		return FilesystemIdentity{}, fmt.Errorf("filesystem identity requires non-zero device and inode")
	}
	return FilesystemIdentity{device: device, inode: inode}, nil
}
func (identity FilesystemIdentity) Device() uint64 { return identity.device }
func (identity FilesystemIdentity) Inode() uint64  { return identity.inode }

type OptionalFilesystemIdentity struct {
	value FilesystemIdentity
	known bool
}

func KnownFilesystemIdentity(identity FilesystemIdentity) OptionalFilesystemIdentity {
	if identity.device == 0 || identity.inode == 0 {
		return UnknownFilesystemIdentity()
	}
	return OptionalFilesystemIdentity{value: identity, known: true}
}
func UnknownFilesystemIdentity() OptionalFilesystemIdentity { return OptionalFilesystemIdentity{} }
func (identity OptionalFilesystemIdentity) Value() (FilesystemIdentity, bool) {
	return identity.value, identity.known
}

type Allocation struct {
	bytes uint64
	known bool
}

func KnownAllocation(bytes uint64) Allocation       { return Allocation{bytes: bytes, known: true} }
func UnknownAllocation() Allocation                 { return Allocation{} }
func (allocation Allocation) Bytes() (uint64, bool) { return allocation.bytes, allocation.known }

type EstimateBasis string

const BasisRegularFileStatSize, BasisStatBlocks512 EstimateBasis = "regular_file_stat_size", "stat_blocks_512"

type EstimateCompleteness string

const (
	CompletenessComplete, CompletenessIncomplete EstimateCompleteness = "complete", "incomplete"
	CompletenessUnknown                          EstimateCompleteness = "unknown"
)

type ValueEstimate struct {
	knownBytes   uint64
	basis        EstimateBasis
	completeness EstimateCompleteness
}

func NewValueEstimate(bytes uint64, basis EstimateBasis, completeness EstimateCompleteness) (ValueEstimate, error) {
	validBasis := basis == BasisRegularFileStatSize || basis == BasisStatBlocks512
	validCompleteness := completeness == CompletenessComplete || completeness == CompletenessIncomplete || completeness == CompletenessUnknown
	if !validBasis || !validCompleteness || (basis == BasisRegularFileStatSize && completeness == CompletenessUnknown) || (completeness == CompletenessUnknown && bytes != 0) {
		return ValueEstimate{}, fmt.Errorf("invalid value estimate")
	}
	return ValueEstimate{bytes, basis, completeness}, nil
}
func (value ValueEstimate) KnownBytes() uint64                 { return value.knownBytes }
func (value ValueEstimate) Basis() EstimateBasis               { return value.basis }
func (value ValueEstimate) Completeness() EstimateCompleteness { return value.completeness }

type Estimate struct {
	logical, allocated    ValueEstimate
	deduplicationComplete bool
}

func NewEstimate(logical, allocated ValueEstimate, deduplicationComplete bool) (Estimate, error) {
	if logical.basis != BasisRegularFileStatSize || allocated.basis != BasisStatBlocks512 ||
		logical.completeness == CompletenessUnknown ||
		(logical.completeness == CompletenessIncomplete && allocated.completeness == CompletenessComplete) {
		return Estimate{}, fmt.Errorf("invalid estimate")
	}
	return Estimate{logical, allocated, deduplicationComplete}, nil
}
func (estimate Estimate) Logical() ValueEstimate      { return estimate.logical }
func (estimate Estimate) Allocated() ValueEstimate    { return estimate.allocated }
func (estimate Estimate) DeduplicationComplete() bool { return estimate.deduplicationComplete }

type Accounting struct {
	measured, attributed, alreadyAccounted, hardLinks, unknownIdentity uint64
}

func NewAccounting(measured, attributed, alreadyAccounted, hardLinks, unknownIdentity uint64) (Accounting, error) {
	if attributed > measured || alreadyAccounted != measured-attributed ||
		hardLinks > measured || unknownIdentity > attributed {
		return Accounting{}, fmt.Errorf("invalid accounting counts")
	}
	return Accounting{measured, attributed, alreadyAccounted, hardLinks, unknownIdentity}, nil
}
func (accounting Accounting) MeasuredObjects() uint64        { return accounting.measured }
func (accounting Accounting) AttributedObjects() uint64      { return accounting.attributed }
func (accounting Accounting) AlreadyAccountedPaths() uint64  { return accounting.alreadyAccounted }
func (accounting Accounting) HardLinkObservations() uint64   { return accounting.hardLinks }
func (accounting Accounting) UnknownIdentityObjects() uint64 { return accounting.unknownIdentity }

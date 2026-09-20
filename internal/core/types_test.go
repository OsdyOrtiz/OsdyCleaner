package core

import "testing"

func TestParseAreaID(t *testing.T) {
	cases := []struct {
		text string
		want AreaID
		rank int
	}{
		{"npm-cache", AreaNPMCache, 0},
		{"homebrew-cache", AreaHomebrewCache, 1},
		{"gradle-caches", AreaGradleCaches, 2},
		{"xcode-derived-data", AreaXcodeDerivedData, 3},
		{"core-simulator", AreaCoreSimulator, 4},
	}
	for _, tt := range cases {
		t.Run(tt.text, func(t *testing.T) {
			got, err := ParseAreaID(tt.text)
			if err != nil || got != tt.want {
				t.Fatalf("ParseAreaID(%q) = %q, %v; want %q, nil", tt.text, got, err, tt.want)
			}
			if got := AreaRank(tt.want); got != tt.rank {
				t.Fatalf("AreaRank(%q) = %d; want %d", tt.want, got, tt.rank)
			}
		})
	}
	for _, value := range []string{"", "NPM-cache", "other"} {
		assertRejected(t, value, ParseAreaID)
		if rank := AreaRank(AreaID(value)); rank != -1 {
			t.Errorf("AreaRank(%q) = %d; want -1", value, rank)
		}
	}
}

func TestParseRootStatus(t *testing.T) {
	assertParsed(t, []RootStatus{
		RootScanned, RootMissing, RootInaccessible, RootSkipped,
		RootBoundaryLimited, RootPartial, RootCancelled,
	}, ParseRootStatus)

	for _, tt := range []struct{ got, want string }{
		{string(OutcomeComplete), "complete"},
		{string(OutcomePartial), "partial"},
		{string(OutcomeCancelled), "cancelled"},
		{string(RiskManualReview), "manual_review"},
	} {
		if tt.got != tt.want {
			t.Errorf("named value = %q; want %q", tt.got, tt.want)
		}
	}
	for _, value := range []string{"", "SCANNED", "complete"} {
		assertRejected(t, value, ParseRootStatus)
	}
}

func TestParseRootReasonCode(t *testing.T) {
	assertParsed(t, []RootReasonCode{
		ReasonCompleted, ReasonNotFound, ReasonRootInspectionFailed,
		ReasonRootSymlink, ReasonRootDeviceBoundary, ReasonRootNotDirectory,
		ReasonCancelledBeforeStart, ReasonDeviceBoundary,
		ReasonEntryVisibilityGap, ReasonCancelledDuringScan,
	}, ParseRootReasonCode)
	for _, value := range []string{"", "cancelled", "not-found"} {
		assertRejected(t, value, ParseRootReasonCode)
	}
}

func TestParseWarningCode(t *testing.T) {
	cases := []struct {
		code    WarningCode
		message string
	}{
		{WarningRootInaccessible, "Built-in root could not be inspected."},
		{WarningRootSymlink, "Built-in root is a symbolic link and was not followed."},
		{WarningRootDeviceBoundary, "Built-in root is on another device and was not scanned."},
		{WarningRootNotDirectory, "Built-in root is not a directory and was not scanned."},
		{WarningEntryInaccessible, "Entry could not be inspected."},
		{WarningEntryChanged, "Entry changed during inspection and was not measured."},
		{WarningSymlinkSkipped, "Symbolic link was skipped and its target was not measured."},
		{WarningDeviceBoundary, "Entry is on another device and was not scanned."},
		{WarningUnsupportedEntryType, "Unsupported entry type was not measured."},
		{WarningIdentityUnavailable, "Filesystem identity was unavailable; deduplication is incomplete."},
		{WarningAllocationUnavailable, "Allocated size was unavailable and remains unknown."},
		{WarningHardLinkObserved, "Multiple paths share one observed filesystem identity."},
		{WarningIdentityAlreadyAccounted, "Filesystem identity was already included in totals."},
		{WarningEntryLimitReached, "Entry limit was reached before the root was fully inspected."},
		{WarningPathBudgetReached, "Path budget was reached before the root was fully inspected."},
		{WarningCancelledDuringRoot, "Scan was cancelled while inspecting this root."},
		{WarningCancelledBeforeRoot, "Scan was cancelled before this root was inspected."},
		{WarningAPFSEstimateUncertain, "APFS allocation is an estimate and not guaranteed reclaimable space."},
		{WarningCoreSimulatorSemantics, "CoreSimulator data requires product-aware manual review."},
	}
	for _, tt := range cases {
		t.Run(string(tt.code), func(t *testing.T) {
			got, err := ParseWarningCode(string(tt.code))
			if err != nil || got != tt.code {
				t.Fatalf("ParseWarningCode(%q) = %q, %v; want %q, nil", tt.code, got, err, tt.code)
			}
			if got := tt.code.Message(); got != tt.message {
				t.Fatalf("%q.Message() = %q; want %q", tt.code, got, tt.message)
			}
		})
	}
	for _, value := range []string{"", "ROOT_INACCESSIBLE", "unknown_warning"} {
		assertRejected(t, value, ParseWarningCode)
	}
}

func TestAllocation(t *testing.T) {
	for _, tt := range []struct {
		name  string
		bytes uint64
	}{{"zero", 0}, {"nonzero", 8192}, {"maximum", ^uint64(0)}} {
		t.Run(tt.name, func(t *testing.T) {
			got, known := KnownAllocation(tt.bytes).Bytes()
			if !known || got != tt.bytes {
				t.Fatalf("KnownAllocation(%d).Bytes() = %d, %t", tt.bytes, got, known)
			}
		})
	}
	const logicalSize = uint64(4096)
	if bytes, known := UnknownAllocation().Bytes(); known || bytes == logicalSize {
		t.Fatalf("unknown allocation = %d, %t; must not fall back to logical size", bytes, known)
	}
	for _, tt := range []struct{ got, want string }{
		{string(BasisRegularFileStatSize), "regular_file_stat_size"},
		{string(BasisStatBlocks512), "stat_blocks_512"},
		{string(CompletenessComplete), "complete"},
		{string(CompletenessIncomplete), "incomplete"},
		{string(CompletenessUnknown), "unknown"},
	} {
		if tt.got != tt.want {
			t.Errorf("estimate named value = %q; want %q", tt.got, tt.want)
		}
	}
}

func TestFilesystemIdentity(t *testing.T) {
	identity, err := NewFilesystemIdentity(7, 11)
	if err != nil {
		t.Fatalf("NewFilesystemIdentity(7, 11): %v", err)
	}
	if identity.Device() != 7 || identity.Inode() != 11 {
		t.Fatalf("identity = (%d, %d); want (7, 11)", identity.Device(), identity.Inode())
	}
	got, known := KnownFilesystemIdentity(identity).Value()
	if !known || got != identity {
		t.Fatalf("known identity = %#v, %t; want %#v, true", got, known, identity)
	}
	if _, known := KnownFilesystemIdentity(FilesystemIdentity{}).Value(); known {
		t.Fatal("KnownFilesystemIdentity(FilesystemIdentity{}).Value() reported a known identity")
	}
	for _, invalid := range []FilesystemIdentity{{device: 7}, {inode: 11}} {
		if _, known := KnownFilesystemIdentity(invalid).Value(); known {
			t.Errorf("KnownFilesystemIdentity(%#v).Value() reported a known identity", invalid)
		}
	}
	if _, known := UnknownFilesystemIdentity().Value(); known {
		t.Fatal("UnknownFilesystemIdentity().Value() reported a known identity")
	}
	for _, tt := range []struct{ device, inode uint64 }{{0, 11}, {7, 0}, {0, 0}} {
		if _, err := NewFilesystemIdentity(tt.device, tt.inode); err == nil {
			t.Errorf("NewFilesystemIdentity(%d, %d) succeeded; want error", tt.device, tt.inode)
		}
	}
	if _, err := NewFilesystemIdentity(^uint64(0), ^uint64(0)); err != nil {
		t.Errorf("maximum identity values were rejected: %v", err)
	}
}

func assertParsed[T ~string](t *testing.T, values []T, parse func(string) (T, error)) {
	t.Helper()
	for _, want := range values {
		t.Run(string(want), func(t *testing.T) {
			if got, err := parse(string(want)); err != nil || got != want {
				t.Errorf("parse(%q) = %q, %v; want %q, nil", want, got, err, want)
			}
		})
	}
}

func assertRejected[T ~string](t *testing.T, value string, parse func(string) (T, error)) {
	t.Helper()
	if _, err := parse(value); err == nil {
		t.Errorf("parser accepted invalid value %q", value)
	}
}

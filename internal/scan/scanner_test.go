package scan

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/osdy/OsdyCleaner/internal/core"
)

type scannerHome struct {
	home  string
	calls int
}

func (h *scannerHome) Home() (string, error) {
	h.calls++
	return h.home, nil
}

type scannerWalker struct {
	results     map[core.AreaID]error
	closeErrors map[core.AreaID]error
	calls       []string
	events      []string
	closes      int
	onCall      func(core.AreaID)
}

func (w *scannerWalker) WalkFacts(context.Context, TrustedRoot, func(DirectoryFact) error) error {
	return nil
}

func (w *scannerWalker) AcquireRoot(_ context.Context, home, root AbsoluteComponents, _ WalkLimits) (TrustedRoot, error) {
	area, err := areaForRoot(root.String())
	if err != nil {
		return nil, err
	}
	w.calls = append(w.calls, home.String()+"|"+root.String())
	if w.onCall != nil {
		w.onCall(area)
	}
	if err := w.results[area]; err != nil {
		return nil, err
	}
	facts, _ := NewRootFacts(Identity{Device: 7, Inode: uint64(core.AreaRank(area) + 1)}, true)
	return NewTrustedRoot(facts, func() error {
		w.events = append(w.events, "close:"+string(area))
		w.closes++
		return w.closeErrors[area]
	})
}

func areaForRoot(root string) (core.AreaID, error) {
	for _, definition := range builtinDefinitions {
		if root == "/fixture/home/"+definition.RelativePath {
			return definition.AreaID, nil
		}
	}
	return "", ErrInvalidMetadata
}

func scannerDefinitions(t *testing.T, home string) []BuiltinDefinition {
	t.Helper()
	definitions, err := ResolveBuiltins(&scannerHome{home: home}, &recordingHomeInspector{})
	if err != nil {
		t.Fatalf("ResolveBuiltins(%q): %v", home, err)
	}
	return definitions
}

func TestScannerCanonicalPreflight(t *testing.T) {
	for _, tc := range []struct {
		name        string
		definitions []BuiltinDefinition
	}{
		{"reordered", []BuiltinDefinition{scannerDefinitions(t, "/fixture/home")[1], scannerDefinitions(t, "/fixture/home")[0], scannerDefinitions(t, "/fixture/home")[2], scannerDefinitions(t, "/fixture/home")[3], scannerDefinitions(t, "/fixture/home")[4]}},
		{"duplicate", append(scannerDefinitions(t, "/fixture/home"), scannerDefinitions(t, "/fixture/home")[0])},
		{"unclean relative", func() []BuiltinDefinition {
			d := scannerDefinitions(t, "/fixture/home")
			d[0].RelativePath = "bad/../path"
			return d
		}()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			walker := &scannerWalker{}
			_, err := NewScanner(&scannerHome{home: "/fixture/home"}, tc.definitions, walker, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}).Scan(context.Background())
			if !errors.Is(err, ErrInvalidComponents) || len(walker.calls) != 0 {
				t.Fatalf("Scan() error=%v calls=%v", err, walker.calls)
			}
		})
	}
}

func TestScannerSerialRoots(t *testing.T) {
	walker := &scannerWalker{}
	walker.onCall = func(area core.AreaID) {
		if got, want := walker.closes, core.AreaRank(area); got != want {
			t.Errorf("root %s began with %d prior closes, want %d", area, got, want)
		}
	}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}).Scan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"/fixture/home|/fixture/home/.npm",
		"/fixture/home|/fixture/home/Library/Caches/Homebrew",
		"/fixture/home|/fixture/home/.gradle/caches",
		"/fixture/home|/fixture/home/Library/Developer/Xcode/DerivedData",
		"/fixture/home|/fixture/home/Library/Developer/CoreSimulator",
	}
	if !reflect.DeepEqual(walker.calls, want) || walker.closes != 5 || len(observations) != 5 {
		t.Fatalf("calls=%v closes=%d observations=%d", walker.calls, walker.closes, len(observations))
	}
}

func TestScannerAcquisitionStates(t *testing.T) {
	for _, tc := range []struct {
		name     string
		err      error
		closeErr error
		status   core.RootStatus
		reason   core.RootReasonCode
		warnings []core.WarningCode
	}{
		{"missing", ErrMissing, nil, core.RootMissing, core.ReasonNotFound, nil},
		{"inaccessible", ErrInaccessible, nil, core.RootInaccessible, core.ReasonRootInspectionFailed, []core.WarningCode{core.WarningRootInaccessible}},
		{"symlink", ErrSymlink, nil, core.RootSkipped, core.ReasonRootSymlink, []core.WarningCode{core.WarningRootSymlink}},
		{"not directory", ErrNotDirectory, nil, core.RootSkipped, core.ReasonRootNotDirectory, []core.WarningCode{core.WarningRootNotDirectory}},
		{"device", ErrDeviceBoundary, nil, core.RootSkipped, core.ReasonRootDeviceBoundary, []core.WarningCode{core.WarningRootDeviceBoundary}},
		{"non local", ErrNonLocal, nil, core.RootSkipped, core.ReasonRootDeviceBoundary, []core.WarningCode{core.WarningRootDeviceBoundary}},
		{"combined device and inaccessible", errors.Join(ErrDeviceBoundary, ErrInaccessible), nil, core.RootSkipped, core.ReasonRootDeviceBoundary, []core.WarningCode{core.WarningRootDeviceBoundary}},
		{"close error", nil, ErrInaccessible, core.RootInaccessible, core.ReasonRootInspectionFailed, []core.WarningCode{core.WarningRootInaccessible}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			walker := &scannerWalker{results: map[core.AreaID]error{core.AreaNPMCache: tc.err}, closeErrors: map[core.AreaID]error{core.AreaNPMCache: tc.closeErr}}
			observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}).Scan(context.Background())
			if err != nil || observations[0].Status() != tc.status || observations[0].Reason() != tc.reason || !reflect.DeepEqual(observations[0].WarningCodes(), tc.warnings) || len(walker.calls) != 5 || (tc.closeErr != nil && walker.closes != 5) {
				t.Fatalf("observations=%v err=%v calls=%v closes=%d", observations, err, walker.calls, walker.closes)
			}
		})
	}
}

func TestScannerFiveTransitionalObservations(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	walker := &scannerWalker{onCall: func(area core.AreaID) {
		if area == core.AreaHomebrewCache {
			cancel()
		}
	}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}).Scan(ctx)
	if err != nil || len(observations) != 5 || walker.closes != 2 {
		t.Fatalf("observations=%v err=%v closes=%d", observations, err, walker.closes)
	}
	want := []core.RootStatus{core.RootScanned, core.RootCancelled, core.RootSkipped, core.RootSkipped, core.RootSkipped}
	for i, observation := range observations {
		if observation.AreaID() != core.AreaID([]string{"npm-cache", "homebrew-cache", "gradle-caches", "xcode-derived-data", "core-simulator"}[i]) || observation.Status() != want[i] || observation.Estimate().Logical().Completeness() != core.CompletenessIncomplete || observation.Estimate().Logical().KnownBytes() != 0 {
			t.Fatalf("observation %d = %#v", i, observation)
		}
	}
}

func TestScannerNoRealHome(t *testing.T) {
	home := &scannerHome{home: "/fixture/home"}
	walker := &scannerWalker{}
	_, err := NewScanner(home, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}).Scan(context.Background())
	if err != nil || home.calls != 1 || len(walker.calls) != 5 {
		t.Fatalf("err=%v home-calls=%d walker-calls=%d", err, home.calls, len(walker.calls))
	}
}

type directoryScriptWalker struct {
	*scannerWalker
	scripts map[core.AreaID][]DirectoryFact
	errs    map[core.AreaID]error
	visits  int
}

func (w *directoryScriptWalker) WalkFacts(_ context.Context, root TrustedRoot, visit func(DirectoryFact) error) error {
	area := core.AreaID("")
	for _, definition := range builtinDefinitions {
		if root.Facts().Identity.Inode == uint64(core.AreaRank(definition.AreaID)+1) {
			area = definition.AreaID
		}
	}
	for _, fact := range w.scripts[area] {
		w.visits++
		if err := visit(fact); err != nil {
			return err
		}
	}
	return w.errs[area]
}

func directoryFact(t *testing.T, components []string, kind EntryKind, identity, enumerated Identity, local bool) DirectoryFact {
	t.Helper()
	fact, err := NewDirectoryFact(components, kind, &enumerated, kind, &identity, local, nil)
	if err != nil {
		t.Fatal(err)
	}
	return fact
}

func TestScannerDirectoryFacts(t *testing.T) {
	walker := &directoryScriptWalker{scannerWalker: &scannerWalker{}, scripts: map[core.AreaID][]DirectoryFact{
		core.AreaNPMCache: {directoryFact(t, []string{"cache"}, EntryDirectory, Identity{Device: 7, Inode: 10}, Identity{Device: 7, Inode: 10}, true)},
	}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	// Active design § Exact status: a fully traversed directory-only root is complete.
	if got := observations[0]; got.Status() != core.RootScanned || got.Reason() != core.ReasonCompleted || len(got.WarningCodes()) != 0 {
		t.Fatalf("directory-only root = %#v", got)
	}
	if len(walker.calls) != 5 || walker.closes != 5 {
		t.Fatalf("calls=%v closes=%d", walker.calls, walker.closes)
	}
	testScannerDirectoryPermutationsAndContainment(t)
}

func TestScannerDirectoryChanges(t *testing.T) {
	base := directoryFact(t, []string{"safe"}, EntryDirectory, Identity{Device: 7, Inode: 11}, Identity{Device: 7, Inode: 11}, true)
	changed := directoryFact(t, []string{"changed"}, EntryDirectory, Identity{Device: 7, Inode: 12}, Identity{Device: 7, Inode: 13}, true)
	walker := &directoryScriptWalker{scannerWalker: &scannerWalker{}, scripts: map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {changed, base}}, errs: map[core.AreaID]error{core.AreaNPMCache: ErrMissing}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	got := observations[0]
	if got.Status() != core.RootPartial || !reflect.DeepEqual(got.WarningCodes(), []core.WarningCode{core.WarningEntryChanged}) || len(walker.calls) != 5 {
		t.Fatalf("changed/disappeared root = %#v calls=%v", got, walker.calls)
	}
}

func TestScannerDirectoryBoundaries(t *testing.T) {
	boundary := directoryFact(t, []string{"outside"}, EntryDirectory, Identity{Device: 8, Inode: 20}, Identity{Device: 8, Inode: 20}, false)
	symlink := directoryFact(t, []string{"link"}, EntrySymlink, Identity{}, Identity{}, true)
	special := directoryFact(t, []string{"fifo"}, EntrySpecial, Identity{}, Identity{}, true)
	walker := &directoryScriptWalker{scannerWalker: &scannerWalker{}, scripts: map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {boundary, symlink, special}}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	got := observations[0]
	want := []core.WarningCode{core.WarningDeviceBoundary, core.WarningSymlinkSkipped, core.WarningUnsupportedEntryType}
	if got.Status() != core.RootPartial || !reflect.DeepEqual(got.WarningCodes(), want) {
		t.Fatalf("boundary/symlink/special = %#v", got)
	}
}

func TestScannerRootStatusMapping(t *testing.T) {
	for _, tc := range []struct {
		name string
		err  error
		want core.RootStatus
	}{
		{"boundary", ErrDeviceBoundary, core.RootBoundaryLimited},
		{"inaccessible wins", errors.Join(ErrDeviceBoundary, ErrInaccessible), core.RootPartial},
		{"changed wins", ErrInvalidMetadata, core.RootPartial},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := directoryStatus(tc.err, true); got != tc.want {
				t.Fatalf("directoryStatus(%v)=%s; want %s", tc.err, got, tc.want)
			}
		})
	}
}

func testScannerDirectoryPermutationsAndContainment(t *testing.T) {
	valid := directoryFact(t, []string{"safe"}, EntryDirectory, Identity{Device: 7, Inode: 40}, Identity{Device: 7, Inode: 40}, true)
	changed := directoryFact(t, []string{"changed"}, EntryDirectory, Identity{Device: 7, Inode: 41}, Identity{Device: 7, Inode: 42}, true)
	boundary := directoryFact(t, []string{"outside"}, EntryDirectory, Identity{Device: 8, Inode: 43}, Identity{Device: 8, Inode: 43}, false)
	unsafe := DirectoryFact{components: []string{"..", "target"}, enumeratedKind: EntryDirectory, openedKind: EntryDirectory, enumeratedIdentity: Identity{Device: 7, Inode: 44}, openedIdentity: Identity{Device: 7, Inode: 44}, enumeratedHasIdentity: true, openedHasIdentity: true, local: true}
	var signatures [][]core.WarningCode
	for _, script := range [][]DirectoryFact{{valid, changed, boundary, unsafe}, {unsafe, boundary, changed, valid}} {
		walker := &directoryScriptWalker{scannerWalker: &scannerWalker{}, scripts: map[core.AreaID][]DirectoryFact{core.AreaNPMCache: script}}
		observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
		if err != nil || walker.visits != 4 || len(walker.calls) != 5 {
			t.Fatalf("err=%v visits=%d calls=%v", err, walker.visits, walker.calls)
		}
		got := observations[0]
		if got.Status() != core.RootPartial || got.Reason() != core.ReasonEntryVisibilityGap {
			t.Fatalf("unsafe callback was accepted: %#v", got)
		}
		signatures = append(signatures, got.WarningCodes())
	}
	want := []core.WarningCode{core.WarningDeviceBoundary, core.WarningEntryChanged}
	if !reflect.DeepEqual(signatures[0], want) || !reflect.DeepEqual(signatures[1], want) {
		t.Fatalf("warning order changed under callback permutation: %v", signatures)
	}
}

func TestScannerFiveFinalObservations(t *testing.T) {
	walker := &directoryScriptWalker{scannerWalker: &scannerWalker{}, scripts: map[core.AreaID][]DirectoryFact{
		core.AreaNPMCache: {directoryFact(t, []string{"safe"}, EntryDirectory, Identity{Device: 7, Inode: 30}, Identity{Device: 7, Inode: 30}, true)},
	}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil || len(observations) != 5 {
		t.Fatalf("observations=%v err=%v", observations, err)
	}
	for _, observation := range observations {
		if observation.Status() != core.RootScanned || observation.Reason() != core.ReasonCompleted {
			t.Fatalf("fully traversed directory-only observation was not complete: %#v", observation)
		}
	}
}

type factPortWalker struct {
	*scannerWalker
	facts map[core.AreaID][]DirectoryFact
	calls map[core.AreaID]int
}

func (w *factPortWalker) WalkFacts(_ context.Context, root TrustedRoot, visit func(DirectoryFact) error) error {
	area := core.AreaID("")
	for _, definition := range builtinDefinitions {
		if root.Facts().Identity.Inode == uint64(core.AreaRank(definition.AreaID)+1) {
			area = definition.AreaID
		}
	}
	w.calls[area]++
	w.events = append(w.events, "walk-begin:"+string(area))
	defer func() { w.events = append(w.events, "walk-end:"+string(area)) }()
	for _, fact := range w.facts[area] {
		if err := visit(fact); err != nil {
			return err
		}
	}
	return nil
}

func TestScannerProductionFactPort(t *testing.T) {
	stable, err := NewDirectoryFact([]string{"safe"}, EntryDirectory, &Identity{Device: 7, Inode: 9}, EntryDirectory, &Identity{Device: 7, Inode: 9}, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	walker := &factPortWalker{scannerWalker: &scannerWalker{}, facts: map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {stable}}, calls: map[core.AreaID]int{}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil || len(observations) != 5 || walker.closes != 5 {
		t.Fatalf("observations=%d err=%v closes=%d", len(observations), err, walker.closes)
	}
	for _, definition := range scannerDefinitions(t, "/fixture/home") {
		if walker.calls[definition.AreaID] != 1 {
			t.Fatalf("%s fact calls=%d", definition.AreaID, walker.calls[definition.AreaID])
		}
	}
}

func TestScannerFactKindIdentityPairs(t *testing.T) {
	changed, err := NewDirectoryFact([]string{"changed"}, EntryDirectory, &Identity{Device: 7, Inode: 10}, EntrySymlink, nil, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	walker := &factPortWalker{scannerWalker: &scannerWalker{}, facts: map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {changed}}, calls: map[core.AreaID]int{}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil || observations[0].WarningCodes()[0] != core.WarningEntryChanged {
		t.Fatalf("observations=%v err=%v", observations, err)
	}
}

func TestScannerStableSkippedKinds(t *testing.T) {
	walker := &factPortWalker{scannerWalker: &scannerWalker{}, facts: map[core.AreaID][]DirectoryFact{
		core.AreaNPMCache: {
			directoryFact(t, []string{"link"}, EntrySymlink, Identity{}, Identity{}, true),
			directoryFact(t, []string{"fifo"}, EntrySpecial, Identity{}, Identity{}, true),
		},
	}, calls: map[core.AreaID]int{}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	codes := observations[0].WarningCodes()
	want := []core.WarningCode{core.WarningSymlinkSkipped, core.WarningUnsupportedEntryType}
	if !reflect.DeepEqual(codes, want) {
		t.Fatalf("stable skipped kinds = %v, want %v", codes, want)
	}
}

func TestScannerFiveProductionShapedTraversals(t *testing.T) {
	walker := &factPortWalker{scannerWalker: &scannerWalker{}, facts: map[core.AreaID][]DirectoryFact{}, calls: map[core.AreaID]int{}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil || len(observations) != 5 || walker.closes != 5 {
		t.Fatalf("observations=%d err=%v closes=%d", len(observations), err, walker.closes)
	}
	for _, definition := range scannerDefinitions(t, "/fixture/home") {
		if walker.calls[definition.AreaID] != 1 {
			t.Fatalf("%s traversals=%d, want one", definition.AreaID, walker.calls[definition.AreaID])
		}
	}
}

type lifecycleFilePort struct {
	started chan FileJob
	seen    chan error
	release <-chan struct{}
}

func (p lifecycleFilePort) Inspect(ctx context.Context, job FileJob) FileResult {
	p.started <- job
	var cause error
	select {
	case <-ctx.Done():
		cause = ctx.Err()
		p.seen <- cause
	case <-p.release:
	}
	result, _ := NewFileResult(job, EntryRegular, job.FinalIdentity(), uint64(job.FinalIdentity().Inode), 1, cause)
	return result
}

func receiveFileJob(t *testing.T, ch <-chan FileJob, what string) FileJob {
	t.Helper()
	select {
	case job := <-ch:
		return job
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for %s", what)
		return FileJob{}
	}
}

func receiveError(t *testing.T, ch <-chan error, what string) error {
	t.Helper()
	select {
	case err := <-ch:
		return err
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for %s", what)
		return nil
	}
}

func TestScannerBoundedFileJobs(t *testing.T) {
	release := make(chan struct{})
	port := lifecycleFilePort{started: make(chan FileJob, 2), seen: make(chan error, 2), release: release}
	pipeline, err := NewFilePipeline(port, 1, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if err := pipeline.Submit(context.Background(), mustFileJob(t, "one", 1)); err != nil {
		t.Fatal(err)
	}
	receiveFileJob(t, port.started, "active inspect")
	if err := pipeline.Submit(context.Background(), mustFileJob(t, "two", 2)); err != nil {
		t.Fatal(err)
	}
	submitEntered := make(chan struct{})
	submitDone := make(chan error, 1)
	go func() {
		close(submitEntered)
		submitDone <- pipeline.Submit(context.Background(), mustFileJob(t, "three", 3))
	}()
	<-submitEntered // worker is blocked in Inspect and the one-slot job queue is full.
	closeReturned := make(chan struct{})
	go func() { pipeline.CloseJobs(); close(closeReturned) }()
	select {
	case <-closeReturned:
	case <-time.After(time.Second):
		t.Fatal("CloseJobs blocked behind Submit")
	}
	if err := receiveError(t, submitDone, "blocked submit release"); !errors.Is(err, ErrPipelineClosed) {
		t.Fatalf("blocked submit=%v, want closed", err)
	}
	close(release) // Normal close releases submitters but preserves accepted caller contexts.
	results := pipeline.Drain()
	if len(results) != 2 {
		t.Fatalf("accepted facts=%v", results)
	}
	if receiveFileJob(t, port.started, "queued accepted inspect").Components()[0] != "two" {
		t.Fatal("accepted queued job did not drain")
	}
}

func TestFilePipelineContinuesAfterCancelledJob(t *testing.T) {
	release := make(chan struct{})
	port := lifecycleFilePort{started: make(chan FileJob, 2), seen: make(chan error, 1), release: release}
	pipeline, err := NewFilePipeline(port, 1, 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := pipeline.Submit(ctx, mustFileJob(t, "cancelled", 1)); err != nil {
		t.Fatal(err)
	}
	receiveFileJob(t, port.started, "cancelled inspect")
	cancel()
	if err := receiveError(t, port.seen, "cancelled inspect result"); !errors.Is(err, context.Canceled) {
		t.Fatalf("Inspect context=%v", err)
	}
	close(release)
	if err := pipeline.Submit(context.Background(), mustFileJob(t, "later", 2)); err != nil {
		t.Fatal(err)
	}
	pipeline.CloseJobs()
	results := pipeline.Drain()
	if len(results) != 1 || !reflect.DeepEqual(results[0].Components(), []string{"later"}) {
		t.Fatalf("results=%#v", results)
	}
}

func TestScannerFileBackpressure(t *testing.T) {
	testFilePipelineNilContext(t)
	port := lifecycleFilePort{started: make(chan FileJob, 1), seen: make(chan error, 1)}
	pipeline, _ := NewFilePipeline(port, 1, 1, 1)
	caller, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := pipeline.Submit(caller, mustFileJob(t, "caller", 1)); err != nil {
		t.Fatal(err)
	}
	receiveFileJob(t, port.started, "caller inspect")
	cancel()
	if err := receiveError(t, port.seen, "caller cancellation at Inspect"); !errors.Is(err, context.Canceled) {
		t.Fatalf("Inspect context=%v", err)
	}
	pipeline.CloseJobs()
	if got := pipeline.Drain(); len(got) != 0 {
		t.Fatalf("cancelled results=%v", got)
	}
	if _, err := NewFilePipeline(port, 1, 0, 1); !errors.Is(err, ErrInvalidPipeline) {
		t.Fatal(err)
	}
}

func testFilePipelineNilContext(t *testing.T) {
	job := mustFileJob(t, "nil-context", 12)
	port := immediateFilePort{inspected: make(chan FileJob, 1)}
	pipeline, err := NewFilePipeline(port, 1, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name string
		job  FileJob
	}{
		{"valid job", job},
		{"invalid job", FileJob{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := pipeline.Submit(nil, tc.job); !errors.Is(err, ErrInvalidPipeline) {
				t.Fatalf("Submit(nil, %#v) = %v, want ErrInvalidPipeline", tc.job, err)
			}
		})
	}
	if len(port.inspected) != 0 {
		t.Fatal("nil context accepted work")
	}
	if err := pipeline.Submit(context.Background(), job); err != nil {
		t.Fatalf("pipeline unusable after nil context: %v", err)
	}
	if got := receiveFileJob(t, port.inspected, "valid inspect after nil context"); got.FinalIdentity() != job.FinalIdentity() {
		t.Fatalf("inspected %#v, want %#v", got, job)
	}
	pipeline.CloseJobs()
	if got := pipeline.Drain(); len(got) != 1 || !reflect.DeepEqual(got[0].Components(), job.Components()) {
		t.Fatalf("drained results = %#v", got)
	}
	if err := pipeline.Submit(nil, job); !errors.Is(err, ErrInvalidPipeline) {
		t.Fatalf("Submit(nil, validJob) after Close = %v, want ErrInvalidPipeline", err)
	}
}

type immediateFilePort struct{ inspected chan FileJob }

func (p immediateFilePort) Inspect(_ context.Context, job FileJob) FileResult {
	p.inspected <- job
	result, _ := NewFileResult(job, EntryRegular, job.FinalIdentity(), 1, 1, nil)
	return result
}

func TestScannerRegularOnlyFacts(t *testing.T) {
	job := mustFileJob(t, "regular", 3)
	regular, _ := NewFileResult(job, EntryRegular, job.FinalIdentity(), 3, 2, nil)
	directory, _ := NewFileResult(job, EntryDirectory, job.FinalIdentity(), 3, 2, nil)
	changed, _ := NewFileResult(job, EntryRegular, Identity{Device: 7, Inode: 4}, 3, 2, nil)
	facts := normalizeFileResults([]FileResult{directory, changed, regular})
	if len(facts) != 1 || facts[0].LogicalBytes() != 3 || facts[0].AllocationBytes() != 2 {
		t.Fatalf("facts=%v", facts)
	}
}

func TestScannerFileEntryChanges(t *testing.T) {
	for i := 0; i < 25; i++ {
		release := make(chan struct{})
		port := lifecycleFilePort{started: make(chan FileJob, 1), seen: make(chan error, 1), release: release}
		pipeline, _ := NewFilePipeline(port, 1, 1, 1)
		done := make(chan error, 1)
		go func() { done <- pipeline.Submit(context.Background(), mustFileJob(t, "race", uint64(i+1))) }()
		receiveFileJob(t, port.started, "race inspect")
		closeDone := make(chan struct{})
		go func() { pipeline.CloseJobs(); close(closeDone) }()
		pipeline.CloseJobs()
		<-closeDone
		if err := receiveError(t, done, "race submit"); err != nil && !errors.Is(err, ErrPipelineClosed) {
			t.Fatalf("submit=%v", err)
		}
		close(release)
		pipeline.Drain()
	}
}

func mustFileJob(t *testing.T, component string, inode uint64) FileJob {
	t.Helper()
	job, err := NewFileJob([]string{component}, nil, Identity{Device: 7, Inode: inode})
	if err != nil {
		t.Fatal(err)
	}
	return job
}

type filePortRoot struct {
	facts     RootFacts
	area      core.AreaID
	events    *[]string
	jobs      []FileJob
	outcomes  map[string]fileOutcome
	contexts  []string
	onClose   func()
	started   chan struct{}
	cancelled chan struct{}
	block     bool
	releases  map[string]<-chan struct{}
	mu        sync.Mutex
}

type fileOutcome struct {
	kind  EntryKind
	cause error
}

func (r *filePortRoot) Facts() RootFacts { return r.facts }
func (r *filePortRoot) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	*r.events = append(*r.events, "close:"+string(r.area))
	if r.onClose != nil {
		r.onClose()
	}
	return nil
}
func (r *filePortRoot) Inspect(ctx context.Context, job FileJob) FileResult {
	r.mu.Lock()
	r.jobs = append(r.jobs, job)
	if value, ok := ctx.Value(scannerContextKey{}).(string); ok {
		r.contexts = append(r.contexts, value)
	}
	*r.events = append(*r.events, "inspect:"+strings.Join(job.Components(), "/"))
	if r.started != nil {
		select {
		case r.started <- struct{}{}:
		default:
		}
	}
	name := strings.Join(job.Components(), "/")
	release := r.releases[name]
	cancelled := r.cancelled
	block := r.block
	outcome := r.outcomes[name]
	r.mu.Unlock()
	if cancelled != nil {
		<-ctx.Done()
		cancelled <- struct{}{}
	}
	if release != nil {
		<-release
	} else if block {
		<-ctx.Done()
	}
	kind := outcome.kind
	if kind == "" {
		kind = EntryRegular
	}
	result, err := NewFileResult(job, kind, job.FinalIdentity(), 10, 5, outcome.cause)
	if err != nil {
		panic(err)
	}
	return result
}

type filePortWiringWalker struct {
	roots  map[core.AreaID]*filePortRoot
	facts  map[core.AreaID][]DirectoryFact
	errs   map[core.AreaID]error
	events *[]string
}

func (w *filePortWiringWalker) AcquireRoot(_ context.Context, _ AbsoluteComponents, root AbsoluteComponents, _ WalkLimits) (TrustedRoot, error) {
	area, err := areaForRoot(root.String())
	if err != nil {
		return nil, err
	}
	return w.roots[area], nil
}
func (w *filePortWiringWalker) WalkFacts(_ context.Context, root TrustedRoot, visit func(DirectoryFact) error) error {
	fileRoot := root.(*filePortRoot)
	for _, fact := range w.facts[fileRoot.area] {
		if err := visit(fact); err != nil {
			return err
		}
	}
	return w.errs[fileRoot.area]
}

func TestScannerFilePortWiring(t *testing.T) {
	var events []string
	roots := make(map[core.AreaID]*filePortRoot)
	for _, definition := range builtinDefinitions {
		facts, err := NewRootFacts(Identity{Device: 7, Inode: uint64(core.AreaRank(definition.AreaID) + 1)}, true)
		if err != nil {
			t.Fatal(err)
		}
		roots[definition.AreaID] = &filePortRoot{facts: facts, area: definition.AreaID, events: &events}
	}
	regular := func(components []string, identity Identity) DirectoryFact {
		t.Helper()
		fact, err := NewDirectoryFact(components, EntryRegular, &identity, EntryRegular, &identity, true, nil)
		if err != nil {
			t.Fatal(err)
		}
		return fact
	}
	directory := func(components []string, identity Identity) DirectoryFact {
		t.Helper()
		fact, err := NewDirectoryFact(components, EntryDirectory, &identity, EntryDirectory, &identity, true, nil)
		if err != nil {
			t.Fatal(err)
		}
		return fact
	}
	missing, err := NewDirectoryFact([]string{"gone"}, EntryUnknown, nil, EntryUnknown, nil, true, ErrMissing)
	if err != nil {
		t.Fatal(err)
	}
	walker := &filePortWiringWalker{roots: roots, events: &events, facts: map[core.AreaID][]DirectoryFact{
		core.AreaNPMCache: {
			directory([]string{"dir"}, Identity{Device: 7, Inode: 10}),
			regular([]string{"dir", "z"}, Identity{Device: 7, Inode: 12}),
			regular([]string{"a"}, Identity{Device: 7, Inode: 11}),
			// The third job exceeds either one-slot pipeline queue while a result is
			// pending, proving the scanner drains under bounded submit pressure.
			regular([]string{"b"}, Identity{Device: 7, Inode: 15}),
			directoryFact(t, []string{"changed"}, EntryDirectory, Identity{Device: 7, Inode: 13}, Identity{Device: 7, Inode: 14}, true),
			missing,
			directoryFact(t, []string{"special"}, EntrySpecial, Identity{}, Identity{}, true),
		},
	}}
	scanner := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2})
	if _, err := scanner.Scan(context.Background()); err != nil {
		t.Fatal(err)
	}
	root := roots[core.AreaNPMCache]
	if len(root.jobs) != 3 {
		t.Fatalf("jobs=%v, want only accepted regular facts", root.jobs)
	}
	if got, want := root.jobs[0].Components(), []string{"dir", "z"}; !reflect.DeepEqual(got, want) || !reflect.DeepEqual(root.jobs[0].Ancestors(), []Identity{{Device: 7, Inode: 10}}) || root.jobs[0].FinalIdentity() != (Identity{Device: 7, Inode: 12}) {
		t.Fatalf("nested job=%#v", root.jobs[0])
	}
	for index, want := range []struct {
		components []string
		identity   Identity
	}{{[]string{"a"}, Identity{Device: 7, Inode: 11}}, {[]string{"b"}, Identity{Device: 7, Inode: 15}}} {
		job := root.jobs[index+1]
		if got := job.Components(); !reflect.DeepEqual(got, want.components) || len(job.Ancestors()) != 0 || job.FinalIdentity() != want.identity {
			t.Fatalf("root job=%#v", job)
		}
	}
	accepted := scanner.AcceptedFileResults()
	if got := []string{strings.Join(accepted[0].Components(), "/"), strings.Join(accepted[1].Components(), "/"), strings.Join(accepted[2].Components(), "/")}; !reflect.DeepEqual(got, []string{"a", "b", "dir/z"}) {
		t.Fatalf("accepted results=%v", got)
	}
}

type scannerContextKey struct{}

func TestScannerAcceptedFileResults(t *testing.T) {
	var events []string
	facts, err := NewRootFacts(Identity{Device: 7, Inode: 1}, true)
	if err != nil {
		t.Fatal(err)
	}
	root := &filePortRoot{facts: facts, area: core.AreaNPMCache, events: &events, outcomes: map[string]fileOutcome{
		"changed": {cause: ErrInvalidMetadata},
		"missing": {cause: ErrMissing},
		"special": {kind: EntrySpecial},
	}}
	regular := func(component string, inode uint64) DirectoryFact {
		fact, err := NewDirectoryFact([]string{component}, EntryRegular, &Identity{Device: 7, Inode: inode}, EntryRegular, &Identity{Device: 7, Inode: inode}, true, nil)
		if err != nil {
			t.Fatal(err)
		}
		return fact
	}
	roots := map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}
	for _, definition := range builtinDefinitions[1:] {
		facts, err := NewRootFacts(Identity{Device: 7, Inode: uint64(core.AreaRank(definition.AreaID) + 1)}, true)
		if err != nil {
			t.Fatal(err)
		}
		roots[definition.AreaID] = &filePortRoot{facts: facts, area: definition.AreaID, events: &events}
	}
	walker := &filePortWiringWalker{roots: roots, events: &events, facts: map[core.AreaID][]DirectoryFact{
		core.AreaNPMCache: {regular("z", 4), regular("changed", 2), regular("a", 3), regular("missing", 5), regular("special", 6)},
	}}
	scanner := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2})
	var resultsAtClose []FileResult
	root.onClose = func() { resultsAtClose = scanner.AcceptedFileResults() }
	ctx := context.WithValue(context.Background(), scannerContextKey{}, "caller-owned")
	observations, err := scanner.Scan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := observations[0].WarningCodes(), []core.WarningCode{core.WarningEntryChanged}; !reflect.DeepEqual(got, want) {
		t.Fatalf("rejected-result warnings=%v, want %v", got, want)
	}
	got := scanner.AcceptedFileResults()
	if len(got) != 2 || !reflect.DeepEqual(got[0].Components(), []string{"a"}) || !reflect.DeepEqual(got[1].Components(), []string{"z"}) {
		t.Fatalf("accepted results=%#v", got)
	}
	if !reflect.DeepEqual(root.contexts, []string{"caller-owned", "caller-owned", "caller-owned", "caller-owned", "caller-owned"}) {
		t.Fatalf("inspection contexts=%v", root.contexts)
	}
	got[0].job.components[0] = "mutated"
	if again := scanner.AcceptedFileResults(); !reflect.DeepEqual(again[0].Components(), []string{"a"}) {
		t.Fatalf("caller mutation changed scanner result: %#v", again)
	}
	if got := []string{strings.Join(resultsAtClose[0].Components(), "/"), strings.Join(resultsAtClose[1].Components(), "/")}; !reflect.DeepEqual(got, []string{"a", "z"}) {
		t.Fatalf("results at root close=%v events=%v", got, events)
	}
}

func TestScannerProductionTraversalCompletesBeforeRootClose(t *testing.T) {
	walker := &factPortWalker{scannerWalker: &scannerWalker{}, facts: map[core.AreaID][]DirectoryFact{}, calls: map[core.AreaID]int{}}
	_, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).Scan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, definition := range scannerDefinitions(t, "/fixture/home") {
		area := string(definition.AreaID)
		want := []string{"walk-begin:" + area, "walk-end:" + area, "close:" + area}
		start := core.AreaRank(definition.AreaID) * 3
		if got := walker.events[start : start+3]; !reflect.DeepEqual(got, want) {
			t.Fatalf("%s event order=%v, want %v", area, got, want)
		}
	}
}

func TestScannerFileCompletionPermutations(t *testing.T) {
	first := runFinalizedFiles(t, []string{"z", "a"})
	second := runFinalizedFiles(t, []string{"a", "z"})
	if !reflect.DeepEqual(factsSignature(first), factsSignature(second)) {
		t.Fatalf("completion permutations changed finalized facts: %v != %v", factsSignature(first), factsSignature(second))
	}
}

func TestScannerWalkerLimitAfterAcceptedEvidenceIsPartial(t *testing.T) {
	scanner := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "accepted", 10)}})
	walker := scanner.walker.(*filePortWiringWalker)
	walker.errs = map[core.AreaID]error{core.AreaNPMCache: ErrLimit}
	observations, err := scanner.Scan(context.Background())
	if err != nil || observations[0].Status() != core.RootPartial || observations[0].Reason() != core.ReasonEntryVisibilityGap || !reflect.DeepEqual(observations[0].WarningCodes(), []core.WarningCode{core.WarningEntryLimitReached}) || observations[0].Estimate().Logical().Completeness() != core.CompletenessIncomplete {
		t.Fatalf("observations=%#v err=%v", observations, err)
	}
}

func TestScannerPermittedSymlinksDoNotMakeTraversalPartial(t *testing.T) {
	for _, tc := range []struct {
		name    string
		facts   []DirectoryFact
		walkErr error
	}{
		{"walker classification", []DirectoryFact{regularFact(t, "accepted", 10)}, ErrSymlink},
		{"callback symlink only", []DirectoryFact{directoryFact(t, []string{"link"}, EntrySymlink, Identity{}, Identity{}, true)}, nil},
		{"accepted and callback symlink", []DirectoryFact{regularFact(t, "accepted", 10), directoryFact(t, []string{"link"}, EntrySymlink, Identity{}, Identity{}, true)}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			scanner := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: tc.facts})
			scanner.walker.(*filePortWiringWalker).errs = map[core.AreaID]error{core.AreaNPMCache: tc.walkErr}
			observations, err := scanner.Scan(context.Background())
			if err != nil || observations[0].Status() != core.RootScanned || observations[0].Reason() != core.ReasonCompleted || !reflect.DeepEqual(observations[0].WarningCodes(), []core.WarningCode{core.WarningSymlinkSkipped}) {
				t.Fatalf("observations=%#v err=%v", observations, err)
			}
		})
	}
}

func TestScannerUnsafeFileReplacementSymlinkIsPartial(t *testing.T) {
	root := scannerRoot(t, core.AreaNPMCache)
	root.outcomes = map[string]fileOutcome{"rejected": {cause: ErrSymlink}}
	scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "accepted", 10), regularFact(t, "rejected", 11)}})
	observations, err := scanner.Scan(context.Background())
	if err != nil || observations[0].Status() != core.RootPartial || observations[0].Reason() != core.ReasonEntryVisibilityGap || !reflect.DeepEqual(observations[0].WarningCodes(), []core.WarningCode{core.WarningSymlinkSkipped}) || observations[0].Estimate().Logical().Completeness() != core.CompletenessIncomplete {
		t.Fatalf("observations=%#v err=%v", observations, err)
	}
}

func TestScannerTraversalDeterminesRootCompleteness(t *testing.T) {
	for _, tc := range []struct {
		name     string
		facts    []DirectoryFact
		status   core.RootStatus
		complete core.EstimateCompleteness
		findings uint64
	}{
		{"accepted only", []DirectoryFact{regularFact(t, "accepted", 10)}, core.RootScanned, core.CompletenessComplete, 1},
		{"accepted with gap", []DirectoryFact{regularFact(t, "accepted", 10), directoryFact(t, []string{"special"}, EntrySpecial, Identity{}, Identity{}, true)}, core.RootPartial, core.CompletenessIncomplete, 1},
		// Active design § Finding granularity: empty and directory-only roots are complete.
		{"no accepted evidence", []DirectoryFact{directoryFact(t, []string{"directory"}, EntryDirectory, Identity{Device: 7, Inode: 11}, Identity{Device: 7, Inode: 11}, true)}, core.RootScanned, core.CompletenessComplete, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			scanner := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: tc.facts})
			observations, err := scanner.Scan(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			facts, err := scanner.FinalizedFileFacts()
			if err != nil {
				t.Fatal(err)
			}
			if got := observations[0].Status(); got != tc.status {
				t.Fatalf("Scan root status = %s, want %s", got, tc.status)
			}
			root := facts.Roots()[0]
			if root.Status() != tc.status || root.Estimate().Logical().Completeness() != tc.complete || root.FindingCount() != tc.findings {
				t.Fatalf("finalized root = %#v, want status=%s completeness=%s findings=%d", root, tc.status, tc.complete, tc.findings)
			}
		})
	}
}

func TestScannerEmptyRootFinalizesCompleteZeroEstimate(t *testing.T) {
	scanner := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{})
	observations, err := scanner.Scan(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	facts, err := scanner.FinalizedFileFacts()
	if err != nil {
		t.Fatal(err)
	}
	if got := observations[0]; got.Status() != core.RootScanned || got.Reason() != core.ReasonCompleted {
		t.Fatalf("empty observation = %#v", got)
	}
	root := facts.Roots()[0]
	if root.Status() != core.RootScanned || root.Reason() != core.ReasonCompleted || root.Estimate().Logical().KnownBytes() != 0 || root.Estimate().Logical().Completeness() != core.CompletenessComplete || root.FindingCount() != 0 || len(facts.Findings()) != 0 {
		t.Fatalf("empty finalized root = %#v findings=%#v", root, facts.Findings())
	}
}

func TestScannerFinalizedFactsPreserveRootObservations(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	release := make(chan struct{})
	root := scannerRoot(t, core.AreaNPMCache)
	root.outcomes = map[string]fileOutcome{"accepted": {}}
	root.started = make(chan struct{}, 1)
	root.cancelled = make(chan struct{}, 1)
	root.releases = map[string]<-chan struct{}{"accepted": release}
	roots := map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}
	scanner := scannerWithRoots(t, roots, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "accepted", 10)}})
	done := make(chan error, 1)
	go func() {
		_, err := scanner.Scan(ctx)
		done <- err
	}()
	<-root.started
	cancel()
	<-root.cancelled // The active inspection observed cancellation before release/drain/root close.
	close(release)
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	facts, err := scanner.FinalizedFileFacts()
	if err != nil {
		t.Fatal(err)
	}
	got := facts.Roots()
	if len(got) != len(builtinDefinitions) || got[0].Status() != core.RootCancelled || got[0].Reason() != core.ReasonCancelledDuringScan || got[1].Status() != core.RootSkipped || got[1].Reason() != core.ReasonCancelledBeforeStart {
		t.Fatalf("finalized roots lost scan state: %#v", got)
	}
	if len(facts.Findings()) != 1 || facts.Findings()[0].Estimate().Logical().KnownBytes() != 10 {
		t.Fatalf("finalized findings=%#v", facts.Findings())
	}
	got[0] = core.RootObservation{}
	if facts.Roots()[0].Status() != core.RootCancelled {
		t.Fatal("Roots accessor leaked mutable slice")
	}
}

func TestScannerWorkerCompletionPermutations(t *testing.T) {
	run := func(first string) finalizedFacts {
		releaseA, releaseB := make(chan struct{}), make(chan struct{})
		root := scannerRoot(t, core.AreaNPMCache)
		root.started = make(chan struct{}, 2)
		root.releases = map[string]<-chan struct{}{"a": releaseA, "b": releaseB}
		scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "a", 10), regularFact(t, "b", 11)}})
		scanner.fileWorkers = 2
		done := make(chan struct{})
		go func() { scanner.Scan(context.Background()); close(done) }()
		for range 2 {
			<-root.started
		}
		if first == "a" {
			close(releaseA)
			close(releaseB)
		} else {
			close(releaseB)
			close(releaseA)
		}
		<-done
		facts, err := scanner.FinalizedFileFacts()
		if err != nil {
			t.Fatal(err)
		}
		return facts.facts
	}
	if !reflect.DeepEqual(factsSignature(run("a")), factsSignature(run("b"))) {
		t.Fatal("worker completion order changed finalization")
	}
}

func TestScannerHardLinkFinalization(t *testing.T) {
	shared := regularFact(t, "shared", 10)
	roots := map[core.AreaID]*filePortRoot{core.AreaNPMCache: scannerRoot(t, core.AreaNPMCache), core.AreaHomebrewCache: scannerRoot(t, core.AreaHomebrewCache)}
	scanner := scannerWithRoots(t, roots, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {shared}, core.AreaHomebrewCache: {shared}})
	if _, err := scanner.Scan(context.Background()); err != nil {
		t.Fatal(err)
	}
	facts, err := scanner.FinalizedFileFacts()
	if err != nil {
		t.Fatal(err)
	}
	if findingByPath(t, facts.Findings(), "~/.npm/shared").Estimate().Logical().KnownBytes() != 10 || findingByPath(t, facts.Findings(), "~/Library/Caches/Homebrew/shared").Estimate().Logical().KnownBytes() != 0 {
		t.Fatal("same identity was not canonically attributed across areas")
	}
	assertRelatedWarning(t, facts.Warnings(), "~/Library/Caches/Homebrew/shared", core.WarningIdentityAlreadyAccounted, "~/.npm/shared")

	within := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "a", 20), regularFact(t, "b", 20)}})
	if _, err := within.Scan(context.Background()); err != nil {
		t.Fatal(err)
	}
	withinFacts, err := within.FinalizedFileFacts()
	if err != nil {
		t.Fatal(err)
	}
	if findingByPath(t, withinFacts.Findings(), "~/.npm/b").Estimate().Logical().KnownBytes() != 0 {
		t.Fatal("same-area hard link was double-counted")
	}
}

func TestScannerCancellationBeforeEnqueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	root := scannerRoot(t, core.AreaNPMCache)
	scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "never-queued", 10)}})
	observations, err := scanner.Scan(ctx)
	if err != nil || len(root.jobs) != 0 || observations[0].Reason() != core.ReasonCancelledBeforeStart {
		t.Fatalf("err=%v jobs=%v observations=%v", err, root.jobs, observations)
	}
}

func TestScannerCancellationStopDrainJoin(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started := make(chan struct{})
	root := cancellationRoot(t, started)
	scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "active", 10), regularFact(t, "queued", 11)}})
	done := make(chan []core.RootObservation, 1)
	go func() { observations, _ := scanner.Scan(ctx); done <- observations }()
	<-started
	cancel()
	observations := <-done
	if len(root.jobs) != 1 || observations[0].Status() != core.RootCancelled || observations[1].Status() != core.RootSkipped {
		t.Fatalf("jobs=%d observations=%v", len(root.jobs), observations)
	}
}

func TestScannerNoPipelineLeaks(t *testing.T) {
	for range 3 {
		ctx, cancel := context.WithCancel(context.Background())
		started := make(chan struct{})
		root := cancellationRoot(t, started)
		scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "active", 10), regularFact(t, "queued", 11)}})
		done := make(chan struct{})
		go func() { scanner.Scan(ctx); close(done) }()
		<-started
		cancel()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("scan did not drain and join workers")
		}
	}
}

func TestScannerDeterministicWarnings(t *testing.T) {
	for _, facts := range [][]DirectoryFact{{regularFact(t, "changed", 10), regularFact(t, "good", 11)}, {regularFact(t, "good", 11), regularFact(t, "changed", 10)}} {
		root := scannerRoot(t, core.AreaNPMCache)
		root.outcomes = map[string]fileOutcome{"changed": {cause: ErrMissing}}
		scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: facts})
		observations, err := scanner.Scan(context.Background())
		if err != nil || !reflect.DeepEqual(observations[0].WarningCodes(), []core.WarningCode{core.WarningEntryChanged}) {
			t.Fatalf("observations=%v err=%v", observations, err)
		}
	}
}

func runFinalizedFiles(t *testing.T, names []string) finalizedFacts {
	t.Helper()
	facts := make([]DirectoryFact, 0, len(names))
	for index, name := range names {
		facts = append(facts, regularFact(t, name, uint64(index+10)))
	}
	scanner := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: facts})
	if _, err := scanner.Scan(context.Background()); err != nil {
		t.Fatal(err)
	}
	result, err := scanner.FinalizedFileFacts()
	if err != nil {
		t.Fatal(err)
	}
	return result.facts
}

func regularFact(t *testing.T, name string, inode uint64) DirectoryFact {
	t.Helper()
	identity := Identity{Device: 7, Inode: inode}
	fact, err := NewDirectoryFact([]string{name}, EntryRegular, &identity, EntryRegular, &identity, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	return fact
}

func cancellationRoot(t *testing.T, started chan struct{}) *filePortRoot {
	t.Helper()
	root := scannerRoot(t, core.AreaNPMCache)
	root.started = started
	root.block = true
	return root
}

func scannerWithRoots(t *testing.T, roots map[core.AreaID]*filePortRoot, facts map[core.AreaID][]DirectoryFact) Scanner {
	t.Helper()
	if roots == nil {
		roots = map[core.AreaID]*filePortRoot{}
	}
	var events []string
	for _, definition := range builtinDefinitions {
		if roots[definition.AreaID] == nil {
			roots[definition.AreaID] = scannerRoot(t, definition.AreaID)
		}
		roots[definition.AreaID].events = &events
	}
	return NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), &filePortWiringWalker{roots: roots, facts: facts, events: &events}, WalkLimits{MaxDepth: 2, MaxDescriptors: 2})
}

func scannerRoot(t *testing.T, area core.AreaID) *filePortRoot {
	t.Helper()
	facts, err := NewRootFacts(Identity{Device: 7, Inode: uint64(core.AreaRank(area) + 1)}, true)
	if err != nil {
		t.Fatal(err)
	}
	return &filePortRoot{facts: facts, area: area}
}

func TestScannerFinalizedFactsRetainAcceptedTotalsForIncompleteRoots(t *testing.T) {
	scanner := Scanner{results: &scannerFileResults{}}
	definitions := scannerDefinitions(t, "/fixture/home")
	partial := rootObservation(definitions[0], core.RootPartial, core.ReasonEntryVisibilityGap, []core.WarningCode{core.WarningEntryInaccessible})
	boundary := rootObservation(definitions[1], core.RootBoundaryLimited, core.ReasonDeviceBoundary, []core.WarningCode{core.WarningDeviceBoundary})
	inaccessible := inaccessible(definitions[2])
	cancelled := cancelledDuring(definitions[3])
	skipped := cancelledBefore(definitions[4])
	scanner.results.roots = []core.RootObservation{partial, boundary, inaccessible, cancelled, skipped}
	scanner.results.raw = []rawRoot{
		{area: definitions[0].AreaID, displayName: definitions[0].DisplayName, displayPath: definitions[0].DisplayPath, observations: []rawObservation{{displayPath: definitions[0].DisplayPath + "/partial", reason: definitions[0].FindingReason, logicalBytes: 10, allocation: core.KnownAllocation(5), identity: knownIdentity(t, 7, 10)}}},
		{area: definitions[1].AreaID, displayName: definitions[1].DisplayName, displayPath: definitions[1].DisplayPath, observations: []rawObservation{{displayPath: definitions[1].DisplayPath + "/boundary", reason: definitions[1].FindingReason, logicalBytes: 20, allocation: core.KnownAllocation(10), identity: knownIdentity(t, 7, 20)}}},
		{area: definitions[3].AreaID, displayName: definitions[3].DisplayName, displayPath: definitions[3].DisplayPath, observations: []rawObservation{{displayPath: definitions[3].DisplayPath + "/cancelled", reason: definitions[3].FindingReason, logicalBytes: 30, allocation: core.KnownAllocation(15), identity: knownIdentity(t, 7, 30)}}},
	}

	facts, err := scanner.FinalizedFileFacts()
	if err != nil {
		t.Fatal(err)
	}
	for index, want := range []struct {
		status core.RootStatus
		reason core.RootReasonCode
		codes  []core.WarningCode
		bytes  uint64
		count  uint64
	}{
		{core.RootPartial, core.ReasonEntryVisibilityGap, []core.WarningCode{core.WarningEntryInaccessible}, 10, 1},
		{core.RootBoundaryLimited, core.ReasonDeviceBoundary, []core.WarningCode{core.WarningDeviceBoundary}, 20, 1},
		{core.RootInaccessible, core.ReasonRootInspectionFailed, []core.WarningCode{core.WarningRootInaccessible}, 0, 0},
		{core.RootCancelled, core.ReasonCancelledDuringScan, []core.WarningCode{core.WarningCancelledDuringRoot}, 30, 1},
		{core.RootSkipped, core.ReasonCancelledBeforeStart, []core.WarningCode{core.WarningCancelledBeforeRoot}, 0, 0},
	} {
		got := facts.Roots()[index]
		if got.Status() != want.status || got.Reason() != want.reason || !reflect.DeepEqual(got.WarningCodes(), want.codes) || got.Estimate().Logical().KnownBytes() != want.bytes || got.FindingCount() != want.count {
			t.Fatalf("root %d = status=%s reason=%s warnings=%v bytes=%d count=%d; want status=%s reason=%s warnings=%v bytes=%d count=%d", index, got.Status(), got.Reason(), got.WarningCodes(), got.Estimate().Logical().KnownBytes(), got.FindingCount(), want.status, want.reason, want.codes, want.bytes, want.count)
		}
		if want.status != core.RootScanned && want.status != core.RootMissing && got.Estimate().Logical().Completeness() != core.CompletenessIncomplete {
			t.Fatalf("root %d claimed complete aggregate despite %s status", index, got.Status())
		}
	}
}

type stoppingFilePort struct {
	started  chan string
	releases map[string]<-chan struct{}
	mu       sync.Mutex
	inspects []string
}

func (p *stoppingFilePort) Inspect(_ context.Context, job FileJob) FileResult {
	name := strings.Join(job.Components(), "/")
	p.mu.Lock()
	p.inspects = append(p.inspects, name)
	p.mu.Unlock()
	p.started <- name
	<-p.releases[name]
	result, err := NewFileResult(job, EntryRegular, job.FinalIdentity(), 10, 5, nil)
	if err != nil {
		panic(err)
	}
	return result
}

func (p *stoppingFilePort) inspected() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return append([]string(nil), p.inspects...)
}

func TestScannerCancellationAfterWorkerAcceptanceAndWhileBlocked(t *testing.T) {
	t.Run("after worker acceptance drains active result", func(t *testing.T) {
		release := make(chan struct{})
		port := &stoppingFilePort{started: make(chan string, 1), releases: map[string]<-chan struct{}{"active": release}}
		pipeline, err := NewFilePipeline(port, 1, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
		active := mustFileJob(t, "active", 41)
		if err := pipeline.Submit(context.Background(), active); err != nil {
			t.Fatal(err)
		}
		if got := <-port.started; got != "active" {
			t.Fatalf("accepted worker job = %q", got)
		}
		pipeline.Stop()
		close(release)
		results := pipeline.Drain()
		if len(results) != 1 || !results[0].Accepted() || !reflect.DeepEqual(results[0].Components(), active.Components()) || !reflect.DeepEqual(port.inspected(), []string{"active"}) {
			t.Fatalf("active result or joined worker lost: results=%#v inspected=%v", results, port.inspected())
		}
	})

	t.Run("while blocked cancellation prevents queued inspection", func(t *testing.T) {
		release := make(chan struct{})
		port := &stoppingFilePort{started: make(chan string, 2), releases: map[string]<-chan struct{}{"active": release, "queued": make(chan struct{})}}
		pipeline, err := NewFilePipeline(port, 1, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
		active, queued := mustFileJob(t, "active", 51), mustFileJob(t, "queued", 52)
		if err := pipeline.Submit(context.Background(), active); err != nil {
			t.Fatal(err)
		}
		if got := <-port.started; got != "active" {
			t.Fatalf("active worker job = %q", got)
		}
		if err := pipeline.Submit(context.Background(), queued); err != nil {
			t.Fatal(err)
		}
		pipeline.Stop()
		close(release)
		results := pipeline.Drain()
		if len(results) != 1 || !results[0].Accepted() || !reflect.DeepEqual(results[0].Components(), active.Components()) || !reflect.DeepEqual(port.inspected(), []string{"active"}) {
			t.Fatalf("post-cancel inspection or active drain failure: results=%#v inspected=%v", results, port.inspected())
		}
		select {
		case got := <-port.started:
			t.Fatalf("post-cancel Inspect ran for %q", got)
		default:
		}
	})
}

type limitWalker struct {
	limits []WalkLimits
	facts  []DirectoryFact
	visits int
}

func (w *limitWalker) AcquireRoot(_ context.Context, _ AbsoluteComponents, _ AbsoluteComponents, limits WalkLimits) (TrustedRoot, error) {
	w.limits = append(w.limits, limits)
	facts, _ := NewRootFacts(Identity{Device: 7, Inode: uint64(len(w.limits))}, true)
	return NewTrustedRoot(facts, func() error { return nil })
}

func (w *limitWalker) WalkFacts(_ context.Context, root TrustedRoot, visit func(DirectoryFact) error) error {
	if root.Facts().Identity.Inode != 1 {
		return nil
	}
	for _, fact := range w.facts {
		w.visits++
		if err := visit(fact); err != nil {
			return err
		}
	}
	return nil
}

func TestScannerStatusPrecedence(t *testing.T) {
	for _, tc := range []struct {
		name                         string
		cancelled, partial, boundary bool
		want                         core.RootStatus
	}{
		{"scanned", false, false, false, core.RootScanned},
		{"boundary limited", false, false, true, core.RootBoundaryLimited},
		{"partial over boundary", false, true, true, core.RootPartial},
		{"cancelled over partial", true, true, true, core.RootCancelled},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := rootStatus(tc.cancelled, tc.partial, tc.boundary); got != tc.want {
				t.Fatalf("rootStatus() = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestScannerCancellationBeforeRoots(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	walker := &scannerWalker{}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}).Scan(ctx)
	if err != nil || len(walker.calls) != 0 || len(observations) != len(builtinDefinitions) {
		t.Fatalf("err=%v calls=%v observations=%v", err, walker.calls, observations)
	}
	for _, observation := range observations {
		if observation.Status() != core.RootSkipped || observation.Reason() != core.ReasonCancelledBeforeStart {
			t.Fatalf("pre-cancelled root = %#v", observation)
		}
	}
}

func TestScannerCancellationDuringRoot(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	release := make(chan struct{})
	root := scannerRoot(t, core.AreaNPMCache)
	root.started = make(chan struct{}, 1)
	root.releases = map[string]<-chan struct{}{"active": release}
	boundary := directoryFact(t, []string{"outside"}, EntryDirectory, Identity{Device: 8, Inode: 9}, Identity{Device: 8, Inode: 9}, false)
	scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {boundary, regularFact(t, "active", 10), regularFact(t, "queued", 11)}})
	done := make(chan []core.RootObservation, 1)
	go func() { observations, _ := scanner.Scan(ctx); done <- observations }()
	<-root.started
	cancel()
	close(release)
	observations := <-done
	if observations[0].Status() != core.RootCancelled || observations[0].Reason() != core.ReasonCancelledDuringScan || observations[1].Reason() != core.ReasonCancelledBeforeStart || !reflect.DeepEqual(observations[0].WarningCodes(), []core.WarningCode{core.WarningCancelledDuringRoot, core.WarningDeviceBoundary}) {
		t.Fatalf("cancellation observations = %#v", observations)
	}
	if got := scanner.AcceptedFileResults(); len(got) != 1 || got[0].Components()[0] != "active" {
		t.Fatalf("accepted active result was not drained: %#v", got)
	}
}

func TestScannerCancellationAfterRoots(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	walker := &scannerWalker{onCall: func(area core.AreaID) {
		if area == core.AreaHomebrewCache {
			cancel()
		}
	}}
	observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}).Scan(ctx)
	if err != nil || len(walker.calls) != 2 || observations[0].Status() != core.RootScanned || observations[1].Status() != core.RootCancelled {
		t.Fatalf("err=%v calls=%v observations=%#v", err, walker.calls, observations)
	}
	for _, observation := range observations[2:] {
		if observation.Reason() != core.ReasonCancelledBeforeStart {
			t.Fatalf("later root started after cancellation: %#v", observation)
		}
	}
}

func TestScannerNoLeakedWork(t *testing.T) {
	for range 3 {
		ctx, cancel := context.WithCancel(context.Background())
		started := make(chan struct{}, 1)
		root := cancellationRoot(t, started)
		scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "active", 10), regularFact(t, "queued", 11)}})
		done := make(chan struct{})
		go func() { scanner.Scan(ctx); close(done) }()
		<-started
		cancel()
		<-done
	}
}

func TestScannerCompletionOrderStable(t *testing.T) {
	if first, second := runFinalizedFiles(t, []string{"z", "a"}), runFinalizedFiles(t, []string{"a", "z"}); !reflect.DeepEqual(factsSignature(first), factsSignature(second)) {
		t.Fatalf("completion order changed finalization: %v != %v", factsSignature(first), factsSignature(second))
	}
}

func TestScannerDefaultPolicy(t *testing.T) {
	walker := &limitWalker{}
	limits := DefaultWalkLimits()
	scanner := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, limits)
	if got := scanner.Policy(); got != limits {
		t.Fatalf("Policy() = %+v, want %+v", got, limits)
	}
	if _, err := scanner.Scan(context.Background()); err != nil || len(walker.limits) != 5 {
		t.Fatalf("Scan() err=%v limits=%v", err, walker.limits)
	}
	for _, got := range walker.limits {
		if got != limits {
			t.Fatalf("walker limit = %+v, want unchanged %+v", got, limits)
		}
	}
}

func TestScannerEntryLimit(t *testing.T) {
	for _, tc := range []struct {
		name, wantWarning string
		entries           int
		visits            int
	}{
		{"empty root", "", 0, 0},
		{"one below boundary", "", 1, 1},
		{"exact boundary", "", 2, 2},
		{"first over stops", string(core.WarningEntryLimitReached), 4, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			limits := DefaultWalkLimits()
			limits.MaxEntries = 2
			walker := &limitWalker{}
			for i := 0; i < tc.entries; i++ {
				identity := Identity{Device: 7, Inode: uint64(i + 10)}
				walker.facts = append(walker.facts, directoryFact(t, []string{string(rune('a' + i))}, EntryDirectory, identity, identity, true))
			}
			observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, limits).Scan(context.Background())
			if err != nil || walker.visits != tc.visits || len(walker.limits) != 5 {
				t.Fatalf("err=%v visits=%d limits=%d", err, walker.visits, len(walker.limits))
			}
			codes := observations[0].WarningCodes()
			if tc.wantWarning == "" && len(codes) != 0 {
				t.Fatalf("exact boundary warnings=%v", codes)
			}
			if tc.wantWarning != "" && !reflect.DeepEqual(codes, []core.WarningCode{core.WarningEntryLimitReached}) {
				t.Fatalf("over-boundary warnings=%v", codes)
			}
		})
	}
}

func TestScannerPathBudget(t *testing.T) {
	utf8Path := directoryFact(t, []string{"é", "文"}, EntryDirectory, Identity{Device: 7, Inode: 10}, Identity{Device: 7, Inode: 10}, true)
	next := directoryFact(t, []string{"x"}, EntryDirectory, Identity{Device: 7, Inode: 11}, Identity{Device: 7, Inode: 11}, true)
	a := directoryFact(t, []string{"a"}, EntryDirectory, Identity{Device: 7, Inode: 12}, Identity{Device: 7, Inode: 12}, true)
	b := directoryFact(t, []string{"b"}, EntryDirectory, Identity{Device: 7, Inode: 13}, Identity{Device: 7, Inode: 13}, true)
	for _, tc := range []struct {
		name        string
		maxEntries  int
		maxPath     int
		facts       []DirectoryFact
		visits      int
		wantWarning core.WarningCode
	}{
		{"utf8 exact boundary", 2, len("é/文"), []DirectoryFact{utf8Path}, 1, ""},
		{"utf8 first byte over", 2, len("é/文"), []DirectoryFact{utf8Path, next, next}, 2, core.WarningPathBudgetReached},
		{"combined exhaustion selects entry first", 1, len("é/文"), []DirectoryFact{utf8Path, next, next}, 2, core.WarningEntryLimitReached},
		{"combined exhaustion stays entry-first when reordered", 1, 1, []DirectoryFact{a, b}, 2, core.WarningEntryLimitReached},
		{"combined exhaustion stays entry-first in reverse order", 1, 1, []DirectoryFact{b, a}, 2, core.WarningEntryLimitReached},
	} {
		t.Run(tc.name, func(t *testing.T) {
			limits := DefaultWalkLimits()
			limits.MaxEntries, limits.MaxPathBytes = tc.maxEntries, tc.maxPath
			walker := &limitWalker{facts: tc.facts}
			observations, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, limits).Scan(context.Background())
			if err != nil || walker.visits != tc.visits {
				t.Fatalf("err=%v visits=%d", err, walker.visits)
			}
			if tc.wantWarning == "" && len(observations[0].WarningCodes()) != 0 {
				t.Fatalf("exact boundary warnings=%v", observations[0].WarningCodes())
			}
			if tc.wantWarning != "" && !reflect.DeepEqual(observations[0].WarningCodes(), []core.WarningCode{tc.wantWarning}) {
				t.Fatalf("warnings=%v, want %s", observations[0].WarningCodes(), tc.wantWarning)
			}
		})
	}
}

func TestScannerProgressEvents(t *testing.T) {
	for _, tc := range []struct {
		name    string
		acquire error
		facts   []DirectoryFact
		want    []ProgressEventKind
		status  core.RootStatus
	}{
		{"complete", nil, nil, []ProgressEventKind{ProgressScanStarted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressScanFinished}, core.RootScanned},
		{"missing", ErrMissing, nil, []ProgressEventKind{ProgressScanStarted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressScanFinished}, core.RootMissing},
		{"inaccessible", ErrInaccessible, nil, []ProgressEventKind{ProgressScanStarted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressScanFinished}, core.RootInaccessible},
		{"partial", nil, []DirectoryFact{directoryFact(t, []string{"special"}, EntrySpecial, Identity{}, Identity{}, true)}, []ProgressEventKind{ProgressScanStarted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressScanFinished}, core.RootPartial},
	} {
		t.Run(tc.name, func(t *testing.T) {
			walker := &directoryScriptWalker{scannerWalker: &scannerWalker{results: map[core.AreaID]error{core.AreaNPMCache: tc.acquire}}, scripts: map[core.AreaID][]DirectoryFact{core.AreaNPMCache: tc.facts}}
			var events []ProgressEvent
			roots, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, WalkLimits{MaxDepth: 2, MaxDescriptors: 2}).ScanWithProgress(context.Background(), func(event ProgressEvent) { events = append(events, event) })
			if err != nil || len(events) != len(tc.want) || roots[0].Status() != tc.status {
				t.Fatalf("roots=%v events=%v err=%v", roots, events, err)
			}
			for i, want := range tc.want {
				got := events[i]
				if got.Kind() != want || got.Validate() != nil || got.Total() != len(builtinDefinitions) || (want == ProgressCategoryCompleted && got.Status() != roots[got.Processed()-1].Status()) {
					t.Fatalf("event %d = %#v, want kind=%s", i, got, want)
				}
				if want == ProgressCategoryStarted || want == ProgressCategoryCompleted {
					index := got.Processed()
					if want == ProgressCategoryCompleted {
						index--
					}
					if got.CategoryID() != builtinDefinitions[index].AreaID || got.CategoryDisplayName() != builtinDefinitions[index].DisplayName {
						t.Fatalf("event %d category=%s/%q", i, got.CategoryID(), got.CategoryDisplayName())
					}
				}
			}
			if ProgressEventKind("invalid").Validate() == nil {
				t.Fatal("invalid progress kind was accepted")
			}
		})
	}
}

func TestScannerProgressCancellationAndEquivalence(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	root := cancellationRoot(t, make(chan struct{}, 1))
	scanner := scannerWithRoots(t, map[core.AreaID]*filePortRoot{core.AreaNPMCache: root}, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "active", 10)}})
	var events []ProgressEvent
	done := make(chan []core.RootObservation, 1)
	go func() {
		roots, _ := scanner.ScanWithProgress(ctx, func(event ProgressEvent) { events = append(events, event) })
		done <- roots
	}()
	<-root.started
	cancel()
	roots := <-done
	if got := []ProgressEventKind{events[0].Kind(), events[1].Kind(), events[2].Kind(), events[3].Kind()}; !reflect.DeepEqual(got, []ProgressEventKind{ProgressScanStarted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressScanCancelled}) || roots[0].Status() != core.RootCancelled || events[2].Status() != core.RootCancelled || events[3].Processed() != 1 {
		t.Fatalf("roots=%v events=%v", roots, events)
	}
	plain := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "same", 11)}})
	withProgress := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "same", 11)}})
	wantRoots, _ := plain.Scan(context.Background())
	gotRoots, _ := withProgress.ScanWithProgress(context.Background(), nil)
	wantFacts, _ := plain.FinalizedFileFacts()
	gotFacts, _ := withProgress.FinalizedFileFacts()
	if !reflect.DeepEqual(wantRoots, gotRoots) || !reflect.DeepEqual(wantFacts.facts, gotFacts.facts) {
		t.Fatalf("progress changed scan result")
	}

	preCancelled, stop := context.WithCancel(context.Background())
	stop()
	var preEvents []ProgressEvent
	preRoots, err := plain.ScanWithProgress(preCancelled, func(event ProgressEvent) { preEvents = append(preEvents, event) })
	if err != nil || len(preRoots) != len(builtinDefinitions) || !reflect.DeepEqual([]ProgressEventKind{preEvents[0].Kind(), preEvents[1].Kind()}, []ProgressEventKind{ProgressScanStarted, ProgressScanCancelled}) || preEvents[1].Processed() != 0 {
		t.Fatalf("pre-cancelled roots=%v events=%v err=%v", preRoots, preEvents, err)
	}

	concurrent := scannerWithRoots(t, nil, map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {regularFact(t, "a", 12), regularFact(t, "b", 13)}})
	concurrent.fileWorkers = 2
	var mu sync.Mutex
	observing, overlap := false, false
	_, err = concurrent.ScanWithProgress(context.Background(), func(ProgressEvent) {
		mu.Lock()
		if observing {
			overlap = true
		}
		observing = true
		mu.Unlock()
		time.Sleep(time.Millisecond)
		mu.Lock()
		observing = false
		mu.Unlock()
	})
	if err != nil || overlap {
		t.Fatalf("observer overlapped with %d workers: err=%v overlap=%t", concurrent.fileWorkers, err, overlap)
	}
}

func TestScannerLimitsStopRegularJobs(t *testing.T) {
	for _, tc := range []struct {
		name     string
		maxEntry int
		maxPath  int
	}{
		{"entry exhaustion", 1, 64},
		{"path exhaustion", 64, len("first")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var events []string
			roots := make(map[core.AreaID]*filePortRoot)
			for _, definition := range builtinDefinitions {
				facts, _ := NewRootFacts(Identity{Device: 7, Inode: uint64(core.AreaRank(definition.AreaID) + 1)}, true)
				roots[definition.AreaID] = &filePortRoot{facts: facts, area: definition.AreaID, events: &events}
			}
			walker := &filePortWiringWalker{roots: roots, facts: map[core.AreaID][]DirectoryFact{core.AreaNPMCache: {
				directoryFact(t, []string{"first"}, EntryRegular, Identity{Device: 7, Inode: 10}, Identity{Device: 7, Inode: 10}, true),
				directoryFact(t, []string{"second"}, EntryRegular, Identity{Device: 7, Inode: 11}, Identity{Device: 7, Inode: 11}, true),
			}}, events: &events}
			limits := DefaultWalkLimits()
			limits.MaxEntries, limits.MaxPathBytes = tc.maxEntry, tc.maxPath
			if _, err := NewScanner(&scannerHome{home: "/fixture/home"}, scannerDefinitions(t, "/fixture/home"), walker, limits).Scan(context.Background()); err != nil {
				t.Fatal(err)
			}
			jobs := roots[core.AreaNPMCache].jobs
			if len(jobs) != 1 || strings.Join(jobs[0].Components(), "/") != "first" {
				t.Fatalf("regular jobs after limit = %#v", jobs)
			}
		})
	}
}

package scan

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestDescriptorComponents(t *testing.T) {
	for _, tt := range []struct {
		name string
		path string
		ok   bool
	}{
		{"absolute", "/home/alice/.npm", true},
		{"root", "/", true},
		{"relative", "home/alice", false},
		{"repeated", "/home//alice", false},
		{"trailing", "/home/alice/", false},
		{"dot", "/home/./alice", false},
		{"parent", "/home/../alice", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAbsoluteComponents(tt.path)
			if tt.ok {
				if err != nil || got.String() != tt.path {
					t.Fatalf("NewAbsoluteComponents(%q) = %q, %v", tt.path, got.String(), err)
				}
				return
			}
			if !errors.Is(err, ErrInvalidComponents) {
				t.Fatalf("NewAbsoluteComponents(%q) error = %v", tt.path, err)
			}
		})
	}
	got, _ := NewAbsoluteComponents("/home/alice")
	copy := got.Components()
	copy[0] = "changed"
	if want := []string{"home", "alice"}; !reflect.DeepEqual(got.Components(), want) {
		t.Fatal(got.Components())
	}
}

func TestDescriptorRootPolicy(t *testing.T) {
	anchor, _ := NewAbsoluteComponents("/")
	if anchor.String() != "/" || !errors.Is(anchor.ValidateScanRoot(), ErrInvalidComponents) {
		t.Fatal("root policy")
	}
}

func TestDescriptorRelativeChains(t *testing.T) {
	input := []string{"Library", "Caches"}
	got, err := NewRelativeComponents(input)
	if err != nil {
		t.Fatal(err)
	}
	input[0] = "changed"
	copy := got.Components()
	copy[1] = "changed"
	if want := []string{"Library", "Caches"}; !reflect.DeepEqual(got.Components(), want) {
		t.Fatal(got.Components())
	}
	for _, chain := range [][]string{nil, {}, {""}, {"."}, {".."}, {"a/b"}, {"a\x00b"}} {
		mustInvalid(t, chain)
	}
}

func TestDescriptorLimits(t *testing.T) {
	if err := (WalkLimits{MaxDepth: 64, MaxDescriptors: 1}).Validate(); err != nil {
		t.Fatalf("independent descriptor cap: %v", err)
	}
}

func TestDescriptorLimitSemantics(t *testing.T) {
	for _, tt := range []struct {
		name   string
		limits WalkLimits
		valid  bool
	}{
		{"zero depth", WalkLimits{MaxDescriptors: 1}, false},
		{"zero descriptors", WalkLimits{MaxDepth: 1}, false},
		{"minimum independent limits", WalkLimits{MaxDepth: 1, MaxDescriptors: 1}, true},
		{"depth-heavy descriptor-light", WalkLimits{MaxDepth: 64, MaxDescriptors: 1}, true},
		{"depth-light descriptor-heavy", WalkLimits{MaxDepth: 1, MaxDescriptors: 64}, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.limits.Validate()
			if tt.valid && err != nil {
				t.Fatalf("Validate(%+v): %v", tt.limits, err)
			}
			if !tt.valid && !errors.Is(err, ErrInvalidLimits) {
				t.Fatalf("Validate(%+v) = %v, want ErrInvalidLimits", tt.limits, err)
			}
		})
	}
}

func TestDescriptorFactsAndIdentity(t *testing.T) {
	identity := Identity{Device: 7, Inode: 9}
	facts, err := NewRootFacts(identity, true)
	if err != nil || facts.Identity != identity || !facts.Local {
		t.Fatal(facts, err)
	}
	if !errors.Is((Identity{}).Validate(), ErrInvalidMetadata) {
		t.Fatal("zero identity")
	}
	if _, err := NewRootFacts(identity, false); !errors.Is(err, ErrNonLocal) {
		t.Fatal(err)
	}
}

func TestDescriptorPathErrorRejectsInvalidContext(t *testing.T) {
	for _, tt := range []struct {
		name      string
		operation string
		path      string
		class     error
	}{
		{"empty operation", "", ".npm", ErrSymlink},
		{"empty path", "openat", "", ErrSymlink},
		{"nil class", "openat", ".npm", nil},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPathError(tt.operation, tt.path, tt.class)
			if got != nil || !errors.Is(err, ErrInvalidMetadata) {
				t.Fatalf("NewPathError(%q, %q, %v) = %v, %v", tt.operation, tt.path, tt.class, got, err)
			}
		})
	}

	inner, err := NewPathError("fstat", ".npm", ErrDeviceBoundary)
	if err != nil {
		t.Fatal(err)
	}
	wrapped, err := NewPathError("openat", ".npm", inner)
	if err != nil || wrapped.Operation() != "openat" || wrapped.Path() != ".npm" || !errors.Is(wrapped, ErrDeviceBoundary) {
		t.Fatalf("wrapped path error = %v, %v", wrapped, err)
	}
	var typed *PathError
	if !errors.As(wrapped, &typed) || typed != wrapped {
		t.Fatalf("wrapped path error type = %v", wrapped)
	}
}

func TestDescriptorPathError(t *testing.T) {
	inner, err := NewPathError("fstat", ".npm", ErrDeviceBoundary)
	if err != nil {
		t.Fatal(err)
	}
	got, err := NewPathError("openat", ".npm", inner)
	if err != nil {
		t.Fatal(err)
	}
	if !errors.Is(got, ErrDeviceBoundary) || got.Operation() != "openat" || got.Path() != ".npm" {
		t.Fatal(got)
	}
	var typed *PathError
	if !errors.As(got, &typed) {
		t.Fatal("not typed")
	}
}

func TestDescriptorOwnershipTransfer(t *testing.T) {
	facts := testRootFacts(t)
	closeResult := errors.New("close failed")
	root, err := NewTrustedRoot(facts, func() error { return closeResult })
	if err != nil {
		t.Fatal(err)
	}
	walker := fakeDescriptorWalker{root: root}
	home, _ := NewAbsoluteComponents("/home/alice")
	scanRoot, _ := NewAbsoluteComponents("/home/alice/.npm")

	capability, err := walker.AcquireRoot(context.Background(), home, scanRoot, WalkLimits{MaxDepth: 1, MaxDescriptors: 1})
	if err != nil || capability != root || walker.acquires != 1 || capability.Facts() != facts {
		t.Fatalf("AcquireRoot() = %v, %v; acquires=%d", capability, err, walker.acquires)
	}
	if _, err := walker.AcquireRoot(context.Background(), home, scanRoot, WalkLimits{MaxDepth: 1, MaxDescriptors: 1}); !errors.Is(err, ErrInvalidMetadata) || walker.acquires != 1 {
		t.Fatalf("second AcquireRoot() err=%v acquires=%d", err, walker.acquires)
	}
	if got := capability.Close(); got != closeResult {
		t.Fatalf("transferred close result = %v", got)
	}
}

func TestDescriptorCloseIdempotence(t *testing.T) {
	closeResult := &PathError{operation: "close", path: ".npm", err: ErrInaccessible}
	closes := 0
	root, err := NewTrustedRoot(testRootFacts(t), func() error {
		closes++
		return closeResult
	})
	if err != nil {
		t.Fatal(err)
	}

	first := root.Close()
	second := root.Close()
	if closes != 1 || first != second || !errors.Is(second, ErrInaccessible) {
		t.Fatalf("closes=%d first=%v second=%v", closes, first, second)
	}
	var typed *PathError
	if !errors.As(second, &typed) || typed != closeResult {
		t.Fatalf("cached typed close result = %v", second)
	}

	successCloses := 0
	success, err := NewTrustedRoot(testRootFacts(t), func() error { successCloses++; return nil })
	if err != nil || success.Close() != nil || success.Close() != nil || successCloses != 1 {
		t.Fatalf("successful cached close: err=%v closes=%d", err, successCloses)
	}
}

func mustInvalid(t *testing.T, chain []string) {
	t.Helper()
	if _, err := NewRelativeComponents(chain); !errors.Is(err, ErrInvalidComponents) {
		t.Fatal(chain, err)
	}
}

func testRootFacts(t *testing.T) RootFacts {
	t.Helper()
	facts, err := NewRootFacts(Identity{Device: 7, Inode: 9}, true)
	if err != nil {
		t.Fatal(err)
	}
	return facts
}

func TestFileJobConstructor(t *testing.T) {
	ancestors := []Identity{{Device: 7, Inode: 8}}
	job, err := NewFileJob([]string{"cache", "item"}, ancestors, Identity{Device: 7, Inode: 9})
	if err != nil || !reflect.DeepEqual(job.Components(), []string{"cache", "item"}) || !reflect.DeepEqual(job.Ancestors(), ancestors) || job.FinalIdentity() != (Identity{Device: 7, Inode: 9}) {
		t.Fatalf("job=%#v err=%v", job, err)
	}
	copy := job.Components()
	copy[0] = "changed"
	ancestors[0].Inode = 99
	if job.Components()[0] != "cache" || job.Ancestors()[0].Inode != 8 {
		t.Fatal("job retained caller storage")
	}
	for _, components := range [][]string{nil, {""}, {"."}, {".."}, {"a/b"}, {"a\x00b"}} {
		if _, err := NewFileJob(components, nil, Identity{Device: 7, Inode: 9}); !errors.Is(err, ErrInvalidMetadata) {
			t.Fatalf("components %q: %v", components, err)
		}
	}
	for _, bad := range [][]Identity{nil, {{Device: 7, Inode: 8}, {Device: 7, Inode: 9}}, {{}}} {
		if _, err := NewFileJob([]string{"parent", "item"}, bad, Identity{Device: 7, Inode: 9}); !errors.Is(err, ErrInvalidMetadata) {
			t.Fatalf("chain %#v: %v", bad, err)
		}
	}
}

func TestFileResultConstructor(t *testing.T) {
	job, _ := NewFileJob([]string{"item"}, nil, Identity{Device: 7, Inode: 9})
	good, err := NewFileResult(job, EntryRegular, Identity{Device: 7, Inode: 9}, 12, 8, nil)
	if err != nil || !good.Accepted() || good.LogicalBytes() != 12 || good.AllocationBytes() != 8 {
		t.Fatalf("result=%#v err=%v", good, err)
	}
	for _, tc := range []struct {
		kind     EntryKind
		identity Identity
		cause    error
	}{{EntryDirectory, Identity{Device: 7, Inode: 9}, nil}, {EntryRegular, Identity{Device: 7, Inode: 10}, nil}, {EntryRegular, Identity{Device: 7, Inode: 9}, ErrInaccessible}} {
		got, err := NewFileResult(job, tc.kind, tc.identity, 12, 8, tc.cause)
		if err != nil || got.Accepted() || got.LogicalBytes() != 0 || got.AllocationBytes() != 0 {
			t.Fatalf("%#v => %#v, %v", tc, got, err)
		}
	}
}

func TestDescriptorEnumerationFacts(t *testing.T) {
	id := Identity{Device: 7, Inode: 9}
	fact, err := NewEnumerationFact("cache", EntryDirectory, &id, true, nil)
	if err != nil || fact.Name() != "cache" || fact.Kind() != EntryDirectory || !fact.Eligible() || fact.Identity() != id || fact.SkipClass() != nil {
		t.Fatalf("fact=%#v err=%v", fact, err)
	}
	for _, name := range []string{"", ".", "..", "a/b", "a\x00b"} {
		if _, err := NewEnumerationFact(name, EntryDirectory, &id, true, nil); !errors.Is(err, ErrInvalidMetadata) {
			t.Fatalf("name %q: %v", name, err)
		}
	}
}

func TestDescriptorEnumerationOwnership(t *testing.T) {
	id := Identity{Device: 7, Inode: 9}
	fact, err := NewEnumerationFact("safe", EntryDirectory, &id, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	id.Inode = 10
	if got := fact.Identity(); got != (Identity{Device: 7, Inode: 9}) {
		t.Fatalf("identity alias: %#v", got)
	}
	changed, err := NewEnumerationFact("changed", EntryUnknown, nil, false, ErrInvalidMetadata)
	if err != nil || changed.Eligible() || !errors.Is(changed.SkipClass(), ErrInvalidMetadata) {
		t.Fatalf("skip=%#v err=%v", changed, err)
	}
}

type fakeDescriptorWalker struct {
	root     TrustedRoot
	acquires int
}

func (w *fakeDescriptorWalker) WalkFacts(context.Context, TrustedRoot, func(DirectoryFact) error) error {
	return nil
}

func (w *fakeDescriptorWalker) AcquireRoot(_ context.Context, _ AbsoluteComponents, _ AbsoluteComponents, _ WalkLimits) (TrustedRoot, error) {
	if w.acquires != 0 {
		return nil, ErrInvalidMetadata
	}
	w.acquires++
	return w.root, nil
}

func TestDefaultScanLimits(t *testing.T) {
	limits := DefaultWalkLimits()
	if got, want := limits, (WalkLimits{MaxDepth: 64, MaxDescriptors: 96, Workers: 8, JobCapacity: 256, ResultCapacity: 256, ConcurrentRoots: 1, MaxEntries: 250000, MaxPathBytes: 67108864}); got != want {
		t.Fatalf("DefaultWalkLimits() = %+v, want %+v", got, want)
	}
	if _, err := NewWalkLimits(limits); err != nil {
		t.Fatalf("NewWalkLimits(default) = %v", err)
	}
	for _, name := range []string{"Workers", "JobCapacity", "ResultCapacity", "ConcurrentRoots", "MaxEntries", "MaxPathBytes"} {
		invalid := limits
		switch name {
		case "Workers":
			invalid.Workers = 0
		case "JobCapacity":
			invalid.JobCapacity = -1
		case "ResultCapacity":
			invalid.ResultCapacity = 0
		case "ConcurrentRoots":
			invalid.ConcurrentRoots = -1
		case "MaxEntries":
			invalid.MaxEntries = 0
		case "MaxPathBytes":
			invalid.MaxPathBytes = -1
		}
		if _, err := NewWalkLimits(invalid); !errors.Is(err, ErrInvalidLimits) {
			t.Fatalf("%s invalid limits error = %v", name, err)
		}
	}
}

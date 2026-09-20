//go:build darwin

package macos

import (
	"context"
	"encoding/binary"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"testing"

	"github.com/osdy/OsdyCleaner/internal/scan"
)

func TestWalkerDescriptorEnumeration(t *testing.T) {
	walker, record := scriptedWalker()
	record.directory = dirRecord(9, directoryType, "z")
	record.directory = append(record.directory, dirRecord(8, directoryType, "a")...)
	record.stats[1], record.stats[2] = statFor(7, 8), statFor(7, 9)
	got, err := walker.enumerateDirectory(context.Background(), 41, testWalkerFacts(t))
	if err != nil || len(got) != 2 || got[0].Name() != "a" || got[1].Name() != "z" || !got[0].Eligible() || !got[1].Eligible() {
		t.Fatalf("facts=%v err=%v", got, err)
	}
	if record.dirFD != 41 {
		t.Fatalf("read fd=%d", record.dirFD)
	}
}

func TestWalkerPreCancelledEnumerationSkipsReadDir(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	walker, record := scriptedWalker()
	facts, err := walker.enumerateDirectory(ctx, 41, testWalkerFacts(t))
	if facts != nil || !errors.Is(err, scan.ErrInaccessible) || record.readDirs != 0 || record.openCount != 0 {
		t.Fatalf("facts=%v err=%v reads=%d opens=%d", facts, err, record.readDirs, record.openCount)
	}
}

func TestWalkerDirectoryRecords(t *testing.T) {
	short := dirRecord(1, directoryType, "a")
	binary.LittleEndian.PutUint16(short[16:18], 0)
	overflow := dirRecord(1, directoryType, "a")
	binary.LittleEndian.PutUint16(overflow[16:18], 99)
	for _, raw := range [][]byte{{1}, short, overflow, dirRecord(1, directoryType, ""), dirRecord(1, directoryType, "a/b"), dirRecord(1, directoryType, "a\x00b")} {
		if _, err := parseDirectoryRecords(raw); !errors.Is(err, scan.ErrInvalidMetadata) {
			t.Fatalf("raw=%v err=%v", raw, err)
		}
	}
	got, err := parseDirectoryRecords(append(append(dirRecord(1, directoryType, "."), dirRecord(2, directoryType, "..")...), dirRecord(3, directoryType, "...")...))
	if err != nil || len(got) != 1 || got[0].name != "..." {
		t.Fatalf("records=%v err=%v", got, err)
	}
	if got, err := parseDirectoryRecords(nil); err != nil || len(got) != 0 {
		t.Fatalf("empty records=%v err=%v", got, err)
	}
}

func TestWalkerDirectoryRecordTermination(t *testing.T) {
	missingTerminator := dirRecord(9, directoryType, "child")
	missingTerminator = missingTerminator[:len(missingTerminator)-1]
	binary.LittleEndian.PutUint16(missingTerminator[16:18], uint16(len(missingTerminator)))
	nonNULTerminator := dirRecord(9, directoryType, "child")
	nonNULTerminator[direntHeader+len("child")] = 'x'
	truncatedTerminator := dirRecord(9, directoryType, "child")
	binary.LittleEndian.PutUint16(truncatedTerminator[16:18], uint16(len(truncatedTerminator)+1))
	declaredLengthOverflow := dirRecord(9, directoryType, "child")
	binary.LittleEndian.PutUint16(declaredLengthOverflow[16:18], 0xffff)

	for _, tt := range []struct {
		name string
		raw  []byte
	}{
		{"missing terminator", missingTerminator},
		{"non-NUL terminator", nonNULTerminator},
		{"truncated terminator", truncatedTerminator},
		{"declared length overflow", declaredLengthOverflow},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			record.directory = tt.raw
			facts, err := walker.enumerateDirectory(context.Background(), 41, testWalkerFacts(t))
			if !errors.Is(err, scan.ErrInvalidMetadata) || facts != nil {
				t.Fatalf("facts=%v err=%v", facts, err)
			}
			if record.openCount != 0 {
				t.Fatalf("malformed record opened %d children", record.openCount)
			}
		})
	}

	maxName := strings.Repeat("m", int(^uint16(0))-direntHeader-1)
	padded := append(dirRecord(9, directoryType, "padded"), 'g', 'a', 'r', 'b', 'a', 'g', 'e')
	binary.LittleEndian.PutUint16(padded[16:18], uint16(len(padded)))
	for _, tt := range []struct {
		name, want string
		raw        []byte
	}{
		{"minimal terminated", "a", dirRecord(9, directoryType, "a")},
		{"padding after terminator", "padded", padded},
		{"maximum representable name", maxName, dirRecord(9, directoryType, maxName)},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			record.directory = tt.raw
			facts, err := walker.enumerateDirectory(context.Background(), 41, testWalkerFacts(t))
			if err != nil || len(facts) != 1 || facts[0].Name() != tt.want || record.openCount != 1 {
				t.Fatalf("facts=%v err=%v opens=%d", facts, err, record.openCount)
			}
		})
	}

	walker, record := scriptedWalker()
	record.directory = append(dirRecord(9, directoryType, "safe"), missingTerminator...)
	facts, err := walker.enumerateDirectory(context.Background(), 41, testWalkerFacts(t))
	if !errors.Is(err, scan.ErrInvalidMetadata) || facts != nil || record.openCount != 0 {
		t.Fatalf("malformed batch facts=%v err=%v opens=%d", facts, err, record.openCount)
	}
}

func TestWalkerChildOpenParentage(t *testing.T) {
	walker, record := scriptedWalker()
	record.directory = dirRecord(9, directoryType, "child")
	_, _ = walker.enumerateDirectory(context.Background(), 41, testWalkerFacts(t))
	if len(record.parents) != 1 || record.parents[0] != 41 || record.flags[len(record.flags)-1] != walkerFlags {
		t.Fatalf("parents=%v flags=%v", record.parents, record.flags)
	}
}

func TestWalkerChildChanges(t *testing.T) {
	for _, stat := range []syscall.Stat_t{statFor(7, 10), {Dev: 7, Ino: 9, Nlink: 1, Mode: syscall.S_IFREG}} {
		walker, record := scriptedWalker()
		record.directory, record.stats[1] = dirRecord(9, directoryType, "changed"), stat
		got, err := walker.enumerateDirectory(context.Background(), 41, testWalkerFacts(t))
		if err != nil || len(got) != 1 || got[0].Eligible() || !errors.Is(got[0].SkipClass(), scan.ErrInvalidMetadata) {
			t.Fatalf("facts=%v err=%v", got, err)
		}
	}
}

func TestWalkerBoundaryGates(t *testing.T) {
	for _, tt := range []struct {
		stat     syscall.Stat_t
		nonLocal bool
		class    error
	}{{statFor(8, 9), false, scan.ErrDeviceBoundary}, {statFor(7, 9), true, scan.ErrNonLocal}} {
		walker, record := scriptedWalker()
		record.directory, record.stats[1], record.nonLocal[1] = dirRecord(9, directoryType, "other"), tt.stat, tt.nonLocal
		got, err := walker.enumerateDirectory(context.Background(), 41, testWalkerFacts(t))
		if err != nil || len(got) != 1 || got[0].Eligible() || !errors.Is(got[0].SkipClass(), tt.class) {
			t.Fatalf("facts=%v err=%v", got, err)
		}
	}
}

func dirRecord(inode uint64, kind uint8, name string) []byte {
	raw := make([]byte, direntHeader+len(name)+1)
	binary.LittleEndian.PutUint64(raw, inode)
	binary.LittleEndian.PutUint16(raw[16:18], uint16(len(raw)))
	binary.LittleEndian.PutUint16(raw[18:20], uint16(len(name)))
	raw[20] = kind
	copy(raw[direntHeader:], name)
	return raw
}
func testWalkerFacts(t *testing.T) scan.RootFacts {
	t.Helper()
	facts, err := scan.NewRootFacts(scan.Identity{Device: 7, Inode: 1}, true)
	if err != nil {
		t.Fatal(err)
	}
	return facts
}

func TestNewDescriptorWalker(t *testing.T) {
	if NewDescriptorWalker() == nil {
		t.Fatal("NewDescriptorWalker returned nil")
	}
}

func TestWalkerRootHomeAncestryPreFSValidation(t *testing.T) {
	walker, record := scriptedWalker()
	validHome := absolute(t, "/home/alice")
	validRoot := absolute(t, "/home/alice/.npm")
	for _, tt := range []struct {
		name       string
		ctx        context.Context
		home, root scan.AbsoluteComponents
		limits     scan.WalkLimits
	}{
		{"nil context", nil, validHome, validRoot, scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2}},
		{"root outside home", context.Background(), validHome, absolute(t, "/var/cache"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2}},
		{"root is home", context.Background(), validHome, validHome, scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2}},
		{"zero descriptors", context.Background(), validHome, validRoot, scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := walker.AcquireRoot(tt.ctx, tt.home, tt.root, tt.limits)
			if !errors.Is(err, scan.ErrInvalidMetadata) || record.openCount != 0 {
				t.Fatalf("AcquireRoot() = %v, opens=%d", err, record.openCount)
			}
		})
	}
}

func TestWalkerAcquisitionOrderAndFlags(t *testing.T) {
	walker, record := scriptedWalker()
	home := absolute(t, "/Users/alice")
	root := absolute(t, "/Users/alice/Library/Caches")
	trusted, err := walker.AcquireRoot(context.Background(), home, root, scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if err != nil {
		t.Fatal(err)
	}
	wantNames := []string{"/", "Users", "alice", "Library", "Caches"}
	if !reflect.DeepEqual(record.names, wantNames) {
		t.Fatalf("opens = %v, want %v", record.names, wantNames)
	}
	wantFlags := syscall.O_RDONLY | syscall.O_NOFOLLOW | syscall.O_DIRECTORY | syscall.O_CLOEXEC | syscall.O_NONBLOCK
	for _, got := range record.flags {
		if got != wantFlags {
			t.Fatalf("flag = %#x, want %#x", got, wantFlags)
		}
	}
	if err := trusted.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestWalkerComponentFacts(t *testing.T) {
	walker, record := scriptedWalker()
	home := absolute(t, "/home/alice")
	root := absolute(t, "/home/alice/.npm")
	record.stats[4] = statFor(7, 44)
	trusted, err := walker.AcquireRoot(context.Background(), home, root, scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if err != nil {
		t.Fatal(err)
	}
	if got := trusted.Facts(); got.Device != 7 || got.Identity != (scan.Identity{Device: 7, Inode: 44}) || !got.Local {
		t.Fatalf("facts = %#v", got)
	}
	if record.fstats != len(record.names) || record.statfs != len(record.names) {
		t.Fatalf("facts calls = %d/%d for %d opens", record.fstats, record.statfs, len(record.names))
	}
}

func TestWalkerEveryPostHomeComponentGate(t *testing.T) {
	for _, tt := range []struct {
		name  string
		stat  syscall.Stat_t
		local bool
		class error
	}{
		{"wrong device", statFor(9, 3), true, scan.ErrDeviceBoundary}, {"non-local", statFor(7, 3), false, scan.ErrNonLocal},
		{"not directory", syscall.Stat_t{Dev: 7, Ino: 3, Nlink: 1, Mode: syscall.S_IFREG}, true, scan.ErrNotDirectory},
		{"invalid identity", syscall.Stat_t{Dev: 7, Ino: 0, Nlink: 1, Mode: syscall.S_IFDIR}, true, scan.ErrInvalidMetadata},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			record.stats[4], record.nonLocal[4] = tt.stat, !tt.local
			_, err := walker.AcquireRoot(context.Background(), absolute(t, "/home/alice"), absolute(t, "/home/alice/one/two"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
			if !errors.Is(err, tt.class) || len(record.names) != 4 {
				t.Fatalf("err=%v opens=%v", err, record.names)
			}
		})
	}
}

func TestWalkerCoreErrorClassification(t *testing.T) {
	for _, tt := range []struct {
		name       string
		err, class error
	}{
		{"missing", syscall.ENOENT, scan.ErrMissing}, {"symlink", syscall.ELOOP, scan.ErrSymlink}, {"non-directory", syscall.ENOTDIR, scan.ErrNotDirectory}, {"inaccessible", syscall.EACCES, scan.ErrInaccessible},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			record.openErr[2] = tt.err
			_, err := walker.AcquireRoot(context.Background(), absolute(t, "/home"), absolute(t, "/home/cache"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
			var typed *scan.PathError
			if !errors.Is(err, tt.class) || !errors.As(err, &typed) || typed.Path() != "/home" {
				t.Fatalf("error = %v", err)
			}
		})
	}
}

func TestWalkerUnsupportedBoundary(t *testing.T) {
	if NewDescriptorWalker() == nil {
		t.Fatal("Darwin constructor must remain constructible")
	}
}

func TestWalkerIntermediateDevicePoison(t *testing.T) {
	walker, record := scriptedWalker()
	record.stats[4] = statFor(9, 4)
	_, err := walker.AcquireRoot(context.Background(), absolute(t, "/home/alice"), absolute(t, "/home/alice/poison/child"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if !errors.Is(err, scan.ErrDeviceBoundary) || len(record.names) != 4 {
		t.Fatalf("err=%v opens=%v", err, record.names)
	}
}

func TestWalkerAcquisitionHandoffCloseFailure(t *testing.T) {
	for _, tt := range []struct {
		name, home, root, wantPath string
		failedCurrent              int
		wantOpened                 []string
	}{
		{"root to first home", "/home/alice", "/home/alice/cache", "/", 1, []string{"/", "home"}},
		{"post home", "/home", "/home/cache/deeper", "/home", 2, []string{"/", "home", "cache"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			record.closeErr[tt.failedCurrent] = syscall.EIO
			trusted, err := walker.AcquireRoot(context.Background(), absolute(t, tt.home), absolute(t, tt.root), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
			var pathErr *scan.PathError
			if trusted != nil || !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &pathErr) || pathErr.Operation() != "close" || pathErr.Path() != tt.wantPath {
				t.Fatalf("trusted=%v err=%v", trusted, err)
			}
			if !reflect.DeepEqual(record.names, tt.wantOpened) || record.closes[tt.failedCurrent] != 1 || record.closes[tt.failedCurrent+1] != 1 {
				t.Fatalf("opens=%v closes=%v", record.names, record.closes)
			}
		})
	}
}

func TestWalkerAcquisitionHandoffPrimaryPrecedence(t *testing.T) {
	walker, record := scriptedWalker()
	record.fstatErr[2], record.closeErr[1], record.closeErr[2] = syscall.EIO, syscall.EIO, syscall.EIO
	trusted, err := walker.AcquireRoot(context.Background(), absolute(t, "/home"), absolute(t, "/home/cache"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	var pathErr *scan.PathError
	if trusted != nil || !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &pathErr) || pathErr.Operation() != "fstat" || pathErr.Path() != "/home" {
		t.Fatalf("trusted=%v err=%v", trusted, err)
	}
	if record.closes[1] != 1 || record.closes[2] != 1 {
		t.Fatalf("closes=%v", record.closes)
	}
}

func TestWalkerAcquisitionCancellationReachability(t *testing.T) {
	stages := []string{"after-open", "after-fstat", "after-fstatfs", "after-validation"}
	for _, stage := range stages {
		for occurrence := 1; occurrence <= 4; occurrence++ {
			t.Run(stage+" component "+string(rune('0'+occurrence)), func(t *testing.T) {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				walker, record := scriptedWalker()
				seen := 0
				record.checkpoint = func(got string) {
					if got == stage {
						seen++
						if seen == occurrence {
							record.reached = got
							cancel()
						}
					}
				}
				trusted, err := walker.AcquireRoot(ctx, absolute(t, "/home"), absolute(t, "/home/cache/deep"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
				var pathErr *scan.PathError
				if trusted != nil || !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &pathErr) || pathErr.Operation() != "acquire" || record.reached != stage {
					t.Fatalf("trusted=%v err=%v reached=%q", trusted, err, record.reached)
				}
				wantFstats, wantStatfs := occurrence, occurrence
				if stage == "after-open" {
					wantFstats, wantStatfs = occurrence-1, occurrence-1
				} else if stage == "after-fstat" {
					wantStatfs = occurrence - 1
				}
				if record.openCount != occurrence || record.fstats != wantFstats || record.statfs != wantStatfs {
					t.Fatalf("opens=%d fstats=%d statfs=%d", record.openCount, record.fstats, record.statfs)
				}
				for fd := 1; fd <= record.openCount; fd++ {
					if record.closes[fd] != 1 {
						t.Fatalf("fd %d closes=%d", fd, record.closes[fd])
					}
				}
			})
		}
	}
	for _, tt := range []struct {
		checkpoint string
		wantOpen   int
	}{
		{"before", 0},
		{"before-handoff", 4},
	} {
		t.Run(tt.checkpoint, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			walker, record := scriptedWalker()
			record.checkpoint = func(got string) {
				if got == tt.checkpoint {
					record.reached = got
					cancel()
				}
			}
			trusted, err := walker.AcquireRoot(ctx, absolute(t, "/home"), absolute(t, "/home/cache/deep"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
			var pathErr *scan.PathError
			if trusted != nil || !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &pathErr) || pathErr.Operation() != "acquire" || record.reached != tt.checkpoint || record.openCount != tt.wantOpen {
				t.Fatalf("trusted=%v err=%v reached=%q opens=%d", trusted, err, record.reached, record.openCount)
			}
			for fd := 1; fd <= record.openCount; fd++ {
				if record.closes[fd] != 1 {
					t.Fatalf("fd %d closes=%d", fd, record.closes[fd])
				}
			}
		})
	}
}

func TestWalkerAcquisitionLimitSemantics(t *testing.T) {
	for _, tt := range []struct {
		name        string
		limits      scan.WalkLimits
		wantOpen    int
		wantCloses  int
		wantTrusted bool
	}{
		{"one descriptor stops before second capability", scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 1}, 1, 1, false},
		{"two descriptors is exact boundary", scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2}, 3, 3, true},
		{"sixty-four descriptors succeeds", scan.WalkLimits{MaxDepth: 64, MaxDescriptors: 64}, 3, 3, true},
		{"depth does not charge acquisition", scan.WalkLimits{MaxDepth: 64, MaxDescriptors: 2}, 3, 3, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			trusted, err := walker.AcquireRoot(context.Background(), absolute(t, "/home"), absolute(t, "/home/cache"), tt.limits)
			if tt.wantTrusted {
				if err != nil || trusted == nil {
					t.Fatalf("trusted=%v err=%v", trusted, err)
				}
				if err := trusted.Close(); err != nil {
					t.Fatal(err)
				}
			} else if trusted != nil || !errors.Is(err, scan.ErrInvalidMetadata) {
				t.Fatalf("trusted=%v err=%v", trusted, err)
			}
			closes := 0
			for fd := 1; fd <= record.openCount; fd++ {
				closes += record.closes[fd]
			}
			if record.openCount != tt.wantOpen || closes != tt.wantCloses {
				t.Fatalf("opens=%d closes=%v", record.openCount, record.closes)
			}
		})
	}
}

func TestWalkerAcquisitionHandleClose(t *testing.T) {
	walker, record := scriptedWalker()
	trusted, err := walker.AcquireRoot(context.Background(), absolute(t, "/home"), absolute(t, "/home/cache"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if err != nil {
		t.Fatal(err)
	}
	record.closeErr[3] = syscall.EIO
	first, second := trusted.Close(), trusted.Close()
	var typed *scan.PathError
	if !errors.Is(first, scan.ErrInaccessible) || first != second || !errors.As(first, &typed) || typed.Operation() != "close" || record.closes[3] != 1 {
		t.Fatalf("first=%v second=%v closes=%v", first, second, record.closes)
	}
}

func TestWalkerAcquisitionLifecycle(t *testing.T) {
	for _, phase := range []string{"open", "fstat", "fstatfs", "validation", "handoff"} {
		t.Run(phase, func(t *testing.T) {
			walker, record := scriptedWalker()
			record.failAt = phase
			if phase == "validation" {
				record.failAt = ""
				record.stats[2] = syscall.Stat_t{Dev: 7, Ino: 0, Nlink: 1, Mode: syscall.S_IFDIR}
			}
			if phase == "handoff" {
				record.failAt = ""
				record.closeErr[2] = syscall.EIO
			}
			trusted, err := walker.AcquireRoot(context.Background(), absolute(t, "/home"), absolute(t, "/home/cache"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
			if trusted != nil || err == nil {
				t.Fatalf("trusted=%v err=%v", trusted, err)
			}
			for fd := 1; fd <= len(record.names); fd++ {
				if record.closes[fd] != 1 {
					t.Fatalf("fd %d closes=%d", fd, record.closes[fd])
				}
			}
		})
	}
}

func TestWalkerAcquisitionCloseCounts(t *testing.T) {
	walker, record := scriptedWalker()
	trusted, err := walker.AcquireRoot(context.Background(), absolute(t, "/home"), absolute(t, "/home/cache"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if err != nil {
		t.Fatal(err)
	}
	if err := trusted.Close(); err != nil {
		t.Fatal(err)
	}
	for fd := 1; fd <= 3; fd++ {
		if record.closes[fd] != 1 {
			t.Fatalf("fd %d closes=%d", fd, record.closes[fd])
		}
	}
}

func TestWalkerAcquisitionErrorPrecedence(t *testing.T) {
	walker, record := scriptedWalker()
	record.fstatErr[2], record.closeErr[1], record.closeErr[2] = syscall.EIO, syscall.EIO, syscall.EIO
	_, err := walker.AcquireRoot(context.Background(), absolute(t, "/home"), absolute(t, "/home/cache"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	var typed *scan.PathError
	if !errors.As(err, &typed) || typed.Operation() != "fstat" || typed.Path() != "/home" || !errors.Is(err, scan.ErrInaccessible) {
		t.Fatalf("err=%v", err)
	}
}

func TestWalkerAcquisitionLifecycleComponentMatrix(t *testing.T) {
	parts := []string{"/", "home", "alice", "cache", "final"}
	stages := []string{"open", "fstat", "fstatfs", "validation", "handoff"}
	cases := 0
	for position, part := range parts {
		for _, stage := range stages {
			if stage == "handoff" && position == len(parts)-1 {
				continue // The final descriptor transfers to TrustedRoot; its close is tested below.
			}
			cases++
			t.Run(stage+"/"+part, func(t *testing.T) {
				walker, record := scriptedWalker()
				configureLifecyclePrimary(record, stage, position)
				trusted, err := walker.AcquireRoot(context.Background(), matrixHome(t), matrixRoot(t), matrixLimits())
				assertLifecycleFailure(t, record, trusted, err, stage, position, parts)
			})
		}
	}
	if cases != 24 {
		t.Fatalf("case cardinality = %d, want 24", cases)
	}
}

func TestWalkerAcquisitionPrimaryCleanupCartesian(t *testing.T) {
	parts := []string{"/", "home", "alice", "cache", "final"}
	stages := []string{"fstat", "fstatfs", "validation"}
	cases := 0
	for position := 1; position < len(parts); position++ {
		for _, stage := range stages {
			cases++
			t.Run(stage+"/current-close/"+parts[position], func(t *testing.T) {
				walker, record := scriptedWalker()
				configureLifecyclePrimary(record, stage, position)
				currentFD, nextFD := position, position+1
				record.closeErr[currentFD] = syscall.EIO

				trusted, err := walker.AcquireRoot(context.Background(), matrixHome(t), matrixRoot(t), matrixLimits())
				assertLifecycleFailure(t, record, trusted, err, stage, position, parts)

				class := scan.ErrInaccessible
				if stage == "validation" {
					class = scan.ErrInvalidMetadata
				}
				primaryOperation, primaryPath := stage, matrixPath(parts[:position+1])
				if stage == "validation" {
					primaryOperation = "fstat"
				}
				var primary *scan.PathError
				if !errors.Is(err, class) || !errors.As(err, &primary) || primary.Operation() != primaryOperation || primary.Path() != primaryPath {
					t.Fatalf("primary=%v, want %s %s class %v", err, primaryOperation, primaryPath, class)
				}

				currentPath := matrixPath(parts[:position])
				assertSecondaryCloseEvidence(t, err, []string{currentPath})
				if record.closes[nextFD] != 1 || record.closes[currentFD] != 1 {
					t.Fatalf("case %s/%s next fd %d closes=%d current fd %d closes=%d", stage, parts[position], nextFD, record.closes[nextFD], currentFD, record.closes[currentFD])
				}
				if trusted != nil || len(record.names) != position+1 {
					t.Fatalf("case %s/%s trusted=%v opens=%v", stage, parts[position], trusted, record.names)
				}
			})
		}
	}
	if cases != 12 {
		t.Fatalf("case cardinality = %d, want 12", cases)
	}
}

func TestWalkerAcquisitionClosePrecedenceMatrix(t *testing.T) {
	parts := []string{"/", "home", "alice", "cache", "final"}
	stages := []string{"open", "fstat", "fstatfs", "validation", "handoff"}
	cases := 0
	for position := range parts {
		for _, stage := range stages {
			if stage == "handoff" && position == len(parts)-1 {
				continue
			}
			cases++
			t.Run("secondary/"+stage+"/"+parts[position], func(t *testing.T) {
				walker, record := scriptedWalker()
				configureLifecyclePrimary(record, stage, position)
				for _, fd := range lifecycleOwnedFDs(stage, position) {
					record.closeErr[fd] = syscall.EIO
				}
				trusted, err := walker.AcquireRoot(context.Background(), matrixHome(t), matrixRoot(t), matrixLimits())
				assertLifecycleFailure(t, record, trusted, err, stage, position, parts)
				assertSecondaryCloseEvidence(t, err, lifecycleOwnedPaths(stage, position, parts))
			})
		}
	}
	for position := 1; position < len(parts); position++ {
		cases++
		t.Run("multiple-secondary/cancel-after-open/"+parts[position], func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			walker, record := scriptedWalker()
			record.checkpoint = func(point string) {
				if point == "after-open" && record.openCount == position+1 {
					cancel()
				}
			}
			record.closeErr[position], record.closeErr[position+1] = syscall.EIO, syscall.EIO
			trusted, err := walker.AcquireRoot(ctx, matrixHome(t), matrixRoot(t), matrixLimits())
			assertLifecycleFailure(t, record, trusted, err, "acquire", position, parts)
			assertSecondaryCloseEvidence(t, err, []string{matrixPath(parts[:position]), matrixPath(parts[:position+1])})
		})
	}
	for _, closeFails := range []bool{false, true} {
		cases++
		t.Run("transferred-close/"+map[bool]string{false: "success", true: "failure"}[closeFails], func(t *testing.T) {
			walker, record := scriptedWalker()
			trusted, err := walker.AcquireRoot(context.Background(), matrixHome(t), matrixRoot(t), matrixLimits())
			if err != nil {
				t.Fatal(err)
			}
			if closeFails {
				record.closeErr[len(parts)] = syscall.EIO
			}
			first, second := trusted.Close(), trusted.Close()
			if first != second || record.closes[len(parts)] != 1 {
				t.Fatalf("close results=%v/%v closes=%v", first, second, record.closes)
			}
			if want := expectedSuccessTrace(parts); !reflect.DeepEqual(record.trace, want) {
				t.Fatalf("trace=%v, want %v", record.trace, want)
			}
			if closeFails {
				assertSecondaryCloseEvidence(t, first, []string{matrixPath(parts)})
			} else if first != nil {
				t.Fatalf("close = %v, want nil", first)
			}
		})
	}
	if cases != 30 {
		t.Fatalf("case cardinality = %d, want 30", cases)
	}
}

// TestWalkerAcquisitionCheckpointDescriptorLimitMatrix proves operation order rather
// than merely counting rows: an unreachable cancellation cannot mask the one-capability
// limit, while a reached cancellation wins before the limit check.
func TestWalkerAcquisitionCheckpointDescriptorLimitMatrix(t *testing.T) {
	parts := []string{"/", "home", "alice", "cache", "final"}
	checkpoints := []struct {
		name     string
		position []int
	}{
		{"before", []int{0}},
		{"after-open", []int{0, 1, 2, 3, 4}},
		{"after-fstat", []int{0, 1, 2, 3, 4}},
		{"after-fstatfs", []int{0, 1, 2, 3, 4}},
		{"after-validation", []int{0, 1, 2, 3, 4}},
		{"before-handoff", []int{4}},
	}
	cases := 0
	for _, checkpoint := range checkpoints {
		for _, position := range checkpoint.position {
			for _, descriptors := range []int{2, 1} { // MaxDepth is intentionally uncharged here.
				cases++
				t.Run(checkpoint.name+"/"+parts[position]+"/descriptors-"+string(rune('0'+descriptors)), func(t *testing.T) {
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()
					walker, record := scriptedWalker()
					reachable := descriptors == 2 || (position == 0 && checkpoint.name != "before-handoff")
					reached := false
					record.checkpoint = func(point string) {
						if point == checkpoint.name && record.openCount == position+checkpointOpenOffset(checkpoint.name) {
							reached = true
							cancel()
						}
					}
					trusted, err := walker.AcquireRoot(ctx, matrixHome(t), matrixRoot(t), scan.WalkLimits{MaxDepth: 64, MaxDescriptors: descriptors})
					if reached != reachable {
						t.Fatalf("checkpoint reached=%t, want %t; trace=%v", reached, reachable, record.trace)
					}
					if reachable {
						assertCheckpointCancellation(t, record, err, trusted, checkpoint.name, position, parts)
					} else if trusted != nil || !errors.Is(err, scan.ErrInvalidMetadata) {
						t.Fatalf("unreached checkpoint must lose to descriptor limit: trusted=%v err=%v", trusted, err)
					} else {
						var typed *scan.PathError
						wantTrace := []string{"open /", "fstat 1", "fstatfs 1", "close 1"}
						if !errors.As(err, &typed) || typed.Operation() != "acquire" || typed.Path() != "/" || record.openCount != 1 || !reflect.DeepEqual(record.trace, wantTrace) {
							t.Fatalf("limit error=%v opens=%d trace=%v, want %v", err, record.openCount, record.trace, wantTrace)
						}
					}
					for fd := 1; fd <= record.openCount; fd++ {
						if record.closes[fd] != 1 {
							t.Fatalf("fd %d close count=%d", fd, record.closes[fd])
						}
					}
				})
			}
		}
	}
	if cases != 44 {
		t.Fatalf("case cardinality = %d, want 44", cases)
	}
}

func checkpointOpenOffset(checkpoint string) int {
	if checkpoint == "before" {
		return 0
	}
	return 1
}

func assertCheckpointCancellation(t *testing.T, record *walkerRecord, err error, trusted scan.TrustedRoot, checkpoint string, position int, parts []string) {
	t.Helper()
	var typed *scan.PathError
	path := matrixPath(parts[:position+1])
	if checkpoint == "before" || checkpoint == "before-handoff" {
		path = matrixPath(parts)
	}
	wantTrace, wantOpens := expectedCheckpointCancellationTrace(checkpoint, position, parts), position+1
	if checkpoint == "before" {
		wantOpens = 0
	}
	if trusted != nil || !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &typed) || typed.Operation() != "acquire" || typed.Path() != path || record.openCount != wantOpens || !reflect.DeepEqual(record.trace, wantTrace) {
		t.Fatalf("cancellation: trusted=%v err=%v opens=%d trace=%v, want acquire %s opens=%d trace=%v", trusted, err, record.openCount, record.trace, path, wantOpens, wantTrace)
	}
	if position+1 < len(parts) {
		deeper, nextFD := "open "+parts[position+1], "fstat "+string(rune('0'+position+2))
		if containsTrace(record.trace, deeper) || containsTrace(record.trace, nextFD) {
			t.Fatalf("cancellation reached %s at %s but deeper operation/open remains: %v", checkpoint, path, record.trace)
		}
	}
}

func expectedCheckpointCancellationTrace(checkpoint string, position int, parts []string) []string {
	if checkpoint == "before" {
		return nil
	}
	if checkpoint == "before-handoff" {
		return expectedSuccessTrace(parts)
	}
	trace, fd := expectedLifecycleTrace("acquire", position, parts), string(rune('0'+position+1))
	closeAt := len(trace) - 2
	if position == 0 {
		trace = trace[:len(trace)-1]
	}
	if checkpoint != "after-open" {
		trace = append(trace[:closeAt], append([]string{"fstat " + fd}, trace[closeAt:]...)...)
	}
	if checkpoint == "after-fstatfs" || checkpoint == "after-validation" {
		trace = append(trace[:closeAt+1], append([]string{"fstatfs " + fd}, trace[closeAt+1:]...)...)
	}
	return trace
}
func containsTrace(trace []string, want string) bool {
	for _, got := range trace {
		if got == want {
			return true
		}
	}
	return false
}

// TestWalkerAcquisitionRemainingLifecycleMatrix covers its two case-specific rows;
// lifecycleInventory below exhaustively maps the remaining acquisition dimensions.
func TestWalkerAcquisitionRemainingLifecycleMatrix(t *testing.T) {
	assertLifecycleInventory(t)
	for _, tt := range []struct {
		name, stage, primaryPath, secondaryPath string
		position                                int
	}{
		{"open/previous-close", "open", "/home/alice", "/home", 2},
		{"handoff/current-close", "handoff", "/home/alice/cache", "/home/alice/cache/final", 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			configureLifecyclePrimary(record, tt.stage, tt.position)
			for _, fd := range lifecycleOwnedFDs(tt.stage, tt.position) {
				record.closeErr[fd] = syscall.EIO
			}
			trusted, err := walker.AcquireRoot(context.Background(), matrixHome(t), matrixRoot(t), matrixLimits())
			if trusted != nil {
				t.Fatal("failure path transferred root ownership")
			}
			var primary *scan.PathError
			if !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &primary) || primary.Operation() != map[string]string{"open": "openat", "handoff": "close"}[tt.stage] || primary.Path() != tt.primaryPath {
				t.Fatalf("primary=%v, want %s %s", err, tt.stage, tt.primaryPath)
			}
			assertSecondaryCloseEvidence(t, err, []string{tt.secondaryPath})
			for fd := 1; fd <= len(record.names); fd++ {
				if record.closes[fd] != 1 {
					t.Fatalf("fd %d close count=%d trace=%v", fd, record.closes[fd], record.trace)
				}
			}
		})
	}
}

func assertLifecycleInventory(t *testing.T) {
	t.Helper()
	inventory := []struct{ key, evidence string }{{"success", "WU5C1B1A/TestWalkerAcquisitionHandoffCloseFailure"}, {"open", "C2/TestWalkerAcquisitionRemainingLifecycleMatrix/open/previous-close"}, {"fstat", "WU5C1B2B1A/TestWalkerAcquisitionPrimaryCleanupCartesian (excluded 4 current-close rows)"}, {"fstatfs", "WU5C1B2B1A/TestWalkerAcquisitionPrimaryCleanupCartesian (excluded 4 current-close rows)"}, {"validation", "WU5C1B2B1A/TestWalkerAcquisitionPrimaryCleanupCartesian (excluded 4 current-close rows)"}, {"handoff/close-current", "C2/TestWalkerAcquisitionRemainingLifecycleMatrix/handoff/current-close"}, {"cleanup-close-next/current", "WU5C1B2A/TestWalkerAcquisitionClosePrecedenceMatrix"}, {"multiple-cleanup", "WU5C1C1A/TestWalkerAcquisitionClosePrecedenceMatrix/multiple-secondary"}, {"transferred-close-success", "WU5C1B1A/TestWalkerAcquisitionClosePrecedenceMatrix/transferred-close/success"}, {"transferred-close-failure", "WU5C1B1A/TestWalkerAcquisitionClosePrecedenceMatrix/transferred-close/failure"}, {"transferred-close-cache", "WU5C1B2A/TestWalkerAcquisitionHandleClose"}}
	want := map[string]bool{"success": true, "open": true, "fstat": true, "fstatfs": true, "validation": true, "handoff/close-current": true, "cleanup-close-next/current": true, "multiple-cleanup": true, "transferred-close-success": true, "transferred-close-failure": true, "transferred-close-cache": true}
	seen := map[string]string{}
	for _, item := range inventory {
		if !want[item.key] || item.evidence == "" {
			t.Fatalf("unknown or unproven lifecycle key %q", item.key)
		}
		if previous, duplicate := seen[item.key]; duplicate {
			t.Fatalf("duplicate lifecycle key %q: %s and %s", item.key, previous, item.evidence)
		}
		seen[item.key] = item.evidence
	}
	for key := range want {
		if _, present := seen[key]; !present {
			t.Fatalf("missing lifecycle key %q", key)
		}
	}
}

func matrixHome(t *testing.T) scan.AbsoluteComponents { return absolute(t, "/home/alice") }
func matrixRoot(t *testing.T) scan.AbsoluteComponents { return absolute(t, "/home/alice/cache/final") }
func matrixLimits() scan.WalkLimits                   { return scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2} }
func matrixPath(parts []string) string {
	if len(parts) == 1 {
		return "/"
	}
	return "/" + joinComponents(parts[1:])
}
func configureLifecyclePrimary(record *walkerRecord, stage string, position int) {
	fd := position + 1
	switch stage {
	case "open":
		record.openErr[fd] = syscall.EIO
	case "fstat":
		record.fstatErr[fd] = syscall.EIO
	case "fstatfs":
		record.fstatfsErr[fd] = syscall.EIO
	case "validation":
		record.stats[fd] = syscall.Stat_t{Dev: 7, Nlink: 1, Mode: syscall.S_IFDIR}
	case "handoff":
		record.closeErr[fd] = syscall.EIO
	}
}
func lifecycleOwnedFDs(stage string, position int) []int {
	if stage == "open" {
		if position == 0 {
			return nil
		}
		return []int{position}
	}
	if stage == "handoff" {
		return []int{position + 2}
	}
	return []int{position + 1}
}
func lifecycleOwnedPaths(stage string, position int, parts []string) []string {
	fds := lifecycleOwnedFDs(stage, position)
	paths := make([]string, len(fds))
	for i, fd := range fds {
		paths[i] = matrixPath(parts[:fd])
	}
	return paths
}
func assertLifecycleFailure(t *testing.T, record *walkerRecord, trusted scan.TrustedRoot, err error, stage string, position int, parts []string) {
	t.Helper()
	class := scan.ErrInaccessible
	if stage == "validation" {
		class = scan.ErrInvalidMetadata
	}
	if trusted != nil || err == nil || !errors.Is(err, class) {
		t.Fatalf("trusted=%v err=%v", trusted, err)
	}
	operation, path := stage, matrixPath(parts[:position+1])
	if stage == "open" {
		operation = "openat"
		if position == 0 {
			operation = "open"
		}
	} else if stage == "validation" {
		operation = "fstat"
	} else if stage == "handoff" {
		operation, path = "close", matrixPath(parts[:position+1])
	} else if stage == "acquire" {
		operation, path = "acquire", matrixPath(parts[:position+1])
	}
	var typed *scan.PathError
	if !errors.As(err, &typed) || typed.Operation() != operation || typed.Path() != path {
		t.Fatalf("primary=%v, want %s %s", err, operation, path)
	}
	for fd := 1; fd <= len(record.names); fd++ {
		if record.closes[fd] != 1 {
			t.Fatalf("fd %d close count=%d trace=%v", fd, record.closes[fd], record.trace)
		}
	}
	if want := expectedLifecycleTrace(stage, position, parts); !reflect.DeepEqual(record.trace, want) {
		t.Fatalf("trace=%v, want %v", record.trace, want)
	}
	if len(record.names) > position+1 && stage != "handoff" {
		t.Fatalf("deeper open after %s at %s: %v", stage, path, record.names)
	}
}
func expectedSuccessTrace(parts []string) []string {
	trace := make([]string, 0, 4*len(parts))
	for index, part := range parts {
		fd := string(rune('0' + index + 1))
		trace = append(trace, "open "+part, "fstat "+fd, "fstatfs "+fd)
		if index > 0 {
			trace = append(trace, "close "+string(rune('0'+index)))
		}
	}
	return append(trace, "close "+string(rune('0'+len(parts))))
}

func expectedLifecycleTrace(stage string, position int, parts []string) []string {
	trace := make([]string, 0, 4*len(parts))
	open := func(index int) { trace = append(trace, "open "+parts[index]) }
	facts := func(fd int, statfs bool) {
		trace = append(trace, "fstat "+string(rune('0'+fd)))
		if statfs {
			trace = append(trace, "fstatfs "+string(rune('0'+fd)))
		}
	}
	prefix := position
	if stage == "handoff" {
		prefix++
	}
	for index := 0; index < prefix; index++ {
		open(index)
		facts(index+1, true)
		if index > 0 {
			trace = append(trace, "close "+string(rune('0'+index)))
		}
	}
	if stage == "acquire" {
		open(position)
		return append(trace, "close "+string(rune('0'+position+1)), "close "+string(rune('0'+position)))
	}
	if stage == "handoff" {
		open(position + 1)
		facts(position+2, true)
		return append(trace, "close "+string(rune('0'+position+1)), "close "+string(rune('0'+position+2)))
	}
	open(position)
	if stage == "open" {
		if position > 0 {
			trace = append(trace, "close "+string(rune('0'+position)))
		}
		return trace
	}
	facts(position+1, stage != "fstat" && stage != "validation")
	trace = append(trace, "close "+string(rune('0'+position+1)))
	if position > 0 {
		trace = append(trace, "close "+string(rune('0'+position)))
	}
	return trace
}

func assertSecondaryCloseEvidence(t *testing.T, err error, paths []string) {
	t.Helper()
	found := map[string]bool{}
	var visit func(error)
	visit = func(current error) {
		if current == nil {
			return
		}
		if typedClose, ok := current.(*scan.PathError); ok && typedClose.Operation() == "close" {
			if !errors.Is(typedClose, scan.ErrInaccessible) {
				t.Errorf("close evidence for %q lacks ErrInaccessible: %v", typedClose.Path(), typedClose)
			}
			found[typedClose.Path()] = true
		}
		if joined, ok := current.(interface{ Unwrap() []error }); ok {
			for _, next := range joined.Unwrap() {
				visit(next)
			}
			return
		}
		visit(errors.Unwrap(current))
	}
	visit(err)
	for _, path := range paths {
		if !found[path] {
			t.Fatalf("missing close evidence for %q in %v", path, err)
		}
	}
}

type walkerRecord struct {
	names                                   []string
	parents                                 []int
	directory                               []byte
	directories                             map[int][]byte
	dirFD                                   int
	flags                                   []int
	openCount, fstats, statfs, readDirs     int
	stats                                   map[int]syscall.Stat_t
	nonLocal                                map[int]bool
	openErr, fstatErr, fstatfsErr, closeErr map[int]error
	closes                                  map[int]int
	trace                                   []string
	consumedDevices                         []uint64
	cancel                                  context.CancelFunc
	checkpoint                              func(string)
	reached, cancelAt, failAt               string
}

func scriptedWalker() (trustedRootWalker, *walkerRecord) {
	r := &walkerRecord{stats: map[int]syscall.Stat_t{}, nonLocal: map[int]bool{}, openErr: map[int]error{}, fstatErr: map[int]error{}, fstatfsErr: map[int]error{}, closeErr: map[int]error{}, closes: map[int]int{}}
	next := 0
	ops := walkerOps{
		open: func(parent int, name string, flags int, _ uint32) (int, error) {
			r.openCount++
			r.parents = append(r.parents, parent)
			r.trace = append(r.trace, "open "+name)
			if r.cancelAt == "open" {
				r.cancel()
			}
			if r.failAt == "open" {
				return -1, syscall.EIO
			}
			if err := r.openErr[r.openCount]; err != nil {
				return -1, err
			}
			next++
			r.names = append(r.names, name)
			r.flags = append(r.flags, flags)
			return next, nil
		},
		fstat: func(fd int, out *syscall.Stat_t) error {
			r.fstats++
			r.trace = append(r.trace, "fstat "+string(rune('0'+fd)))
			if r.cancelAt == "fstat" {
				r.cancel()
			}
			if r.failAt == "fstat" {
				return syscall.EIO
			}
			if err := r.fstatErr[fd]; err != nil {
				return err
			}
			*out = r.stats[fd]
			if out.Dev != 0 {
				r.consumedDevices = append(r.consumedDevices, uint64(out.Dev))
			}
			if out.Dev == 0 {
				*out = statFor(7, uint64(fd))
			}
			return nil
		},
		fstatfs: func(fd int, out *syscall.Statfs_t) error {
			r.statfs++
			r.trace = append(r.trace, "fstatfs "+string(rune('0'+fd)))
			if r.cancelAt == "fstatfs" {
				r.cancel()
			}
			if r.failAt == "fstatfs" {
				return syscall.EIO
			}
			if err := r.fstatfsErr[fd]; err != nil {
				return err
			}
			if !r.nonLocal[fd] {
				out.Flags = localFilesystemFlag
			}
			return nil
		},
		readDir: func(fd int) ([]byte, error) {
			r.dirFD = fd
			r.readDirs++
			if r.directories != nil {
				return r.directories[fd], nil
			}
			return r.directory, nil
		},
		close: func(fd int) error {
			r.closes[fd]++
			r.trace = append(r.trace, "close "+string(rune('0'+fd)))
			return r.closeErr[fd]
		},
		checkpoint: func(point string) {
			if r.checkpoint != nil {
				r.checkpoint(point)
			}
		},
	}
	return newTrustedRootWalker(ops), r
}
func statFor(dev, ino uint64) syscall.Stat_t {
	return syscall.Stat_t{Dev: int32(dev), Ino: ino, Nlink: 1, Mode: syscall.S_IFDIR}
}
func absolute(t *testing.T, path string) scan.AbsoluteComponents {
	t.Helper()
	got, err := scan.NewAbsoluteComponents(path)
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func TestWalkerAcquisitionRealAncestorReplacement(t *testing.T) {
	fixture := newRealWalkerFixture(t)
	originalRoot := realIdentity(t, fixture.root)
	target := filepath.Join(fixture.temp, "ancestor-poison")
	mustMkdir(t, target)
	targetIdentity := realIdentity(t, target)
	fixture.afterOpen = fixture.home
	fixture.mutate = func() {
		moved := filepath.Join(fixture.temp, "home-retained")
		if err := os.Rename(fixture.home, moved); err != nil {
			t.Fatal(err)
		}
		fixture.renamedRoot = filepath.Join(moved, "root")
		if err := os.Symlink(target, fixture.home); err != nil {
			t.Fatal(err)
		}
	}
	trusted, err := fixture.acquire()
	if err != nil {
		t.Fatal(err)
	}
	if got := trusted.Facts().Identity; got != originalRoot {
		t.Fatalf("retained root identity = %#v, want %#v", got, originalRoot)
	}
	if fixture.saw(targetIdentity) {
		t.Fatalf("replacement target facts were accepted: %#v", targetIdentity)
	}
	assertRetainedIdentity(t, trusted, fixture.renamedRoot, target)
	fixture.assertSafeBound(t, trusted)
}

func TestWalkerAcquisitionRealFinalReplacement(t *testing.T) {
	fixture := newRealWalkerFixture(t)
	originalRoot := realIdentity(t, fixture.root)
	target := filepath.Join(fixture.temp, "final-poison")
	mustMkdir(t, target)
	targetIdentity := realIdentity(t, target)
	fixture.afterOpen = fixture.root
	fixture.mutate = func() {
		moved := filepath.Join(fixture.temp, "root-retained")
		if err := os.Rename(fixture.root, moved); err != nil {
			t.Fatal(err)
		}
		fixture.renamedRoot = moved
		if err := os.Symlink(target, fixture.root); err != nil {
			t.Fatal(err)
		}
	}
	trusted, err := fixture.acquire()
	if err != nil {
		t.Fatal(err)
	}
	if got := trusted.Facts().Identity; got != originalRoot {
		t.Fatalf("final descriptor identity = %#v, want %#v", got, originalRoot)
	}
	if fixture.saw(targetIdentity) {
		t.Fatalf("replacement target facts were accepted: %#v", targetIdentity)
	}
	assertRetainedIdentity(t, trusted, fixture.renamedRoot, target)
	fixture.assertSafeBound(t, trusted)
}

func TestWalkerAcquisitionRealNoFollowBinding(t *testing.T) {
	for _, replace := range []string{"ancestor", "final"} {
		t.Run(replace+" symlink", func(t *testing.T) {
			fixture := newRealWalkerFixture(t)
			originalRoot := realIdentity(t, fixture.root)
			target := filepath.Join(fixture.temp, "symlink-target")
			mustMkdir(t, target)
			targetIdentity := realIdentity(t, target)
			if replace == "ancestor" {
				fixture.afterOpen = fixture.home
			} else {
				fixture.afterOpen = fixture.root
			}
			fixture.mutate = func() {
				old := fixture.home
				moved := filepath.Join(fixture.temp, "retained-home")
				fixture.renamedRoot = filepath.Join(moved, "root")
				if replace == "final" {
					old, moved = fixture.root, filepath.Join(fixture.temp, "retained-root")
					fixture.renamedRoot = moved
				}
				if err := os.Rename(old, moved); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(target, old); err != nil {
					t.Fatal(err)
				}
			}
			trusted, err := fixture.acquire()
			if err != nil {
				t.Fatal(err)
			}
			if got := trusted.Facts().Identity; got != originalRoot {
				t.Fatalf("root identity = %#v, want %#v", got, originalRoot)
			}
			if fixture.saw(targetIdentity) {
				t.Fatalf("symlink target facts were accepted: %#v", targetIdentity)
			}
			assertRetainedIdentity(t, trusted, fixture.renamedRoot, target)
			fixture.assertSafeBound(t, trusted)
		})
	}
}

func TestWalkerAcquisitionMountPoisonScript(t *testing.T) {
	walker, record := scriptedWalker()
	devices := []uint64{7, 9, 7}
	if !reflect.DeepEqual(devices, []uint64{7, 9, 7}) {
		t.Fatalf("mount poison script = %v", devices)
	}
	for index, device := range devices {
		record.stats[index+3] = statFor(device, uint64(index+30))
	}
	_, err := walker.AcquireRoot(context.Background(), absolute(t, "/home/alice"), absolute(t, "/home/alice/poison/child"), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if !errors.Is(err, scan.ErrDeviceBoundary) {
		t.Fatalf("AcquireRoot() = %v, want device boundary", err)
	}
	if !reflect.DeepEqual(record.names, []string{"/", "home", "alice", "poison"}) {
		t.Fatalf("acquisition opened %v, want 7→9 prefix only", record.names)
	}
	if record.fstats != 4 || record.statfs != 4 || record.openCount != 4 {
		t.Fatalf("poison child/final facts consumed: opens=%d fstat=%d fstatfs=%d", record.openCount, record.fstats, record.statfs)
	}
	if !reflect.DeepEqual(record.consumedDevices, []uint64{7, 9}) {
		t.Fatalf("consumed devices = %v, want [7 9]", record.consumedDevices)
	}
	if len(devices) != 3 || devices[2] != 7 {
		t.Fatalf("final declared device slot = %v, want unconsumed 7", devices)
	}
	t.Logf("mount devices declared=%v consumed=%v trace=%v", devices, record.consumedDevices, record.trace)
}

type realWalkerFixture struct {
	temp, home, root, afterOpen string
	renamedRoot                 string
	mutate                      func()
	paths                       map[int]string
	opens                       []int
	records                     []*realOpen
	facts                       []scan.Identity
}

type realOpen struct {
	fd     int
	path   string
	closes int
}

func newRealWalkerFixture(t *testing.T) *realWalkerFixture {
	t.Helper()
	temp, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	fixture := &realWalkerFixture{temp: temp, paths: map[int]string{}}
	fixture.home, fixture.root = filepath.Join(fixture.temp, "home"), filepath.Join(fixture.temp, "home", "root")
	mustMkdir(t, fixture.root)
	return fixture
}
func (f *realWalkerFixture) acquire() (scan.TrustedRoot, error) {
	walker := newTrustedRootWalker(walkerOps{
		open: func(parent int, name string, flags int, mode uint32) (int, error) {
			fd, err := walkerOpen(parent, name, flags, mode)
			if err != nil {
				return fd, err
			}
			path := "/"
			if name != "/" {
				path = filepath.Join(f.paths[parent], name)
			}
			f.paths[fd], f.opens = path, append(f.opens, flags)
			f.records = append(f.records, &realOpen{fd: fd, path: path})
			if path == f.afterOpen && f.mutate != nil {
				f.mutate()
				f.mutate = nil
			}
			return fd, nil
		},
		fstat: func(fd int, stat *syscall.Stat_t) error {
			err := syscall.Fstat(fd, stat)
			if err == nil {
				f.facts = append(f.facts, scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)})
			}
			return err
		},
		fstatfs: syscall.Fstatfs,
		close: func(fd int) error {
			for index := len(f.records) - 1; index >= 0; index-- {
				if f.records[index].fd == fd && f.records[index].closes == 0 {
					f.records[index].closes++
					break
				}
			}
			return syscall.Close(fd)
		},
	})
	return walker.AcquireRoot(context.Background(), absoluteForReal(f.home), absoluteForReal(f.root), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
}
func assertRetainedIdentity(t *testing.T, trusted scan.TrustedRoot, renamedOriginal, replacement string) {
	t.Helper()
	renamedIdentity := statIdentity(t, renamedOriginal)
	replacementIdentity := statIdentity(t, replacement)
	if renamedIdentity == replacementIdentity {
		t.Fatalf("renamed original identity = replacement identity = %#v", renamedIdentity)
	}
	if got := trusted.Facts().Identity; got != renamedIdentity {
		t.Fatalf("TrustedRoot identity = %#v, want renamed original %#v", got, renamedIdentity)
	}
	if got := trusted.Facts().Identity; got == replacementIdentity {
		t.Fatalf("TrustedRoot accepted replacement identity %#v", replacementIdentity)
	}
	t.Logf("TrustedRoot=%#v renamed-original=%#v replacement=%#v", trusted.Facts().Identity, renamedIdentity, replacementIdentity)
}

func (f *realWalkerFixture) saw(identity scan.Identity) bool {
	for _, got := range f.facts {
		if got == identity {
			return true
		}
	}
	return false
}
func (f *realWalkerFixture) assertSafeBound(t *testing.T, trusted scan.TrustedRoot) {
	t.Helper()
	if err := trusted.Close(); err != nil {
		t.Fatal(err)
	}
	for _, flags := range f.opens {
		if flags != walkerFlags {
			t.Fatalf("open flags = %#x, want %#x", flags, walkerFlags)
		}
	}
	for _, record := range f.records {
		if record.closes != 1 {
			t.Fatalf("fd %d (%s) closes=%d, want 1", record.fd, record.path, record.closes)
		}
	}
}
func mustMkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
}
func realIdentity(t *testing.T, path string) scan.Identity {
	t.Helper()
	var stat syscall.Stat_t
	if err := syscall.Lstat(path, &stat); err != nil {
		t.Fatal(err)
	}
	return scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
}
func statIdentity(t *testing.T, path string) scan.Identity {
	t.Helper()
	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		t.Fatal(err)
	}
	return scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
}

func absoluteForReal(path string) scan.AbsoluteComponents {
	components, err := scan.NewAbsoluteComponents(path)
	if err != nil {
		panic(err)
	}
	return components
}

func TestWalkerDFSDepth(t *testing.T) {
	for _, tt := range []struct {
		name      string
		maxDepth  int
		wantOpens []string
		wantLimit bool
	}{
		{"root only", 1, []string{"one"}, true},
		{"exact maximum", 2, []string{"one", "two"}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := dfsWalker(map[int][]byte{41: dirRecord(1, directoryType, "one"), 1: dirRecord(2, directoryType, "two"), 2: nil})
			record.stats[1], record.stats[2] = statFor(7, 1), statFor(7, 2)
			err := walker.walkDirectory(context.Background(), 41, testWalkerFacts(t), scan.WalkLimits{MaxDepth: tt.maxDepth, MaxDescriptors: 3})
			if !reflect.DeepEqual(record.names, tt.wantOpens) || errors.Is(err, scan.ErrLimit) != tt.wantLimit {
				t.Fatalf("opens=%v err=%v", record.names, err)
			}
			assertEachClosedOnce(t, record)
		})
	}
}

func TestWalkerDescriptorCapacity(t *testing.T) {
	for _, tt := range []struct {
		name      string
		limit     int
		wantOpens []string
	}{
		{"root consumes one capability", 1, nil},
		{"exact child boundary", 2, []string{"one"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := dfsWalker(map[int][]byte{41: dirRecord(1, directoryType, "one"), 1: dirRecord(2, directoryType, "two")})
			record.stats[1], record.stats[2] = statFor(7, 1), statFor(7, 2)
			err := walker.walkDirectory(context.Background(), 41, testWalkerFacts(t), scan.WalkLimits{MaxDepth: 2, MaxDescriptors: tt.limit})
			if !errors.Is(err, scan.ErrLimit) || !reflect.DeepEqual(record.names, tt.wantOpens) {
				t.Fatalf("opens=%v err=%v", record.names, err)
			}
			assertEachClosedOnce(t, record)
		})
	}
}

func TestWalkerDirectoryLifecycle(t *testing.T) {
	walker, record := dfsWalker(map[int][]byte{41: dirRecord(1, directoryType, "one")})
	record.fstatErr[1], record.closeErr[1] = syscall.EIO, syscall.EIO
	err := walker.walkDirectory(context.Background(), 41, testWalkerFacts(t), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	var primary *scan.PathError
	if !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &primary) || primary.Operation() != "fstat" || primary.Path() != "one" {
		t.Fatalf("primary=%v", err)
	}
	assertSecondaryCloseEvidence(t, err, []string{"one"})
	assertEachClosedOnce(t, record)

	walker, record = dfsWalker(map[int][]byte{41: dirRecord(1, directoryType, "one"), 1: dirRecord(2, directoryType, "two")})
	record.stats[1], record.fstatErr[2], record.closeErr[1], record.closeErr[2] = statFor(7, 1), syscall.EIO, syscall.EIO, syscall.EIO
	err = walker.walkDirectory(context.Background(), 41, testWalkerFacts(t), scan.WalkLimits{MaxDepth: 2, MaxDescriptors: 3})
	if !errors.As(err, &primary) || primary.Operation() != "fstat" || primary.Path() != "two" || !reflect.DeepEqual(record.trace, []string{"open one", "fstat 1", "fstatfs 1", "open two", "fstat 2", "close 2", "close 1"}) {
		t.Fatalf("multi-close lifecycle err=%v trace=%v", err, record.trace)
	}
	assertSecondaryCloseEvidence(t, err, []string{"one", "two"})
	assertEachClosedOnce(t, record)

	walker, record = dfsWalker(map[int][]byte{41: dirRecord(1, directoryType, "one"), 1: nil})
	record.stats[1], record.closeErr[1] = statFor(7, 1), syscall.EIO
	err = walker.walkDirectory(context.Background(), 41, testWalkerFacts(t), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &primary) || primary.Operation() != "close" || primary.Path() != "one" {
		t.Fatalf("successful operation close=%v", err)
	}
	assertEachClosedOnce(t, record)
}

func TestWalkerSiblingContinuation(t *testing.T) {
	walker, record := dfsWalker(map[int][]byte{41: append(dirRecord(1, directoryType, "bad"), dirRecord(2, directoryType, "safe")...), 2: nil})
	record.openErr[1], record.stats[1] = syscall.EIO, statFor(7, 2)
	err := walker.walkDirectory(context.Background(), 41, testWalkerFacts(t), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if !errors.Is(err, scan.ErrInaccessible) || !reflect.DeepEqual(record.trace, []string{"open bad", "open safe", "fstat 1", "fstatfs 1", "close 1"}) || record.closes[1] != 1 {
		t.Fatalf("err=%v trace=%v closes=%v", err, record.trace, record.closes)
	}

	walker, record = dfsWalker(map[int][]byte{41: append(dirRecord(1, directoryType, "changed"), dirRecord(2, directoryType, "safe")...), 2: nil})
	record.stats[1], record.stats[2] = statFor(7, 9), statFor(7, 2)
	err = walker.walkDirectory(context.Background(), 41, testWalkerFacts(t), scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 2})
	if !errors.Is(err, scan.ErrInvalidMetadata) || !reflect.DeepEqual(record.trace, []string{"open changed", "fstat 1", "fstatfs 1", "close 1", "open safe", "fstat 2", "fstatfs 2", "close 2"}) {
		t.Fatalf("changed child err=%v trace=%v", err, record.trace)
	}
	assertEachClosedOnce(t, record)
}

func dfsWalker(directories map[int][]byte) (trustedRootWalker, *walkerRecord) {
	walker, record := scriptedWalker()
	record.directories = directories
	return walker, record
}

func assertEachClosedOnce(t *testing.T, record *walkerRecord) {
	t.Helper()
	for fd := 1; fd <= len(record.names); fd++ {
		if record.closes[fd] != 1 {
			t.Fatalf("fd %d closes=%d trace=%v", fd, record.closes[fd], record.trace)
		}
	}
}

func TestWalkerRealDirectoryReplacementBinding(t *testing.T) {
	fixture := newRealEnumerationFixture(t, realChildSymlink)
	rootFD := fixture.openRoot(t)
	defer func() { _ = syscall.Close(rootFD) }()

	openedOriginal := fstatIdentity(t, rootFD)
	if err := fixture.walk(rootFD, scan.WalkLimits{MaxDepth: 2, MaxDescriptors: 3}); !errors.Is(err, scan.ErrInvalidMetadata) {
		t.Fatalf("walk error = %v, want changed-child classification", err)
	}
	retained := statIdentity(t, fixture.retained)
	replacement := statIdentity(t, fixture.root)
	if openedOriginal != retained || openedOriginal == replacement {
		t.Fatalf("opened=%#v retained=%#v replacement=%#v", openedOriginal, retained, replacement)
	}
	fixture.assertIdentityLifecycle(t, rootFD, true)
}

func TestWalkerRealChangedChildNoDescent(t *testing.T) {
	for _, kind := range []realChildChange{realChildVanished, realChildSymlink} {
		t.Run(string(kind), func(t *testing.T) {
			fixture := newRealEnumerationFixture(t, kind)
			rootFD := fixture.openRoot(t)
			defer func() { _ = syscall.Close(rootFD) }()

			err := fixture.walk(rootFD, scan.WalkLimits{MaxDepth: 2, MaxDescriptors: 3})
			want := scan.ErrMissing
			if kind == realChildSymlink {
				want = scan.ErrInvalidMetadata
			}
			if !errors.Is(err, want) {
				t.Fatalf("walk error = %v, want %v", err, want)
			}
			fixture.assertIdentityLifecycle(t, rootFD, true)
		})
	}
}

func TestWalkerRealDescriptorAdversarialCombinations(t *testing.T) {
	fixture := newRealEnumerationFixture(t, realChildSymlink)
	rootFD := fixture.openRoot(t)
	defer func() { _ = syscall.Close(rootFD) }()

	err := fixture.walk(rootFD, scan.WalkLimits{MaxDepth: 1, MaxDescriptors: 1})
	if !errors.Is(err, scan.ErrLimit) {
		t.Fatalf("walk error = %v, want descriptor limit", err)
	}
	fixture.assertDescriptorLimitTrace(t, rootFD)
}

type realChildChange string

const (
	realChildVanished realChildChange = "vanished"
	realChildSymlink  realChildChange = "symlink-replaced"
)

type realTraceEvent struct {
	op, name   string
	parent, fd int
	result     string
}
type realChildOpen struct {
	parent, fd            int
	name                  string
	identity              scan.Identity
	fstat, fstatfs, close int
}
type realEnumerationFixture struct {
	temp, root, retained, poison string
	change                       realChildChange
	mutated                      bool
	trace                        []realTraceEvent
	opened                       map[int]*realChildOpen
}

func newRealEnumerationFixture(t *testing.T, change realChildChange) *realEnumerationFixture {
	t.Helper()
	temp := t.TempDir()
	root := filepath.Join(temp, "root")
	mustMkdir(t, filepath.Join(root, "changed"))
	mustMkdir(t, filepath.Join(root, "safe"))
	poison := filepath.Join(temp, "poison")
	mustMkdir(t, filepath.Join(poison, "target-child"))
	return &realEnumerationFixture{temp: temp, root: root, poison: poison, change: change, opened: map[int]*realChildOpen{}}
}

func (f *realEnumerationFixture) event(op, name string, parent, fd int, result string) {
	f.trace = append(f.trace, realTraceEvent{op: op, name: name, parent: parent, fd: fd, result: result})
}
func (f *realEnumerationFixture) openRoot(t *testing.T) int {
	t.Helper()
	fd, err := syscall.Open(f.root, walkerFlags, 0)
	if err != nil {
		t.Fatal(err)
	}
	return fd
}
func (f *realEnumerationFixture) walk(rootFD int, limits scan.WalkLimits) error {
	rootFacts, err := scan.NewRootFacts(fstatIdentityNoFail(rootFD), true)
	if err != nil {
		return err
	}
	walker := newTrustedRootWalker(walkerOps{
		open: func(parent int, name string, flags int, mode uint32) (int, error) {
			f.event("open-attempt", name, parent, -1, "")
			fd, err := walkerOpen(parent, name, flags, mode)
			if err != nil {
				f.event("open-result", name, parent, -1, err.Error())
				return -1, err
			}
			f.opened[fd] = &realChildOpen{parent: parent, fd: fd, name: name}
			f.event("open-result", name, parent, fd, "ok")
			return fd, nil
		},
		fstat: func(fd int, stat *syscall.Stat_t) error {
			err := syscall.Fstat(fd, stat)
			f.event("fstat", "", -1, fd, resultOf(err))
			if open := f.opened[fd]; open != nil {
				open.fstat++
				if err == nil {
					open.identity = scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
				}
			}
			return err
		},
		fstatfs: func(fd int, stat *syscall.Statfs_t) error {
			err := syscall.Fstatfs(fd, stat)
			f.event("fstatfs", "", -1, fd, resultOf(err))
			if open := f.opened[fd]; open != nil {
				open.fstatfs++
			}
			return err
		},
		close: func(fd int) error {
			err := syscall.Close(fd)
			f.event("close", "", -1, fd, resultOf(err))
			if open := f.opened[fd]; open != nil {
				open.close++
			}
			return err
		},
		readDir: func(fd int) ([]byte, error) {
			f.event("read-dir", "", -1, fd, "")
			raw, err := walkerReadDir(fd)
			if err == nil && fd == rootFD && !f.mutated {
				f.mutateEnumeratedChild()
			}
			return raw, err
		},
	})
	err = walker.walkDirectory(context.Background(), rootFD, rootFacts, limits)
	f.event("walk-return", "", -1, rootFD, resultOf(err))
	return err
}
func resultOf(err error) string {
	if err != nil {
		return err.Error()
	}
	return "ok"
}
func (f *realEnumerationFixture) mutateEnumeratedChild() {
	f.mutated = true
	f.event("mutation", "", -1, -1, "rename-and-replace")
	f.retained = filepath.Join(f.temp, "retained")
	if err := os.Rename(f.root, f.retained); err != nil {
		panic(err)
	}
	mustMkdirPanic(f.root)
	mustMkdirPanic(filepath.Join(f.root, "safe"))
	mustMkdirPanic(filepath.Join(f.root, "poison-replacement-child"))
	changed := filepath.Join(f.retained, "changed")
	if err := os.RemoveAll(changed); err != nil {
		panic(err)
	}
	if f.change == realChildSymlink {
		if err := os.Symlink(f.poison, changed); err != nil {
			panic(err)
		}
	}
}
func (f *realEnumerationFixture) assertIdentityLifecycle(t *testing.T, rootFD int, changedFails bool) {
	t.Helper()
	if !f.mutated {
		t.Fatal("enumerated root was not mutated")
	}
	safe := f.openedByName("safe")
	if safe == nil || safe.parent != rootFD {
		t.Fatalf("safe open=%#v rootFD=%d trace=%#v", safe, rootFD, f.trace)
	}
	wantSafe := statIdentity(t, filepath.Join(f.retained, "safe"))
	replacementSafe := statIdentity(t, filepath.Join(f.root, "safe"))
	if wantSafe == replacementSafe {
		t.Fatalf("retained safe identity = replacement safe identity = %#v", wantSafe)
	}
	if safe.identity != wantSafe || safe.identity == replacementSafe {
		t.Fatalf("safe fd identity=%#v retained=%#v replacement=%#v", safe.identity, wantSafe, replacementSafe)
	}
	for _, child := range f.opened {
		if child.parent != rootFD || child.fstat != 1 || child.fstatfs != 1 || child.close != 1 {
			t.Fatalf("child lifecycle=%#v trace=%#v", child, f.trace)
		}
		open, fstat, fstatfs, read, close := f.indicesFor(child.fd)
		if !(open < fstat && fstat < fstatfs && fstatfs < read && read < close && close < f.index("walk-return", rootFD)) {
			t.Fatalf("child trace ordering=%d/%d/%d/%d/%d trace=%#v", open, fstat, fstatfs, read, close, f.trace)
		}
	}
	for _, event := range f.trace {
		if strings.Contains(event.name, "poison") || event.name == "target-child" || event.name == "poison-replacement-child" || (event.name == "safe" && event.parent != rootFD) {
			t.Fatalf("replacement target or safe-tree trace=%#v", f.trace)
		}
	}
	if changedFails && (f.openedByName("changed") != nil || f.closeCountForName("changed") != 0 || f.resultFD("changed") != -1) {
		t.Fatalf("changed child owned fd/close trace=%#v", f.trace)
	}
}
func (f *realEnumerationFixture) assertDescriptorLimitTrace(t *testing.T, rootFD int) {
	t.Helper()
	if !f.mutated || !f.has("read-dir", rootFD) || !f.has("mutation", -1) {
		t.Fatalf("limit lacked read/mutation trace=%#v", f.trace)
	}
	for _, event := range f.trace {
		if event.op == "open-attempt" || event.op == "open-result" || event.op == "close" {
			t.Fatalf("limit performed child lifecycle operation=%#v trace=%#v", event, f.trace)
		}
	}
}
func (f *realEnumerationFixture) openedByName(name string) *realChildOpen {
	for _, open := range f.opened {
		if open.name == name {
			return open
		}
	}
	return nil
}
func (f *realEnumerationFixture) closeCountForName(name string) int {
	if open := f.openedByName(name); open != nil {
		return open.close
	}
	return 0
}
func (f *realEnumerationFixture) has(op string, fd int) bool { return f.index(op, fd) >= 0 }
func (f *realEnumerationFixture) index(op string, fd int) int {
	for i, event := range f.trace {
		if event.op == op && event.fd == fd {
			return i
		}
	}
	return -1
}
func (f *realEnumerationFixture) resultFD(name string) int {
	for _, event := range f.trace {
		if event.op == "open-result" && event.name == name {
			return event.fd
		}
	}
	return -1
}
func (f *realEnumerationFixture) indicesFor(fd int) (int, int, int, int, int) {
	out := [5]int{-1, -1, -1, -1, -1}
	for i, e := range f.trace {
		if e.fd != fd {
			continue
		}
		switch e.op {
		case "open-result":
			out[0] = i
		case "fstat":
			out[1] = i
		case "fstatfs":
			out[2] = i
		case "read-dir":
			out[3] = i
		case "close":
			out[4] = i
		}
	}
	return out[0], out[1], out[2], out[3], out[4]
}

func fstatIdentity(t *testing.T, fd int) scan.Identity {
	t.Helper()
	identity, err := fstatIdentityChecked(fd)
	if err != nil {
		t.Fatal(err)
	}
	return identity
}

func fstatIdentityNoFail(fd int) scan.Identity {
	identity, err := fstatIdentityChecked(fd)
	if err != nil {
		panic(err)
	}
	return identity
}

func fstatIdentityChecked(fd int) (scan.Identity, error) {
	var stat syscall.Stat_t
	if err := syscall.Fstat(fd, &stat); err != nil {
		return scan.Identity{}, err
	}
	return scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}, nil
}

func mustMkdirPanic(path string) {
	if err := os.MkdirAll(path, 0o755); err != nil {
		panic(err)
	}
}
func TestWalkerPreCancelledDirectorySkipsReadAndChildWork(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	walker, record := dfsWalker(map[int][]byte{41: dirRecord(9, directoryType, "child")})
	root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
	err := walker.WalkFacts(ctx, root, func(scan.DirectoryFact) error { t.Fatal("pre-cancelled walker emitted a child"); return nil })
	if !errors.Is(err, scan.ErrInaccessible) || record.readDirs != 0 || record.openCount != 0 {
		t.Fatalf("err=%v reads=%d opens=%d", err, record.readDirs, record.openCount)
	}
	if err := root.Close(); err != nil || record.closes[41] != 1 {
		t.Fatalf("root close=%v count=%d", err, record.closes[41])
	}
}

func TestWalkerFactPortTraversal(t *testing.T) {
	walker, record := scriptedWalker()
	record.directory = append(dirRecord(9, directoryType, "dir"), dirRecord(0, 10, "link")...)
	record.directory = append(record.directory, dirRecord(0, 1, "special")...)
	record.directories = map[int][]byte{41: record.directory, 1: nil}
	record.stats[1] = statFor(7, 9)
	root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
	var facts []scan.DirectoryFact
	if err := walker.WalkFacts(context.Background(), root, func(fact scan.DirectoryFact) error { facts = append(facts, fact); return nil }); err != nil {
		t.Fatal(err)
	}
	if len(facts) != 3 || facts[0].EnumeratedKind() != scan.EntryDirectory || facts[0].OpenedKind() != scan.EntryDirectory || facts[1].EnumeratedKind() != scan.EntrySymlink || facts[1].OpenedKind() != scan.EntrySymlink || facts[2].EnumeratedKind() != scan.EntrySpecial || facts[2].OpenedKind() != scan.EntrySpecial {
		t.Fatalf("facts=%#v", facts)
	}
	if err := root.Close(); err != nil || record.closes[41] != 1 || record.readDirs != 2 {
		t.Fatalf("close=%v closes=%v reads=%d", err, record.closes, record.readDirs)
	}
}

func TestWalkerFactPortNestedFacts(t *testing.T) {
	walker, record := dfsWalker(map[int][]byte{
		41: dirRecord(1, directoryType, "pkg"),
		1:  dirRecord(2, 8, "data"),
	})
	record.stats[1] = statFor(7, 1)
	record.stats[2] = syscall.Stat_t{Dev: 7, Ino: 2, Nlink: 1, Mode: syscall.S_IFREG}
	root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
	var got []scan.DirectoryFact
	if err := walker.WalkFacts(context.Background(), root, func(fact scan.DirectoryFact) error { got = append(got, fact); return nil }); err != nil {
		t.Fatal(err)
	}
	if components := [][]string{got[0].Components(), got[1].Components()}; !reflect.DeepEqual(components, [][]string{{"pkg"}, {"pkg", "data"}}) {
		t.Fatalf("components=%v", components)
	}
	enumerated, enumOK := got[1].EnumeratedIdentity()
	opened, openOK := got[1].OpenedIdentity()
	if got[1].EnumeratedKind() != scan.EntryRegular || got[1].OpenedKind() != scan.EntryRegular || !enumOK || !openOK || enumerated != opened || got[1].SkipClass() != nil {
		t.Fatalf("nested regular fact=%#v", got[1])
	}
}

func TestWalkerFactPortOwnership(t *testing.T) {
	walker, record := scriptedWalker()
	record.directory = dirRecord(9, directoryType, "child")
	record.stats[1] = statFor(7, 9)
	root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
	if err := walker.WalkFacts(context.Background(), root, func(scan.DirectoryFact) error { return nil }); err != nil {
		t.Fatal(err)
	}
	if record.closes[41] != 0 {
		t.Fatalf("WalkFacts closed scanner-owned root %d times", record.closes[41])
	}
	if err := root.Close(); err != nil || record.closes[41] != 1 {
		t.Fatalf("root close=%v count=%d", err, record.closes[41])
	}
}

func TestWalkerFactPortCallbackFailureJoinsChildCloseFailure(t *testing.T) {
	walker, record := dfsWalker(map[int][]byte{41: dirRecord(9, directoryType, "child"), 1: nil})
	record.stats[1], record.closeErr[1] = statFor(7, 9), syscall.EIO
	root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
	callbackClass := errors.New("callback failed")
	callbackErr, invalid := scan.NewPathError("emit", "child", callbackClass)
	if invalid != nil {
		t.Fatal(invalid)
	}

	err := walker.WalkFacts(context.Background(), root, func(scan.DirectoryFact) error { return callbackErr })
	var primary *scan.PathError
	if !errors.Is(err, callbackClass) || !errors.Is(err, scan.ErrInaccessible) || !errors.As(err, &primary) || primary != callbackErr || primary.Operation() != "emit" || primary.Path() != "child" {
		t.Fatalf("error precedence=%v primary=%v", err, primary)
	}
	assertSecondaryCloseEvidence(t, err, []string{"child"})
	if record.closes[1] != 1 || record.closes[41] != 0 || !reflect.DeepEqual(record.trace, []string{"open child", "fstat 1", "fstatfs 1", "close 1"}) {
		t.Fatalf("closes=%v trace=%v", record.closes, record.trace)
	}
}

func TestWalkerFactPortKindIdentityPairs(t *testing.T) {
	walker, record := scriptedWalker()
	record.directory = dirRecord(9, directoryType, "changed")
	record.stats[1] = syscall.Stat_t{Dev: 7, Ino: 10, Nlink: 1, Mode: syscall.S_IFREG}
	root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
	var got scan.DirectoryFact
	if err := walker.WalkFacts(context.Background(), root, func(fact scan.DirectoryFact) error { got = fact; return nil }); err != nil {
		t.Fatal(err)
	}
	enumerated, enumOK := got.EnumeratedIdentity()
	opened, openOK := got.OpenedIdentity()
	if got.EnumeratedKind() != scan.EntryDirectory || got.OpenedKind() != scan.EntryRegular || !enumOK || !openOK || enumerated.Inode != 9 || opened.Inode != 10 || !errors.Is(got.SkipClass(), scan.ErrInvalidMetadata) {
		t.Fatalf("changed pair=%#v", got)
	}
}

func TestWalkerRegularFactEvidence(t *testing.T) {
	for _, tt := range []struct {
		name                         string
		stat                         syscall.Stat_t
		openErr, fstatErr, statfsErr error
		closeErr                     error
		nonLocal                     bool
		opened                       bool
		closes                       int
		class                        error
	}{
		{"stable", syscall.Stat_t{Dev: 7, Ino: 9, Nlink: 1, Mode: syscall.S_IFREG}, nil, nil, nil, nil, false, true, 1, nil},
		{"replacement", syscall.Stat_t{Dev: 7, Ino: 10, Nlink: 1, Mode: syscall.S_IFREG}, nil, nil, nil, nil, false, true, 1, scan.ErrInvalidMetadata},
		{"symlink", syscall.Stat_t{}, syscall.ELOOP, nil, nil, nil, false, false, 0, scan.ErrInvalidMetadata},
		{"missing", syscall.Stat_t{}, syscall.ENOENT, nil, nil, nil, false, false, 0, scan.ErrMissing},
		{"type", statFor(7, 9), nil, nil, nil, nil, false, true, 1, scan.ErrInvalidMetadata},
		{"device", syscall.Stat_t{Dev: 8, Ino: 9, Nlink: 1, Mode: syscall.S_IFREG}, nil, nil, nil, nil, false, true, 1, scan.ErrDeviceBoundary},
		{"nonlocal", syscall.Stat_t{Dev: 7, Ino: 9, Nlink: 1, Mode: syscall.S_IFREG}, nil, nil, nil, nil, true, true, 1, scan.ErrNonLocal},
		{"fstat", syscall.Stat_t{}, nil, syscall.EIO, nil, nil, false, false, 1, scan.ErrInaccessible},
		{"fstatfs", syscall.Stat_t{Dev: 7, Ino: 9, Nlink: 1, Mode: syscall.S_IFREG}, nil, nil, syscall.EIO, nil, false, true, 1, scan.ErrInaccessible},
		{"close", syscall.Stat_t{Dev: 7, Ino: 9, Nlink: 1, Mode: syscall.S_IFREG}, nil, nil, nil, syscall.EIO, false, true, 1, scan.ErrInaccessible},
	} {
		t.Run(tt.name, func(t *testing.T) {
			walker, record := scriptedWalker()
			record.directory = dirRecord(9, 8, "data")
			record.stats[1], record.openErr[1], record.fstatErr[1], record.fstatfsErr[1], record.closeErr[1], record.nonLocal[1] = tt.stat, tt.openErr, tt.fstatErr, tt.statfsErr, tt.closeErr, tt.nonLocal
			root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
			var got scan.DirectoryFact
			if err := walker.WalkFacts(context.Background(), root, func(fact scan.DirectoryFact) error { got = fact; return nil }); err != nil {
				t.Fatal(err)
			}
			enumerated, enumOK := got.EnumeratedIdentity()
			opened, openOK := got.OpenedIdentity()
			if got.EnumeratedKind() != scan.EntryRegular || !enumOK || enumerated != (scan.Identity{Device: 7, Inode: 9}) || openOK != tt.opened || !errors.Is(got.SkipClass(), tt.class) {
				t.Fatalf("fact=%#v enum=%#v/%t opened=%#v/%t", got, enumerated, enumOK, opened, openOK)
			}
			if tt.class == nil && (got.OpenedKind() != scan.EntryRegular || opened != enumerated || !got.Local()) {
				t.Fatalf("stable fact=%#v enum=%#v opened=%#v", got, enumerated, opened)
			}
			wantFlags := []int(nil)
			if tt.closes == 1 {
				wantFlags = []int{fileOpenFlags}
			}
			if !reflect.DeepEqual(record.parents, []int{41}) || !reflect.DeepEqual(record.flags, wantFlags) || record.closes[1] != tt.closes {
				t.Fatalf("parents=%v flags=%#x closes=%v", record.parents, record.flags, record.closes)
			}
			if gotPath, ok := got.Components()[0], len(got.Components()) == 1; !ok || gotPath != "data" {
				t.Fatalf("components=%v", got.Components())
			}
		})
	}
}

func TestWalkerFactPortNoDuplicateTraversal(t *testing.T) {
	walker, record := dfsWalker(map[int][]byte{41: dirRecord(1, directoryType, "one"), 1: nil})
	record.stats[1] = statFor(7, 1)
	root := &acquiredRoot{facts: testWalkerFacts(t), fd: 41, close: func() error { return walker.ops.close(41) }}
	if err := walker.WalkFacts(context.Background(), root, func(scan.DirectoryFact) error { return nil }); err != nil {
		t.Fatal(err)
	}
	if record.readDirs != 2 || len(record.names) != 1 {
		t.Fatalf("directory reads=%d opens=%v, want one root and one child traversal", record.readDirs, record.names)
	}
}

func TestWalkerRelativeFileReopen(t *testing.T) {
	fixture := newRealWalkerFixture(t)
	mustMkdir(t, filepath.Join(fixture.root, "nested"))
	file := filepath.Join(fixture.root, "nested", "file")
	if err := os.WriteFile(file, []byte("metadata only"), 0o600); err != nil {
		t.Fatal(err)
	}
	root, err := fixture.acquire()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := root.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	job, err := scan.NewFileJob([]string{"nested", "file"}, []scan.Identity{statIdentity(t, filepath.Join(fixture.root, "nested"))}, statIdentity(t, file))
	if err != nil {
		t.Fatal(err)
	}
	result := root.(*acquiredRoot).Inspect(context.Background(), job)
	if !result.Accepted() || result.LogicalBytes() != uint64(len("metadata only")) || result.AllocationBytes() == 0 {
		t.Fatalf("result=%#v", result)
	}
	if got := fixture.opens[len(fixture.opens)-2:]; !reflect.DeepEqual(got, []int{walkerFlags, fileOpenFlags}) {
		t.Fatalf("reopen flags=%#x", got)
	}
	if fixture.paths[fixture.records[len(fixture.records)-1].fd] != file {
		t.Fatal("file reopen did not remain relative to retained root")
	}
}

func TestWalkerFileIdentityChain(t *testing.T) {
	for _, change := range []string{"identity", "rename", "symlink", "missing"} {
		t.Run(change, func(t *testing.T) {
			fixture := newRealWalkerFixture(t)
			nested := filepath.Join(fixture.root, "nested")
			mustMkdir(t, nested)
			file := filepath.Join(nested, "file")
			if err := os.WriteFile(file, []byte("x"), 0o600); err != nil {
				t.Fatal(err)
			}
			root, err := fixture.acquire()
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = root.Close() }()
			ancestor, final := statIdentity(t, nested), statIdentity(t, file)
			if change == "rename" {
				if err := os.Rename(nested, filepath.Join(fixture.temp, "retained")); err != nil {
					t.Fatal(err)
				}
				mustMkdir(t, nested)
				_ = os.WriteFile(filepath.Join(nested, "file"), []byte("replacement"), 0o600)
			}
			if change == "symlink" {
				if err := os.RemoveAll(nested); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(fixture.temp, nested); err != nil {
					t.Fatal(err)
				}
			}
			if change == "missing" {
				if err := os.Remove(file); err != nil {
					t.Fatal(err)
				}
			}
			if change == "identity" {
				ancestor = scan.Identity{Device: 9, Inode: 9}
			}
			job, _ := scan.NewFileJob([]string{"nested", "file"}, []scan.Identity{ancestor}, final)
			result := root.(*acquiredRoot).Inspect(context.Background(), job)
			if result.Accepted() || result.LogicalBytes() != 0 || result.AllocationBytes() != 0 || result.Error() == nil {
				t.Fatalf("changed result=%#v err=%v", result, result.Error())
			}
		})
	}
}

func TestWalkerRegularFileFacts(t *testing.T) {
	fixture := newRealWalkerFixture(t)
	for _, name := range []string{"directory", "file"} {
		mustMkdir(t, filepath.Join(fixture.root, name))
	}
	root, err := fixture.acquire()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = root.Close() }()
	directory := filepath.Join(fixture.root, "directory")
	job, _ := scan.NewFileJob([]string{"directory"}, nil, statIdentity(t, directory))
	result := root.(*acquiredRoot).Inspect(context.Background(), job)
	if result.Accepted() || result.LogicalBytes() != 0 || result.AllocationBytes() != 0 || !errors.Is(result.Error(), scan.ErrInvalidMetadata) {
		t.Fatalf("nonregular result=%#v err=%v", result, result.Error())
	}
	fifo := filepath.Join(fixture.root, "fifo")
	if err := syscall.Mkfifo(fifo, 0o600); err != nil {
		t.Fatal(err)
	}
	job, _ = scan.NewFileJob([]string{"fifo"}, nil, statIdentity(t, fifo))
	result = root.(*acquiredRoot).Inspect(context.Background(), job)
	if result.Accepted() || result.LogicalBytes() != 0 || !errors.Is(result.Error(), scan.ErrInvalidMetadata) {
		t.Fatalf("FIFO result=%#v err=%v", result, result.Error())
	}

	boundary := newRealWalkerFixture(t)
	file := filepath.Join(boundary.root, "file")
	if err := os.WriteFile(file, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	trusted, err := boundary.acquire()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = trusted.Close() }()
	port := trusted.(*acquiredRoot)
	port.ops.fstatfs = func(_ int, volume *syscall.Statfs_t) error { volume.Flags = 0; return nil }
	job, _ = scan.NewFileJob([]string{"file"}, nil, statIdentity(t, file))
	result = port.Inspect(context.Background(), job)
	if result.Accepted() || !errors.Is(result.Error(), scan.ErrNonLocal) || result.LogicalBytes() != 0 {
		t.Fatalf("boundary result=%#v err=%v", result, result.Error())
	}
}

func TestWalkerFileReopenLifecycle(t *testing.T) {
	fixture := newRealWalkerFixture(t)
	mustMkdir(t, filepath.Join(fixture.root, "nested"))
	file := filepath.Join(fixture.root, "nested", "file")
	if err := os.WriteFile(file, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	root, err := fixture.acquire()
	if err != nil {
		t.Fatal(err)
	}
	job, _ := scan.NewFileJob([]string{"nested", "file"}, []scan.Identity{statIdentity(t, filepath.Join(fixture.root, "nested"))}, statIdentity(t, file))
	result := root.(*acquiredRoot).Inspect(context.Background(), job)
	if !result.Accepted() {
		t.Fatalf("result=%#v err=%v", result, result.Error())
	}
	if err := root.Close(); err != nil {
		t.Fatal(err)
	}
	for _, record := range fixture.records {
		if record.closes != 1 {
			t.Fatalf("fd %d closes=%d", record.fd, record.closes)
		}
	}

	closes := []int{}
	r := &acquiredRoot{facts: scan.RootFacts{Device: 7, Identity: scan.Identity{Device: 7, Inode: 100}, Local: true}, fd: 100}
	r.ops = walkerOps{
		open: func(parent int, name string, flags int, _ uint32) (int, error) {
			if parent == 100 && name == "dir" && flags == walkerFlags {
				return 1, nil
			}
			if parent == 1 && name == "file" && flags == fileOpenFlags {
				return 2, nil
			}
			return -1, syscall.EIO
		},
		fstat: func(fd int, stat *syscall.Stat_t) error {
			if fd == 2 {
				return syscall.EIO
			}
			*stat = statFor(7, 1)
			return nil
		},
		fstatfs: func(_ int, volume *syscall.Statfs_t) error { volume.Flags = localFilesystemFlag; return nil },
		close:   func(fd int) error { closes = append(closes, fd); return syscall.EIO },
	}
	job, _ = scan.NewFileJob([]string{"dir", "file"}, []scan.Identity{{Device: 7, Inode: 1}}, scan.Identity{Device: 7, Inode: 2})
	result = r.Inspect(context.Background(), job)
	var primary *scan.PathError
	if result.Accepted() || !errors.Is(result.Error(), scan.ErrInaccessible) || !errors.As(result.Error(), &primary) || primary.Operation() != "fstat" || !reflect.DeepEqual(closes, []int{2, 1}) {
		t.Fatalf("lifecycle err=%v closes=%v", result.Error(), closes)
	}

	closes = nil
	r.ops.fstat = func(_ int, stat *syscall.Stat_t) error { *stat = statFor(7, 1); return nil }
	r.ops.fstatfs = func(_ int, _ *syscall.Statfs_t) error { return syscall.EIO }
	r.ops.close = func(fd int) error { closes = append(closes, fd); return nil }
	result = r.Inspect(context.Background(), job)
	if result.Accepted() || !errors.As(result.Error(), &primary) || primary.Operation() != "fstatfs" || !reflect.DeepEqual(closes, []int{1}) {
		t.Fatalf("fstatfs err=%v closes=%v", result.Error(), closes)
	}
}

//go:build darwin

package macos

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
	"testing"
	"time"

	"github.com/osdy/OsdyCleaner/internal/core"
	"golang.org/x/sys/unix"
)

func TestDarwinMetadataDisposable(t *testing.T) {
	temp := physicalTemp(t)
	file := filepath.Join(temp, "fixture")
	if err := os.WriteFile(file, []byte("fixture content"), 0o600); err != nil {
		t.Fatal(err)
	}
	stamp := time.Unix(1_700_000_000, 0)
	if err := os.Chtimes(file, stamp, stamp); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(temp, "fixture-link")
	if err := os.Link(file, link); err != nil {
		t.Fatal(err)
	}
	beforeContent, beforeInfo := readFixture(t, file)
	stat := lstatFixture(t, temp)
	adapter, err := NewMetadataAdapter(uint64(stat.Dev))
	if err != nil {
		t.Fatal(err)
	}

	fileMetadata, err := adapter.Inspect(file, false)
	if err != nil {
		t.Fatal(err)
	}
	linkMetadata, err := adapter.Inspect(link, false)
	if err != nil {
		t.Fatal(err)
	}
	directoryMetadata, err := adapter.Inspect(temp, true)
	if err != nil {
		t.Fatal(err)
	}
	if !fileMetadata.Mode().IsRegular() || !directoryMetadata.Mode().IsDir() {
		t.Fatalf("unexpected modes: file=%v directory=%v", fileMetadata.Mode(), directoryMetadata.Mode())
	}
	if fileMetadata.LogicalSize() != uint64(len(beforeContent)) {
		t.Fatalf("logical size = %d, want %d", fileMetadata.LogicalSize(), len(beforeContent))
	}
	if _, known := fileMetadata.Allocation().Bytes(); !known {
		t.Fatal("allocation unexpectedly unknown for disposable regular file")
	}
	fileIdentity, fileKnown := fileMetadata.Identity().Value()
	linkIdentity, linkKnown := linkMetadata.Identity().Value()
	if !fileKnown || !linkKnown || fileIdentity != linkIdentity {
		t.Fatalf("hard links did not retain one identity: %v/%v", fileKnown, linkKnown)
	}
	if fileMetadata.LinkCount() < 2 || linkMetadata.LinkCount() < 2 {
		t.Fatalf("link counts = %d and %d, want at least 2", fileMetadata.LinkCount(), linkMetadata.LinkCount())
	}
	assertFixtureUnchanged(t, file, beforeContent, beforeInfo)
}

func TestDarwinMetadataProductionSeams(t *testing.T) {
	typeOfOps := reflect.TypeOf(componentOps{})
	for _, forbidden := range []string{"statType", "revalidate"} {
		if _, exists := typeOfOps.FieldByName(forbidden); exists {
			t.Fatalf("componentOps exposes unreachable %s hook", forbidden)
		}
	}
}

func TestDarwinMetadataDescriptorStability(t *testing.T) {
	temp := physicalTemp(t)
	path := filepath.Join(temp, "opened")
	preserved := filepath.Join(temp, "preserved")
	replacement := filepath.Join(temp, "replacement")
	if err := os.WriteFile(path, []byte("opened-descriptor"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Link(path, preserved); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(replacement, make([]byte, 4096), 0o644); err != nil {
		t.Fatal(err)
	}
	want := *lstatFixture(t, preserved)
	current := lstatFixture(t, replacement)
	if want.Ino == current.Ino || want.Size == current.Size || want.Mode == current.Mode {
		t.Fatalf("replacement is not distinguishable: opened=%+v replacement=%+v", want, current)
	}

	replaced := false
	var accepted syscall.Stat_t
	fstatfsCalls := 0
	adapter := newComponentMetadataAdapter(uint64(want.Dev), componentOps{
		openRoot: unix.Open,
		openat:   unix.Openat,
		fstat: func(fd int, out *syscall.Stat_t) error {
			if !replaced {
				replaced = true
				if err := os.Rename(replacement, path); err != nil {
					return err
				}
			}
			if err := syscall.Fstat(fd, out); err != nil {
				return err
			}
			accepted = *out
			return nil
		},
		fstatfs: func(fd int, out *syscall.Statfs_t) error {
			fstatfsCalls++
			return syscall.Fstatfs(fd, out)
		},
		close: syscall.Close,
	})
	got, err := adapter.Inspect(path, false)
	if err != nil {
		t.Fatal(err)
	}
	if !replaced || fstatfsCalls != 1 {
		t.Fatalf("replacement=%v fstatfs calls=%d", replaced, fstatfsCalls)
	}
	identity, known := got.Identity().Value()
	if !known || got.Device() != uint64(accepted.Dev) || got.LogicalSize() != uint64(accepted.Size) || got.LinkCount() != uint64(accepted.Nlink) || identity != mustIdentity(t, accepted) {
		t.Fatalf("descriptor facts=%+v identity=%v, want accepted descriptor=%+v", got, identity, accepted)
	}
	if allocation, known := got.Allocation().Bytes(); !known || allocation != uint64(accepted.Blocks)*512 {
		t.Fatalf("allocation=%d known=%v, want blocks=%d", allocation, known, accepted.Blocks)
	}
	if got.Mode().Perm() != fs.FileMode(accepted.Mode&0o777) || !got.Mode().IsRegular() {
		t.Fatalf("mode=%v, want accepted descriptor regular mode %#o", got.Mode(), accepted.Mode)
	}
	if after := lstatFixture(t, path); after.Ino != current.Ino || after.Size != current.Size {
		t.Fatalf("path still identifies original descriptor: after=%+v replacement=%+v", after, current)
	}
}

func TestDarwinMetadataClassification(t *testing.T) {
	temp := physicalTemp(t)
	file := filepath.Join(temp, "file")
	if err := os.WriteFile(file, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(temp, "link")
	if err := os.Symlink(file, link); err != nil {
		t.Fatal(err)
	}
	device := uint64(lstatFixture(t, temp).Dev)
	adapter, err := NewMetadataAdapter(device)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		name      string
		path      string
		directory bool
		want      error
	}{
		{"missing", filepath.Join(temp, "missing"), true, ErrMetadataMissing},
		{"symlink", link, false, ErrMetadataSymlink},
		{"non-directory", file, true, ErrMetadataNotDirectory},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := adapter.Inspect(tt.path, tt.directory)
			if !errors.Is(err, tt.want) {
				t.Fatalf("Inspect() error = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestDarwinMetadataRegularOnlyBytes(t *testing.T) {
	temp := physicalTemp(t)
	file := filepath.Join(temp, "file")
	if err := os.WriteFile(file, []byte("bytes"), 0o600); err != nil {
		t.Fatal(err)
	}
	fifo := filepath.Join(temp, "pipe")
	if err := syscall.Mkfifo(fifo, 0o600); err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct{ name, path string }{{"regular", file}, {"directory", temp}, {"fifo", fifo}} {
		t.Run(tt.name, func(t *testing.T) {
			adapter, err := NewMetadataAdapter(uint64(lstatFixture(t, temp).Dev))
			if err != nil {
				t.Fatal(err)
			}
			got, err := adapter.Inspect(tt.path, false)
			if err != nil {
				t.Fatal(err)
			}
			if tt.name == "regular" && got.LogicalSize() == 0 {
				t.Fatal("regular logical size is zero")
			}
			if tt.name != "regular" {
				if got.LogicalSize() != 0 {
					t.Fatalf("special logical size = %d", got.LogicalSize())
				}
				if _, known := got.Allocation().Bytes(); known {
					t.Fatal("special allocation is known")
				}
			}
		})
	}
}

func TestDarwinMetadataComponentParsing(t *testing.T) {
	for _, path := range []string{"relative", "/fixture/./file", "/fixture/../file", "/fixture//file", "/fixture/file/"} {
		if _, err := parseCleanAbsolutePath(path); !errors.Is(err, ErrMetadataInaccessible) {
			t.Fatalf("parseCleanAbsolutePath(%q) = %v, want rejection", path, err)
		}
	}
	parts, err := parseCleanAbsolutePath("/fixture/child/file")
	if err != nil || len(parts) != 3 || parts[0] != "fixture" || parts[2] != "file" {
		t.Fatalf("parts=%v err=%v", parts, err)
	}
}

func TestDarwinMetadataOpenatFlags(t *testing.T) {
	var flags []int
	adapter := newComponentMetadataAdapter(1, componentOps{openRoot: func(_ string, got int, _ uint32) (int, error) {
		flags = append(flags, got)
		return 10, nil
	}, openat: func(_ int, _ string, got int, _ uint32) (int, error) {
		flags = append(flags, got)
		return 11 + len(flags), nil
	}, fstat: validFstat, fstatfs: localFstatfs, close: func(int) error { return nil }})
	if _, err := adapter.Inspect("/first/second/file", false); err != nil {
		t.Fatal(err)
	}
	base := syscall.O_RDONLY | syscall.O_NOFOLLOW | syscall.O_CLOEXEC | syscall.O_NONBLOCK
	if len(flags) != 4 || flags[0] != base|syscall.O_DIRECTORY || flags[1] != base|syscall.O_DIRECTORY || flags[2] != base|syscall.O_DIRECTORY || flags[3] != base {
		t.Fatalf("flags=%#v", flags)
	}
}

func TestDarwinMetadataDescriptorChain(t *testing.T) {
	var closes []int
	adapter := newComponentMetadataAdapter(1, componentOps{openRoot: func(string, int, uint32) (int, error) { return 10, nil }, openat: func(parent int, _ string, _ int, _ uint32) (int, error) { return parent + 1, nil }, fstat: validFstat, fstatfs: localFstatfs, close: func(fd int) error { closes = append(closes, fd); return nil }})
	if _, err := adapter.Inspect("/one/two/file", false); err != nil {
		t.Fatal(err)
	}
	if len(closes) != 4 {
		t.Fatalf("close count=%d, want root/intermediates/final", len(closes))
	}
}

func TestDarwinMetadataNoFollowComponents(t *testing.T) {
	for _, name := range []string{"ancestor", "final"} {
		t.Run(name, func(t *testing.T) {
			adapter := newComponentMetadataAdapter(1, componentOps{openRoot: func(string, int, uint32) (int, error) { return 10, nil }, openat: func(_ int, _ string, flags int, _ uint32) (int, error) {
				if flags&syscall.O_NOFOLLOW == 0 {
					t.Fatal("openat followed link")
				}
				return -1, syscall.ELOOP
			}, close: func(int) error { return nil }})
			if _, err := adapter.Inspect("/one/file", false); !errors.Is(err, ErrMetadataSymlink) {
				t.Fatalf("Inspect() error=%v", err)
			}
		})
	}
}

func TestDarwinMetadataAcceptedFacts(t *testing.T) {
	adapter := newComponentMetadataAdapter(1, componentOps{openRoot: func(string, int, uint32) (int, error) { return 10, nil }, openat: func(int, string, int, uint32) (int, error) { return 11, nil }, fstat: validFstat, fstatfs: localFstatfs, close: func(int) error { return nil }})
	got, err := adapter.Inspect("/file", false)
	if err != nil || !got.Mode().IsRegular() || got.Device() != 1 || got.LogicalSize() != 7 {
		t.Fatalf("metadata=%+v err=%v", got, err)
	}
}

func TestDarwinMetadataRevalidation(t *testing.T) {
	adapter := newComponentMetadataAdapter(2, componentOps{openRoot: func(string, int, uint32) (int, error) { return 10, nil }, openat: func(int, string, int, uint32) (int, error) { return 11, nil }, fstat: validFstat, fstatfs: localFstatfs, close: func(int) error { return nil }})
	if _, err := adapter.Inspect("/file", false); !errors.Is(err, ErrMetadataDeviceMismatch) {
		t.Fatalf("device mismatch=%v", err)
	}
}

func TestDarwinMetadataResolverErrors(t *testing.T) {
	for _, tt := range []struct {
		name string
		ops  componentOps
		want error
	}{
		{"missing", componentOps{openRoot: func(string, int, uint32) (int, error) { return -1, fs.ErrNotExist }}, ErrMetadataMissing},
		{"open", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return -1, syscall.EACCES }, close: func(int) error { return nil }}, ErrMetadataInaccessible},
		{"fstat", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: func(int, *syscall.Stat_t) error { return syscall.EIO }, close: func(int) error { return nil }}, ErrMetadataInaccessible},
		{"fstatfs", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: validFstat, fstatfs: func(int, *syscall.Statfs_t) error { return syscall.EIO }, close: func(int) error { return nil }}, ErrMetadataInaccessible},
	} {
		t.Run(tt.name, func(t *testing.T) {
			adapter := newComponentMetadataAdapter(1, tt.ops)
			if _, err := adapter.Inspect("/file", false); !errors.Is(err, tt.want) {
				t.Fatalf("Inspect() error=%v, want %v", err, tt.want)
			}
		})
	}
}

func TestDarwinMetadataErrorClasses(t *testing.T) {
	for _, tt := range []struct {
		name string
		ops  componentOps
		want error
	}{
		{"missing", componentOps{openRoot: func(string, int, uint32) (int, error) { return -1, fs.ErrNotExist }}, ErrMetadataMissing},
		{"symlink", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return -1, syscall.ELOOP }, close: func(int) error { return nil }}, ErrMetadataSymlink},
		{"not-directory", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: validFstat, fstatfs: localFstatfs, close: func(int) error { return nil }}, ErrMetadataNotDirectory},
		{"device", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: validFstat, fstatfs: localFstatfs, close: func(int) error { return nil }}, ErrMetadataDeviceMismatch},
		{"non-local", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: validFstat, fstatfs: func(int, *syscall.Statfs_t) error { return nil }, close: func(int) error { return nil }}, ErrMetadataNonLocalFilesystem},
		{"inaccessible", componentOps{openRoot: func(string, int, uint32) (int, error) { return -1, syscall.EACCES }}, ErrMetadataInaccessible},
		{"invalid", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: invalidFstat, close: func(int) error { return nil }}, ErrMetadataInvalid},
	} {
		t.Run(tt.name, func(t *testing.T) {
			device := uint64(1)
			if tt.name == "device" {
				device = 2
			}
			adapter := newComponentMetadataAdapter(device, tt.ops)
			_, err := adapter.Inspect("/file", tt.name == "not-directory")
			if !errors.Is(err, tt.want) {
				t.Fatalf("Inspect() error = %v, want %v", err, tt.want)
			}
			var pathErr *MetadataPathError
			if !errors.As(err, &pathErr) || pathErr.Operation == "" || pathErr.Path != "/file" {
				t.Fatalf("Inspect() error = %v, want MetadataPathError context", err)
			}
		})
	}
}

func TestDarwinMetadataInvalidFacts(t *testing.T) {
	for _, tt := range []struct {
		name string
		stat syscall.Stat_t
	}{
		{"zero-device", syscall.Stat_t{Ino: 1, Nlink: 1, Mode: syscall.S_IFREG}},
		{"zero-inode", syscall.Stat_t{Dev: 1, Nlink: 1, Mode: syscall.S_IFREG}},
		{"zero-links", syscall.Stat_t{Dev: 1, Ino: 1, Mode: syscall.S_IFREG}},
		{"negative-size", syscall.Stat_t{Dev: 1, Ino: 1, Nlink: 1, Mode: syscall.S_IFREG, Size: -1}},
		{"negative-blocks", syscall.Stat_t{Dev: 1, Ino: 1, Nlink: 1, Mode: syscall.S_IFREG, Blocks: -1}},
		{"overflow-blocks", syscall.Stat_t{Dev: 1, Ino: 1, Nlink: 1, Mode: syscall.S_IFREG, Blocks: int64(^uint64(0)/512 + 1)}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			adapter := lifecycleAdapter(1, func(_ int, out *syscall.Stat_t) error { *out = tt.stat; return nil }, func(int) error { return nil })
			if _, err := adapter.Inspect("/file", false); !errors.Is(err, ErrMetadataInvalid) {
				t.Fatalf("Inspect() error = %v, want invalid metadata", err)
			}
		})
	}

}

func TestDarwinMetadataCloseFailurePrecedence(t *testing.T) {
	for _, tt := range []struct {
		name  string
		fstat func(int, *syscall.Stat_t) error
		close func(int) error
		want  error
	}{
		{"success-close-failure", validFstat, func(int) error { return syscall.EIO }, ErrMetadataInaccessible},
		{"fstat-primary", func(int, *syscall.Stat_t) error { return syscall.EPERM }, func(int) error { return syscall.EIO }, ErrMetadataInaccessible},
	} {
		t.Run(tt.name, func(t *testing.T) {
			adapter := lifecycleAdapter(1, tt.fstat, tt.close)
			got, err := adapter.Inspect("/file", false)
			if !errors.Is(err, tt.want) || got != (Metadata{}) {
				t.Fatalf("Inspect() metadata=%+v error=%v, want %v and no metadata", got, err, tt.want)
			}
		})
	}
}

func TestDarwinMetadataAdversarialLifecycle(t *testing.T) {
	for _, tt := range []struct {
		name       string
		openatFail bool
		fstatFail  bool
		wantCloses int
	}{
		{"open-intermediate", true, false, 2},
		{"fstat-final", false, true, 4},
		{"successful-chain", false, false, 4},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closes := 0
			opens := 0
			ops := componentOps{openRoot: func(string, int, uint32) (int, error) { return 10, nil }, close: func(int) error { closes++; return nil }, fstatfs: localFstatfs}
			ops.openat = func(parent int, _ string, _ int, _ uint32) (int, error) {
				opens++
				if tt.openatFail && opens == 2 {
					return -1, syscall.EIO
				}
				return parent + 1, nil
			}
			ops.fstat = validFstat
			if tt.fstatFail {
				ops.fstat = func(int, *syscall.Stat_t) error { return syscall.EIO }
			}
			adapter := newComponentMetadataAdapter(1, ops)
			_, _ = adapter.Inspect("/one/two/file", false)
			if closes != tt.wantCloses {
				t.Fatalf("close count = %d, want %d", closes, tt.wantCloses)
			}
		})
	}
}

func TestDarwinMetadataLifecycleCombinations(t *testing.T) {
	for _, tt := range []struct {
		name string
		ops  componentOps
		want error
	}{
		{"fstatfs-primary", componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: validFstat, fstatfs: func(int, *syscall.Statfs_t) error { return syscall.EIO }, close: func(int) error { return syscall.EIO }}, ErrMetadataInaccessible},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closes := 0
			originalClose := tt.ops.close
			tt.ops.close = func(fd int) error { closes++; return originalClose(fd) }
			adapter := newComponentMetadataAdapter(1, tt.ops)
			if _, err := adapter.Inspect("/file", false); !errors.Is(err, tt.want) || closes != 2 {
				t.Fatalf("Inspect() error=%v closes=%d, want %v and 2", err, closes, tt.want)
			}
		})
	}

	closes := 0
	adapter := lifecycleAdapter(1, validFstat, func(int) error { closes++; return syscall.EIO })
	if got, err := adapter.Inspect("/file", false); !errors.Is(err, ErrMetadataInaccessible) || got != (Metadata{}) || closes != 2 {
		t.Fatalf("close-only failure metadata=%+v error=%v closes=%d", got, err, closes)
	}
}

func lifecycleAdapter(device uint64, fstat func(int, *syscall.Stat_t) error, close func(int) error) MetadataAdapter {
	return newComponentMetadataAdapter(device, componentOps{openRoot: func(string, int, uint32) (int, error) { return 1, nil }, openat: func(int, string, int, uint32) (int, error) { return 2, nil }, fstat: fstat, fstatfs: localFstatfs, close: close})
}

func mustIdentity(t *testing.T, stat syscall.Stat_t) core.FilesystemIdentity {
	t.Helper()
	identity, err := core.NewFilesystemIdentity(uint64(stat.Dev), uint64(stat.Ino))
	if err != nil {
		t.Fatal(err)
	}
	return identity
}

func invalidFstat(_ int, out *syscall.Stat_t) error {
	*out = syscall.Stat_t{Dev: 1, Ino: 1, Nlink: 0, Mode: syscall.S_IFREG}
	return nil
}

func validFstat(_ int, out *syscall.Stat_t) error {
	*out = syscall.Stat_t{Dev: 1, Ino: 1, Nlink: 1, Mode: syscall.S_IFREG, Size: 7, Blocks: 1}
	return nil
}

func localFstatfs(_ int, out *syscall.Statfs_t) error { out.Flags = localFilesystemFlag; return nil }

func readFixture(t *testing.T, path string) ([]byte, os.FileInfo) {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Lstat(path)
	if err != nil {
		t.Fatal(err)
	}
	return content, info
}

func assertFixtureUnchanged(t *testing.T, path string, wantContent []byte, wantInfo os.FileInfo) {
	t.Helper()
	content, info := readFixture(t, path)
	if string(content) != string(wantContent) || info.Mode() != wantInfo.Mode() || !info.ModTime().Equal(wantInfo.ModTime()) {
		t.Fatalf("fixture changed: content=%q mode=%v mtime=%v", content, info.Mode(), info.ModTime())
	}
}

func physicalTemp(t *testing.T) string {
	t.Helper()
	path, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	return path
}

func fixtureInfo(t *testing.T, path string) os.FileInfo {
	t.Helper()
	info, err := os.Lstat(path)
	if err != nil {
		t.Fatal(err)
	}
	return info
}

func lstatFixture(t *testing.T, path string) *syscall.Stat_t {
	t.Helper()
	info := fixtureInfo(t, path)
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		t.Fatalf("stat type = %T", info.Sys())
	}
	return stat
}

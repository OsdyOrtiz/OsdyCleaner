//go:build darwin

package macos

import (
	"context"
	"encoding/binary"
	"errors"
	"sort"
	"sync"
	"syscall"

	"github.com/osdy/OsdyCleaner/internal/scan"
	"golang.org/x/sys/unix"
)

const walkerFlags = syscall.O_RDONLY | syscall.O_NOFOLLOW | syscall.O_DIRECTORY | syscall.O_CLOEXEC | syscall.O_NONBLOCK
const fileOpenFlags = syscall.O_RDONLY | syscall.O_NOFOLLOW | syscall.O_CLOEXEC | syscall.O_NONBLOCK

type walkerOps struct {
	open       func(int, string, int, uint32) (int, error)
	fstat      func(int, *syscall.Stat_t) error
	fstatfs    func(int, *syscall.Statfs_t) error
	close      func(int) error
	readDir    func(int) ([]byte, error)
	checkpoint func(string)
}
type trustedRootWalker struct{ ops walkerOps }

type acquiredRoot struct {
	facts  scan.RootFacts
	fd     int
	limits scan.WalkLimits
	ops    walkerOps
	close  func() error
	once   sync.Once
	err    error
}

func (r *acquiredRoot) Facts() scan.RootFacts { return r.facts }
func (r *acquiredRoot) Close() error          { r.once.Do(func() { r.err = r.close() }); return r.err }

// NewDescriptorWalker constructs the Darwin-only descriptor acquisition boundary.
func NewDescriptorWalker() scan.DescriptorWalker {
	return newTrustedRootWalker(walkerOps{open: walkerOpen, fstat: syscall.Fstat, fstatfs: syscall.Fstatfs, close: syscall.Close, readDir: walkerReadDir})
}
func newTrustedRootWalker(ops walkerOps) trustedRootWalker { return trustedRootWalker{ops: ops} }

const (
	directoryType = 4
	direntHeader  = 21
)

type directoryRecord struct {
	inode uint64
	kind  uint8
	name  string
}

func parseDirectoryRecords(raw []byte) ([]directoryRecord, error) {
	if len(raw) == 0 {
		return []directoryRecord{}, nil
	}
	var records []directoryRecord
	for len(raw) > 0 {
		if len(raw) < direntHeader {
			return nil, scan.ErrInvalidMetadata
		}
		recordLength := int(binary.LittleEndian.Uint16(raw[16:18]))
		nameLength := int(binary.LittleEndian.Uint16(raw[18:20]))
		if recordLength < direntHeader || recordLength > len(raw) || nameLength == 0 || nameLength > recordLength-direntHeader-1 {
			return nil, scan.ErrInvalidMetadata
		}
		nameEnd := direntHeader + nameLength
		if raw[nameEnd] != 0 {
			return nil, scan.ErrInvalidMetadata
		}
		name := string(raw[direntHeader:nameEnd])
		if name != "." && name != ".." {
			if name == "" || len(name) != nameLength || containsNULOrSlash(name) {
				return nil, scan.ErrInvalidMetadata
			}
			records = append(records, directoryRecord{binary.LittleEndian.Uint64(raw[:8]), raw[20], name})
		}
		raw = raw[recordLength:]
	}
	sort.Slice(records, func(i, j int) bool { return records[i].name < records[j].name })
	return records, nil
}
func containsNULOrSlash(name string) bool {
	for _, c := range name {
		if c == 0 || c == '/' {
			return true
		}
	}
	return false
}
func recordKind(kind uint8) scan.EntryKind {
	if kind == directoryType {
		return scan.EntryDirectory
	}
	if kind == 10 {
		return scan.EntrySymlink
	}
	if kind == 8 {
		return scan.EntryRegular
	}
	return scan.EntrySpecial
}

// walkDirectory preserves the legacy no-fact traversal contract for its direct
// DFS tests. It delegates to the same recursive engine as WalkFacts.
func (w trustedRootWalker) walkDirectory(ctx context.Context, rootFD int, root scan.RootFacts, limits scan.WalkLimits) error {
	if ctx == nil || limits.Validate() != nil {
		return scan.ErrInvalidMetadata
	}
	return w.walkDirectoryFrame(ctx, rootFD, root, limits, 0, 1, nil, nil)
}

func joinWalkerErrors(primary, next error) error {
	if primary == nil {
		return next
	}
	if next == nil {
		return primary
	}
	return errors.Join(primary, next)
}

func (w trustedRootWalker) enumerateDirectory(ctx context.Context, parent int, root scan.RootFacts) ([]scan.EnumerationFact, error) {
	if ctx == nil || w.ops.readDir == nil {
		return nil, scan.ErrInvalidMetadata
	}
	if ctx.Err() != nil {
		return nil, walkerPathError("enumerate", ".", scan.ErrInaccessible)
	}
	raw, err := w.ops.readDir(parent)
	if err != nil {
		return nil, walkerPathError("getdirentries", ".", scan.ErrInaccessible)
	}
	records, err := parseDirectoryRecords(raw)
	if err != nil {
		return nil, err
	}
	facts := make([]scan.EnumerationFact, 0, len(records))
	for _, record := range records {
		kind := recordKind(record.kind)
		if kind != scan.EntryDirectory {
			fact, _ := scan.NewEnumerationFact(record.name, kind, nil, false, scan.ErrInvalidMetadata)
			facts = append(facts, fact)
			continue
		}
		fd, openErr := w.ops.open(parent, record.name, walkerFlags, 0)
		if openErr != nil {
			fact, _ := scan.NewEnumerationFact(record.name, kind, nil, false, classifyWalkerError("openat", record.name, openErr))
			facts = append(facts, fact)
			continue
		}
		var stat syscall.Stat_t
		factErr := w.ops.fstat(fd, &stat)
		var volume syscall.Statfs_t
		if factErr == nil {
			factErr = w.ops.fstatfs(fd, &volume)
		}
		closeErr := w.ops.close(fd)
		identity := scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
		class := error(nil)
		if factErr != nil || closeErr != nil {
			class = scan.ErrInaccessible
		} else if stat.Mode&syscall.S_IFMT != syscall.S_IFDIR || identity.Validate() != nil || record.inode == 0 || identity.Inode != record.inode {
			class = scan.ErrInvalidMetadata
		} else if identity.Device != root.Device {
			class = scan.ErrDeviceBoundary
		} else if volume.Flags&localFilesystemFlag == 0 {
			class = scan.ErrNonLocal
		}
		fact, _ := scan.NewEnumerationFact(record.name, kind, &identity, class == nil, class)
		facts = append(facts, fact)
	}
	return facts, nil
}

func (w trustedRootWalker) AcquireRoot(ctx context.Context, home, root scan.AbsoluteComponents, limits scan.WalkLimits) (scan.TrustedRoot, error) {
	if ctx == nil || limits.Validate() != nil || home.ValidateScanRoot() != nil || root.ValidateScanRoot() != nil || !underHome(home.Components(), root.Components()) {
		return nil, walkerPathError("acquire", root.String(), scan.ErrInvalidMetadata)
	}
	check := func(checkpoint, path string) error {
		if w.ops.checkpoint != nil {
			w.ops.checkpoint(checkpoint)
		}
		if ctx.Err() != nil {
			return walkerPathError("acquire", path, scan.ErrInaccessible)
		}
		return nil
	}
	if err := check("before", root.String()); err != nil {
		return nil, err
	}

	fd, err := w.ops.open(-1, "/", walkerFlags, 0)
	if err != nil {
		return nil, classifyWalkerError("open", "/", err)
	}
	type ownedDescriptor struct {
		fd   int
		path string
	}
	current, currentPath := fd, "/"
	closeAll := func(primary error, descriptors ...ownedDescriptor) error {
		secondary := []error{primary}
		for _, descriptor := range descriptors {
			if closeErr := w.ops.close(descriptor.fd); closeErr != nil {
				secondary = append(secondary, walkerPathError("close", descriptor.path, scan.ErrInaccessible))
			}
		}
		if primary == nil {
			secondary = secondary[1:]
		}
		if len(secondary) == 0 {
			return nil
		}
		return errors.Join(secondary...)
	}
	if err := check("after-open", "/"); err != nil {
		return nil, closeAll(err, ownedDescriptor{current, currentPath})
	}
	if _, err = w.facts(ctx, current, "/", check); err != nil {
		return nil, closeAll(err, ownedDescriptor{current, currentPath})
	}
	if limits.MaxDescriptors < 2 {
		return nil, closeAll(walkerPathError("acquire", "/", scan.ErrInvalidMetadata), ownedDescriptor{current, currentPath})
	}

	homeParts, rootParts := home.Components(), root.Components()
	var homeFacts, rootFacts scan.RootFacts
	for i, part := range rootParts {
		path := accumulated(rootParts[:i+1])
		if err := check("before-open", path); err != nil {
			return nil, closeAll(err, ownedDescriptor{current, currentPath})
		}
		next, openErr := w.ops.open(current, part, walkerFlags, 0)
		if openErr != nil {
			return nil, closeAll(classifyWalkerError("openat", path, openErr), ownedDescriptor{current, currentPath})
		}
		if err := check("after-open", path); err != nil {
			return nil, closeAll(err, ownedDescriptor{next, path}, ownedDescriptor{current, currentPath})
		}
		facts, factErr := w.facts(ctx, next, path, check)
		if factErr != nil {
			return nil, closeAll(factErr, ownedDescriptor{next, path}, ownedDescriptor{current, currentPath})
		}
		rootFacts = facts
		if i+1 == len(homeParts) {
			homeFacts = facts
		} else if i+1 > len(homeParts) && facts.Device != homeFacts.Device {
			return nil, closeAll(walkerPathError("fstat", path, scan.ErrDeviceBoundary), ownedDescriptor{next, path}, ownedDescriptor{current, currentPath})
		}
		if err := w.ops.close(current); err != nil {
			primary := walkerPathError("close", currentPath, scan.ErrInaccessible)
			return nil, closeAll(primary, ownedDescriptor{next, path})
		}
		current, currentPath = next, path
	}
	if err := check("before-handoff", root.String()); err != nil {
		return nil, closeAll(err, ownedDescriptor{current, currentPath})
	}
	if rootFacts.Device != homeFacts.Device {
		return nil, closeAll(walkerPathError("fstat", root.String(), scan.ErrDeviceBoundary), ownedDescriptor{current, currentPath})
	}
	return &acquiredRoot{facts: rootFacts, fd: current, limits: limits, ops: w.ops, close: func() error {
		if closeErr := w.ops.close(current); closeErr != nil {
			return walkerPathError("close", root.String(), scan.ErrInaccessible)
		}
		return nil
	}}, nil
}

// WalkFacts runs the one descriptor-owned recursive engine. It passes copied
// facts synchronously and leaves the retained root close boundary to Scanner.
// Inspect reopens only job components below the retained trusted-root descriptor.
// It never exposes that descriptor or constructs a descendant absolute pathname.
func (r *acquiredRoot) Inspect(ctx context.Context, job scan.FileJob) scan.FileResult {
	result := func(kind scan.EntryKind, identity scan.Identity, cause error) scan.FileResult {
		out, err := scan.NewFileResult(job, kind, identity, 0, 0, cause)
		if err != nil {
			panic(err)
		}
		return out
	}
	if ctx == nil || ctx.Err() != nil {
		return result(scan.EntryUnknown, job.FinalIdentity(), scan.ErrInaccessible)
	}
	current := r.fd
	currentOwned := false
	closeCurrent := func(primary error) error {
		if !currentOwned {
			return primary
		}
		if err := r.ops.close(current); err != nil {
			closeErr := walkerPathError("close", joinComponents(job.Components()), scan.ErrInaccessible)
			if primary == nil {
				return closeErr
			}
			return joinWalkerErrors(primary, closeErr)
		}
		return primary
	}
	for i, name := range job.Components()[:len(job.Components())-1] {
		if ctx.Err() != nil {
			return result(scan.EntryUnknown, job.FinalIdentity(), closeCurrent(scan.ErrInaccessible))
		}
		next, err := r.ops.open(current, name, walkerFlags, 0)
		if err != nil {
			return result(scan.EntryUnknown, job.FinalIdentity(), closeCurrent(classifyChangedChildOpen(name, err)))
		}
		var stat syscall.Stat_t
		var volume syscall.Statfs_t
		cause := error(nil)
		if err := r.ops.fstat(next, &stat); err != nil {
			cause = walkerPathError("fstat", name, scan.ErrInaccessible)
		} else if err := r.ops.fstatfs(next, &volume); err != nil {
			cause = walkerPathError("fstatfs", name, scan.ErrInaccessible)
		}
		identity := scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
		if cause == nil && (entryKindFromMode(stat.Mode) != scan.EntryDirectory || identity.Validate() != nil || identity != job.Ancestors()[i]) {
			cause = walkerPathError("fstat", name, scan.ErrInvalidMetadata)
		} else if cause == nil && identity.Device != r.facts.Device {
			cause = walkerPathError("fstat", name, scan.ErrDeviceBoundary)
		} else if cause == nil && volume.Flags&localFilesystemFlag == 0 {
			cause = walkerPathError("fstatfs", name, scan.ErrNonLocal)
		}
		if cause != nil {
			if err := r.ops.close(next); err != nil {
				cause = joinWalkerErrors(cause, walkerPathError("close", name, scan.ErrInaccessible))
			}
			return result(scan.EntryUnknown, job.FinalIdentity(), closeCurrent(cause))
		}
		if err := closeCurrent(nil); err != nil {
			_ = r.ops.close(next)
			return result(scan.EntryUnknown, job.FinalIdentity(), err)
		}
		current, currentOwned = next, true
	}
	name := job.Components()[len(job.Components())-1]
	fd, err := r.ops.open(current, name, fileOpenFlags, 0)
	if err != nil {
		return result(scan.EntryUnknown, job.FinalIdentity(), closeCurrent(classifyChangedChildOpen(name, err)))
	}
	var stat syscall.Stat_t
	var volume syscall.Statfs_t
	cause := error(nil)
	if err := r.ops.fstat(fd, &stat); err != nil {
		cause = walkerPathError("fstat", name, scan.ErrInaccessible)
	} else if err := r.ops.fstatfs(fd, &volume); err != nil {
		cause = walkerPathError("fstatfs", name, scan.ErrInaccessible)
	}
	identity := scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
	if cause == nil && (entryKindFromMode(stat.Mode) != scan.EntryRegular || identity.Validate() != nil || identity != job.FinalIdentity() || stat.Size < 0 || stat.Blocks < 0 || uint64(stat.Blocks) > ^uint64(0)/512) {
		cause = walkerPathError("fstat", name, scan.ErrInvalidMetadata)
	} else if cause == nil && identity.Device != r.facts.Device {
		cause = walkerPathError("fstat", name, scan.ErrDeviceBoundary)
	} else if cause == nil && volume.Flags&localFilesystemFlag == 0 {
		cause = walkerPathError("fstatfs", name, scan.ErrNonLocal)
	}
	if closeErr := r.ops.close(fd); closeErr != nil {
		cause = joinWalkerErrors(cause, walkerPathError("close", name, scan.ErrInaccessible))
	}
	cause = closeCurrent(cause)
	if cause != nil {
		return result(scan.EntryUnknown, job.FinalIdentity(), cause)
	}
	out, err := scan.NewFileResult(job, scan.EntryRegular, identity, uint64(stat.Size), uint64(stat.Blocks)*512, nil)
	if err != nil {
		return result(scan.EntryUnknown, job.FinalIdentity(), scan.ErrInvalidMetadata)
	}
	return out
}

func (w trustedRootWalker) WalkFacts(ctx context.Context, root scan.TrustedRoot, visit func(scan.DirectoryFact) error) error {
	retained, ok := root.(*acquiredRoot)
	if !ok || visit == nil {
		return scan.ErrInvalidMetadata
	}
	limits := retained.limits
	if limits.Validate() != nil { // Test-only retained roots use this safe default.
		limits = scan.WalkLimits{MaxDepth: 64, MaxDescriptors: 96}
	}
	return w.walkDirectoryFrame(ctx, retained.fd, retained.facts, limits, 0, 1, nil, visit)
}

// walkDirectoryFrame is the sole recursive descriptor engine. A nil visitor
// preserves legacy behavior; a visitor receives synchronous scanner facts.
func (w trustedRootWalker) walkDirectoryFrame(ctx context.Context, parent int, root scan.RootFacts, limits scan.WalkLimits, depth, active int, prefix []string, visit func(scan.DirectoryFact) error) error {
	if ctx.Err() != nil {
		return walkerPathError("walk", ".", scan.ErrInaccessible)
	}
	raw, err := w.ops.readDir(parent)
	if err != nil {
		return walkerPathError("getdirentries", ".", scan.ErrInaccessible)
	}
	records, err := parseDirectoryRecords(raw)
	if err != nil {
		return err
	}
	var result error
	for _, record := range records {
		if ctx.Err() != nil {
			return joinWalkerErrors(result, walkerPathError("walk", record.name, scan.ErrInaccessible))
		}
		kind := recordKind(record.kind)
		if kind != scan.EntryDirectory && visit == nil {
			continue
		}
		components := append(append([]string(nil), prefix...), record.name)
		var enumerated *scan.Identity
		if record.inode != 0 {
			identity := scan.Identity{Device: root.Device, Inode: record.inode}
			enumerated = &identity
		}
		if kind == scan.EntryRegular {
			fd, openErr := w.ops.open(parent, record.name, fileOpenFlags, 0)
			if openErr != nil {
				if err := w.emitDirectoryFact(components, kind, enumerated, scan.EntryUnknown, nil, true, classifyChangedChildOpen(record.name, openErr), visit); err != nil {
					return err
				}
				continue
			}
			openedKind, opened, local, class := w.openedRegularFact(fd, record, root)
			if closeErr := w.ops.close(fd); closeErr != nil {
				class = joinWalkerErrors(class, walkerPathError("close", record.name, scan.ErrInaccessible))
			}
			if err := w.emitDirectoryFact(components, kind, enumerated, openedKind, opened, local, class, visit); err != nil {
				return err
			}
			continue
		}
		if kind != scan.EntryDirectory {
			if err := w.emitDirectoryFact(components, kind, enumerated, kind, nil, true, nil, visit); err != nil {
				return err
			}
			continue
		}
		if depth+1 > limits.MaxDepth || active >= limits.MaxDescriptors {
			limit := walkerPathError("openat", record.name, scan.ErrLimit)
			if visit == nil {
				result = joinWalkerErrors(result, limit)
			} else if err := w.emitDirectoryFact(components, kind, enumerated, scan.EntryUnknown, nil, true, scan.ErrLimit, visit); err != nil {
				return err
			}
			continue
		}
		fd, openErr := w.ops.open(parent, record.name, walkerFlags, 0)
		if openErr != nil {
			class := classifyChangedChildOpen(record.name, openErr)
			if visit == nil {
				result = joinWalkerErrors(result, class)
			} else if err := w.emitDirectoryFact(components, kind, enumerated, scan.EntryUnknown, nil, true, class, visit); err != nil {
				return err
			}
			continue
		}
		openedKind, opened, local, class := w.openedDirectoryFact(fd, record, root)
		if visit != nil {
			if err := w.emitDirectoryFact(components, kind, enumerated, openedKind, opened, local, class, visit); err != nil {
				if closeErr := w.ops.close(fd); closeErr != nil {
					return joinWalkerErrors(err, walkerPathError("close", record.name, scan.ErrInaccessible))
				}
				return err
			}
		}
		var childErr error
		if class == nil {
			childErr = w.walkDirectoryFrame(ctx, fd, root, limits, depth+1, active+1, components, visit)
		} else if visit == nil {
			childErr = class
		}
		if closeErr := w.ops.close(fd); closeErr != nil {
			childErr = joinWalkerErrors(childErr, walkerPathError("close", record.name, scan.ErrInaccessible))
		}
		result = joinWalkerErrors(result, childErr)
	}
	return result
}

func (w trustedRootWalker) emitDirectoryFact(components []string, enumeratedKind scan.EntryKind, enumerated *scan.Identity, openedKind scan.EntryKind, opened *scan.Identity, local bool, class error, visit func(scan.DirectoryFact) error) error {
	fact, err := scan.NewDirectoryFact(components, enumeratedKind, enumerated, openedKind, opened, local, class)
	if err != nil {
		return err
	}
	return visit(fact)
}

func (w trustedRootWalker) openedRegularFact(fd int, record directoryRecord, root scan.RootFacts) (scan.EntryKind, *scan.Identity, bool, error) {
	var stat syscall.Stat_t
	if err := w.ops.fstat(fd, &stat); err != nil {
		return scan.EntryUnknown, nil, true, walkerPathError("fstat", record.name, scan.ErrInaccessible)
	}
	identity := scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
	var opened *scan.Identity
	if identity.Validate() == nil {
		opened = &identity
	}
	kind := entryKindFromMode(stat.Mode)
	var volume syscall.Statfs_t
	if err := w.ops.fstatfs(fd, &volume); err != nil {
		return kind, opened, true, walkerPathError("fstatfs", record.name, scan.ErrInaccessible)
	}
	local := volume.Flags&localFilesystemFlag != 0
	switch {
	case kind != scan.EntryRegular || opened == nil || record.inode == 0 || identity.Inode != record.inode:
		return kind, opened, local, walkerPathError("fstat", record.name, scan.ErrInvalidMetadata)
	case identity.Device != root.Device:
		return kind, opened, local, walkerPathError("fstat", record.name, scan.ErrDeviceBoundary)
	case !local:
		return kind, opened, local, walkerPathError("fstatfs", record.name, scan.ErrNonLocal)
	default:
		return kind, opened, local, nil
	}
}

func (w trustedRootWalker) openedDirectoryFact(fd int, record directoryRecord, root scan.RootFacts) (scan.EntryKind, *scan.Identity, bool, error) {
	var stat syscall.Stat_t
	if err := w.ops.fstat(fd, &stat); err != nil {
		return scan.EntryUnknown, nil, true, walkerPathError("fstat", record.name, scan.ErrInaccessible)
	}
	identity := scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
	opened := &identity
	kind := entryKindFromMode(stat.Mode)
	var volume syscall.Statfs_t
	if err := w.ops.fstatfs(fd, &volume); err != nil {
		return kind, opened, true, walkerPathError("fstatfs", record.name, scan.ErrInaccessible)
	}
	local := volume.Flags&localFilesystemFlag != 0
	switch {
	case kind != scan.EntryDirectory || identity.Validate() != nil || record.inode == 0 || identity.Inode != record.inode:
		return kind, opened, local, walkerPathError("fstat", record.name, scan.ErrInvalidMetadata)
	case identity.Device != root.Device:
		return kind, opened, local, walkerPathError("fstat", record.name, scan.ErrDeviceBoundary)
	case !local:
		return kind, opened, local, walkerPathError("fstatfs", record.name, scan.ErrNonLocal)
	default:
		return kind, opened, local, nil
	}
}

func entryKindFromMode(mode uint16) scan.EntryKind {
	switch mode & syscall.S_IFMT {
	case syscall.S_IFDIR:
		return scan.EntryDirectory
	case syscall.S_IFREG:
		return scan.EntryRegular
	case syscall.S_IFLNK:
		return scan.EntrySymlink
	default:
		return scan.EntrySpecial
	}
}

func (w trustedRootWalker) facts(ctx context.Context, fd int, path string, check func(string, string) error) (scan.RootFacts, error) {
	if err := check("before-fstat", path); err != nil {
		return scan.RootFacts{}, err
	}
	if ctx.Err() != nil {
		return scan.RootFacts{}, walkerPathError("acquire", path, scan.ErrInaccessible)
	}
	var stat syscall.Stat_t
	if err := w.ops.fstat(fd, &stat); err != nil {
		return scan.RootFacts{}, walkerPathError("fstat", path, scan.ErrInaccessible)
	}
	if err := check("after-fstat", path); err != nil {
		return scan.RootFacts{}, err
	}
	if stat.Mode&syscall.S_IFMT != syscall.S_IFDIR {
		return scan.RootFacts{}, walkerPathError("fstat", path, scan.ErrNotDirectory)
	}
	identity := scan.Identity{Device: uint64(stat.Dev), Inode: uint64(stat.Ino)}
	if stat.Nlink == 0 || identity.Validate() != nil {
		return scan.RootFacts{}, walkerPathError("fstat", path, scan.ErrInvalidMetadata)
	}
	var volume syscall.Statfs_t
	if err := w.ops.fstatfs(fd, &volume); err != nil {
		return scan.RootFacts{}, walkerPathError("fstatfs", path, scan.ErrInaccessible)
	}
	if err := check("after-fstatfs", path); err != nil {
		return scan.RootFacts{}, err
	}
	if volume.Flags&localFilesystemFlag == 0 {
		return scan.RootFacts{}, walkerPathError("fstatfs", path, scan.ErrNonLocal)
	}
	facts, err := scan.NewRootFacts(identity, true)
	if err != nil {
		return scan.RootFacts{}, walkerPathError("fstat", path, scan.ErrInvalidMetadata)
	}
	if err := check("after-validation", path); err != nil {
		return scan.RootFacts{}, err
	}
	return facts, nil
}
func walkerReadDir(fd int) ([]byte, error) {
	buffer := make([]byte, 8192)
	var records []byte
	for {
		count, err := unix.Getdirentries(fd, buffer, nil)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			return records, nil
		}
		records = append(records, buffer[:count]...)
	}
}
func walkerOpen(parent int, name string, flags int, mode uint32) (int, error) {
	if name == "/" {
		return unix.Open(name, flags, mode)
	}
	return unix.Openat(parent, name, flags, mode)
}
func underHome(home, root []string) bool {
	if len(root) <= len(home) {
		return false
	}
	for i := range home {
		if home[i] != root[i] {
			return false
		}
	}
	return true
}
func accumulated(parts []string) string { return "/" + joinComponents(parts) }
func joinComponents(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	out := parts[0]
	for _, part := range parts[1:] {
		out += "/" + part
	}
	return out
}
func walkerPathError(operation, path string, class error) error {
	err, invalid := scan.NewPathError(operation, path, class)
	if invalid != nil {
		return scan.ErrInvalidMetadata
	}
	return err
}
func classifyWalkerError(operation, path string, err error) error {
	class := scan.ErrInaccessible
	switch {
	case errors.Is(err, syscall.ENOENT):
		class = scan.ErrMissing
	case errors.Is(err, syscall.ELOOP), errors.Is(err, syscall.EMLINK):
		class = scan.ErrSymlink
	case errors.Is(err, syscall.ENOTDIR):
		class = scan.ErrNotDirectory
	}
	return walkerPathError(operation, path, class)
}

func classifyChangedChildOpen(name string, err error) error {
	if errors.Is(err, syscall.ELOOP) || errors.Is(err, syscall.ENOTDIR) {
		return walkerPathError("openat", name, scan.ErrInvalidMetadata)
	}
	return classifyWalkerError("openat", name, err)
}

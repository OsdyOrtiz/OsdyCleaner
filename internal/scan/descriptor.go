// Package scan owns platform-neutral read-only traversal contracts.
package scan

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	ErrInvalidComponents = errors.New("invalid descriptor components")
	ErrInvalidLimits     = errors.New("invalid descriptor limits")
	ErrLimit             = errors.New("scan traversal limit reached")
	ErrMissing           = errors.New("scan path is missing")
	ErrSymlink           = errors.New("scan path is a symbolic link")
	ErrNotDirectory      = errors.New("scan path is not a directory")
	ErrInaccessible      = errors.New("scan path is inaccessible")
	ErrDeviceBoundary    = errors.New("scan path crosses a device boundary")
	ErrNonLocal          = errors.New("scan path is not on a local filesystem")
	ErrInvalidMetadata   = errors.New("scan path metadata is invalid")
	ErrUnsupported       = errors.New("scan platform is unsupported")
	ErrInvalidPipeline   = errors.New("invalid file pipeline")
	ErrPipelineClosed    = errors.New("file pipeline is closed")
)

type AbsoluteComponents struct {
	path       string
	components []string
}

func NewAbsoluteComponents(path string) (AbsoluteComponents, error) {
	if path == "/" {
		return AbsoluteComponents{path: path}, nil
	}
	if len(path) < 2 || path[0] != '/' || path[len(path)-1] == '/' {
		return AbsoluteComponents{}, ErrInvalidComponents
	}
	parts := strings.Split(path[1:], "/")
	if _, err := NewRelativeComponents(parts); err != nil {
		return AbsoluteComponents{}, err
	}
	return AbsoluteComponents{path: path, components: append([]string(nil), parts...)}, nil
}
func (p AbsoluteComponents) String() string       { return p.path }
func (p AbsoluteComponents) Components() []string { return append([]string(nil), p.components...) }

func (p AbsoluteComponents) ValidateScanRoot() error {
	if p.path == "" || p.path == "/" {
		return ErrInvalidComponents
	}
	return nil
}

type RelativeComponents struct{ components []string }

func NewRelativeComponents(parts []string) (RelativeComponents, error) {
	if len(parts) == 0 {
		return RelativeComponents{}, ErrInvalidComponents
	}
	for _, part := range parts {
		if part == "" || part == "." || part == ".." || strings.ContainsAny(part, "/\x00") {
			return RelativeComponents{}, ErrInvalidComponents
		}
	}
	return RelativeComponents{components: append([]string(nil), parts...)}, nil
}
func (p RelativeComponents) Components() []string { return append([]string(nil), p.components...) }

type WalkLimits struct {
	// MaxDepth is the maximum DFS component depth below an acquired scan root.
	MaxDepth int
	// MaxDescriptors is the maximum simultaneously active opaque capabilities.
	MaxDescriptors                                        int
	Workers, JobCapacity, ResultCapacity, ConcurrentRoots int
	MaxEntries, MaxPathBytes                              int
}

// DefaultWalkLimits is the single scanner policy default. Its fields are passed
// unchanged to descriptor traversal, even where a platform does not use one yet.
func DefaultWalkLimits() WalkLimits {
	return WalkLimits{MaxDepth: 64, MaxDescriptors: 96, Workers: 8, JobCapacity: 256, ResultCapacity: 256, ConcurrentRoots: 1, MaxEntries: 250000, MaxPathBytes: 67108864}
}

// NewWalkLimits rejects incomplete policy input. Validate retains legacy two-field
// literals until their platform call sites can be migrated without widening WU6A.
func NewWalkLimits(limits WalkLimits) (WalkLimits, error) {
	if limits.Validate() != nil || limits.Workers < 1 || limits.JobCapacity < 1 || limits.ResultCapacity < 1 || limits.ConcurrentRoots != 1 || limits.MaxEntries < 1 || limits.MaxPathBytes < 1 {
		return WalkLimits{}, ErrInvalidLimits
	}
	return limits, nil
}

func (l WalkLimits) Validate() error {
	if l.MaxDepth < 1 || l.MaxDescriptors < 1 {
		return ErrInvalidLimits
	}
	policyFields := []int{l.Workers, l.JobCapacity, l.ResultCapacity, l.ConcurrentRoots, l.MaxEntries, l.MaxPathBytes}
	allZero := true
	for _, field := range policyFields {
		if field != 0 {
			allZero = false
		}
		if field < 0 {
			return ErrInvalidLimits
		}
	}
	if allZero {
		return nil
	}
	if _, err := validateFullWalkLimits(l); err != nil {
		return err
	}
	return nil
}

func validateFullWalkLimits(l WalkLimits) (WalkLimits, error) {
	if l.Workers < 1 || l.JobCapacity < 1 || l.ResultCapacity < 1 || l.ConcurrentRoots != 1 || l.MaxEntries < 1 || l.MaxPathBytes < 1 {
		return WalkLimits{}, ErrInvalidLimits
	}
	return l, nil
}

type Identity struct{ Device, Inode uint64 }

func (i Identity) Validate() error {
	if i.Device == 0 || i.Inode == 0 {
		return ErrInvalidMetadata
	}
	return nil
}

type EntryKind string

const (
	EntryDirectory EntryKind = "directory"
	EntryRegular   EntryKind = "regular"
	EntrySymlink   EntryKind = "symlink"
	EntrySpecial   EntryKind = "special"
	EntryUnknown   EntryKind = "unknown"
)

type EnumerationFact struct {
	name                  string
	kind                  EntryKind
	identity              Identity
	hasIdentity, eligible bool
	skip                  error
}

func NewEnumerationFact(name string, kind EntryKind, identity *Identity, eligible bool, skip error) (EnumerationFact, error) {
	if name == "" || name == "." || name == ".." || strings.ContainsAny(name, "/\x00") || (kind != EntryDirectory && kind != EntryRegular && kind != EntrySymlink && kind != EntrySpecial && kind != EntryUnknown) || (eligible && (kind != EntryDirectory || identity == nil || identity.Validate() != nil)) || (!eligible && skip == nil) {
		return EnumerationFact{}, ErrInvalidMetadata
	}
	fact := EnumerationFact{name: name, kind: kind, eligible: eligible, skip: skip}
	if identity != nil {
		fact.identity, fact.hasIdentity = *identity, true
	}
	return fact, nil
}
func (f EnumerationFact) Name() string       { return f.name }
func (f EnumerationFact) Kind() EntryKind    { return f.kind }
func (f EnumerationFact) Identity() Identity { return f.identity }
func (f EnumerationFact) HasIdentity() bool  { return f.hasIdentity }
func (f EnumerationFact) Eligible() bool     { return f.eligible }
func (f EnumerationFact) SkipClass() error   { return f.skip }

// DirectoryFact is synchronous opaque traversal evidence. Enumerated and opened
// facts are deliberately separate because a directory record never authorizes
// accepting the object eventually opened from that bound descriptor.
type DirectoryFact struct {
	components                               []string
	enumeratedKind, openedKind               EntryKind
	enumeratedIdentity, openedIdentity       Identity
	enumeratedHasIdentity, openedHasIdentity bool
	local                                    bool
	skip                                     error
}

func NewDirectoryFact(components []string, enumeratedKind EntryKind, enumeratedIdentity *Identity, openedKind EntryKind, openedIdentity *Identity, local bool, skip error) (DirectoryFact, error) {
	if _, err := NewRelativeComponents(components); err != nil || !validEntryKind(enumeratedKind) || !validEntryKind(openedKind) {
		return DirectoryFact{}, ErrInvalidMetadata
	}
	fact := DirectoryFact{components: append([]string(nil), components...), enumeratedKind: enumeratedKind, openedKind: openedKind, local: local, skip: skip}
	if enumeratedIdentity != nil {
		fact.enumeratedIdentity, fact.enumeratedHasIdentity = *enumeratedIdentity, true
	}
	if openedIdentity != nil {
		fact.openedIdentity, fact.openedHasIdentity = *openedIdentity, true
	}
	return fact, nil
}
func validEntryKind(kind EntryKind) bool {
	return kind == EntryDirectory || kind == EntryRegular || kind == EntrySymlink || kind == EntrySpecial || kind == EntryUnknown
}
func (f DirectoryFact) Components() []string      { return append([]string(nil), f.components...) }
func (f DirectoryFact) EnumeratedKind() EntryKind { return f.enumeratedKind }
func (f DirectoryFact) OpenedKind() EntryKind     { return f.openedKind }
func (f DirectoryFact) EnumeratedIdentity() (Identity, bool) {
	return f.enumeratedIdentity, f.enumeratedHasIdentity
}
func (f DirectoryFact) OpenedIdentity() (Identity, bool) {
	return f.openedIdentity, f.openedHasIdentity
}
func (f DirectoryFact) Local() bool      { return f.local }
func (f DirectoryFact) SkipClass() error { return f.skip }

type RootFacts struct {
	Device   uint64
	Identity Identity
	Local    bool
}

func NewRootFacts(identity Identity, local bool) (RootFacts, error) {
	facts := RootFacts{Device: identity.Device, Identity: identity, Local: local}
	return facts, facts.Validate()
}

func (f RootFacts) Validate() error {
	if err := f.Identity.Validate(); err != nil || f.Device == 0 || f.Device != f.Identity.Device {
		return ErrInvalidMetadata
	}
	if !f.Local {
		return ErrNonLocal
	}
	return nil
}

type TrustedRoot interface {
	Facts() RootFacts
	Close() error
}
type DescriptorWalker interface {
	AcquireRoot(context.Context, AbsoluteComponents, AbsoluteComponents, WalkLimits) (TrustedRoot, error)
	WalkFacts(context.Context, TrustedRoot, func(DirectoryFact) error) error
}

type trustedRoot struct {
	facts RootFacts
	close func() error
	once  sync.Once
	err   error
}

// NewTrustedRoot transfers one opaque root capability and its close boundary.
// The callback may retain platform resources but never exposes them to scan consumers.
func NewTrustedRoot(facts RootFacts, close func() error) (TrustedRoot, error) {
	if err := facts.Validate(); err != nil || close == nil {
		return nil, ErrInvalidMetadata
	}
	return &trustedRoot{facts: facts, close: close}, nil
}

func (r *trustedRoot) Facts() RootFacts { return r.facts }

func (r *trustedRoot) Close() error {
	r.once.Do(func() { r.err = r.close() })
	return r.err
}

// FileJob is a descriptor-free, relative file reopen request. Ancestors name
// the directories before the final component and must match that depth.
type FileJob struct {
	components []string
	ancestors  []Identity
	final      Identity
}

func NewFileJob(components []string, ancestors []Identity, final Identity) (FileJob, error) {
	if _, err := NewRelativeComponents(components); err != nil || len(ancestors) != len(components)-1 || final.Validate() != nil {
		return FileJob{}, ErrInvalidMetadata
	}
	for _, identity := range ancestors {
		if identity.Validate() != nil {
			return FileJob{}, ErrInvalidMetadata
		}
	}
	return FileJob{components: append([]string(nil), components...), ancestors: append([]Identity(nil), ancestors...), final: final}, nil
}
func (j FileJob) Components() []string    { return append([]string(nil), j.components...) }
func (j FileJob) Ancestors() []Identity   { return append([]Identity(nil), j.ancestors...) }
func (j FileJob) FinalIdentity() Identity { return j.final }

// FileResult carries metadata only; only an unchanged regular result is accepted.
type FileResult struct {
	job                 FileJob
	kind                EntryKind
	identity            Identity
	logical, allocation uint64
	err                 error
}

func NewFileResult(job FileJob, kind EntryKind, identity Identity, logical, allocation uint64, cause error) (FileResult, error) {
	if _, err := NewFileJob(job.components, job.ancestors, job.final); err != nil || !validEntryKind(kind) || identity.Validate() != nil {
		return FileResult{}, ErrInvalidMetadata
	}
	result := FileResult{job: job, kind: kind, identity: identity, err: cause}
	if cause == nil && kind == EntryRegular && identity == job.final {
		result.logical, result.allocation = logical, allocation
	}
	return result, nil
}
func (r FileResult) Components() []string    { return r.job.Components() }
func (r FileResult) LogicalBytes() uint64    { return r.logical }
func (r FileResult) AllocationBytes() uint64 { return r.allocation }
func (r FileResult) Error() error            { return r.err }
func (r FileResult) Accepted() bool {
	return r.err == nil && r.kind == EntryRegular && r.identity == r.job.final
}

type FilePort interface {
	Inspect(context.Context, FileJob) FileResult
}

type PathError struct {
	operation, path string
	err             error
}

func NewPathError(operation, path string, class error) (*PathError, error) {
	if operation == "" || path == "" || class == nil {
		return nil, ErrInvalidMetadata
	}
	return &PathError{operation: operation, path: path, err: class}, nil
}
func (e *PathError) Error() string     { return fmt.Sprintf("scan %s %q: %v", e.operation, e.path, e.err) }
func (e *PathError) Unwrap() error     { return e.err }
func (e *PathError) Operation() string { return e.operation }
func (e *PathError) Path() string      { return e.path }

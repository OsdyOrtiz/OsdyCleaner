package scan

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// fileEnvelope retains a submitter context without exposing it through FileJob.
type fileEnvelope struct {
	job FileJob
	ctx context.Context
}

// FilePipeline owns a bounded, fixed worker pool over opaque file jobs.
type FilePipeline struct {
	port                                 FilePort
	workers, jobCapacity, resultCapacity int
	jobs                                 chan fileEnvelope
	results                              chan FileResult
	workersWG, submittersWG              sync.WaitGroup
	closeOnce                            sync.Once
	mu                                   sync.Mutex
	closed                               bool
	submitCtx, stopCtx                   context.Context
	cancelSubmit, cancelStop             context.CancelFunc
	stopOnContextCancel                  bool
	stopOnce                             sync.Once
}

func NewFilePipeline(port FilePort, workers, jobCapacity, resultCapacity int) (*FilePipeline, error) {
	return newFilePipeline(port, workers, jobCapacity, resultCapacity, false)
}

func newFilePipeline(port FilePort, workers, jobCapacity, resultCapacity int, stopOnContextCancel bool) (*FilePipeline, error) {
	if port == nil || workers < 1 || jobCapacity < 1 || resultCapacity < 1 {
		return nil, ErrInvalidPipeline
	}
	submitCtx, cancelSubmit := context.WithCancel(context.Background())
	stopCtx, cancelStop := context.WithCancel(context.Background())
	p := &FilePipeline{port: port, workers: workers, jobCapacity: jobCapacity, resultCapacity: resultCapacity, jobs: make(chan fileEnvelope, jobCapacity), results: make(chan FileResult, resultCapacity), submitCtx: submitCtx, cancelSubmit: cancelSubmit, stopCtx: stopCtx, cancelStop: cancelStop, stopOnContextCancel: stopOnContextCancel}
	p.workersWG.Add(workers)
	for range workers {
		go func() {
			defer p.workersWG.Done()
			for {
				if p.stopCtx.Err() != nil {
					return
				}
				select {
				case <-p.stopCtx.Done():
					return
				case envelope, ok := <-p.jobs:
					// A stop racing a ready job must win before inspection.
					if !ok || p.stopCtx.Err() != nil {
						return
					}
					result := p.port.Inspect(envelope.ctx, envelope.job)
					p.results <- result
					if p.stopCtx.Err() != nil || p.stopOnContextCancel && envelope.ctx.Err() != nil {
						return
					}
				}
			}
		}()
	}
	return p, nil
}
func (p *FilePipeline) Workers() int        { return p.workers }
func (p *FilePipeline) JobCapacity() int    { return p.jobCapacity }
func (p *FilePipeline) ResultCapacity() int { return p.resultCapacity }
func (p *FilePipeline) Submit(ctx context.Context, job FileJob) error {
	if ctx == nil {
		return ErrInvalidPipeline
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if _, err := NewFileJob(job.components, job.ancestors, job.final); err != nil {
		return err
	}
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return ErrPipelineClosed
	}
	p.submittersWG.Add(1)
	p.mu.Unlock()
	defer p.submittersWG.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-p.submitCtx.Done():
		return ErrPipelineClosed
	case p.jobs <- fileEnvelope{job: job, ctx: ctx}:
		return nil
	}
}

// CloseJobs prevents new submissions and asynchronously closes owned channels.
// Drain must consume results so workers blocked on result pressure can join.
func (p *FilePipeline) CloseJobs() {
	p.closeOnce.Do(func() {
		p.mu.Lock()
		p.closed = true
		p.cancelSubmit()
		p.mu.Unlock()
		go func() {
			p.submittersWG.Wait()
			close(p.jobs)
			p.workersWG.Wait()
			close(p.results)
		}()
	})
}

// Stop prevents workers from accepting queued work while preserving any active result.
func (p *FilePipeline) Stop() {
	p.stopOnce.Do(p.cancelStop)
	p.CloseJobs()
}

func (p *FilePipeline) Drain() []FileResult {
	return normalizeFileResults(p.drain())
}

func (p *FilePipeline) drain() []FileResult {
	var results []FileResult
	for result := range p.results {
		results = append(results, result)
	}
	return results
}
func normalizeFileResults(results []FileResult) []FileResult {
	accepted := results[:0]
	for _, result := range results {
		if result.Accepted() {
			accepted = append(accepted, result)
		}
	}
	sort.Slice(accepted, func(i, j int) bool {
		return strings.Join(accepted[i].Components(), "/") < strings.Join(accepted[j].Components(), "/")
	})
	return accepted
}

type scannerFileResults struct {
	mu       sync.RWMutex
	accepted []FileResult
	raw      []rawRoot
	roots    []core.RootObservation
}

func (r *scannerFileResults) reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.accepted = nil
	r.raw = nil
	r.roots = nil
}

func (r *scannerFileResults) store(definition BuiltinDefinition, results []FileResult) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.accepted = append(r.accepted, copyFileResults(results)...)
	r.accepted = normalizeFileResults(r.accepted)
	r.raw = append(r.raw, rawRoot{area: definition.AreaID, displayName: definition.DisplayName, displayPath: definition.DisplayPath, observations: rawFileObservations(definition, results)})
}

func (r *scannerFileResults) setRoots(roots []core.RootObservation) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.roots = append([]core.RootObservation(nil), roots...)
}

func (r *scannerFileResults) finalized() (FinalizedFileFacts, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	facts, err := finalizeRawRoots(append([]rawRoot(nil), r.raw...))
	if err != nil {
		return FinalizedFileFacts{}, err
	}
	generated := make(map[core.AreaID]core.RootObservation, len(facts.roots))
	for _, root := range facts.roots {
		generated[root.AreaID()] = root
	}
	merged := make([]core.RootObservation, 0, len(r.roots))
	for _, observed := range r.roots {
		if finalized, ok := generated[observed.AreaID()]; ok {
			estimate := finalized.Estimate()
			if observed.Status() != core.RootScanned {
				estimate, err = incompleteEstimate(estimate)
				if err != nil {
					return FinalizedFileFacts{}, err
				}
			}
			root, err := core.NewRootObservation(observed.AreaID(), observed.DisplayName(), observed.DisplayPath(), observed.Status(), observed.Reason(), estimate, finalized.FindingCount(), observed.WarningCodes())
			if err != nil {
				return FinalizedFileFacts{}, err
			}
			merged = append(merged, root)
			continue
		}
		merged = append(merged, observed)
	}
	facts.roots = merged
	return FinalizedFileFacts{facts: facts}, nil
}

func incompleteEstimate(estimate core.Estimate) (core.Estimate, error) {
	logical, err := core.NewValueEstimate(estimate.Logical().KnownBytes(), core.BasisRegularFileStatSize, core.CompletenessIncomplete)
	if err != nil {
		return core.Estimate{}, err
	}
	allocatedCompleteness := core.CompletenessIncomplete
	allocatedBytes := estimate.Allocated().KnownBytes()
	if estimate.Allocated().Completeness() == core.CompletenessUnknown {
		allocatedCompleteness, allocatedBytes = core.CompletenessUnknown, 0
	}
	allocated, err := core.NewValueEstimate(allocatedBytes, core.BasisStatBlocks512, allocatedCompleteness)
	if err != nil {
		return core.Estimate{}, err
	}
	return core.NewEstimate(logical, allocated, estimate.DeduplicationComplete())
}

func rawFileObservations(definition BuiltinDefinition, results []FileResult) []rawObservation {
	accepted := normalizeFileResults(append([]FileResult(nil), results...))
	observations := make([]rawObservation, 0, len(accepted))
	for _, result := range accepted {
		identity, err := core.NewFilesystemIdentity(result.identity.Device, result.identity.Inode)
		if err != nil {
			continue
		}
		observations = append(observations, rawObservation{displayPath: definition.DisplayPath + "/" + strings.Join(result.Components(), "/"), reason: definition.FindingReason, logicalBytes: result.LogicalBytes(), allocation: core.KnownAllocation(result.AllocationBytes()), identity: core.KnownFilesystemIdentity(identity)})
	}
	return observations
}

func (r *scannerFileResults) copy() []FileResult {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return copyFileResults(r.accepted)
}

func copyFileResults(results []FileResult) []FileResult {
	out := make([]FileResult, len(results))
	for i, result := range results {
		out[i] = result
		out[i].job.components = append([]string(nil), result.job.components...)
		out[i].job.ancestors = append([]Identity(nil), result.job.ancestors...)
	}
	return out
}

// Scanner consumes the fixed built-in roots through the opaque platform walker.
// Accepted file results are scanner-owned evidence until a later finalizer consumes them.
type Scanner struct {
	home        HomeResolver
	definitions []BuiltinDefinition
	walker      DescriptorWalker
	limits      WalkLimits
	results     *scannerFileResults
	fileWorkers int
}

func NewScanner(home HomeResolver, definitions []BuiltinDefinition, walker DescriptorWalker, limits WalkLimits) Scanner {
	limits = scannerLimits(limits)
	return Scanner{home: home, definitions: append([]BuiltinDefinition(nil), definitions...), walker: walker, limits: limits, results: &scannerFileResults{}, fileWorkers: limits.Workers}
}

func scannerLimits(limits WalkLimits) WalkLimits {
	if limits.Workers != 0 || limits.JobCapacity != 0 || limits.ResultCapacity != 0 || limits.ConcurrentRoots != 0 || limits.MaxEntries != 0 || limits.MaxPathBytes != 0 {
		return limits
	}
	defaults := DefaultWalkLimits()
	defaults.MaxDepth, defaults.MaxDescriptors = limits.MaxDepth, limits.MaxDescriptors
	// Existing callers supplied only traversal limits before WU6A. Preserve their
	// one-worker queue behavior while giving their traversal bounded defaults.
	defaults.Workers, defaults.JobCapacity, defaults.ResultCapacity = 1, 1, 1
	return defaults
}

// Policy returns the immutable scanner policy passed to descriptor traversal.
func (s Scanner) Policy() WalkLimits { return s.limits }

// AcceptedFileResults returns a deep copy of accepted results from the most recent Scan.
func (s Scanner) AcceptedFileResults() []FileResult {
	if s.results == nil {
		return nil
	}
	return s.results.copy()
}

// FinalizedFileFacts converts retained accepted results through the canonical finalizer.
func (s Scanner) FinalizedFileFacts() (FinalizedFileFacts, error) {
	if s.results == nil {
		return FinalizedFileFacts{}, nil
	}
	return s.results.finalized()
}

func (s Scanner) Scan(ctx context.Context) ([]core.RootObservation, error) {
	if s.results != nil {
		s.results.reset()
	}
	home, definitions, err := s.preflight()
	if err != nil {
		return nil, err
	}
	observations := make([]core.RootObservation, 0, len(definitions))
	for _, definition := range definitions {
		if ctx.Err() != nil {
			observations = append(observations, cancelledBefore(definition))
			continue
		}
		root, err := s.walker.AcquireRoot(ctx, home, absoluteUnder(home.String(), definition.RelativePath), s.limits)
		if err != nil {
			if errors.Is(err, context.Canceled) || ctx.Err() != nil {
				observations = append(observations, cancelledDuring(definition))
				continue
			}
			observations = append(observations, acquisitionObservation(definition, err))
			continue
		}
		observation := s.consumeDirectories(ctx, definition, root)
		closeErr := root.Close()
		if ctx.Err() != nil {
			codes := observation.WarningCodes()
			if closeErr != nil {
				codes = append(codes, core.WarningRootInaccessible)
			}
			observations = append(observations, cancelledDuringWithWarnings(definition, codes))
			continue
		}
		if closeErr != nil {
			observations = append(observations, inaccessible(definition))
			continue
		}
		observations = append(observations, observation)
	}
	if s.results != nil {
		s.results.setRoots(observations)
	}
	return observations, nil
}

func (s Scanner) consumeDirectories(ctx context.Context, definition BuiltinDefinition, root TrustedRoot) core.RootObservation {
	var codes []core.WarningCode
	partial, boundary := false, false
	add := func(code core.WarningCode) { codes = append(codes, code) }

	// A retained platform root may also provide the opaque relative-file reopen port.
	// Scanner never receives its descriptor or an absolute descendant path.
	port, hasFilePort := root.(FilePort)
	ancestors := make(map[string]Identity)
	var pipeline *FilePipeline
	var drained <-chan []FileResult
	var pipelineDone chan struct{}
	if hasFilePort {
		workers := s.fileWorkers
		if workers < 1 {
			workers = 1
		}
		pipeline, _ = newFilePipeline(port, workers, s.limits.JobCapacity, s.limits.ResultCapacity, true)
		pipelineDone = make(chan struct{})
		go func() {
			select {
			case <-ctx.Done():
				pipeline.Stop()
			case <-pipelineDone:
			}
		}()
		results := make(chan []FileResult, 1)
		go func() { results <- pipeline.drain() }()
		drained = results
	}
	fileJob := func(fact DirectoryFact) (FileJob, error) {
		components := fact.Components()
		enumerated, hasEnumerated := fact.EnumeratedIdentity()
		opened, hasOpened := fact.OpenedIdentity()
		if fact.SkipClass() != nil || fact.EnumeratedKind() != EntryRegular || fact.OpenedKind() != EntryRegular || !hasEnumerated || !hasOpened || enumerated.Validate() != nil || opened.Validate() != nil || enumerated != opened || !fact.Local() || opened.Device != root.Facts().Device {
			return FileJob{}, ErrInvalidMetadata
		}
		chain := make([]Identity, 0, len(components)-1)
		for i := 1; i < len(components); i++ {
			identity, ok := ancestors[strings.Join(components[:i], "/")]
			if !ok {
				return FileJob{}, ErrInvalidMetadata
			}
			chain = append(chain, identity)
		}
		return NewFileJob(components, chain, opened)
	}
	entries, pathBytes := 0, 0
	limitReached := false
	walkErr := s.walker.WalkFacts(ctx, root, func(fact DirectoryFact) error {
		if err := ctx.Err(); err != nil {
			return err
		}
		if s.limits.MaxEntries > 0 && entries >= s.limits.MaxEntries {
			add(core.WarningEntryLimitReached)
			limitReached, partial = true, true
			return ErrLimit
		}
		pathBytes += len(strings.Join(fact.Components(), "/"))
		if s.limits.MaxPathBytes > 0 && pathBytes > s.limits.MaxPathBytes {
			add(core.WarningPathBudgetReached)
			limitReached, partial = true, true
			return ErrLimit
		}
		entries++
		if skip := fact.SkipClass(); skip != nil {
			switch {
			case errors.Is(skip, ErrDeviceBoundary), errors.Is(skip, ErrNonLocal):
				add(core.WarningDeviceBoundary)
				boundary = true
			case errors.Is(skip, ErrSymlink):
				// A skipped entry symlink points outside the observable root scope.
				add(core.WarningSymlinkSkipped)
			case errors.Is(skip, ErrNotDirectory):
				add(core.WarningUnsupportedEntryType)
				partial = true
			default:
				add(core.WarningEntryChanged)
				partial = true
			}
			return nil
		}
		enumerated, hasEnumerated := fact.EnumeratedIdentity()
		opened, hasOpened := fact.OpenedIdentity()
		switch fact.EnumeratedKind() {
		case EntryDirectory:
			if fact.OpenedKind() != EntryDirectory || !hasEnumerated || !hasOpened || enumerated.Validate() != nil || opened.Validate() != nil || enumerated != opened {
				add(core.WarningEntryChanged)
				partial = true
			} else if !fact.Local() || opened.Device != root.Facts().Device {
				add(core.WarningDeviceBoundary)
				boundary = true
			} else {
				ancestors[strings.Join(fact.Components(), "/")] = opened
			}
		case EntryRegular:
			if !hasFilePort {
				add(core.WarningEntryChanged)
				partial = true
				return nil
			}
			job, err := fileJob(fact)
			if err != nil || pipeline.Submit(ctx, job) != nil {
				add(core.WarningEntryChanged)
				partial = true
			}
		case EntrySymlink:
			// A directly enumerated symlink is permitted but always surfaced.
			add(core.WarningSymlinkSkipped)
		case EntrySpecial, EntryUnknown:
			add(core.WarningUnsupportedEntryType)
			partial = true
		default:
			add(core.WarningEntryChanged)
			partial = true
		}
		return nil
	})
	if hasFilePort {
		if ctx.Err() != nil {
			pipeline.Stop()
		} else {
			pipeline.CloseJobs()
		}
		// Draining completes before this method returns, so Scan closes the root
		// only after active work joins and every accepted result is consumed.
		results := <-drained
		close(pipelineDone)
		consumeFileResults(definition, results, s.results, add, &partial, &boundary)
	}
	if walkErr != nil {
		switch {
		case errors.Is(walkErr, ErrLimit):
			if !limitReached {
				add(core.WarningEntryLimitReached)
			}
			partial = true
		case errors.Is(walkErr, ErrDeviceBoundary), errors.Is(walkErr, ErrNonLocal):
			add(core.WarningDeviceBoundary)
			boundary = true
		case errors.Is(walkErr, ErrSymlink):
			// Walker-level symlink classification denotes a permitted skipped entry.
			add(core.WarningSymlinkSkipped)
		case errors.Is(walkErr, ErrMissing), errors.Is(walkErr, ErrInvalidMetadata):
			add(core.WarningEntryChanged)
			partial = true
		case errors.Is(walkErr, ErrNotDirectory):
			add(core.WarningUnsupportedEntryType)
			partial = true
		default:
			add(core.WarningEntryInaccessible)
			partial = true
		}
	}
	sort.Slice(codes, func(i, j int) bool { return codes[i] < codes[j] })
	codes = uniqueWarningCodes(codes)
	status := rootStatus(false, partial, boundary)
	reason := core.ReasonCompleted
	if status == core.RootPartial {
		reason = core.ReasonEntryVisibilityGap
	} else if status == core.RootBoundaryLimited {
		reason = core.ReasonDeviceBoundary
	}
	return finalizeDirectoryObservation(definition, status, reason, codes)
}

func consumeFileResults(definition BuiltinDefinition, results []FileResult, retained *scannerFileResults, add func(core.WarningCode), partial, boundary *bool) {
	accepted := normalizeFileResults(append([]FileResult(nil), results...))
	if retained != nil {
		retained.store(definition, accepted)
	}
	for _, result := range results {
		if result.Accepted() {
			continue
		}
		switch {
		case errors.Is(result.Error(), ErrDeviceBoundary), errors.Is(result.Error(), ErrNonLocal):
			add(core.WarningDeviceBoundary)
			*boundary = true
		case errors.Is(result.Error(), ErrSymlink):
			add(core.WarningSymlinkSkipped)
			*partial = true
		case errors.Is(result.Error(), ErrNotDirectory):
			add(core.WarningUnsupportedEntryType)
			*partial = true
		default:
			add(core.WarningEntryChanged)
			*partial = true
		}
	}
}

func directoryStatus(err error, fileEvidence bool) core.RootStatus {
	boundary := errors.Is(err, ErrDeviceBoundary) || errors.Is(err, ErrNonLocal)
	partial := !fileEvidence || (err != nil && (!boundary || errors.Is(err, ErrInaccessible) || errors.Is(err, ErrMissing) || errors.Is(err, ErrInvalidMetadata) || errors.Is(err, ErrNotDirectory)))
	return directoryStatusFlags(partial, boundary)
}

func rootStatus(cancelled, partial, boundary bool) core.RootStatus {
	if cancelled {
		return core.RootCancelled
	}
	if partial {
		return core.RootPartial
	}
	if boundary {
		return core.RootBoundaryLimited
	}
	return core.RootScanned
}

func directoryStatusFlags(partial, boundary bool) core.RootStatus {
	return rootStatus(false, partial, boundary)
}

func uniqueWarningCodes(codes []core.WarningCode) []core.WarningCode {
	out := codes[:0]
	for _, code := range codes {
		if len(out) == 0 || out[len(out)-1] != code {
			out = append(out, code)
		}
	}
	return out
}

func (s Scanner) preflight() (AbsoluteComponents, []BuiltinDefinition, error) {
	if s.home == nil || s.walker == nil || s.limits.Validate() != nil {
		return AbsoluteComponents{}, nil, ErrInvalidMetadata
	}
	homePath, err := s.home.Home()
	if err != nil {
		return AbsoluteComponents{}, nil, fmt.Errorf("resolve home: %w", err)
	}
	home, err := NewAbsoluteComponents(homePath)
	if err != nil {
		return AbsoluteComponents{}, nil, err
	}
	if err := canonicalDefinitions(s.definitions); err != nil {
		return AbsoluteComponents{}, nil, err
	}
	return home, append([]BuiltinDefinition(nil), s.definitions...), nil
}

func canonicalDefinitions(definitions []BuiltinDefinition) error {
	if len(definitions) != len(builtinDefinitions) {
		return ErrInvalidComponents
	}
	for i, definition := range definitions {
		want := builtinDefinitions[i]
		if definition.AreaID != want.AreaID || definition.DisplayName != want.DisplayName || definition.DisplayPath != want.DisplayPath || definition.RelativePath != want.RelativePath || definition.FindingReason != want.FindingReason {
			return ErrInvalidComponents
		}
		if _, err := NewRelativeComponents(strings.Split(definition.RelativePath, "/")); err != nil {
			return err
		}
	}
	return nil
}

func absoluteUnder(home, relative string) AbsoluteComponents {
	path := home + "/" + relative
	if home == "/" {
		path = "/" + relative
	}
	root, _ := NewAbsoluteComponents(path)
	return root
}

func incompleteZeroEstimate() core.Estimate {
	logical, _ := core.NewValueEstimate(0, core.BasisRegularFileStatSize, core.CompletenessIncomplete)
	allocated, _ := core.NewValueEstimate(0, core.BasisStatBlocks512, core.CompletenessIncomplete)
	estimate, _ := core.NewEstimate(logical, allocated, true)
	return estimate
}

func rootObservation(definition BuiltinDefinition, status core.RootStatus, reason core.RootReasonCode, codes []core.WarningCode) core.RootObservation {
	observation, _ := core.NewRootObservation(definition.AreaID, definition.DisplayName, definition.DisplayPath, status, reason, incompleteZeroEstimate(), 0, codes)
	return observation
}

func cancelledDuring(definition BuiltinDefinition) core.RootObservation {
	return cancelledDuringWithWarnings(definition, nil)
}

func cancelledDuringWithWarnings(definition BuiltinDefinition, codes []core.WarningCode) core.RootObservation {
	codes = append(append([]core.WarningCode(nil), codes...), core.WarningCancelledDuringRoot)
	return rootObservation(definition, core.RootCancelled, core.ReasonCancelledDuringScan, codes)
}

func cancelledBefore(definition BuiltinDefinition) core.RootObservation {
	return rootObservation(definition, core.RootSkipped, core.ReasonCancelledBeforeStart, []core.WarningCode{core.WarningCancelledBeforeRoot})
}

func inaccessible(definition BuiltinDefinition) core.RootObservation {
	return rootObservation(definition, core.RootInaccessible, core.ReasonRootInspectionFailed, []core.WarningCode{core.WarningRootInaccessible})
}

func acquisitionObservation(definition BuiltinDefinition, err error) core.RootObservation {
	switch {
	case errors.Is(err, ErrMissing):
		return rootObservation(definition, core.RootMissing, core.ReasonNotFound, nil)
	case errors.Is(err, ErrSymlink):
		return rootObservation(definition, core.RootSkipped, core.ReasonRootSymlink, []core.WarningCode{core.WarningRootSymlink})
	case errors.Is(err, ErrNotDirectory):
		return rootObservation(definition, core.RootSkipped, core.ReasonRootNotDirectory, []core.WarningCode{core.WarningRootNotDirectory})
	case errors.Is(err, ErrDeviceBoundary), errors.Is(err, ErrNonLocal):
		return rootObservation(definition, core.RootSkipped, core.ReasonRootDeviceBoundary, []core.WarningCode{core.WarningRootDeviceBoundary})
	default:
		return inaccessible(definition)
	}
}

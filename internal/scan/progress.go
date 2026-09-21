package scan

import (
	"fmt"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// ProgressEventKind identifies a scanner lifecycle event.
type ProgressEventKind string

const (
	ProgressScanStarted       ProgressEventKind = "scan-started"
	ProgressCategoryStarted   ProgressEventKind = "category-started"
	ProgressCategoryCompleted ProgressEventKind = "category-completed"
	ProgressScanCancelled     ProgressEventKind = "scan-cancelled"
	ProgressScanFinished      ProgressEventKind = "scan-finished"
)

// Validate reports whether kind is a defined progress event kind.
func (kind ProgressEventKind) Validate() error {
	switch kind {
	case ProgressScanStarted, ProgressCategoryStarted, ProgressCategoryCompleted, ProgressScanCancelled, ProgressScanFinished:
		return nil
	default:
		return fmt.Errorf("invalid progress event kind %q", kind)
	}
}

// ProgressObserver synchronously receives scanner lifecycle events.
type ProgressObserver func(ProgressEvent)

// ProgressEvent is immutable by convention: its values describe only category
// lifecycle state and never expose paths, descriptors, or file facts.
type ProgressEvent struct {
	kind                      ProgressEventKind
	categoryID                core.AreaID
	categoryDisplayName       string
	processed, canonicalTotal int
	status                    core.RootStatus
}

func newProgressEvent(kind ProgressEventKind, definition *BuiltinDefinition, processed, total int, status core.RootStatus) ProgressEvent {
	event := ProgressEvent{kind: kind, processed: processed, canonicalTotal: total, status: status}
	if definition != nil {
		event.categoryID = definition.AreaID
		event.categoryDisplayName = definition.DisplayName
	}
	return event
}

func (event ProgressEvent) Kind() ProgressEventKind     { return event.kind }
func (event ProgressEvent) CategoryID() core.AreaID     { return event.categoryID }
func (event ProgressEvent) CategoryDisplayName() string { return event.categoryDisplayName }
func (event ProgressEvent) Processed() int              { return event.processed }
func (event ProgressEvent) Total() int                  { return event.canonicalTotal }
func (event ProgressEvent) Status() core.RootStatus     { return event.status }

// Validate reports whether event has a defined kind and canonical counters.
func (event ProgressEvent) Validate() error {
	if err := event.kind.Validate(); err != nil {
		return err
	}
	if event.processed < 0 || event.canonicalTotal < 0 || event.processed > event.canonicalTotal {
		return ErrInvalidMetadata
	}
	return nil
}

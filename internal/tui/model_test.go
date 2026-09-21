package tui

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/osdy/OsdyCleaner/internal/core"
	"github.com/osdy/OsdyCleaner/internal/scan"
)

func TestModelInit(t *testing.T) {
	if cmd := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted)).Init(); cmd != nil {
		t.Fatal("Init must not start work")
	}
}

func TestDashboardResultsResponsiveAndSafe(t *testing.T) {
	for _, width := range []int{1, 12, 60, 100} {
		t.Run(fmt.Sprintf("width=%d", width), func(t *testing.T) {
			m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
			m.width, m.height = width, 30
			m.setViewportContent()
			view := m.View().Content
			for _, line := range strings.Split(view, "\n") {
				if got := visibleWidth(line); got > width {
					t.Fatalf("width %d rendered %d columns: %q", width, got, line)
				}
			}
			if width >= 90 {
				for _, want := range []string{"LOGICAL ESTIMATE", "ALLOCATED ESTIMATE", "OUTCOME / COMPLETENESS", "FINDINGS", "WARNINGS", "CATEGORIES", "DETAIL / FINDINGS", "55 B", "105 B", "npm cache", "Homebrew cache", "Gradle caches", "Xcode DerivedData", "CoreSimulator"} {
					if !strings.Contains(view, want) {
						t.Errorf("missing %q: %s", want, view)
					}
				}
			}
		})
	}
}

func TestDashboardEmptyDetailsRemainExplicit(t *testing.T) {
	m := NewModel(snapshotWithoutDetails(t))
	m.width, m.height = 100, 30
	if view := m.View().Content; !strings.Contains(view, "No findings recorded.") {
		t.Fatalf("empty findings are not explicit: %s", view)
	}
	m = update(m, "tab")
	if view := m.View().Content; !strings.Contains(view, "No warnings recorded.") {
		t.Fatalf("empty warnings are not explicit: %s", view)
	}
}

func TestDashboardTabOnlyTogglesFindingsAndWarnings(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	if m.warnings {
		t.Fatal("initial pane must be findings")
	}
	m = update(m, "tab")
	if !m.warnings || !strings.Contains(m.View().Content, "DETAIL / WARNINGS") {
		t.Fatal("Tab must switch to warnings")
	}
	m = update(m, "tab")
	if m.warnings || !strings.Contains(m.View().Content, "DETAIL / FINDINGS") {
		t.Fatal("Tab must return to findings")
	}
}

func TestModelNavigation(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	for range 3 {
		m = update(m, "l")
	}
	if m.selected != 3 || m.selectedArea() != core.AreaXcodeDerivedData {
		t.Fatalf("right navigation lost canonical category order: %#v", m)
	}
	m = update(m, "h")
	if m.selected != 2 {
		t.Fatalf("left navigation did not move category: %#v", m)
	}
	m = update(m, "down")
	if m.selected != 2 || m.detailOffset != 1 {
		t.Fatalf("down must scroll the visible detail pane without moving categories: %#v", m)
	}
}

func TestModelPaging(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m.height = 3
	m = update(m, "pgdown")
	if m.offset == 0 || !m.selectionVisible() {
		t.Fatalf("page down hid the selected category: %#v", m)
	}
	m = update(m, "pgup")
	if m.offset != 0 || !m.selectionVisible() {
		t.Fatalf("page up did not retain a visible selection: %#v", m)
	}
}

func TestModelPagingDetailContextReset(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m.height = 1
	m = update(m, "tab")
	m = update(m, "j")

	m = update(m, "pgup")
	if m.selected != 0 || m.detailOffset != 1 || !m.selectionVisible() {
		t.Fatalf("page up at the first category changed valid detail state: %#v", m)
	}

	m = update(m, "pgdown")
	if m.selected != 1 || m.detailOffset != 0 || !strings.Contains(m.View().Content, "Category: Homebrew cache") {
		t.Fatalf("page down did not reset detail context: %#v view=%q", m, m.View().Content)
	}

	m = update(m, "j")
	m = update(m, "pgup")
	if m.selected != 0 || m.detailOffset != 0 || !strings.Contains(m.View().Content, "Category: npm cache") {
		t.Fatalf("page up did not reset detail context: %#v view=%q", m, m.View().Content)
	}

	for range 10 {
		m = update(m, "pgdown")
	}
	m = update(m, "j")
	m = update(m, "pgdown")
	if m.selected != 4 || m.offset != 4 || m.detailOffset != 1 || !m.selectionVisible() {
		t.Fatalf("page down at the final category broke saturated state: %#v", m)
	}
}

func TestModelDetails(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m = update(m, "tab")
	if !m.details || !m.warnings || !strings.Contains(m.View().Content, "DETAIL / WARNINGS") {
		t.Fatal("Tab must switch the visible read-only pane to warnings")
	}
}

func TestModelDetailsFitTerminalHeight(t *testing.T) {
	for _, tt := range []struct {
		name     string
		status   core.RootStatus
		reason   core.RootReasonCode
		warnings bool
	}{
		{name: "complete details", status: core.RootScanned, reason: core.ReasonCompleted},
		{name: "incomplete details", status: core.RootPartial, reason: core.ReasonEntryVisibilityGap},
		{name: "incomplete warnings", status: core.RootPartial, reason: core.ReasonEntryVisibilityGap, warnings: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := NewModel(snapshot(t, tt.status, tt.reason))
			m.width, m.height = 500, 22
			m.setViewportContent()
			m = update(m, "tab")
			if tt.warnings {
				m = update(m, "tab")
			}

			if got := len(strings.Split(m.View().Content, "\n")); got > m.height {
				t.Fatalf("details view uses %d rows at terminal height %d: %q", got, m.height, m.View().Content)
			}
		})
	}
}

func TestModelDetailContextReset(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m.height = 1
	m = update(m, "tab")
	m = update(m, "j")
	if m.detailOffset != 1 {
		t.Fatalf("setup did not scroll details: %#v", m)
	}

	m = update(m, "l")
	if m.detailOffset != 0 || !strings.Contains(m.View().Content, "Category: Homebrew cache") {
		t.Fatalf("category change did not reset detail context: %#v view=%q", m, m.View().Content)
	}

	m = update(m, "j")
	m = update(m, "tab")
	if m.detailOffset != 0 || m.warnings {
		t.Fatalf("findings switch did not reset detail offset: %#v", m)
	}
}

func TestModelExtremeSizes(t *testing.T) {
	for _, height := range []int{0, -1, -5, math.MinInt} {
		t.Run(fmt.Sprintf("height=%d", height), func(t *testing.T) {
			m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
			m.height, m.selected, m.offset = height, 4, 4
			m.keepSelectedVisible()
			if m.pageSize() != 1 || m.detailPageSize() != 1 {
				t.Fatalf("page sizes at height %d = (%d, %d), want (1, 1)", height, m.pageSize(), m.detailPageSize())
			}
			if !m.selectionVisible() {
				t.Fatalf("selection is not visible at height %d: %#v", height, m)
			}
			_ = m.View()
		})
	}
}

func TestModelResize(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m.selected, m.offset = 4, 3
	next, _ := m.Update(tea.WindowSizeMsg{Width: 70, Height: 3})
	m = next.(Model)
	if m.width != 70 || m.height != 3 || !m.selectionVisible() {
		t.Fatalf("resize did not retain a visible selection: %#v", m)
	}
}

func TestModelQuit(t *testing.T) {
	for _, key := range []string{"q", "esc", "ctrl+c"} {
		m := update(NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted)), key)
		if !m.quitting {
			t.Errorf("%q did not quit", key)
		}
		if _, cmd := m.Update(tea.KeyPressMsg{Text: "j"}); cmd != nil {
			t.Errorf("%q allowed a command after quit", key)
		}
	}
}

func TestModelOutcomes(t *testing.T) {
	for _, tt := range []struct {
		status core.RootStatus
		reason core.RootReasonCode
		want   string
	}{
		{core.RootScanned, core.ReasonCompleted, "OUTCOME / COMPLETENESS: complete / true"},
		{core.RootPartial, core.ReasonEntryVisibilityGap, "INCOMPLETE"},
		{core.RootCancelled, core.ReasonCancelledDuringScan, "CANCELLED"},
	} {
		t.Run(tt.want, func(t *testing.T) {
			view := NewModel(snapshot(t, tt.status, tt.reason)).View().Content
			if !strings.Contains(view, tt.want) || !strings.Contains(view, "READ-ONLY") || !strings.Contains(view, "Manual review required") || !strings.Contains(view, "Estimate—not guaranteed reclaimable space") {
				t.Fatalf("required read-only outcome language missing: %s", view)
			}
			if tt.status != core.RootScanned && strings.Index(view, tt.want) > strings.Index(view, "CATEGORIES") {
				t.Fatal("incomplete state must precede facts")
			}
			for _, forbidden := range []string{"action", "select", "confirm", "delete", "remove", "trash", "cleanup"} {
				if strings.Contains(strings.ToLower(view), forbidden) {
					t.Fatalf("unsafe affordance %q in %q", forbidden, view)
				}
			}
		})
	}
}

func TestModelPresentationFacts(t *testing.T) {
	view := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted)).View().Content
	for _, want := range []string{"LOGICAL ESTIMATE: 55 B", "ALLOCATED ESTIMATE: 105 B", "OUTCOME / COMPLETENESS: complete / true", "FINDINGS: 5", "WARNINGS: 5"} {
		if !strings.Contains(view, want) {
			t.Fatalf("summary omitted %q: %s", want, view)
		}
	}
}

func TestModelCategoryFacts(t *testing.T) {
	view := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted)).View().Content
	for _, want := range []string{"> npm cache | scanned | 11 B | F:1 W:1", "Category: npm cache"} {
		if !strings.Contains(view, want) {
			t.Fatalf("category omitted %q: %s", want, view)
		}
	}
}

func TestModelFindingFacts(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m.width, m.height = 500, 30
	m.setViewportContent()
	view := m.View().Content
	for _, want := range []string{"Findings: npm-cache", "~/.npm/entry-0", "Reason: built_in_npm_cache_entry Risk: manual_review", "Accounting: measured_objects=3 attributed_objects=2 already_accounted_paths=1 hard_link_observations=2 unknown_identity_objects=1", "Warnings: [apfs_estimate_uncertain]"} {
		if !strings.Contains(view, want) {
			t.Fatalf("finding omitted %q: %s", want, view)
		}
	}
}

func TestModelWarningFacts(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m = update(m, "l")
	m = update(m, "l")
	m.width, m.height = 500, 30
	m.setViewportContent()
	m = update(m, "tab")
	view := m.View().Content
	for _, want := range []string{"Warnings: gradle-caches", "Path: ~/.gradle/caches/warning-2", "Code: apfs_estimate_uncertain", "Message: APFS allocation is an estimate and not guaranteed reclaimable space.", "Related path: ~/.gradle/caches/related-entry"} {
		if !strings.Contains(view, want) {
			t.Fatalf("warning omitted %q: %s", want, view)
		}
	}
}

func TestModelViewportScrolling(t *testing.T) {
	m := NewModel(snapshot(t, core.RootScanned, core.ReasonCompleted))
	m.height = 1
	for range m.maxDetailOffset() {
		m = update(m, "j")
	}
	if m.detailOffset != m.maxDetailOffset() || !strings.Contains(m.View().Content, "Warnings: [apfs_estimate_uncertain]") {
		t.Fatalf("j did not visibly saturate full finding content: %#v view=%q", m, m.View().Content)
	}
	for range m.maxDetailOffset() {
		m = update(m, "k")
	}
	if m.detailOffset != 0 || !strings.Contains(m.View().Content, "Category: npm cache") {
		t.Fatalf("k did not reverse full finding content: %#v view=%q", m, m.View().Content)
	}
	m = update(m, "tab")
	if m.detailOffset != 0 || !m.warnings || !strings.Contains(m.View().Content, "DETAIL / WARNINGS") {
		t.Fatalf("warning switch did not reset viewport context: %#v view=%q", m, m.View().Content)
	}
	for range m.maxDetailOffset() + 1 {
		m = update(m, "down")
	}
	if m.detailOffset != m.maxDetailOffset() || !strings.Contains(m.View().Content, "Related path: none") {
		t.Fatalf("down did not visibly saturate full warning content: %#v view=%q", m, m.View().Content)
	}
}

func TestInvalidSnapshot(t *testing.T) {
	m := NewModel(core.Snapshot{})
	m = update(m, "tab")
	if m.selectedArea() != "" || m.details || m.warnings || strings.Contains(m.View().Content, "Details\n") || strings.Contains(m.View().Content, "Warnings\n") {
		t.Fatalf("invalid snapshot must not enter detail state: %#v", m)
	}
	m = update(m, "l")
	m = update(m, "pgdown")
	if !m.selectionVisible() {
		t.Fatalf("empty snapshot selection semantics are invalid: %#v", m)
	}
}

func TestRunRejectsInvalidFiles(t *testing.T) {
	s := snapshot(t, core.RootScanned, core.ReasonCompleted)
	if err := Run(s, nil, os.Stdout); err == nil {
		t.Fatal("runner accepted nil input")
	}
	if err := Run(s, os.Stdin, nil); err == nil {
		t.Fatal("runner accepted nil output")
	}
}

func TestRunUsesAlternateScreen(t *testing.T) {
	old := startProgram
	t.Cleanup(func() { startProgram = old })
	called := false
	startProgram = func(model Model, _ io.Reader, _ io.Writer, options programOptions) program {
		called = options.alternateScreen && model.View().AltScreen
		return fakeProgram{}
	}
	if err := run(snapshot(t, core.RootScanned, core.ReasonCompleted), &terminalBuffer{}, &terminalBuffer{}, func(uintptr) bool { return true }); err != nil || !called {
		t.Fatalf("runner did not configure alternate-screen lifecycle: err=%v called=%v", err, called)
	}
}

func TestRunRejectsNonTerminal(t *testing.T) {
	if err := run(snapshot(t, core.RootScanned, core.ReasonCompleted), bytes.NewBuffer(nil), bytes.NewBuffer(nil), func(uintptr) bool { return false }); err == nil {
		t.Fatal("runner accepted non-terminal streams")
	}
}

func TestScanModelLifecycle(t *testing.T) {
	started := make(chan struct{}, 1)
	finished := make(chan struct{})
	want := snapshot(t, core.RootScanned, core.ReasonCompleted)
	m := newScanModel(context.Background(), func(ctx context.Context, _ scan.ProgressObserver) (core.Snapshot, error) {
		started <- struct{}{}
		<-ctx.Done()
		close(finished)
		return want, nil
	})
	if cap(m.session.messages) != scanMessageCapacity {
		t.Fatalf("message channel capacity = %d, want %d", cap(m.session.messages), scanMessageCapacity)
	}
	start := m.Init()
	if start == nil {
		t.Fatal("scan model Init must start scan session")
	}
	msg := start()
	if _, ok := msg.(scanStartedMsg); !ok {
		t.Fatalf("Init command = %T, want scanStartedMsg", msg)
	}
	<-started
	next, wait := m.Update(msg)
	m = next.(Model)
	if !m.running || wait == nil {
		t.Fatalf("started scan state = %#v, wait=%v", m, wait)
	}
	next, cancel := m.Update(tea.KeyPressMsg{Text: "q"})
	m = next.(Model)
	if !m.cancelling || m.quitting || cancel == nil {
		t.Fatalf("q during scan state = %#v, cancel=%v", m, cancel)
	}
	if _, duplicateCancel := m.Update(tea.KeyPressMsg{Text: "q"}); duplicateCancel != nil {
		t.Fatal("repeated q requested cancellation twice")
	}
	cancelMsg := cancel()
	if _, ok := cancelMsg.(scanCancelMsg); !ok {
		t.Fatalf("q command = %T, want scanCancelMsg", cancelMsg)
	}
	next, duplicateWait := m.Update(cancelMsg)
	m = next.(Model)
	if duplicateWait != nil {
		t.Fatal("cancellation acknowledgement started a second scan listener")
	}
	result := wait()
	next, quit := m.Update(result)
	m = next.(Model)
	<-finished
	if m.running || !m.finished || !m.quitting || quit == nil || m.snapshot.Outcome() != want.Outcome() {
		t.Fatalf("result state = %#v quit=%v", m, quit)
	}
}

func TestScanModelProgressAndFailure(t *testing.T) {
	m := newScanModel(context.Background(), func(context.Context, scan.ProgressObserver) (core.Snapshot, error) { return core.Snapshot{}, nil })
	next, _ := m.Update(scanProgressMsg{processed: 1, total: 5, category: "npm cache"})
	m = next.(Model)
	if m.processed != 1 || m.total != 5 || m.category != "npm cache" {
		t.Fatalf("progress state = %#v", m)
	}
	next, _ = m.Update(scanResultMsg{err: errors.New("scan failed")})
	m = next.(Model)
	if !m.failed || m.running || !strings.Contains(m.View().Content, "scan failed") {
		t.Fatalf("failure state = %#v view=%q", m, m.View().Content)
	}
}

func TestRunScanRejectsTerminalBeforeStarting(t *testing.T) {
	called := false
	_, err := runScan(context.Background(), func(context.Context, scan.ProgressObserver) (core.Snapshot, error) {
		called = true
		return core.Snapshot{}, nil
	}, bytes.NewBuffer(nil), bytes.NewBuffer(nil), func(uintptr) bool { return false })
	if err == nil || called {
		t.Fatalf("non-terminal RunScan err=%v called=%v", err, called)
	}
}

func TestScanInitStartsOnlyOnce(t *testing.T) {
	calls := 0
	m := newScanModel(context.Background(), func(context.Context, scan.ProgressObserver) (core.Snapshot, error) {
		calls++
		return core.Snapshot{}, nil
	})
	m.Init()()
	m.Init()()
	_, err := m.session.waitShutdown()
	if err != nil || calls != 1 {
		t.Fatalf("scan calls=%d err=%v, want one successful call", calls, err)
	}
}

func TestRunScanProgramReturnCancelsAndWaits(t *testing.T) {
	old := startProgram
	t.Cleanup(func() { startProgram = old })
	stopped := make(chan struct{})
	programErr := errors.New("program stopped")
	startProgram = func(m Model, _ io.Reader, _ io.Writer, _ programOptions) program {
		if !m.View().AltScreen {
			t.Fatal("RunScan did not configure alternate screen")
		}
		return fakeProgramFunc(func() (tea.Model, error) {
			m.Init()()
			return nil, programErr
		})
	}
	_, err := runScan(context.Background(), func(ctx context.Context, _ scan.ProgressObserver) (core.Snapshot, error) {
		<-ctx.Done()
		close(stopped)
		return core.Snapshot{}, ctx.Err()
	}, &terminalBuffer{}, &terminalBuffer{}, func(uintptr) bool { return true })
	<-stopped
	if !errors.Is(err, programErr) {
		t.Fatalf("RunScan error = %v, want program error", err)
	}
}

func TestScanProgressView(t *testing.T) {
	for _, tt := range []struct {
		name               string
		width              int
		processed, total   int
		category           string
		cancelling, failed bool
		want               []string
	}{
		{
			name: "wide preparing panel", width: 100, total: 0,
			want: []string{"OSDY", "CLEANER", "READ-ONLY", "Preparing category scan", "0/5 categories processed", "[----------------------------------------]", "Press q to cancel"},
		},
		{
			name: "narrow compact mark", width: 32, processed: 1, total: 5, category: "npm cache",
			want: []string{"OSDY // CLEANER", "1/5 categories processed", "[#####-----------------------]", "Active: npm cache", "READ-ONLY"},
		},
		{
			name: "tiny terminal fallback", width: 12, processed: 1, total: 5, category: "npm cache",
			want: []string{"OSDY", "READ-ONLY", "1/5", "npm cache"},
		},
		{name: "single-column bound", width: 1, total: 0},
		{
			name: "complete categories", width: 80, processed: 5, total: 5, category: "CoreSimulator",
			want: []string{"5/5 categories processed", "[########################################]", "Active: CoreSimulator"},
		},
		{
			name: "cancelling", width: 80, processed: 1, total: 5, category: "npm cache", cancelling: true,
			want: []string{"CANCELLING", "waiting for safe scanner shutdown", "interface will close after shutdown", "READ-ONLY"},
		},
		{
			name: "failed", width: 80, failed: true,
			want: []string{"SCAN FAILED", "scanner unavailable", "READ-ONLY", "Press q to quit"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := newScanModel(context.Background(), func(context.Context, scan.ProgressObserver) (core.Snapshot, error) { return core.Snapshot{}, nil })
			m.width, m.running, m.cancelling, m.failed = tt.width, !tt.failed, tt.cancelling, tt.failed
			m.processed, m.total, m.category = tt.processed, tt.total, tt.category
			if tt.failed {
				m.err = errors.New("scanner unavailable")
			}
			view := m.View().Content
			for _, want := range tt.want {
				if !strings.Contains(view, want) {
					t.Errorf("view missing %q:\n%s", want, view)
				}
			}
			for _, line := range strings.Split(view, "\n") {
				if got := visibleWidth(line); got > tt.width {
					t.Errorf("visible width %d exceeds terminal width %d: %q", got, tt.width, line)
				}
			}
		})
	}
}

func visibleWidth(s string) int {
	for {
		start := strings.Index(s, "\x1b[")
		if start < 0 {
			return len([]rune(s))
		}
		end := start + 2
		for end < len(s) && (s[end] < '@' || s[end] > '~') {
			end++
		}
		if end == len(s) {
			return len([]rune(s[:start]))
		}
		s = s[:start] + s[end+1:]
	}
}

type terminalBuffer struct{ bytes.Buffer }

func (terminalBuffer) Fd() uintptr { return 0 }

type fakeProgram struct{}

func (fakeProgram) Run() (tea.Model, error) { return nil, nil }

type fakeProgramFunc func() (tea.Model, error)

func (f fakeProgramFunc) Run() (tea.Model, error) { return f() }

func update(m Model, key string) Model {
	next, _ := m.Update(tea.KeyPressMsg{Text: key})
	return next.(Model)
}

func snapshotWithoutDetails(t *testing.T) core.Snapshot {
	t.Helper()
	base := snapshot(t, core.RootScanned, core.ReasonCompleted)
	roots := base.Roots()
	for i, root := range roots {
		roots[i], _ = core.NewRootObservation(root.AreaID(), root.DisplayName(), root.DisplayPath(), root.Status(), root.Reason(), root.Estimate(), 0, nil)
	}
	empty, err := core.NewSnapshot(base.RuleSetVersion(), base.ScanPolicy(), roots, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	return empty
}

func snapshot(t *testing.T, status core.RootStatus, reason core.RootReasonCode) core.Snapshot {
	t.Helper()
	policy, _ := core.NewScanPolicy(true, true, 10, 100)
	logical, _ := core.NewValueEstimate(11, core.BasisRegularFileStatSize, core.CompletenessComplete)
	allocated, _ := core.NewValueEstimate(21, core.BasisStatBlocks512, core.CompletenessComplete)
	estimate, _ := core.NewEstimate(logical, allocated, true)
	areas := []core.AreaID{core.AreaNPMCache, core.AreaHomebrewCache, core.AreaGradleCaches, core.AreaXcodeDerivedData, core.AreaCoreSimulator}
	names := []string{"npm cache", "Homebrew cache", "Gradle caches", "Xcode DerivedData", "CoreSimulator"}
	paths := []string{"~/.npm", "~/Library/Caches/Homebrew", "~/.gradle/caches", "~/Library/Developer/Xcode/DerivedData", "~/Library/Developer/CoreSimulator"}
	roots := make([]core.RootObservation, len(areas))
	findings := make([]core.Finding, len(areas))
	warnings := make([]core.Warning, len(areas))
	accounting, _ := core.NewAccounting(3, 2, 1, 2, 1)
	for i := range areas {
		rootStatus, rootReason := core.RootScanned, core.ReasonCompleted
		if i == 0 {
			rootStatus, rootReason = status, reason
		}
		codes := []core.WarningCode{core.WarningAPFSEstimateUncertain}
		roots[i], _ = core.NewRootObservation(areas[i], names[i], paths[i], rootStatus, rootReason, estimate, 1, codes)
		findings[i], _ = core.NewFinding(areas[i], paths[i]+"/entry-"+fmt.Sprint(i), []core.FindingReason{core.FindingBuiltInNPMCacheEntry, core.FindingBuiltInHomebrewCacheEntry, core.FindingBuiltInGradleCacheEntry, core.FindingBuiltInXcodeDerivedDataEntry, core.FindingBuiltInCoreSimulatorEntry}[i], estimate, accounting, codes)
		var related *string
		if i == 2 {
			value := paths[i] + "/related-entry"
			related = &value
		}
		warnings[i], _ = core.NewWarning(areas[i], paths[i]+"/warning-"+fmt.Sprint(i), core.WarningAPFSEstimateUncertain, related)
	}
	s, err := core.NewSnapshot("rules-v1", policy, roots, findings, warnings)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

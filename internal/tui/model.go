// Package tui presents finalized scan snapshots without starting or changing scans.
package tui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// Model contains only presentation state over an immutable finalized snapshot.
type Model struct {
	snapshot                       core.Snapshot
	selected, offset, detailOffset int
	width, height                  int
	details, warnings, quitting    bool
	alternateScreen                bool
	viewport                       viewport.Model
}

// NewModel constructs a viewer for an already finalized snapshot.
func NewModel(snapshot core.Snapshot) Model {
	m := Model{snapshot: snapshot, width: 120, height: 8, viewport: viewport.New()}
	m.setViewportContent()
	return m
}

// Init is deliberately pure: a viewer never starts a scan.
func (Model) Init() tea.Cmd { return nil }

// Update changes only viewer state. It never performs scan or filesystem work.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.quitting {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.setViewportContent()
		m.keepSelectedVisible()
	case tea.KeyPressMsg:
		key := msg.String()
		if msg.Text != "" {
			key = msg.Text
		}
		switch key {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "h", "left":
			m.moveCategory(-1)
		case "l", "right":
			m.moveCategory(1)
		case "j", "down":
			if m.details {
				m.viewport.ScrollDown(1)
				m.detailOffset = m.viewport.YOffset()
			}
		case "k", "up":
			if m.details {
				m.viewport.ScrollUp(1)
				m.detailOffset = m.viewport.YOffset()
			}
		case "pgdown":
			m.pageSelection(m.pageSize())
		case "pgup":
			m.pageSelection(-m.pageSize())
		case "tab":
			if len(m.snapshot.Roots()) != 0 {
				if !m.details {
					m.details = true
				} else if !m.warnings {
					m.warnings = true
				} else {
					m.details, m.warnings = false, false
				}
				m.resetViewportContent()
			}
		}
		m.keepSelectedVisible()
	}
	return m, nil
}

func (m *Model) moveCategory(delta int) {
	selected := m.selected
	m.selected += delta
	m.keepSelectedVisible()
	if m.selected != selected {
		m.resetViewportContent()
	}
}

func (m *Model) pageSelection(delta int) {
	selected := m.selected
	m.selected += delta
	m.keepSelectedVisible()
	if m.selected != selected {
		m.resetViewportContent()
	}
}

func (m Model) selectedArea() core.AreaID {
	roots := m.snapshot.Roots()
	if m.selected < 0 || m.selected >= len(roots) {
		return ""
	}
	return roots[m.selected].AreaID()
}
func (m Model) pageSize() int {
	if m.height <= 2 {
		return 1
	}
	return m.height - 2
}
func (m Model) detailPageSize() int {
	if m.height <= 6 {
		return 1
	}
	return m.height - 6
}
func (m Model) maxDetailOffset() int { return max(0, len(m.detailLines())-m.detailPageSize()) }
func (m Model) detailLines() []string {
	if m.warnings {
		return m.warningLines()
	}
	area := m.selectedArea()
	if area == "" {
		return nil
	}
	lines := []string{fmt.Sprintf("Findings: %s", area)}
	for _, finding := range m.snapshot.Findings() {
		if finding.AreaID() != area {
			continue
		}
		lines = append(lines, "Path: "+finding.DisplayPath(), fmt.Sprintf("Reason: %s Risk: %s", finding.Reason(), finding.Risk()), estimateLine("Estimate", finding.Estimate()))
		accounting := finding.Accounting()
		lines = append(lines, fmt.Sprintf("Accounting: measured_objects=%d attributed_objects=%d already_accounted_paths=%d hard_link_observations=%d unknown_identity_objects=%d", accounting.MeasuredObjects(), accounting.AttributedObjects(), accounting.AlreadyAccountedPaths(), accounting.HardLinkObservations(), accounting.UnknownIdentityObjects()), fmt.Sprintf("Warnings: %v", finding.WarningCodes()))
	}
	if len(lines) == 1 {
		return append(lines, "No findings recorded.")
	}
	return lines
}
func (m Model) warningLines() []string {
	area := m.selectedArea()
	if area == "" {
		return nil
	}
	lines := []string{fmt.Sprintf("Warnings: %s", area)}
	for _, warning := range m.snapshot.Warnings() {
		if warning.AreaID() != area {
			continue
		}
		related := "none"
		if path, ok := warning.RelatedPath(); ok {
			related = path
		}
		lines = append(lines, "Path: "+warning.DisplayPath(), "Code: "+string(warning.Code()), "Message: "+warning.Message(), "Related path: "+related)
	}
	if len(lines) == 1 {
		return append(lines, "No warnings recorded.")
	}
	return lines
}
func (m Model) selectionVisible() bool {
	roots := m.snapshot.Roots()
	return len(roots) == 0 || (m.selected >= m.offset && m.selected < min(len(roots), m.offset+m.pageSize()))
}
func (m *Model) resetViewportContent() {
	m.viewport.GotoTop()
	m.setViewportContent()
}

func (m *Model) setViewportContent() {
	m.viewport.SetWidth(max(1, m.width))
	m.viewport.SetHeight(m.detailPageSize())
	m.viewport.SetContent(strings.Join(m.detailLines(), "\n"))
	m.detailOffset = m.viewport.YOffset()
}
func (m *Model) keepSelectedVisible() {
	roots := m.snapshot.Roots()
	if len(roots) == 0 {
		m.selected, m.offset, m.detailOffset = 0, 0, 0
		return
	}
	m.selected = min(max(0, m.selected), len(roots)-1)
	m.offset = min(max(0, m.offset), max(0, len(roots)-m.pageSize()))
	if m.selected < m.offset {
		m.offset = m.selected
	}
	if last := m.offset + m.pageSize() - 1; m.selected > last {
		m.offset = m.selected - m.pageSize() + 1
	}
	m.detailOffset = m.viewport.YOffset()
}

func estimateLine(label string, estimate core.Estimate) string {
	logical, allocated := estimate.Logical(), estimate.Allocated()
	return fmt.Sprintf("%s: logical=%d basis=%s completeness=%s; allocated=%d basis=%s completeness=%s; deduplication_complete=%t", label, logical.KnownBytes(), logical.Basis(), logical.Completeness(), allocated.KnownBytes(), allocated.Basis(), allocated.Completeness(), estimate.DeduplicationComplete())
}

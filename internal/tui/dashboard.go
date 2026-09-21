package tui

import (
	"fmt"
	"strings"

	"github.com/osdy/OsdyCleaner/internal/core"
)

// dashboardView renders only the finalized snapshot. It never derives scan facts.
func (m Model) dashboardView() string {
	width := max(1, m.width)
	if width < 20 {
		return m.tinyDashboard(width)
	}
	lines := []string{m.dashboardBrand(width), "READ-ONLY // FINALIZED SNAPSHOT", "Manual review required. Estimate—not guaranteed reclaimable space."}
	if !m.snapshot.Complete() {
		state := "INCOMPLETE"
		if m.snapshot.Outcome() == "cancelled" {
			state = "CANCELLED // INCOMPLETE"
		}
		lines = append(lines, state+": available facts may not cover every root.")
	}
	logical, allocated := m.snapshot.Estimate().Logical(), m.snapshot.Estimate().Allocated()
	summary := []string{
		"LOGICAL ESTIMATE: " + iecBytes(logical.KnownBytes()),
		"ALLOCATED ESTIMATE: " + iecBytes(allocated.KnownBytes()),
		fmt.Sprintf("OUTCOME / COMPLETENESS: %s / %t", m.snapshot.Outcome(), m.snapshot.Complete()),
		fmt.Sprintf("FINDINGS: %d", len(m.snapshot.Findings())),
		fmt.Sprintf("WARNINGS: %d", len(m.snapshot.Warnings())),
	}
	if width >= 90 {
		for i := 0; i < len(summary); i += 2 {
			right := ""
			if i+1 < len(summary) {
				right = summary[i+1]
			}
			lines = append(lines, dashboardColumns(summary[i], right, width))
		}
		left := m.categoryLines((width - 3) / 2)
		right := m.detailPaneLines((width - 3) / 2)
		lines = append(lines, dashboardColumns("CATEGORIES", "DETAIL / "+m.detailMode(), width))
		for i := 0; i < max(len(left), len(right)); i++ {
			l, r := "", ""
			if i < len(left) {
				l = left[i]
			}
			if i < len(right) {
				r = right[i]
			}
			lines = append(lines, dashboardColumns(l, r, width))
		}
	} else {
		lines = append(lines, summary...)
		lines = append(lines, "CATEGORIES")
		lines = append(lines, m.categoryLines(width)...)
		lines = append(lines, "DETAIL / "+m.detailMode())
		lines = append(lines, m.detailPaneLines(width)...)
	}
	lines = append(lines, "Navigate: h/l categories · j/k scroll · Page Up/Down page · Tab findings/warnings · q quit")
	return dashboardLines(lines, width)
}

func (m Model) tinyDashboard(width int) string {
	lines := []string{"OSDY", "READ-ONLY", string(m.snapshot.Outcome()), "F:" + fmt.Sprint(len(m.snapshot.Findings())) + " W:" + fmt.Sprint(len(m.snapshot.Warnings()))}
	if !m.snapshot.Complete() {
		lines = append(lines, "INCOMPLETE")
	}
	return dashboardLines(lines, width)
}

func (m Model) dashboardBrand(width int) string {
	if width >= 90 {
		return "OSDY CLEANER // READ-ONLY RESULTS DASHBOARD"
	}
	return "OSDY // CLEANER"
}

func dashboardColumns(left, right string, width int) string {
	column := max(1, (width-3)/2)
	left, right = truncateLine(left, column), truncateLine(right, column)
	return left + strings.Repeat(" ", column-len([]rune(left))) + " │ " + right
}

func dashboardLines(lines []string, width int) string {
	for i, line := range lines {
		line = truncateLine(line, width)
		if strings.HasPrefix(line, "OSDY") || strings.HasPrefix(line, "READ-ONLY") || strings.HasPrefix(line, "CATEGORIES") || strings.HasPrefix(line, "DETAIL") {
			lines[i] = neonMagenta.Render(line)
		} else if strings.HasPrefix(line, "INCOMPLETE") || strings.HasPrefix(line, "CANCELLED") {
			lines[i] = neonFailure.Render(line)
		} else {
			lines[i] = neonMuted.Render(line)
		}
	}
	return strings.Join(lines, "\n")
}

func (m Model) categoryLines(width int) []string {
	roots := m.snapshot.Roots()
	lines := make([]string, 0, len(roots))
	for i, root := range roots {
		marker := " "
		if i == m.selected {
			marker = ">"
		}
		line := fmt.Sprintf("%s %s | %s | %s | F:%d W:%d", marker, root.DisplayName(), root.Status(), iecBytes(root.Estimate().Logical().KnownBytes()), root.FindingCount(), len(root.WarningCodes()))
		lines = append(lines, truncateLine(line, width))
	}
	return lines
}

func (m Model) detailMode() string {
	if m.warnings {
		return "WARNINGS"
	}
	return "FINDINGS"
}

func (m Model) detailPaneLines(width int) []string {
	lines := m.detailLines()
	if m.detailOffset < len(lines) {
		lines = lines[m.detailOffset:]
	} else {
		lines = nil
	}
	limit := max(1, m.detailPageSize())
	if len(lines) > limit {
		lines = lines[:limit]
	}
	for i := range lines {
		lines[i] = truncateLine(lines[i], width)
	}
	return lines
}

func iecBytes(bytes uint64) string {
	units := []string{"B", "KiB", "MiB", "GiB", "TiB"}
	value := float64(bytes)
	unit := 0
	for value >= 1024 && unit < len(units)-1 {
		value /= 1024
		unit++
	}
	if unit == 0 {
		return fmt.Sprintf("%d B", bytes)
	}
	return fmt.Sprintf("%.1f %s", value, units[unit])
}

func dashboardEstimate(label string, estimate core.Estimate) string {
	logical, allocated := estimate.Logical(), estimate.Allocated()
	return fmt.Sprintf("%s: logical=%s basis=%s completeness=%s; allocated=%s basis=%s completeness=%s", label, iecBytes(logical.KnownBytes()), logical.Basis(), logical.Completeness(), iecBytes(allocated.KnownBytes()), allocated.Basis(), allocated.Completeness())
}

package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// View is a pure rendering of scan state and local navigation state.
func (m Model) View() tea.View {
	var out strings.Builder
	if m.failed || m.running || m.cancelling {
		view := tea.NewView(m.scanProgressView())
		view.AltScreen = m.alternateScreen
		return view
	}
	out.WriteString(m.dashboardView())
	view := tea.NewView(out.String())
	view.AltScreen = m.alternateScreen
	return view
}

const categoryCount = 5

func (m Model) scanProgressView() string {
	width := max(1, m.width)
	if width < 20 {
		return m.tinyScanProgressView(width)
	}
	inside := width - 2
	lines := make([]string, 0, 10)
	if width >= 70 {
		lines = append(lines,
			"  OOO   SSS  DDDD  Y   Y   CCCC L     EEEEE  AAA  N   N EEEEE RRRR",
			" O   O S     D   D  Y Y   C     L     E     A   A NN  N E     R   R",
			" O   O  SSS  D   D   Y    C     L     EEEE  AAAAA N N N EEEE  RRRR",
			" O   O     S D   D   Y    C     L     E     A   A N  NN E     R R",
			"  OOO  SSSS  DDDD    Y     CCCC LLLLL EEEEE A   A N   N EEEEE R  RR",
			"OSDY // CLEANER",
		)
	} else {
		lines = append(lines, "OSDY // CLEANER")
	}
	lines = append(lines, "READ-ONLY // PHASE 1 SCAN")
	if m.failed {
		lines = append(lines, "SCAN FAILED", fmt.Sprintf("Scanner: %v", m.err), "Press q to quit")
		return neonPanel(lines, inside, neonFailure)
	}

	processed := min(max(0, m.processed), categoryCount)
	barWidth := min(40, max(10, inside-2))
	filled := processed * barWidth / categoryCount
	bar := "[" + strings.Repeat("#", filled) + strings.Repeat("-", barWidth-filled) + "]"
	if m.total == 0 {
		lines = append(lines, "Preparing category scan", "0/5 categories processed", bar)
	} else {
		lines = append(lines, fmt.Sprintf("%d/5 categories processed", processed), bar)
		if m.category == "" {
			lines = append(lines, "Finalizing snapshot")
		} else {
			lines = append(lines, "Active: "+m.category)
		}
	}
	if m.cancelling {
		lines = append(lines, "CANCELLING // waiting for safe scanner shutdown", "The interface will close after shutdown")
	} else {
		lines = append(lines, "Press q to cancel")
	}
	return neonPanel(lines, inside, neonCyan)
}

func (m Model) tinyScanProgressView(width int) string {
	processed := min(max(0, m.processed), categoryCount)
	lines := []string{"OSDY", "READ-ONLY"}
	switch {
	case m.failed:
		lines = append(lines, "FAILED", fmt.Sprint(m.err), "q quit")
	case m.cancelling:
		lines = append(lines, fmt.Sprintf("%d/5", processed), "CANCELLING", "waiting")
	case m.total == 0:
		lines = append(lines, "PREPARING", "0/5", "q cancel")
	default:
		lines = append(lines, fmt.Sprintf("%d/5", processed), m.category, "q cancel")
	}
	for i, line := range lines {
		line = truncateLine(line, width)
		if i < 2 {
			lines[i] = neonMagenta.Render(line)
		} else {
			lines[i] = neonMuted.Render(line)
		}
	}
	return strings.Join(lines, "\n")
}

func neonPanel(lines []string, inside int, accent lipgloss.Style) string {
	var out strings.Builder
	out.WriteString(accent.Render("┌" + strings.Repeat("─", inside) + "┐"))
	for _, line := range lines {
		out.WriteByte('\n')
		line = truncateLine(line, inside)
		content := neonMuted.Render(line)
		if strings.HasPrefix(line, "READ-ONLY") || strings.HasPrefix(line, "OSDY") {
			content = neonMagenta.Render(line)
		}
		if strings.HasPrefix(line, "SCAN FAILED") {
			content = neonFailure.Render(line)
		}
		if strings.HasPrefix(line, "5/5") {
			content = neonGreen.Render(line)
		}
		if strings.HasPrefix(line, "[") {
			content = accent.Render(line)
		}
		out.WriteString(accent.Render("│"))
		out.WriteString(content)
		out.WriteString(strings.Repeat(" ", inside-len([]rune(line))))
		out.WriteString(accent.Render("│"))
	}
	out.WriteByte('\n')
	out.WriteString(accent.Render("└" + strings.Repeat("─", inside) + "┘"))
	return out.String()
}

func truncateLine(line string, width int) string {
	runes := []rune(line)
	if len(runes) <= width {
		return line
	}
	if width < 2 {
		return string(runes[:width])
	}
	return string(runes[:width-1]) + "…"
}

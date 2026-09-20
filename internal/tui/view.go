package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// View is a pure rendering of the finalized snapshot and local navigation state.
func (m Model) View() tea.View {
	var out strings.Builder
	out.WriteString("Read-only scan\n")
	if !m.snapshot.Complete() {
		if m.snapshot.Outcome() == "cancelled" {
			out.WriteString("CANCELLED: available facts are incomplete.\n")
		} else {
			out.WriteString("INCOMPLETE: available facts may not cover every root.\n")
		}
	}
	fmt.Fprintf(&out, "Outcome: %s\nComplete: %t\n%s\n", m.snapshot.Outcome(), m.snapshot.Complete(), estimateLine("Aggregate estimate", m.snapshot.Estimate()))
	out.WriteString("Manual review required. Estimate—not guaranteed reclaimable space.\n")
	out.WriteString("Logical and allocated sizes are estimates; APFS clones, snapshots, compression, and hard links can make reclaimed space lower or different.\n")
	out.WriteString("CoreSimulator data requires product-aware manual review; no device lifecycle or follow-up semantics are inferred.\n\nCategories:\n")
	roots := m.snapshot.Roots()
	end := min(len(roots), m.offset+m.pageSize())
	for i := m.offset; i < end; i++ {
		marker := " "
		if i == m.selected {
			marker = ">"
		}
		root := roots[i]
		fmt.Fprintf(&out, "%s %s | %s | %s | %s | findings=%d | warning_codes=%v\n", marker, root.AreaID(), root.DisplayName(), root.Status(), root.Reason(), root.FindingCount(), root.WarningCodes())
	}
	if area := m.selectedArea(); area != "" {
		for _, root := range roots {
			if root.AreaID() == area {
				fmt.Fprintf(&out, "\nCategory: %s\nRoot: %s\nRoot estimate: %s\n", root.DisplayName(), root.DisplayPath(), estimateLine("", root.Estimate())[2:])
				break
			}
		}
	}
	if m.details {
		heading := "Details"
		if m.warnings {
			heading = "Warnings"
		}
		fmt.Fprintf(&out, "\n%s\n%s\n", heading, m.viewport.View())
	}
	out.WriteString("\nNavigate: h/l categories · j/k or arrows scroll · Page Up/Down page · Tab details/warnings · q quit")
	view := tea.NewView(out.String())
	view.AltScreen = m.alternateScreen
	return view
}

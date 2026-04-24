package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7B2FBE")).
			Padding(0, 2)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#AFAFAF"))

	rowStyle = lipgloss.NewStyle().
			Padding(0, 1)

	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)
)

// View renders the terminal dashboard as a string.
func (m Model) View() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render(" gRPC Health Probe "))
	sb.WriteString("\n\n")

	if len(m.Targets) == 0 {
		sb.WriteString(rowStyle.Render("No targets configured."))
		sb.WriteString("\n")
		return sb.String()
	}

	sb.WriteString(headerStyle.Render(fmt.Sprintf("%-40s %-12s %s", "TARGET", "STATUS", "MESSAGE")))
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat("─", 70))
	sb.WriteString("\n")

	for _, target := range m.Targets {
		status, ok := m.Statuses[target]
		if !ok {
			status = StatusUnknown
		}

		label := StatusLabel(status)
		color := StatusColor(status)
		styledLabel := lipgloss.NewStyle().Foreground(color).Bold(true).Render(label)

		line := fmt.Sprintf("%-40s %-12s", target, styledLabel)
		sb.WriteString(rowStyle.Render(line))
		sb.WriteString("\n")
	}

	sb.WriteString(footerStyle.Render(fmt.Sprintf("Polling every %s — press q to quit", m.Interval)))
	sb.WriteString("\n")

	return sb.String()
}

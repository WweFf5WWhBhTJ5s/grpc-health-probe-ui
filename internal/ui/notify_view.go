package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// NotifyView renders the notification log panel.
func NotifyView(store *NotifyStore, width int) string {
	if store == nil || store.Len() == 0 {
		return ""
	}

	theme := DefaultTheme()

	infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	warnStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
	errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)

	var sb strings.Builder
	entries := store.All()
	for _, n := range entries {
		label := LevelLabel(n.Level)
		ts := n.CreatedAt.Format("15:04:05")
		line := fmt.Sprintf(" %s %s  %s", ts, label, n.Message)

		var styled string
		switch n.Level {
		case NotifyWarn:
			styled = warnStyle.Render(line)
		case NotifyError:
			styled = errStyle.Render(line)
		default:
			styled = infoStyle.Render(line)
		}

		sb.WriteString(styled)
		sb.WriteRune('\n')
	}

	box := theme.Title.Copy().
		Width(width).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(0, 1)

	return box.Render(sb.String())
}

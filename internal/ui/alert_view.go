package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	alertInfoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	alertWarnStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Bold(true)
	alertErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
	alertPrefixes   = map[AlertLevel]string{
		AlertInfo:  "ℹ ",
		AlertWarn:  "⚠ ",
		AlertError: "✖ ",
	}
)

// AlertBanner renders all active alerts as a stacked banner string.
// Returns an empty string when there are no alerts.
func AlertBanner(queue *AlertQueue, width int) string {
	if queue == nil || queue.Len() == 0 {
		return ""
	}

	var sb strings.Builder
	for _, a := range queue.Active() {
		prefix := alertPrefixes[a.Level]
		msg := fmt.Sprintf("%s%s", prefix, a.Message)
		if width > 0 && len(msg) > width {
			msg = msg[:width-1] + "…"
		}
		var styled string
		switch a.Level {
		case AlertWarn:
			styled = alertWarnStyle.Render(msg)
		case AlertError:
			styled = alertErrorStyle.Render(msg)
		default:
			styled = alertInfoStyle.Render(msg)
		}
		sb.WriteString(styled)
		sb.WriteString("\n")
	}
	return strings.TrimRight(sb.String(), "\n")
}

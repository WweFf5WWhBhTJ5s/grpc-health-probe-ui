package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nicholasgasior/grpc-health-probe-ui/internal/probe"
)

var (
	detailTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12")).
			PaddingBottom(1)

	detailLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Width(14)

	detailValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15"))

	detailBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8")).
			Padding(1, 2)
)

// DetailView renders a detailed panel for a single target.
func DetailView(t probe.Target, history []probe.StatusEntry, width int) string {
	if width <= 0 {
		width = 60
	}

	var sb strings.Builder

	sb.WriteString(detailTitleStyle.Render(t.Name))
	sb.WriteString("\n")
	sb.WriteString(detailRow("Host", t.Host))
	sb.WriteString(detailRow("Status", StatusLabel(t.Status)))
	sb.WriteString(detailRow("Message", emptyFallback(t.Message, "—")))
	sb.WriteString(detailRow("Latency", formatLatency(t.LatencyMs)))
	sb.WriteString(detailRow("Last Check", formatTime(t.LastChecked)))

	if len(history) > 0 {
		sb.WriteString("\n")
		sb.WriteString(detailLabelStyle.Render("History"))
		sb.WriteString(" ")
		sb.WriteString(SparkLine(history, 20))
		sb.WriteString("\n")
		sb.WriteString(detailLabelStyle.Render("Uptime"))
		sb.WriteString(" ")
		sb.WriteString(UptimeLabel(history))
		sb.WriteString("\n")
	}

	inner := sb.String()
	return detailBoxStyle.Width(width - 4).Render(inner)
}

func detailRow(label, value string) string {
	return detailLabelStyle.Render(label+":") + " " + detailValueStyle.Render(value) + "\n"
}

func emptyFallback(s, fallback string) string {
	if strings.TrimSpace(s) == "" {
		return fallback
	}
	return s
}

func formatLatency(ms int64) string {
	if ms < 0 {
		return "—"
	}
	return fmt.Sprintf("%d ms", ms)
}

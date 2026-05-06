package ui

import (
	"fmt"
	"strings"
)

// MetricView renders a MetricSummary as a formatted multi-line string
// suitable for display in the terminal dashboard.
func MetricView(s MetricSummary, width int) string {
	if width <= 0 {
		width = 40
	}

	var b strings.Builder

	title := fmt.Sprintf(" Metrics: %s", s.Target)
	b.WriteString(title + "\n")
	b.WriteString(strings.Repeat("─", min(width, 60)) + "\n")

	rows := []struct {
		label string
		value string
	}{
		{"Total checks", fmt.Sprintf("%d", s.Total)},
		{"Serving", fmt.Sprintf("%d", s.Serving)},
		{"Not Serving", fmt.Sprintf("%d", s.NotServing)},
		{"Unknown", fmt.Sprintf("%d", s.Unknown)},
		{"Uptime", fmt.Sprintf("%.1f%%", s.UptimePct)},
		{"Avg latency", formatLatency(s.AvgLatency)},
		{"Min latency", formatLatency(s.MinLatency)},
		{"Max latency", formatLatency(s.MaxLatency)},
	}

	for _, r := range rows {
		b.WriteString(fmt.Sprintf("  %-16s %s\n", r.label+":", r.value))
	}

	return b.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

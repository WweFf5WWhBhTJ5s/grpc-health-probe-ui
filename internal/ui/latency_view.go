package ui

import (
	"fmt"
	"strings"
)

// LatencyView renders a formatted latency statistics panel.
func LatencyView(stats LatencyStats, width int) string {
	if width <= 0 {
		width = 40
	}

	if stats.Count == 0 {
		return centerPad("no latency data", width)
	}

	rows := []string{
		latencyRow("min", formatLatencyDuration(stats.Min), width),
		latencyRow("mean", formatLatencyDuration(stats.Mean), width),
		latencyRow("p50", formatLatencyDuration(stats.P50), width),
		latencyRow("p95", formatLatencyDuration(stats.P95), width),
		latencyRow("p99", formatLatencyDuration(stats.P99), width),
		latencyRow("max", formatLatencyDuration(stats.Max), width),
		latencyRow("samples", fmt.Sprintf("%d", stats.Count), width),
	}
	return strings.Join(rows, "\n")
}

func latencyRow(label, value string, width int) string {
	const sep = "  "
	left := fmt.Sprintf("  %-8s", label)
	right := value
	pad := width - len(left) - len(right) - len(sep)
	if pad < 0 {
		pad = 0
	}
	return left + sep + strings.Repeat(" ", pad) + right
}

func formatLatencyDuration(d interface{ String() string }) string {
	return d.String()
}

func centerPad(s string, width int) string {
	if len(s) >= width {
		return s
	}
	pad := (width - len(s)) / 2
	return strings.Repeat(" ", pad) + s
}

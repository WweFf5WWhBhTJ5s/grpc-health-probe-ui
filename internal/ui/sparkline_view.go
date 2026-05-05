package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// SparklineBar characters ordered from empty to full.
var sparkBars = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

// SparklineView renders a compact inline sparkline for a named target using
// its recorded history. Width controls the maximum number of bars shown.
func SparklineView(h *History, key string, width int, t Theme) string {
	if h == nil || width <= 0 {
		return ""
	}

	entries := h.Get(key)
	if len(entries) == 0 {
		return t.Dim.Render(strings.Repeat("·", width))
	}

	// Take only the most recent `width` entries.
	if len(entries) > width {
		entries = entries[len(entries)-width:]
	}

	var sb strings.Builder
	for _, e := range entries {
		bar, style := barForEntry(e, t)
		sb.WriteString(style.Render(string(bar)))
	}

	// Pad left with dim dots when fewer entries than width.
	pad := width - len(entries)
	if pad > 0 {
		prefix := t.Dim.Render(strings.Repeat("·", pad))
		return prefix + sb.String()
	}
	return sb.String()
}

// SparklineRow renders a labelled sparkline row suitable for the detail panel.
func SparklineRow(label, key string, h *History, width int, t Theme) string {
	if width < 10 {
		return ""
	}
	barWidth := width - len(label) - 3
	if barWidth < 1 {
		barWidth = 1
	}
	line := SparklineView(h, key, barWidth, t)
	return fmt.Sprintf("%-*s  %s", len(label), label, line)
}

// barForEntry maps a history entry to a sparkline bar character and style.
func barForEntry(e HistoryEntry, t Theme) (rune, lipgloss.Style) {
	switch e.Status {
	case StatusServing:
		idx := latencyIndex(e.LatencyMs)
		return sparkBars[idx], t.Serving
	case StatusNotServing:
		return sparkBars[0], t.NotServing
	default:
		return '?', t.Dim
	}
}

// latencyIndex maps a latency in milliseconds to a bar height index (0–7).
func latencyIndex(ms float64) int {
	switch {
	case ms <= 0:
		return 7 // unknown latency — show full bar
	case ms < 10:
		return 7
	case ms < 25:
		return 6
	case ms < 50:
		return 5
	case ms < 100:
		return 4
	case ms < 200:
		return 3
	case ms < 500:
		return 2
	case ms < 1000:
		return 1
	default:
		return 0
	}
}

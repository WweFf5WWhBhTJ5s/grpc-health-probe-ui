package ui

import (
	"fmt"
	"strings"

	"github.com/user/grpc-health-probe-ui/internal/probe"
)

const (
	// sparkServing is the block character used for SERVING entries.
	sparkServing = "█"
	// sparkDegraded is used for NOT_SERVING / UNKNOWN entries.
	sparkDegraded = "▄"
	// sparkUnreachable is used when the host was unreachable.
	sparkUnreachable = "░"
)

// SparkLine renders a compact, single-line history sparkline for the given
// entries. Each entry is represented by one character. Returns an empty string
// when entries is nil or empty.
func SparkLine(entries []HistoryEntry) string {
	if len(entries) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, e := range entries {
		switch e.Status {
		case probe.StatusServing:
			sb.WriteString(sparkServing)
		case probe.StatusNotServing, probe.StatusUnknown:
			sb.WriteString(sparkDegraded)
		default:
			sb.WriteString(sparkUnreachable)
		}
	}
	return sb.String()
}

// UptimeLabel formats an uptime percentage as a short human-readable string,
// e.g. "100.0%" or " 75.0%" (space-padded to 6 characters).
func UptimeLabel(pct float64) string {
	return fmt.Sprintf("%5.1f%%", pct)
}

// HistoryRow renders a single row combining the sparkline and uptime label
// suitable for embedding in the terminal dashboard view.
// Example output: "█▄░█  75.0%"
func HistoryRow(entries []HistoryEntry, uptime float64) string {
	spark := SparkLine(entries)
	if spark == "" {
		return "no data"
	}
	return fmt.Sprintf("%s  %s", spark, UptimeLabel(uptime))
}

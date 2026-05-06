package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// StatusChange represents a transition in health status for a single target.
type StatusChange struct {
	Target string
	From   grpc_health_v1.HealthCheckResponse_ServingStatus
	To     grpc_health_v1.HealthCheckResponse_ServingStatus
}

// String returns a human-readable description of the status change.
func (sc StatusChange) String() string {
	return fmt.Sprintf("%s: %s → %s", sc.Target, statusLabel(sc.From), statusLabel(sc.To))
}

// DiffResults compares two snapshots of probe results keyed by target name
// and returns a slice of StatusChange for any targets whose status changed.
func DiffResults(
	prev map[string]grpc_health_v1.HealthCheckResponse_ServingStatus,
	next map[string]grpc_health_v1.HealthCheckResponse_ServingStatus,
) []StatusChange {
	var changes []StatusChange
	for key, nextStatus := range next {
		prevStatus, ok := prev[key]
		if !ok {
			// First time we've seen this target — treat UNKNOWN as the prior state.
			prevStatus = grpc_health_v1.HealthCheckResponse_UNKNOWN
		}
		if prevStatus != nextStatus {
			changes = append(changes, StatusChange{
				Target: key,
				From:   prevStatus,
				To:     nextStatus,
			})
		}
	}
	return changes
}

// DiffView renders a compact, styled summary of a slice of StatusChange values
// suitable for embedding in a notification banner or detail panel.
func DiffView(changes []StatusChange, theme Theme, width int) string {
	if len(changes) == 0 {
		return ""
	}

	base := lipgloss.NewStyle().Width(width)
	lines := make([]string, 0, len(changes))
	for _, c := range changes {
		var toStyle lipgloss.Style
		switch c.To {
		case grpc_health_v1.HealthCheckResponse_SERVING:
			toStyle = theme.Serving
		case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
			toStyle = theme.NotServing
		default:
			toStyle = theme.Unknown
		}
		line := fmt.Sprintf("%-30s %s → %s",
			c.Target,
			statusLabel(c.From),
			toStyle.Render(statusLabel(c.To)),
		)
		lines = append(lines, line)
	}

	body := ""
	for _, l := range lines {
		body += l + "\n"
	}
	return base.Render(body)
}

// statusLabel returns a short display label for a serving status.
func statusLabel(s grpc_health_v1.HealthCheckResponse_ServingStatus) string {
	switch s {
	case grpc_health_v1.HealthCheckResponse_SERVING:
		return "SERVING"
	case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
		return "NOT_SERVING"
	case grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN:
		return "SERVICE_UNKNOWN"
	default:
		return "UNKNOWN"
	}
}

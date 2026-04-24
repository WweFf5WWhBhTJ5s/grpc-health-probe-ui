package ui

import (
	"time"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// Status represents the health status of a single gRPC target.
type Status struct {
	Target    string
	Service   string
	State     grpc_health_v1.HealthCheckResponse_ServingStatus
	LastCheck time.Time
	Err       error
}

// StatusLabel returns a human-readable label for the serving status.
func (s Status) StatusLabel() string {
	switch s.State {
	case grpc_health_v1.HealthCheckResponse_SERVING:
		return "SERVING"
	case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
		return "NOT_SERVING"
	case grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN:
		return "UNKNOWN"
	default:
		return "UNREACHABLE"
	}
}

// StatusColor returns an ANSI color code string for terminal rendering.
func (s Status) StatusColor() string {
	switch s.State {
	case grpc_health_v1.HealthCheckResponse_SERVING:
		return "\033[32m" // green
	case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
		return "\033[31m" // red
	default:
		return "\033[33m" // yellow
	}
}

// Model holds the full UI state for the dashboard.
type Model struct {
	Statuses  []Status
	UpdatedAt time.Time
	Interval  time.Duration
}

// NewModel creates an empty Model with the given poll interval.
func NewModel(interval time.Duration) *Model {
	return &Model{
		Interval: interval,
	}
}

// Update replaces the status list and records the update timestamp.
func (m *Model) Update(statuses []Status) {
	m.Statuses = statuses
	m.UpdatedAt = time.Now()
}

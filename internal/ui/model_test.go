package ui

import (
	"errors"
	"testing"
	"time"

	"google.golang.org/grpc/health/grpc_health_v1"
)

func TestStatusLabel(t *testing.T) {
	cases := []struct {
		state    grpc_health_v1.HealthCheckResponse_ServingStatus
		expected string
	}{
		{grpc_health_v1.HealthCheckResponse_SERVING, "SERVING"},
		{grpc_health_v1.HealthCheckResponse_NOT_SERVING, "NOT_SERVING"},
		{grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN, "UNKNOWN"},
		{grpc_health_v1.HealthCheckResponse_UNKNOWN, "UNREACHABLE"},
	}
	for _, tc := range cases {
		s := Status{State: tc.state}
		if got := s.StatusLabel(); got != tc.expected {
			t.Errorf("StatusLabel() for state %v = %q, want %q", tc.state, got, tc.expected)
		}
	}
}

func TestStatusColor(t *testing.T) {
	serving := Status{State: grpc_health_v1.HealthCheckResponse_SERVING}
	if serving.StatusColor() != "\033[32m" {
		t.Errorf("expected green for SERVING")
	}

	notServing := Status{State: grpc_health_v1.HealthCheckResponse_NOT_SERVING}
	if notServing.StatusColor() != "\033[31m" {
		t.Errorf("expected red for NOT_SERVING")
	}

	unknown := Status{State: grpc_health_v1.HealthCheckResponse_UNKNOWN}
	if unknown.StatusColor() != "\033[33m" {
		t.Errorf("expected yellow for UNKNOWN")
	}
}

func TestNewModel(t *testing.T) {
	interval := 5 * time.Second
	m := NewModel(interval)
	if m == nil {
		t.Fatal("NewModel returned nil")
	}
	if m.Interval != interval {
		t.Errorf("Interval = %v, want %v", m.Interval, interval)
	}
	if len(m.Statuses) != 0 {
		t.Errorf("expected empty Statuses, got %d", len(m.Statuses))
	}
}

func TestModelUpdate(t *testing.T) {
	m := NewModel(time.Second)
	before := time.Now()

	statuses := []Status{
		{Target: "localhost:50051", Service: "svc", State: grpc_health_v1.HealthCheckResponse_SERVING},
		{Target: "localhost:50052", Service: "svc", Err: errors.New("connection refused")},
	}
	m.Update(statuses)

	if len(m.Statuses) != 2 {
		t.Errorf("expected 2 statuses, got %d", len(m.Statuses))
	}
	if m.UpdatedAt.Before(before) {
		t.Errorf("UpdatedAt not set correctly")
	}
}

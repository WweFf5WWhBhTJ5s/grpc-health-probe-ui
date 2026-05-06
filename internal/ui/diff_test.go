package ui

import (
	"strings"
	"testing"

	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	serving    = grpc_health_v1.HealthCheckResponse_SERVING
	notServing = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	unknown    = grpc_health_v1.HealthCheckResponse_UNKNOWN
)

func TestDiffResults_NoChanges(t *testing.T) {
	prev := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{
		"svc-a": serving,
	}
	next := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{
		"svc-a": serving,
	}
	changes := DiffResults(prev, next)
	if len(changes) != 0 {
		t.Fatalf("expected 0 changes, got %d", len(changes))
	}
}

func TestDiffResults_DetectsChange(t *testing.T) {
	prev := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{
		"svc-a": serving,
	}
	next := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{
		"svc-a": notServing,
	}
	changes := DiffResults(prev, next)
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].Target != "svc-a" {
		t.Errorf("unexpected target: %s", changes[0].Target)
	}
	if changes[0].From != serving {
		t.Errorf("expected From=SERVING, got %v", changes[0].From)
	}
	if changes[0].To != notServing {
		t.Errorf("expected To=NOT_SERVING, got %v", changes[0].To)
	}
}

func TestDiffResults_NewTargetTreatedAsUnknown(t *testing.T) {
	prev := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{}
	next := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{
		"svc-b": serving,
	}
	changes := DiffResults(prev, next)
	if len(changes) != 1 {
		t.Fatalf("expected 1 change for new target, got %d", len(changes))
	}
	if changes[0].From != unknown {
		t.Errorf("expected From=UNKNOWN for new target, got %v", changes[0].From)
	}
}

func TestDiffResults_MultipleChanges(t *testing.T) {
	prev := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{
		"svc-a": serving,
		"svc-b": notServing,
		"svc-c": serving,
	}
	next := map[string]grpc_health_v1.HealthCheckResponse_ServingStatus{
		"svc-a": notServing,
		"svc-b": notServing, // unchanged
		"svc-c": unknown,
	}
	changes := DiffResults(prev, next)
	if len(changes) != 2 {
		t.Fatalf("expected 2 changes, got %d", len(changes))
	}
}

func TestStatusChange_String(t *testing.T) {
	sc := StatusChange{Target: "my-service", From: serving, To: notServing}
	s := sc.String()
	if !strings.Contains(s, "my-service") {
		t.Errorf("expected target name in string, got: %s", s)
	}
	if !strings.Contains(s, "SERVING") {
		t.Errorf("expected SERVING in string, got: %s", s)
	}
	if !strings.Contains(s, "NOT_SERVING") {
		t.Errorf("expected NOT_SERVING in string, got: %s", s)
	}
}

func TestDiffView_EmptyChanges(t *testing.T) {
	theme := DefaultTheme()
	out := DiffView(nil, theme, 80)
	if out != "" {
		t.Errorf("expected empty output for nil changes, got: %q", out)
	}
}

func TestDiffView_ContainsTargetName(t *testing.T) {
	changes := []StatusChange{
		{Target: "payment-service", From: serving, To: notServing},
	}
	theme := DefaultTheme()
	out := DiffView(changes, theme, 80)
	if !strings.Contains(out, "payment-service") {
		t.Errorf("expected target name in diff view output, got: %q", out)
	}
}

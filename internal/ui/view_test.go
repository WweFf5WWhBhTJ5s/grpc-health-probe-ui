package ui

import (
	"strings"
	"testing"
	"time"
)

func makeModel(targets []string, statuses map[string]Status) Model {
	m := NewModel(targets, time.Second*5)
	for target, status := range statuses {
		m.Statuses[target] = status
	}
	return m
}

func TestView_ContainsTitle(t *testing.T) {
	m := makeModel([]string{"localhost:50051"}, nil)
	view := m.View()
	if !strings.Contains(view, "gRPC Health Probe") {
		t.Errorf("expected view to contain title, got:\n%s", view)
	}
}

func TestView_NoTargets(t *testing.T) {
	m := makeModel([]string{}, nil)
	view := m.View()
	if !strings.Contains(view, "No targets configured") {
		t.Errorf("expected 'No targets configured' message, got:\n%s", view)
	}
}

func TestView_ShowsTargets(t *testing.T) {
	targets := []string{"host1:50051", "host2:50052"}
	m := makeModel(targets, map[string]Status{
		"host1:50051": StatusServing,
		"host2:50052": StatusNotServing,
	})
	view := m.View()
	for _, target := range targets {
		if !strings.Contains(view, target) {
			t.Errorf("expected view to contain target %q", target)
		}
	}
}

func TestView_ShowsStatusLabels(t *testing.T) {
	m := makeModel([]string{"svc:9000"}, map[string]Status{
		"svc:9000": StatusServing,
	})
	view := m.View()
	if !strings.Contains(view, "SERVING") {
		t.Errorf("expected SERVING label in view, got:\n%s", view)
	}
}

func TestView_ShowsFooterWithInterval(t *testing.T) {
	m := makeModel([]string{"svc:9000"}, nil)
	view := m.View()
	if !strings.Contains(view, "5s") {
		t.Errorf("expected interval in footer, got:\n%s", view)
	}
	if !strings.Contains(view, "press q to quit") {
		t.Errorf("expected quit hint in footer, got:\n%s", view)
	}
}

func TestView_UnknownStatusForMissingEntry(t *testing.T) {
	m := makeModel([]string{"svc:9000"}, nil)
	view := m.View()
	if !strings.Contains(view, "UNKNOWN") {
		t.Errorf("expected UNKNOWN status for target with no status entry, got:\n%s", view)
	}
}

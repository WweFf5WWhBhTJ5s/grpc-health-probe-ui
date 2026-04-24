package ui

import (
	"strings"
	"testing"
	"time"

	"github.com/nicholasgasior/grpc-health-probe-ui/internal/probe"
)

func makeDetailTarget() probe.Target {
	return probe.Target{
		Name:        "auth-service",
		Host:        "localhost:50051",
		Status:      probe.StatusServing,
		Message:     "",
		LatencyMs:   42,
		LastChecked: time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC),
	}
}

func TestDetailView_ContainsName(t *testing.T) {
	tgt := makeDetailTarget()
	out := DetailView(tgt, nil, 80)
	if !strings.Contains(out, "auth-service") {
		t.Errorf("expected output to contain target name, got:\n%s", out)
	}
}

func TestDetailView_ContainsHost(t *testing.T) {
	tgt := makeDetailTarget()
	out := DetailView(tgt, nil, 80)
	if !strings.Contains(out, "localhost:50051") {
		t.Errorf("expected output to contain host, got:\n%s", out)
	}
}

func TestDetailView_ContainsStatus(t *testing.T) {
	tgt := makeDetailTarget()
	out := DetailView(tgt, nil, 80)
	if !strings.Contains(out, "SERVING") {
		t.Errorf("expected output to contain status label, got:\n%s", out)
	}
}

func TestDetailView_ContainsLatency(t *testing.T) {
	tgt := makeDetailTarget()
	out := DetailView(tgt, nil, 80)
	if !strings.Contains(out, "42 ms") {
		t.Errorf("expected output to contain latency, got:\n%s", out)
	}
}

func TestDetailView_NegativeLatencyShowsDash(t *testing.T) {
	tgt := makeDetailTarget()
	tgt.LatencyMs = -1
	out := DetailView(tgt, nil, 80)
	if !strings.Contains(out, "—") {
		t.Errorf("expected dash for negative latency, got:\n%s", out)
	}
}

func TestDetailView_EmptyMessageShowsDash(t *testing.T) {
	tgt := makeDetailTarget()
	tgt.Message = ""
	out := DetailView(tgt, nil, 80)
	if !strings.Contains(out, "—") {
		t.Errorf("expected dash for empty message, got:\n%s", out)
	}
}

func TestDetailView_WithHistory(t *testing.T) {
	tgt := makeDetailTarget()
	history := []probe.StatusEntry{
		{Status: probe.StatusServing},
		{Status: probe.StatusNotServing},
		{Status: probe.StatusServing},
	}
	out := DetailView(tgt, history, 80)
	if !strings.Contains(out, "Uptime") {
		t.Errorf("expected uptime row when history provided, got:\n%s", out)
	}
}

func TestDetailView_ZeroWidth(t *testing.T) {
	tgt := makeDetailTarget()
	// should not panic
	out := DetailView(tgt, nil, 0)
	if out == "" {
		t.Error("expected non-empty output even with zero width")
	}
}

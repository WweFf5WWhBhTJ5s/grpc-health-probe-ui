package ui

import (
	"testing"
	"time"

	"github.com/mitchellh/go-grpc-health-probe-ui/internal/probe"
)

func makeResults(statuses []probe.Status, latencies []time.Duration) []probe.Result {
	out := make([]probe.Result, len(statuses))
	for i, st := range statuses {
		l := time.Duration(0)
		if i < len(latencies) {
			l = latencies[i]
		}
		out[i] = probe.Result{Status: st, Latency: l}
	}
	return out
}

func TestComputeMetrics_Empty(t *testing.T) {
	s := ComputeMetrics("svc", nil)
	if s.Total != 0 {
		t.Fatalf("expected 0 total, got %d", s.Total)
	}
	if s.UptimePct != 0 {
		t.Fatalf("expected 0 uptime, got %f", s.UptimePct)
	}
}

func TestComputeMetrics_AllServing(t *testing.T) {
	entries := makeResults(
		[]probe.Status{probe.StatusServing, probe.StatusServing, probe.StatusServing},
		[]time.Duration{10 * time.Millisecond, 20 * time.Millisecond, 30 * time.Millisecond},
	)
	s := ComputeMetrics("svc", entries)
	if s.Serving != 3 {
		t.Fatalf("expected 3 serving, got %d", s.Serving)
	}
	if s.UptimePct != 100.0 {
		t.Fatalf("expected 100%% uptime, got %f", s.UptimePct)
	}
	if s.AvgLatency != 20*time.Millisecond {
		t.Fatalf("expected 20ms avg, got %v", s.AvgLatency)
	}
	if s.MinLatency != 10*time.Millisecond {
		t.Fatalf("expected 10ms min, got %v", s.MinLatency)
	}
	if s.MaxLatency != 30*time.Millisecond {
		t.Fatalf("expected 30ms max, got %v", s.MaxLatency)
	}
}

func TestComputeMetrics_Mixed(t *testing.T) {
	entries := makeResults(
		[]probe.Status{probe.StatusServing, probe.StatusNotServing, probe.StatusUnknown},
		[]time.Duration{5 * time.Millisecond, 0, 0},
	)
	s := ComputeMetrics("svc", entries)
	if s.Total != 3 {
		t.Fatalf("expected 3 total, got %d", s.Total)
	}
	if s.NotServing != 1 {
		t.Fatalf("expected 1 not-serving, got %d", s.NotServing)
	}
	if s.Unknown != 1 {
		t.Fatalf("expected 1 unknown, got %d", s.Unknown)
	}
	expected := 100.0 / 3.0
	if s.UptimePct < expected-0.1 || s.UptimePct > expected+0.1 {
		t.Fatalf("expected ~33.3%% uptime, got %f", s.UptimePct)
	}
}

func TestComputeMetrics_TargetName(t *testing.T) {
	s := ComputeMetrics("my-service", nil)
	if s.Target != "my-service" {
		t.Fatalf("expected target 'my-service', got %q", s.Target)
	}
}

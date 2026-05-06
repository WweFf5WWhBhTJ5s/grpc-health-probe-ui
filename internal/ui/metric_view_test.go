package ui

import (
	"strings"
	"testing"
	"time"

	"github.com/mitchellh/go-grpc-health-probe-ui/internal/probe"
)

func makeMetricSummary() MetricSummary {
	entries := makeResults(
		[]probe.Status{probe.StatusServing, probe.StatusServing, probe.StatusNotServing},
		[]time.Duration{10 * time.Millisecond, 30 * time.Millisecond, 0},
	)
	return ComputeMetrics("my-svc", entries)
}

func TestMetricView_ContainsTarget(t *testing.T) {
	s := makeMetricSummary()
	out := MetricView(s, 60)
	if !strings.Contains(out, "my-svc") {
		t.Fatalf("expected target name in output, got:\n%s", out)
	}
}

func TestMetricView_ContainsUptime(t *testing.T) {
	s := makeMetricSummary()
	out := MetricView(s, 60)
	if !strings.Contains(out, "Uptime") {
		t.Fatalf("expected 'Uptime' in output, got:\n%s", out)
	}
}

func TestMetricView_ContainsLatencyFields(t *testing.T) {
	s := makeMetricSummary()
	out := MetricView(s, 60)
	for _, want := range []string{"Avg latency", "Min latency", "Max latency"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected %q in output, got:\n%s", want, out)
		}
	}
}

func TestMetricView_ZeroWidth_UsesDefault(t *testing.T) {
	s := makeMetricSummary()
	out := MetricView(s, 0)
	if out == "" {
		t.Fatal("expected non-empty output with zero width")
	}
}

func TestMetricView_ServingCount(t *testing.T) {
	s := makeMetricSummary()
	out := MetricView(s, 60)
	if !strings.Contains(out, "Serving") {
		t.Fatalf("expected 'Serving' in output, got:\n%s", out)
	}
}

func TestMetricView_EmptySummary(t *testing.T) {
	s := ComputeMetrics("empty-svc", nil)
	out := MetricView(s, 60)
	if !strings.Contains(out, "0") {
		t.Fatalf("expected '0' counts in output, got:\n%s", out)
	}
}

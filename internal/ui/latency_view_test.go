package ui

import (
	"strings"
	"testing"
	"time"
)

func makeLatencyStats(min, mean, p50, p95, p99, max int, count int) LatencyStats {
	return LatencyStats{
		Min:   time.Duration(min) * time.Millisecond,
		Mean:  time.Duration(mean) * time.Millisecond,
		P50:   time.Duration(p50) * time.Millisecond,
		P95:   time.Duration(p95) * time.Millisecond,
		P99:   time.Duration(p99) * time.Millisecond,
		Max:   time.Duration(max) * time.Millisecond,
		Count: count,
	}
}

func TestLatencyView_NoData(t *testing.T) {
	out := LatencyView(LatencyStats{}, 40)
	if !strings.Contains(out, "no latency data") {
		t.Errorf("expected no-data message, got: %q", out)
	}
}

func TestLatencyView_ContainsLabels(t *testing.T) {
	stats := makeLatencyStats(1, 5, 5, 9, 10, 10, 10)
	out := LatencyView(stats, 60)
	for _, label := range []string{"min", "mean", "p50", "p95", "p99", "max", "samples"} {
		if !strings.Contains(out, label) {
			t.Errorf("expected label %q in output", label)
		}
	}
}

func TestLatencyView_ContainsSampleCount(t *testing.T) {
	stats := makeLatencyStats(1, 5, 5, 9, 10, 10, 42)
	out := LatencyView(stats, 60)
	if !strings.Contains(out, "42") {
		t.Errorf("expected sample count 42 in output, got: %q", out)
	}
}

func TestLatencyView_ZeroWidth_UsesDefault(t *testing.T) {
	stats := makeLatencyStats(1, 5, 5, 9, 10, 10, 5)
	out := LatencyView(stats, 0)
	if out == "" {
		t.Error("expected non-empty output with zero width")
	}
}

func TestLatencyView_ContainsDurationValues(t *testing.T) {
	stats := makeLatencyStats(2, 5, 5, 8, 9, 10, 3)
	out := LatencyView(stats, 60)
	if !strings.Contains(out, "ms") {
		t.Errorf("expected duration units in output, got: %q", out)
	}
}

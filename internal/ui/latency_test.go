package ui

import (
	"testing"
	"time"
)

func makeSamples(ms ...int) []time.Duration {
	out := make([]time.Duration, len(ms))
	for i, m := range ms {
		out[i] = time.Duration(m) * time.Millisecond
	}
	return out
}

func TestComputeLatencyStats_Empty(t *testing.T) {
	stats := ComputeLatencyStats(nil)
	if stats.Count != 0 {
		t.Fatalf("expected count 0, got %d", stats.Count)
	}
}

func TestComputeLatencyStats_Single(t *testing.T) {
	stats := ComputeLatencyStats(makeSamples(42))
	if stats.Count != 1 {
		t.Fatalf("expected count 1, got %d", stats.Count)
	}
	if stats.Min != 42*time.Millisecond {
		t.Errorf("unexpected min: %v", stats.Min)
	}
	if stats.Max != 42*time.Millisecond {
		t.Errorf("unexpected max: %v", stats.Max)
	}
}

func TestComputeLatencyStats_Percentiles(t *testing.T) {
	samples := makeSamples(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	stats := ComputeLatencyStats(samples)

	if stats.Min != 10*time.Millisecond {
		t.Errorf("expected min 10ms, got %v", stats.Min)
	}
	if stats.Max != 100*time.Millisecond {
		t.Errorf("expected max 100ms, got %v", stats.Max)
	}
	if stats.P50 != 50*time.Millisecond {
		t.Errorf("expected p50 50ms, got %v", stats.P50)
	}
	if stats.P95 != 95*time.Millisecond {
		t.Errorf("expected p95 95ms, got %v", stats.P95)
	}
	if stats.P99 != 99*time.Millisecond {
		t.Errorf("expected p99 99ms, got %v", stats.P99)
	}
}

func TestComputeLatencyStats_Mean(t *testing.T) {
	samples := makeSamples(10, 20, 30)
	stats := ComputeLatencyStats(samples)
	if stats.Mean != 20*time.Millisecond {
		t.Errorf("expected mean 20ms, got %v", stats.Mean)
	}
}

func TestComputeLatencyStats_DoesNotMutateInput(t *testing.T) {
	samples := makeSamples(30, 10, 20)
	original := make([]time.Duration, len(samples))
	copy(original, samples)
	ComputeLatencyStats(samples)
	for i, d := range samples {
		if d != original[i] {
			t.Errorf("input mutated at index %d", i)
		}
	}
}

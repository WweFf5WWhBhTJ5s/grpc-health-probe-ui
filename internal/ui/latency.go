package ui

import (
	"math"
	"sort"
	"time"
)

// LatencyStats holds computed percentile and aggregate latency metrics.
type LatencyStats struct {
	Min    time.Duration
	Max    time.Duration
	Mean   time.Duration
	P50    time.Duration
	P95    time.Duration
	P99    time.Duration
	Count  int
}

// ComputeLatencyStats calculates latency statistics from a slice of durations.
// Returns a zero-value LatencyStats if samples is empty.
func ComputeLatencyStats(samples []time.Duration) LatencyStats {
	if len(samples) == 0 {
		return LatencyStats{}
	}

	sorted := make([]time.Duration, len(samples))
	copy(sorted, samples)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	var total time.Duration
	for _, d := range sorted {
		total += d
	}

	return LatencyStats{
		Min:   sorted[0],
		Max:   sorted[len(sorted)-1],
		Mean:  time.Duration(int64(total) / int64(len(sorted))),
		P50:   percentile(sorted, 50),
		P95:   percentile(sorted, 95),
		P99:   percentile(sorted, 99),
		Count: len(sorted),
	}
}

// percentile returns the p-th percentile value from a sorted slice.
func percentile(sorted []time.Duration, p float64) time.Duration {
	if len(sorted) == 0 {
		return 0
	}
	index := int(math.Ceil(float64(len(sorted))*p/100.0)) - 1
	if index < 0 {
		index = 0
	}
	if index >= len(sorted) {
		index = len(sorted) - 1
	}
	return sorted[index]
}

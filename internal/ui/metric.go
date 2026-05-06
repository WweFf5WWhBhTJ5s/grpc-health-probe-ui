package ui

import (
	"math"
	"time"

	"github.com/mitchellh/go-grpc-health-probe-ui/internal/probe"
)

// MetricSummary holds aggregated statistics for a single target's history.
type MetricSummary struct {
	Target      string
	Total       int
	Serving     int
	NotServing  int
	Unknown     int
	UptimePct   float64
	AvgLatency  time.Duration
	MaxLatency  time.Duration
	MinLatency  time.Duration
}

// ComputeMetrics derives a MetricSummary from a slice of history entries.
func ComputeMetrics(target string, entries []probe.Result) MetricSummary {
	if len(entries) == 0 {
		return MetricSummary{Target: target, MinLatency: 0}
	}

	s := MetricSummary{
		Target:     target,
		Total:      len(entries),
		MinLatency: time.Duration(math.MaxInt64),
	}

	var totalLatency time.Duration

	for _, e := range entries {
		switch e.Status {
		case probe.StatusServing:
			s.Serving++
		case probe.StatusNotServing:
			s.NotServing++
		default:
			s.Unknown++
		}

		if e.Latency > s.MaxLatency {
			s.MaxLatency = e.Latency
		}
		if e.Latency > 0 && e.Latency < s.MinLatency {
			s.MinLatency = e.Latency
		}
		totalLatency += e.Latency
	}

	if s.MinLatency == time.Duration(math.MaxInt64) {
		s.MinLatency = 0
	}

	if s.Total > 0 {
		s.UptimePct = float64(s.Serving) / float64(s.Total) * 100.0
		s.AvgLatency = totalLatency / time.Duration(s.Total)
	}

	return s
}

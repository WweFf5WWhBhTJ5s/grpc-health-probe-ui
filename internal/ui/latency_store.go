package ui

import "time"

const defaultLatencyMaxLen = 200

// LatencyStore accumulates raw latency samples per target key and
// provides computed statistics on demand.
type LatencyStore struct {
	samples map[string][]time.Duration
	maxLen  int
}

// NewLatencyStore returns a LatencyStore with the given maximum sample
// window per target. If maxLen <= 0, defaultLatencyMaxLen is used.
func NewLatencyStore(maxLen int) *LatencyStore {
	if maxLen <= 0 {
		maxLen = defaultLatencyMaxLen
	}
	return &LatencyStore{
		samples: make(map[string][]time.Duration),
		maxLen:  maxLen,
	}
}

// Record appends a latency sample for the given key, evicting the oldest
// entry when the window is full.
func (s *LatencyStore) Record(key string, d time.Duration) {
	buf := s.samples[key]
	buf = append(buf, d)
	if len(buf) > s.maxLen {
		buf = buf[len(buf)-s.maxLen:]
	}
	s.samples[key] = buf
}

// Stats returns computed LatencyStats for the given key.
// Returns a zero-value LatencyStats if no samples exist.
func (s *LatencyStore) Stats(key string) LatencyStats {
	return ComputeLatencyStats(s.samples[key])
}

// Clear removes all samples for the given key.
func (s *LatencyStore) Clear(key string) {
	delete(s.samples, key)
}

// Keys returns all target keys that have recorded samples.
func (s *LatencyStore) Keys() []string {
	keys := make([]string, 0, len(s.samples))
	for k := range s.samples {
		keys = append(keys, k)
	}
	return keys
}

// SampleCount returns the number of recorded samples for the given key.
func (s *LatencyStore) SampleCount(key string) int {
	return len(s.samples[key])
}

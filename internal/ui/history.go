package ui

import (
	"sync"
	"time"

	"github.com/user/grpc-health-probe-ui/internal/probe"
)

// HistoryEntry records a single probe result with a timestamp.
type HistoryEntry struct {
	Status    probe.Status
	CheckedAt time.Time
}

// History tracks recent probe results per target key (host:service).
type History struct {
	mu      sync.RWMutex
	entries map[string][]HistoryEntry
	maxLen  int
}

// NewHistory creates a History that retains at most maxLen entries per target.
func NewHistory(maxLen int) *History {
	if maxLen <= 0 {
		maxLen = 10
	}
	return &History{
		entries: make(map[string][]HistoryEntry),
		maxLen:  maxLen,
	}
}

// Record appends a new entry for the given key, evicting the oldest if needed.
func (h *History) Record(key string, status probe.Status, at time.Time) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.entries[key] = append(h.entries[key], HistoryEntry{Status: status, CheckedAt: at})
	if len(h.entries[key]) > h.maxLen {
		h.entries[key] = h.entries[key][len(h.entries[key])-h.maxLen:]
	}
}

// Get returns a copy of the history entries for the given key.
func (h *History) Get(key string) []HistoryEntry {
	h.mu.RLock()
	defer h.mu.RUnlock()

	src := h.entries[key]
	if len(src) == 0 {
		return nil
	}
	out := make([]HistoryEntry, len(src))
	copy(out, src)
	return out
}

// UptimePercent returns the fraction of SERVING entries over all recorded
// entries for key, or 0 if no entries exist.
func (h *History) UptimePercent(key string) float64 {
	h.mu.RLock()
	defer h.mu.RUnlock()

	entries := h.entries[key]
	if len(entries) == 0 {
		return 0
	}
	var serving int
	for _, e := range entries {
		if e.Status == probe.StatusServing {
			serving++
		}
	}
	return float64(serving) / float64(len(entries)) * 100
}

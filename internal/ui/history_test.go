package ui

import (
	"testing"
	"time"

	"github.com/user/grpc-health-probe-ui/internal/probe"
)

func TestNewHistory_DefaultMaxLen(t *testing.T) {
	h := NewHistory(0)
	if h.maxLen != 10 {
		t.Fatalf("expected maxLen 10, got %d", h.maxLen)
	}
}

func TestHistory_RecordAndGet(t *testing.T) {
	h := NewHistory(5)
	now := time.Now()
	h.Record("host:svc", probe.StatusServing, now)
	h.Record("host:svc", probe.StatusNotServing, now.Add(time.Second))

	entries := h.Get("host:svc")
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	if entries[0].Status != probe.StatusServing {
		t.Errorf("expected first entry SERVING")
	}
	if entries[1].Status != probe.StatusNotServing {
		t.Errorf("expected second entry NOT_SERVING")
	}
}

func TestHistory_EvictsOldEntries(t *testing.T) {
	h := NewHistory(3)
	now := time.Now()
	for i := 0; i < 5; i++ {
		h.Record("k", probe.StatusServing, now.Add(time.Duration(i)*time.Second))
	}
	entries := h.Get("k")
	if len(entries) != 3 {
		t.Fatalf("expected 3 entries after eviction, got %d", len(entries))
	}
}

func TestHistory_GetUnknownKey(t *testing.T) {
	h := NewHistory(5)
	if entries := h.Get("missing"); entries != nil {
		t.Errorf("expected nil for unknown key, got %v", entries)
	}
}

func TestHistory_GetReturnsCopy(t *testing.T) {
	h := NewHistory(5)
	now := time.Now()
	h.Record("k", probe.StatusServing, now)

	copy1 := h.Get("k")
	copy1[0].Status = probe.StatusUnknown

	copy2 := h.Get("k")
	if copy2[0].Status != probe.StatusServing {
		t.Errorf("Get should return an independent copy")
	}
}

func TestHistory_UptimePercent(t *testing.T) {
	h := NewHistory(10)
	now := time.Now()
	h.Record("k", probe.StatusServing, now)
	h.Record("k", probe.StatusServing, now.Add(time.Second))
	h.Record("k", probe.StatusNotServing, now.Add(2*time.Second))
	h.Record("k", probe.StatusUnreachable, now.Add(3*time.Second))

	got := h.UptimePercent("k")
	want := 50.0
	if got != want {
		t.Errorf("UptimePercent: got %.1f, want %.1f", got, want)
	}
}

func TestHistory_UptimePercent_NoEntries(t *testing.T) {
	h := NewHistory(5)
	if p := h.UptimePercent("missing"); p != 0 {
		t.Errorf("expected 0 for missing key, got %.1f", p)
	}
}

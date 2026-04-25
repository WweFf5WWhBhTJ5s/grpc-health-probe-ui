package ui

import (
	"strings"
	"testing"
	"time"
)

// TestAlertLifecycle exercises the full push → render → prune cycle.
func TestAlertLifecycle(t *testing.T) {
	now := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	q := NewAlertQueue(5)

	// Push two alerts with different TTLs.
	q.Push(Alert{Message: "short-lived", Level: AlertWarn, CreatedAt: now, TTL: 2 * time.Second})
	q.Push(Alert{Message: "long-lived", Level: AlertInfo, CreatedAt: now, TTL: time.Minute})

	// Before pruning both should appear in the banner.
	banner := AlertBanner(q, 120)
	if !strings.Contains(banner, "short-lived") {
		t.Errorf("expected 'short-lived' in banner before prune")
	}
	if !strings.Contains(banner, "long-lived") {
		t.Errorf("expected 'long-lived' in banner before prune")
	}

	// Advance time past the short-lived TTL and prune.
	q.Prune(now.Add(3 * time.Second))

	if q.Len() != 1 {
		t.Fatalf("expected 1 alert after prune, got %d", q.Len())
	}

	banner = AlertBanner(q, 120)
	if strings.Contains(banner, "short-lived") {
		t.Error("'short-lived' should not appear after prune")
	}
	if !strings.Contains(banner, "long-lived") {
		t.Error("'long-lived' should still appear after prune")
	}
}

// TestAlertQueue_CapacityEnforced ensures the queue never exceeds its max size.
func TestAlertQueue_CapacityEnforced(t *testing.T) {
	const cap = 3
	q := NewAlertQueue(cap)
	for i := 0; i < 10; i++ {
		q.Push(Alert{Message: "msg", Level: AlertInfo, CreatedAt: time.Now(), TTL: time.Minute})
	}
	if q.Len() > cap {
		t.Errorf("queue exceeded capacity: len=%d, max=%d", q.Len(), cap)
	}
}

package ui

import (
	"strings"
	"testing"
	"time"
)

func makeQueue(alerts ...Alert) *AlertQueue {
	q := NewAlertQueue(10)
	for _, a := range alerts {
		q.Push(a)
	}
	return q
}

func TestAlertBanner_Empty(t *testing.T) {
	q := NewAlertQueue(5)
	if out := AlertBanner(q, 80); out != "" {
		t.Errorf("expected empty banner, got %q", out)
	}
}

func TestAlertBanner_Nil(t *testing.T) {
	if out := AlertBanner(nil, 80); out != "" {
		t.Errorf("expected empty banner for nil queue, got %q", out)
	}
}

func TestAlertBanner_ContainsMessage(t *testing.T) {
	q := makeQueue(Alert{
		Message:   "probe failed",
		Level:     AlertError,
		CreatedAt: time.Now(),
		TTL:       time.Minute,
	})
	out := AlertBanner(q, 80)
	if !strings.Contains(out, "probe failed") {
		t.Errorf("banner missing message, got: %q", out)
	}
}

func TestAlertBanner_MultipleAlerts(t *testing.T) {
	now := time.Now()
	q := makeQueue(
		Alert{Message: "alpha", Level: AlertInfo, CreatedAt: now, TTL: time.Minute},
		Alert{Message: "beta", Level: AlertWarn, CreatedAt: now, TTL: time.Minute},
	)
	out := AlertBanner(q, 80)
	if !strings.Contains(out, "alpha") || !strings.Contains(out, "beta") {
		t.Errorf("banner missing one or more messages: %q", out)
	}
}

func TestAlertBanner_TruncatesLongMessage(t *testing.T) {
	long := strings.Repeat("x", 200)
	q := makeQueue(Alert{Message: long, Level: AlertInfo, CreatedAt: time.Now(), TTL: time.Minute})
	out := AlertBanner(q, 40)
	// Strip ANSI for length check
	raw := stripANSI(out)
	if len(raw) > 41 {
		t.Errorf("expected truncation to width 40, got len %d", len(raw))
	}
}

// stripANSI is a minimal helper to remove ANSI escape sequences for length checks.
func stripANSI(s string) string {
	var out strings.Builder
	inEsc := false
	for _, r := range s {
		switch {
		case r == '\x1b':
			inEsc = true
		case inEsc && r == 'm':
			inEsc = false
		case !inEsc:
			out.WriteRune(r)
		}
	}
	return out.String()
}

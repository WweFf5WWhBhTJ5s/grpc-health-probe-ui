package ui

import (
	"strings"
	"testing"
	"time"
)

func TestStatusBar_NeverRefreshed(t *testing.T) {
	sb := StatusBar{
		LastRefresh: time.Time{},
		SortOrder:   SortByName,
		Width:       80,
	}
	out := sb.Render()
	if !strings.Contains(out, "never") {
		t.Errorf("expected 'never' in output, got: %s", out)
	}
}

func TestStatusBar_ShowsRefreshTime(t *testing.T) {
	ts := time.Date(2024, 1, 15, 9, 30, 0, 0, time.UTC)
	sb := StatusBar{
		LastRefresh: ts,
		SortOrder:   SortByName,
		Width:       100,
	}
	out := sb.Render()
	if !strings.Contains(out, "09:30:00") {
		t.Errorf("expected time '09:30:00' in output, got: %s", out)
	}
}

func TestStatusBar_ShowsSortOrder(t *testing.T) {
	cases := []struct {
		order SortOrder
		label string
	}{
		{SortByName, "name"},
		{SortByHost, "host"},
		{SortByStatus, "status"},
	}
	for _, c := range cases {
		sb := StatusBar{SortOrder: c.order, Width: 100}
		out := sb.Render()
		if !strings.Contains(out, c.label) {
			t.Errorf("expected sort label %q in output for order %v", c.label, c.order)
		}
	}
}

func TestStatusBar_ContainsKeyHints(t *testing.T) {
	sb := StatusBar{Width: 120}
	out := sb.Render()
	for _, hint := range []string{"quit", "refresh", "sort", "scroll"} {
		if !strings.Contains(out, hint) {
			t.Errorf("expected key hint %q in status bar output", hint)
		}
	}
}

func TestStatusBar_ZeroWidth(t *testing.T) {
	sb := StatusBar{Width: 0, LastRefresh: time.Now(), SortOrder: SortByName}
	// Should not panic
	_ = sb.Render()
}

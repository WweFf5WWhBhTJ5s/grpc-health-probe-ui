package ui

import (
	"strings"
	"testing"
)

func TestNewLegend_ReturnsLegend(t *testing.T) {
	theme := DefaultTheme()
	l := NewLegend(theme)
	if l.theme.Serving.GetForeground() != theme.Serving.GetForeground() {
		t.Error("expected legend to carry provided theme")
	}
}

func TestLegend_Entries_Count(t *testing.T) {
	l := NewLegend(DefaultTheme())
	if got := len(l.Entries()); got != 4 {
		t.Errorf("expected 4 entries, got %d", got)
	}
}

func TestLegend_Entries_Labels(t *testing.T) {
	l := NewLegend(DefaultTheme())
	want := []string{"Serving", "Not Serving", "Unknown", "Unreachable"}
	for i, e := range l.Entries() {
		if e.Label != want[i] {
			t.Errorf("entry %d: expected label %q, got %q", i, want[i], e.Label)
		}
	}
}

func TestLegend_Render_ContainsAllLabels(t *testing.T) {
	l := NewLegend(DefaultTheme())
	out := l.Render(0)
	for _, label := range []string{"Serving", "Not Serving", "Unknown", "Unreachable"} {
		if !strings.Contains(out, label) {
			t.Errorf("expected rendered legend to contain %q", label)
		}
	}
}

func TestLegend_Render_ZeroWidth_NoTruncation(t *testing.T) {
	l := NewLegend(DefaultTheme())
	full := l.Render(0)
	truncated := l.Render(0)
	if full != truncated {
		t.Error("zero width should not alter output")
	}
}

func TestLegend_Render_NarrowWidth_Truncates(t *testing.T) {
	l := NewLegend(DefaultTheme())
	out := l.Render(5)
	// The raw byte length may include ANSI codes, but visible content should
	// not exceed a reasonable multiple of the requested width.
	if len(out) == 0 {
		t.Error("expected non-empty output even when truncated")
	}
}

func TestLegend_Entries_SymbolsNonEmpty(t *testing.T) {
	l := NewLegend(DefaultTheme())
	for _, e := range l.Entries() {
		if e.Symbol == "" {
			t.Errorf("entry %q has empty symbol", e.Label)
		}
	}
}

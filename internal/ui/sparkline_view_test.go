package ui

import (
	"strings"
	"testing"
)

func makeSparkHistory(entries []HistoryEntry) *History {
	h := NewHistory(20)
	const key = "svc"
	for _, e := range entries {
		h.Record(key, e)
	}
	return h
}

func TestSparklineView_NilHistory(t *testing.T) {
	out := SparklineView(nil, "svc", 10, DefaultTheme())
	if out != "" {
		t.Errorf("expected empty string for nil history, got %q", out)
	}
}

func TestSparklineView_ZeroWidth(t *testing.T) {
	h := NewHistory(10)
	out := SparklineView(h, "svc", 0, DefaultTheme())
	if out != "" {
		t.Errorf("expected empty string for zero width, got %q", out)
	}
}

func TestSparklineView_NoEntries_PadsWithDots(t *testing.T) {
	h := NewHistory(10)
	out := SparklineView(h, "unknown-key", 5, DefaultTheme())
	// Strip ANSI — just check plain dot characters are present.
	stripped := stripANSI(out)
	if !strings.Contains(stripped, "·") {
		t.Errorf("expected padding dots, got %q", stripped)
	}
}

func TestSparklineView_ServingEntries_ContainsBars(t *testing.T) {
	entries := []HistoryEntry{
		{Status: StatusServing, LatencyMs: 5},
		{Status: StatusServing, LatencyMs: 50},
		{Status: StatusServing, LatencyMs: 300},
	}
	h := makeSparkHistory(entries)
	out := stripANSI(SparklineView(h, "svc", 10, DefaultTheme()))
	for _, bar := range []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'} {
		if strings.ContainsRune(out, bar) {
			return // at least one bar character found
		}
	}
	t.Errorf("expected sparkline bar characters in output, got %q", out)
}

func TestSparklineView_TruncatesToWidth(t *testing.T) {
	entries := make([]HistoryEntry, 20)
	for i := range entries {
		entries[i] = HistoryEntry{Status: StatusServing, LatencyMs: 10}
	}
	h := makeSparkHistory(entries)
	out := stripANSI(SparklineView(h, "svc", 8, DefaultTheme()))
	// Count runes (each bar is one rune).
	count := []rune(out)
	if len(count) != 8 {
		t.Errorf("expected 8 runes, got %d: %q", len(count), out)
	}
}

func TestSparklineRow_ZeroWidth(t *testing.T) {
	h := NewHistory(10)
	out := SparklineRow("Latency", "svc", h, 5, DefaultTheme())
	if out != "" {
		t.Errorf("expected empty string for narrow width, got %q", out)
	}
}

func TestSparklineRow_ContainsLabel(t *testing.T) {
	h := NewHistory(10)
	out := SparklineRow("Latency", "svc", h, 40, DefaultTheme())
	if !strings.Contains(out, "Latency") {
		t.Errorf("expected label in sparkline row, got %q", out)
	}
}

func TestLatencyIndex_Boundaries(t *testing.T) {
	cases := []struct {
		ms  float64
		want int
	}{
		{0, 7},
		{5, 7},
		{15, 6},
		{30, 5},
		{75, 4},
		{150, 3},
		{300, 2},
		{750, 1},
		{2000, 0},
	}
	for _, tc := range cases {
		got := latencyIndex(tc.ms)
		if got != tc.want {
			t.Errorf("latencyIndex(%.0f) = %d, want %d", tc.ms, got, tc.want)
		}
	}
}

// stripANSI removes ANSI escape sequences for plain-text assertions.
func stripANSI(s string) string {
	out := strings.Builder{}
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

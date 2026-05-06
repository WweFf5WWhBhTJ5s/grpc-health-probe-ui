package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// LegendEntry describes a single status symbol and its meaning.
type LegendEntry struct {
	Symbol string
	Label  string
	Style  lipgloss.Style
}

// Legend renders a compact status-symbol legend suitable for display at the
// bottom of the terminal dashboard.
type Legend struct {
	theme Theme
}

// NewLegend constructs a Legend using the provided Theme.
func NewLegend(t Theme) Legend {
	return Legend{theme: t}
}

// Entries returns the ordered slice of LegendEntry values for all known
// health statuses.
func (l Legend) Entries() []LegendEntry {
	return []LegendEntry{
		{Symbol: "●", Label: "Serving", Style: l.theme.Serving},
		{Symbol: "●", Label: "Not Serving", Style: l.theme.NotServing},
		{Symbol: "●", Label: "Unknown", Style: l.theme.Unknown},
		{Symbol: "✕", Label: "Unreachable", Style: l.theme.Unreachable},
	}
}

// Render returns a single-line string representation of the legend, truncated
// to fit within width characters. If width is zero or negative the full string
// is returned without truncation.
func (l Legend) Render(width int) string {
	parts := make([]string, 0, len(l.Entries()))
	for _, e := range l.Entries() {
		symbol := e.Style.Render(e.Symbol)
		parts = append(parts, fmt.Sprintf("%s %s", symbol, e.Label))
	}
	line := strings.Join(parts, "  ")
	if width > 0 && lipgloss.Width(line) > width {
		line = line[:width]
	}
	return line
}

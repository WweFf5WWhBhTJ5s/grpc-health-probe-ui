package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// StatusBar renders a bottom status bar with last-refresh time, sort order, and key hints.
type StatusBar struct {
	LastRefresh time.Time
	SortOrder   SortOrder
	Width       int
}

var (
	statusBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("236")).
			Foreground(lipgloss.Color("250")).
			Padding(0, 1)

	statusHighlight = lipgloss.NewStyle().
			Background(lipgloss.Color("236")).
			Foreground(lipgloss.Color("86")).
			Bold(true)
)

// Render returns the status bar string for the given width.
func (s StatusBar) Render() string {
	refreshStr := "never"
	if !s.LastRefresh.IsZero() {
		refreshStr = s.LastRefresh.Format("15:04:05")
	}

	left := statusBarStyle.Render(
		fmt.Sprintf("refreshed: %s", refreshStr),
	)

	center := statusBarStyle.Render(
		fmt.Sprintf("sort: %s", statusHighlight.Render(SortOrderLabel(s.SortOrder))),
	)

	right := statusBarStyle.Render("q quit  r refresh  s sort  ↑↓ scroll")

	avail := s.Width - lipgloss.Width(left) - lipgloss.Width(center) - lipgloss.Width(right)
	if avail < 0 {
		avail = 0
	}
	pad := statusBarStyle.Render(fmt.Sprintf("%*s", avail, ""))

	return left + pad + center + right
}

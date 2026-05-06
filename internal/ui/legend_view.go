package ui

import "github.com/charmbracelet/lipgloss"

// LegendView renders the status legend inside a styled container that matches
// the dashboard footer area. It is intended to be composed alongside the
// StatusBar at the bottom of the View.
func LegendView(l Legend, width int) string {
	if width <= 0 {
		return ""
	}

	containerStyle := lipgloss.NewStyle().
		Width(width).
		PaddingLeft(1).
		PaddingRight(1)

	content := l.Render(width - 2) // account for horizontal padding
	return containerStyle.Render(content)
}

package ui

import "github.com/charmbracelet/lipgloss"

// Theme holds all styled renderers used across the UI.
type Theme struct {
	Title       lipgloss.Style
	Serving     lipgloss.Style
	NotServing  lipgloss.Style
	Unknown     lipgloss.Style
	Highlight   lipgloss.Style
	Dim         lipgloss.Style
	Border      lipgloss.Style
	StatusBar   lipgloss.Style
	AlertInfo   lipgloss.Style
	AlertWarn   lipgloss.Style
}

// DefaultTheme returns the default colour scheme for the dashboard.
func DefaultTheme() Theme {
	return Theme{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#5C5FE0")).
			Padding(0, 1),

		Serving: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00D26A")).
			Bold(true),

		NotServing: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF4D4D")).
			Bold(true),

		Unknown: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")),

		Highlight: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true),

		Dim: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")),

		Border: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#5C5FE0")).
			Padding(0, 1),

		StatusBar: lipgloss.NewStyle().
			Background(lipgloss.Color("#2A2A2A")).
			Foreground(lipgloss.Color("#CCCCCC")).
			Padding(0, 1),

		AlertInfo: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00BFFF")).
			Bold(true),

		AlertWarn: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500")).
			Bold(true),
	}
}

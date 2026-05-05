package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// TagBadge renders a single tag as a styled pill.
func TagBadge(theme Theme, tag string) string {
	if tag == "" {
		return ""
	}
	style := lipgloss.NewStyle().
		Background(theme.UnknownColor).
		Foreground(lipgloss.Color("#ffffff")).
		Padding(0, 1).
		MarginRight(1)
	return style.Render(tag)
}

// TagRow renders all tags for a target as a horizontal row of badges.
// Returns an empty string when tags is nil or empty.
func TagRow(theme Theme, tags []string) string {
	if len(tags) == 0 {
		return ""
	}
	badges := make([]string, 0, len(tags))
	for _, t := range tags {
		badges = append(badges, TagBadge(theme, t))
	}
	return strings.Join(badges, "")
}

// TagListView renders a summary of all tags used across targets.
// Intended for a sidebar or overlay panel.
func TagListView(theme Theme, store *TagStore) string {
	if store == nil {
		return ""
	}
	all := store.AllTags()
	if len(all) == 0 {
		return lipgloss.NewStyle().
			Foreground(theme.UnknownColor).
			Italic(true).
			Render("no tags defined")
	}
	header := lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Render("Tags")
	lines := make([]string, 0, len(all)+1)
	lines = append(lines, header)
	for _, t := range all {
		lines = append(lines, "  "+TagBadge(theme, t))
	}
	return strings.Join(lines, "\n")
}

package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// BookmarkBadge returns a short styled indicator when a target is bookmarked.
func BookmarkBadge(bookmarked bool, theme Theme) string {
	if !bookmarked {
		return ""
	}
	style := lipgloss.NewStyle().
		Foreground(theme.ServingColor).
		Bold(true)
	return style.Render("★ ")
}

// BookmarkListView renders a simple list of all bookmarked names.
// Returns a placeholder string when the store is empty.
func BookmarkListView(store *BookmarkStore, theme Theme) string {
	if store == nil || store.Count() == 0 {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("No bookmarks")
	}

	titleStyle := lipgloss.NewStyle().Bold(true).Underline(true)
	rowStyle := lipgloss.NewStyle().PaddingLeft(2)

	var sb strings.Builder
	sb.WriteString(titleStyle.Render("Bookmarks"))
	sb.WriteString("\n")
	for i, name := range store.All() {
		sb.WriteString(rowStyle.Render(fmt.Sprintf("%d. %s", i+1, name)))
		sb.WriteString("\n")
	}
	return strings.TrimRight(sb.String(), "\n")
}

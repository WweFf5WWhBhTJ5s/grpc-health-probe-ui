package ui

import "fmt"

const noteBadgeSymbol = "✎"

// NoteBadge returns a short inline indicator when a note exists for the target.
// Returns an empty string when no note is present.
func NoteBadge(store *NoteStore, target string) string {
	if store == nil {
		return ""
	}
	if store.Has(target) {
		return noteBadgeSymbol
	}
	return ""
}

// NoteRow renders a labelled row suitable for embedding in DetailView.
// It returns an empty string when no note is set for the target.
func NoteRow(store *NoteStore, target string) string {
	if store == nil {
		return ""
	}
	n, ok := store.Get(target)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%-12s %s", "Note:", n)
}

// NoteListView renders all stored notes as a labelled list.
// Returns a placeholder string when the store is nil or empty.
func NoteListView(store *NoteStore) string {
	if store == nil || store.Len() == 0 {
		return "No notes."
	}
	out := ""
	for _, target := range store.All() {
		n, _ := store.Get(target)
		out += fmt.Sprintf("%s %s: %s\n", noteBadgeSymbol, target, n)
	}
	return out
}

package ui

import "sort"

// BookmarkStore holds a set of bookmarked target names.
type BookmarkStore struct {
	entries map[string]struct{}
}

// NewBookmarkStore returns an empty BookmarkStore.
func NewBookmarkStore() *BookmarkStore {
	return &BookmarkStore{entries: make(map[string]struct{})}
}

// Toggle adds the name if absent, removes it if present.
// Returns true if the name is now bookmarked.
func (b *BookmarkStore) Toggle(name string) bool {
	if _, ok := b.entries[name]; ok {
		delete(b.entries, name)
		return false
	}
	b.entries[name] = struct{}{}
	return true
}

// IsBookmarked reports whether name is bookmarked.
func (b *BookmarkStore) IsBookmarked(name string) bool {
	_, ok := b.entries[name]
	return ok
}

// All returns a sorted slice of all bookmarked names.
func (b *BookmarkStore) All() []string {
	out := make([]string, 0, len(b.entries))
	for k := range b.entries {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// Count returns the number of bookmarks.
func (b *BookmarkStore) Count() int {
	return len(b.entries)
}

// Clear removes all bookmarks.
func (b *BookmarkStore) Clear() {
	b.entries = make(map[string]struct{})
}

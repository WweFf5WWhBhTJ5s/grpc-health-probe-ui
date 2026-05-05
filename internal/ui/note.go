package ui

import (
	"sort"
	"strings"
	"sync"
)

// NoteStore holds user-defined text notes keyed by target name.
type NoteStore struct {
	mu    sync.RWMutex
	notes map[string]string
}

// NewNoteStore returns an initialised, empty NoteStore.
func NewNoteStore() *NoteStore {
	return &NoteStore{notes: make(map[string]string)}
}

// Set stores or replaces a note for the given target.
// Passing an empty string removes the note.
func (s *NoteStore) Set(target, text string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	text = strings.TrimSpace(text)
	if text == "" {
		delete(s.notes, target)
		return
	}
	s.notes[target] = text
}

// Get returns the note for the given target and whether one exists.
func (s *NoteStore) Get(target string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n, ok := s.notes[target]
	return n, ok
}

// Has reports whether a note exists for the given target.
func (s *NoteStore) Has(target string) bool {
	_, ok := s.Get(target)
	return ok
}

// All returns a sorted slice of targets that have notes.
func (s *NoteStore) All() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]string, 0, len(s.notes))
	for k := range s.notes {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// Clear removes all notes.
func (s *NoteStore) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.notes = make(map[string]string)
}

// Len returns the number of notes stored.
func (s *NoteStore) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.notes)
}

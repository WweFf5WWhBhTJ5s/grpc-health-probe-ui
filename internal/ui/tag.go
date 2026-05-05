package ui

import "sort"

// TagStore manages a set of string tags associated with named targets.
type TagStore struct {
	tags map[string]map[string]struct{}
}

// NewTagStore returns an empty TagStore.
func NewTagStore() *TagStore {
	return &TagStore{tags: make(map[string]map[string]struct{})}
}

// Add adds a tag to the given target. Tags are case-sensitive and trimmed.
func (s *TagStore) Add(target, tag string) {
	tag = trimString(tag)
	if tag == "" {
		return
	}
	if s.tags[target] == nil {
		s.tags[target] = make(map[string]struct{})
	}
	s.tags[target][tag] = struct{}{}
}

// Remove removes a tag from the given target. No-op if absent.
func (s *TagStore) Remove(target, tag string) {
	if s.tags[target] == nil {
		return
	}
	delete(s.tags[target], tag)
	if len(s.tags[target]) == 0 {
		delete(s.tags, target)
	}
}

// Get returns a sorted slice of tags for the given target.
func (s *TagStore) Get(target string) []string {
	set := s.tags[target]
	if len(set) == 0 {
		return nil
	}
	out := make([]string, 0, len(set))
	for t := range set {
		out = append(out, t)
	}
	sort.Strings(out)
	return out
}

// Has reports whether the given target has the given tag.
func (s *TagStore) Has(target, tag string) bool {
	if s.tags[target] == nil {
		return false
	}
	_, ok := s.tags[target][tag]
	return ok
}

// AllTags returns a sorted, deduplicated list of every tag across all targets.
func (s *TagStore) AllTags() []string {
	seen := make(map[string]struct{})
	for _, set := range s.tags {
		for t := range set {
			seen[t] = struct{}{}
		}
	}
	out := make([]string, 0, len(seen))
	for t := range seen {
		out = append(out, t)
	}
	sort.Strings(out)
	return out
}

// ClearTarget removes all tags for a given target.
func (s *TagStore) ClearTarget(target string) {
	delete(s.tags, target)
}

// trimString trims leading/trailing whitespace.
func trimString(s string) string {
	result := []rune{}
	start, end := 0, len([]rune(s))-1
	runes := []rune(s)
	for start <= end && (runes[start] == ' ' || runes[start] == '\t') {
		start++
	}
	for end >= start && (runes[end] == ' ' || runes[end] == '\t') {
		end--
	}
	result = runes[start : end+1]
	return string(result)
}

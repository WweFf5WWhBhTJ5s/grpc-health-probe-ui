package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
)

// SearchState holds the state of the interactive search/filter input.
type SearchState struct {
	Active bool
	input  textinput.Model
}

// NewSearchState returns a SearchState with a configured text input.
func NewSearchState() SearchState {
	ti := textinput.New()
	ti.Placeholder = "filter targets…"
	ti.CharLimit = 64
	ti.Width = 40
	return SearchState{input: tti}
}

// Activate enables the search input and focuses it.
func (s *SearchState) Activate() {
	s.Active = true
	s.input.Focus()
}

// Deactivate disables the search input and clears the query.
func (s *SearchState) Deactivate() {
	s.Active = false
	s.input.Blur()
	s.input.SetValue("")
}

// Query returns the current (trimmed) search string.
func (s *SearchState) Query() string {
	return strings.TrimSpace(s.input.Value())
}

// SetQuery sets the input value directly (useful for testing).
func (s *SearchState) SetQuery(q string) {
	s.input.SetValue(q)
}

// Update forwards a bubbletea message to the underlying text input.
// Returns the updated SearchState and any command produced.
func (s SearchState) Update(msg tea.Msg) (SearchState, tea.Cmd) {
	if !s.Active {
		return s, nil
	}
	updated, cmd := s.input.Update(msg)
	s.input = updated
	return s, cmd
}

// View renders the search input bar, or an empty string when inactive.
func (s SearchState) View() string {
	if !s.Active {
		return ""
	}
	return "Search: " + s.input.View()
}

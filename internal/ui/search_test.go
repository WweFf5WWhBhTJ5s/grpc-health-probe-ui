package ui

import (
	"testing"

	"github.com/charmbracelet/bubbletea"
)

func TestNewSearchState_NotActive(t *testing.T) {
	s := NewSearchState()
	if s.Active {
		t.Fatal("expected search to be inactive by default")
	}
}

func TestNewSearchState_EmptyQuery(t *testing.T) {
	s := NewSearchState()
	if q := s.Query(); q != "" {
		t.Fatalf("expected empty query, got %q", q)
	}
}

func TestActivate_SetsActive(t *testing.T) {
	s := NewSearchState()
	s.Activate()
	if !s.Active {
		t.Fatal("expected search to be active after Activate()")
	}
}

func TestDeactivate_ClearsActive(t *testing.T) {
	s := NewSearchState()
	s.Activate()
	s.SetQuery("hello")
	s.Deactivate()
	if s.Active {
		t.Fatal("expected search to be inactive after Deactivate()")
	}
}

func TestDeactivate_ClearsQuery(t *testing.T) {
	s := NewSearchState()
	s.Activate()
	s.SetQuery("my-service")
	s.Deactivate()
	if q := s.Query(); q != "" {
		t.Fatalf("expected empty query after deactivate, got %q", q)
	}
}

func TestSetQuery_TrimsSpace(t *testing.T) {
	s := NewSearchState()
	s.SetQuery("  svc  ")
	if q := s.Query(); q != "svc" {
		t.Fatalf("expected trimmed query %q, got %q", "svc", q)
	}
}

func TestView_InactiveReturnsEmpty(t *testing.T) {
	s := NewSearchState()
	if v := s.View(); v != "" {
		t.Fatalf("expected empty view when inactive, got %q", v)
	}
}

func TestView_ActiveContainsPrefix(t *testing.T) {
	s := NewSearchState()
	s.Activate()
	v := s.View()
	if len(v) == 0 {
		t.Fatal("expected non-empty view when active")
	}
	// The rendered view should contain the "Search:" label.
	if !containsSubstring(v, "Search:") {
		t.Fatalf("expected view to contain 'Search:', got %q", v)
	}
}

func TestUpdate_InactiveIsNoop(t *testing.T) {
	s := NewSearchState()
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}
	s2, cmd := s.Update(msg)
	if s2.Active {
		t.Fatal("inactive search should remain inactive after Update")
	}
	if cmd != nil {
		t.Fatal("expected nil cmd for inactive search update")
	}
}

// containsSubstring is a helper to avoid importing strings in tests.
func containsSubstring(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > 0 && searchContains(s, sub))
}

func searchContains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

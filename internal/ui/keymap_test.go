package ui

import (
	"testing"

	"github.com/charmbracelet/bubbles/key"
)

func TestDefaultKeyMap_Quit(t *testing.T) {
	km := DefaultKeyMap()
	if !key.Matches(key.NewBinding(key.WithKeys("q")), km.Quit) {
		t.Error("expected 'q' to match Quit binding")
	}
}

func TestDefaultKeyMap_Refresh(t *testing.T) {
	km := DefaultKeyMap()
	if !key.Matches(key.NewBinding(key.WithKeys("r")), km.Refresh) {
		t.Error("expected 'r' to match Refresh binding")
	}
}

func TestDefaultKeyMap_Sort(t *testing.T) {
	km := DefaultKeyMap()
	if !key.Matches(key.NewBinding(key.WithKeys("s")), km.Sort) {
		t.Error("expected 's' to match Sort binding")
	}
}

func TestDefaultKeyMap_Navigation(t *testing.T) {
	km := DefaultKeyMap()
	if !key.Matches(key.NewBinding(key.WithKeys("up")), km.Up) {
		t.Error("expected 'up' to match Up binding")
	}
	if !key.Matches(key.NewBinding(key.WithKeys("down")), km.Down) {
		t.Error("expected 'down' to match Down binding")
	}
}

func TestKeyMap_ShortHelp(t *testing.T) {
	km := DefaultKeyMap()
	short := km.ShortHelp()
	if len(short) != 3 {
		t.Errorf("expected 3 short help bindings, got %d", len(short))
	}
}

func TestKeyMap_FullHelp(t *testing.T) {
	km := DefaultKeyMap()
	full := km.FullHelp()
	if len(full) == 0 {
		t.Error("expected non-empty full help")
	}
	total := 0
	for _, group := range full {
		total += len(group)
	}
	if total < 5 {
		t.Errorf("expected at least 5 total bindings in full help, got %d", total)
	}
}

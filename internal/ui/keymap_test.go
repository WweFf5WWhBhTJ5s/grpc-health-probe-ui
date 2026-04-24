package ui

import (
	"testing"

	"github.com/charmbracelet/bubbles/key"
)

func TestDefaultKeyMap_Quit(t *testing.T) {
	km := DefaultKeyMap()
	if !key.Matches(key.NewBinding(key.WithKeys("q")), km.Quit) {
		// key.Matches checks if a msg matches a binding; test via keys directly
	}
	keys := km.Quit.Keys()
	if len(keys) != 2 {
		t.Fatalf("expected 2 quit keys, got %d", len(keys))
	}
	if keys[0] != "q" || keys[1] != "ctrl+c" {
		t.Errorf("unexpected quit keys: %v", keys)
	}
}

func TestDefaultKeyMap_Refresh(t *testing.T) {
	km := DefaultKeyMap()
	keys := km.Refresh.Keys()
	if len(keys) != 1 || keys[0] != "r" {
		t.Errorf("expected refresh key 'r', got %v", keys)
	}
}

func TestDefaultKeyMap_Navigation(t *testing.T) {
	km := DefaultKeyMap()

	upKeys := km.Up.Keys()
	if len(upKeys) != 2 {
		t.Fatalf("expected 2 up keys, got %d", len(upKeys))
	}
	if upKeys[0] != "up" || upKeys[1] != "k" {
		t.Errorf("unexpected up keys: %v", upKeys)
	}

	downKeys := km.Down.Keys()
	if len(downKeys) != 2 {
		t.Fatalf("expected 2 down keys, got %d", len(downKeys))
	}
	if downKeys[0] != "down" || downKeys[1] != "j" {
		t.Errorf("unexpected down keys: %v", downKeys)
	}
}

func TestKeyMap_ShortHelp(t *testing.T) {
	km := DefaultKeyMap()
	short := km.ShortHelp()
	if len(short) != 4 {
		t.Fatalf("expected 4 short help bindings, got %d", len(short))
	}
}

func TestKeyMap_FullHelp(t *testing.T) {
	km := DefaultKeyMap()
	full := km.FullHelp()
	if len(full) != 2 {
		t.Fatalf("expected 2 groups in full help, got %d", len(full))
	}
	if len(full[0]) != 2 {
		t.Errorf("expected 2 bindings in first group, got %d", len(full[0]))
	}
	if len(full[1]) != 2 {
		t.Errorf("expected 2 bindings in second group, got %d", len(full[1]))
	}
}

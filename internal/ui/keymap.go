package ui

import (
	"github.com/charmbracelet/bubbles/key"
)

// KeyMap defines all key bindings for the dashboard.
type KeyMap struct {
	Quit    key.Binding
	Refresh key.Binding
	Sort    key.Binding
	Up      key.Binding
	Down    key.Binding
	Filter  key.Binding
	Escape  key.Binding
}

// DefaultKeyMap returns the default keyboard bindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Refresh: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "refresh"),
		),
		Sort: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "sort"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "filter"),
		),
		Escape: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "clear filter"),
		),
	}
}

// ShortHelp returns a compact list of key bindings for the help view.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Refresh, k.Sort, k.Filter, k.Quit}
}

// FullHelp returns all key bindings grouped for the full help view.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Refresh, k.Sort},
		{k.Filter, k.Escape},
		{k.Quit},
	}
}

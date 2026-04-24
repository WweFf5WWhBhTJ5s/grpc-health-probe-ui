package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

// DetailKeyMap holds key bindings used in the detail panel.
type DetailKeyMap struct {
	Open  key.Binding
	Close key.Binding
	Up    key.Binding
	Down  key.Binding
}

// DefaultDetailKeyMap returns the default key bindings for the detail panel.
func DefaultDetailKeyMap() DetailKeyMap {
	return DetailKeyMap{
		Open: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "open detail"),
		),
		Close: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "close detail"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "prev target"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "next target"),
		),
	}
}

// ShortHelp returns abbreviated key hints for the detail panel.
func (k DetailKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Open, k.Close, k.Up, k.Down}
}

// FullHelp returns the full set of key bindings for the detail panel.
func (k DetailKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

var detailHintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

// DetailHints renders a compact one-line hint string for the detail panel.
func DetailHints(km DetailKeyMap) string {
	var hints []string
	for _, b := range km.ShortHelp() {
		if h := b.Help(); h.Key != "" {
			hints = append(hints, h.Key+" "+h.Desc)
		}
	}
	result := ""
	for i, h := range hints {
		if i > 0 {
			result += "  "
		}
		result += detailHintStyle.Render(h)
	}
	return result
}

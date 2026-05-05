package ui

import "github.com/charmbracelet/bubbles/key"

// BookmarkKeyMap defines keybindings related to bookmark actions.
type BookmarkKeyMap struct {
	Toggle   key.Binding
	ShowList key.Binding
	Clear    key.Binding
}

// DefaultBookmarkKeyMap returns the default bookmark keybindings.
func DefaultBookmarkKeyMap() BookmarkKeyMap {
	return BookmarkKeyMap{
		Toggle: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "toggle bookmark"),
		),
		ShowList: key.NewBinding(
			key.WithKeys("B"),
			key.WithHelp("B", "show bookmarks"),
		),
		Clear: key.NewBinding(
			key.WithKeys("ctrl+b"),
			key.WithHelp("ctrl+b", "clear bookmarks"),
		),
	}
}

// ShortHelp returns a subset of bindings shown in the compact help bar.
func (k BookmarkKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Toggle, k.ShowList}
}

// FullHelp returns all bookmark bindings.
func (k BookmarkKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Toggle, k.ShowList, k.Clear}}
}

package ui

import "github.com/charmbracelet/bubbles/key"

// GroupKeyMap holds key bindings related to group management.
type GroupKeyMap struct {
	AddToGroup    key.Binding
	RemoveGroup   key.Binding
	ListGroups    key.Binding
}

// DefaultGroupKeyMap returns a GroupKeyMap with sensible defaults.
func DefaultGroupKeyMap() GroupKeyMap {
	return GroupKeyMap{
		AddToGroup: key.NewBinding(
			key.WithKeys("g"),
			key.WithHelp("g", "add to group"),
		),
		RemoveGroup: key.NewBinding(
			key.WithKeys("G"),
			key.WithHelp("G", "remove from group"),
		),
		ListGroups: key.NewBinding(
			key.WithKeys("ctrl+g"),
			key.WithHelp("ctrl+g", "list groups"),
		),
	}
}

// ShortHelp returns a minimal set of bindings for the help bar.
func (k GroupKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.AddToGroup, k.ListGroups}
}

// FullHelp returns all group-related bindings.
func (k GroupKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.AddToGroup, k.RemoveGroup, k.ListGroups}}
}

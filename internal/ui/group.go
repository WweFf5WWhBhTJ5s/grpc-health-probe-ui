package ui

import "sort"

// GroupStore manages named groups that targets can be assigned to.
type GroupStore struct {
	groups map[string]map[string]struct{} // group name -> set of target keys
}

// NewGroupStore returns an initialised GroupStore.
func NewGroupStore() *GroupStore {
	return &GroupStore{groups: make(map[string]map[string]struct{})}
}

// Add assigns a target key to a group, creating the group if necessary.
func (g *GroupStore) Add(group, key string) {
	if group == "" || key == "" {
		return
	}
	if _, ok := g.groups[group]; !ok {
		g.groups[group] = make(map[string]struct{})
	}
	g.groups[group][key] = struct{}{}
}

// Remove removes a target key from a group.
func (g *GroupStore) Remove(group, key string) {
	if m, ok := g.groups[group]; ok {
		delete(m, key)
		if len(m) == 0 {
			delete(g.groups, group)
		}
	}
}

// GroupsFor returns the sorted list of group names that contain key.
func (g *GroupStore) GroupsFor(key string) []string {
	var out []string
	for name, members := range g.groups {
		if _, ok := members[key]; ok {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}

// Members returns the sorted list of target keys in a group.
func (g *GroupStore) Members(group string) []string {
	m, ok := g.groups[group]
	if !ok {
		return nil
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// All returns all group names in sorted order.
func (g *GroupStore) All() []string {
	out := make([]string, 0, len(g.groups))
	for name := range g.groups {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// Clear removes all groups.
func (g *GroupStore) Clear() {
	g.groups = make(map[string]map[string]struct{})
}

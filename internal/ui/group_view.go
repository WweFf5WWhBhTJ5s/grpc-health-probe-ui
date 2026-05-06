package ui

import (
	"fmt"
	"strings"
)

// GroupBadge renders a compact inline badge listing the groups for a target.
// Returns an empty string when the target belongs to no groups.
func GroupBadge(store *GroupStore, key string) string {
	if store == nil {
		return ""
	}
	groups := store.GroupsFor(key)
	if len(groups) == 0 {
		return ""
	}
	return fmt.Sprintf("[%s]", strings.Join(groups, ","))
}

// GroupListView renders a multi-line summary of all groups and their members.
// Returns a placeholder when the store is nil or empty.
func GroupListView(store *GroupStore) string {
	if store == nil {
		return "(no groups)"
	}
	names := store.All()
	if len(names) == 0 {
		return "(no groups)"
	}
	var sb strings.Builder
	for _, name := range names {
		members := store.Members(name)
		sb.WriteString(fmt.Sprintf("%-20s %s\n", name, strings.Join(members, ", ")))
	}
	return strings.TrimRight(sb.String(), "\n")
}

package ui

import (
	"sort"

	"github.com/your-org/grpc-health-probe-ui/internal/probe"
)

// SortOrder defines how targets are sorted in the dashboard.
type SortOrder int

const (
	SortByName SortOrder = iota
	SortByStatus
	SortByHost
)

// SortTargets returns a sorted copy of the given target slice.
func SortTargets(targets []probe.Target, order SortOrder) []probe.Target {
	copy_ := make([]probe.Target, len(targets))
	copy(copy_, targets)

	switch order {
	case SortByName:
		sort.Slice(copy_, func(i, j int) bool {
			return copy_[i].Name < copy_[j].Name
		})
	case SortByStatus:
		sort.Slice(copy_, func(i, j int) bool {
			return copy_[i].LastStatus < copy_[j].LastStatus
		})
	case SortByHost:
		sort.Slice(copy_, func(i, j int) bool {
			return copy_[i].Address < copy_[j].Address
		})
	}

	return copy_
}

// NextSortOrder cycles through available sort orders.
func NextSortOrder(current SortOrder) SortOrder {
	return (current + 1) % 3
}

// SortOrderLabel returns a human-readable label for the sort order.
func SortOrderLabel(order SortOrder) string {
	switch order {
	case SortByName:
		return "name"
	case SortByStatus:
		return "status"
	case SortByHost:
		return "host"
	default:
		return "unknown"
	}
}

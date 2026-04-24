package ui

import (
	"strings"

	"github.com/jhump/grpc-health-probe-ui/internal/probe"
)

// FilterTargets returns a subset of targets whose Name or Host contains query
// (case-insensitive). If query is empty, all targets are returned.
func FilterTargets(targets []probe.Target, query string) []probe.Target {
	if query == "" {
		return targets
	}
	q := strings.ToLower(query)
	out := make([]probe.Target, 0, len(targets))
	for _, t := range targets {
		if strings.Contains(strings.ToLower(t.Name), q) ||
			strings.Contains(strings.ToLower(t.Host), q) {
			out = append(out, t)
		}
	}
	return out
}

// FilterQuery holds the current filter state for the UI model.
type FilterQuery struct {
	Active bool
	Query  string
}

// Toggle activates or deactivates filtering, resetting the query on deactivation.
func (f *FilterQuery) Toggle() {
	if f.Active {
		f.Active = false
		f.Query = ""
	} else {
		f.Active = true
	}
}

// Append adds a rune to the current query.
func (f *FilterQuery) Append(r rune) {
	if f.Active {
		f.Query += string(r)
	}
}

// Backspace removes the last character from the query.
func (f *FilterQuery) Backspace() {
	if f.Active && len(f.Query) > 0 {
		f.Query = f.Query[:len(f.Query)-1]
	}
}

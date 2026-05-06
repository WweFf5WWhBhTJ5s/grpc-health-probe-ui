package ui

import "sort"

// PinStore tracks which targets are pinned to the top of the list.
type PinStore struct {
	pinned map[string]bool
}

// NewPinStore returns an empty PinStore.
func NewPinStore() *PinStore {
	return &PinStore{pinned: make(map[string]bool)}
}

// Toggle pins the key if it is not pinned, or unpins it if it is.
func (p *PinStore) Toggle(key string) {
	if key == "" {
		return
	}
	if p.pinned[key] {
		delete(p.pinned, key)
	} else {
		p.pinned[key] = true
	}
}

// IsPinned reports whether key is currently pinned.
func (p *PinStore) IsPinned(key string) bool {
	return p.pinned[key]
}

// All returns a sorted slice of all pinned keys.
func (p *PinStore) All() []string {
	keys := make([]string, 0, len(p.pinned))
	for k := range p.pinned {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Clear removes all pins.
func (p *PinStore) Clear() {
	p.pinned = make(map[string]bool)
}

// Count returns the number of pinned targets.
func (p *PinStore) Count() int {
	return len(p.pinned)
}

// ApplyPins reorders targets so that pinned entries appear first,
// preserving relative order within each group.
func ApplyPins(targets []string, store *PinStore) []string {
	if store == nil || store.Count() == 0 {
		return targets
	}
	pinned := make([]string, 0, len(targets))
	unpinned := make([]string, 0, len(targets))
	for _, t := range targets {
		if store.IsPinned(t) {
			pinned = append(pinned, t)
		} else {
			unpinned = append(unpinned, t)
		}
	}
	return append(pinned, unpinned...)
}

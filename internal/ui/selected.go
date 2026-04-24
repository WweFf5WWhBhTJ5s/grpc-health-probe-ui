package ui

// SelectedState tracks which target (if any) is focused in the detail panel.
type SelectedState struct {
	index    int
	active   bool
}

// NewSelectedState returns a SelectedState with no selection.
func NewSelectedState() SelectedState {
	return SelectedState{index: 0, active: false}
}

// Select activates the selection at index i.
func (s SelectedState) Select(i int) SelectedState {
	if i < 0 {
		i = 0
	}
	return SelectedState{index: i, active: true}
}

// Deselect clears the active selection.
func (s SelectedState) Deselect() SelectedState {
	return SelectedState{index: s.index, active: false}
}

// MoveUp decrements the selection index, clamped at 0.
func (s SelectedState) MoveUp() SelectedState {
	if !s.active {
		return s
	}
	next := s.index - 1
	if next < 0 {
		next = 0
	}
	return SelectedState{index: next, active: true}
}

// MoveDown increments the selection index, clamped at max.
func (s SelectedState) MoveDown(max int) SelectedState {
	if !s.active {
		return s
	}
	next := s.index + 1
	if next > max {
		next = max
	}
	return SelectedState{index: next, active: true}
}

// Index returns the current selection index.
func (s SelectedState) Index() int { return s.index }

// Active reports whether a target is currently selected.
func (s SelectedState) Active() bool { return s.active }

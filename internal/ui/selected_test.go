package ui

import "testing"

func TestNewSelectedState_NotActive(t *testing.T) {
	s := NewSelectedState()
	if s.Active() {
		t.Error("expected new state to be inactive")
	}
	if s.Index() != 0 {
		t.Errorf("expected index 0, got %d", s.Index())
	}
}

func TestSelect_SetsActiveAndIndex(t *testing.T) {
	s := NewSelectedState().Select(3)
	if !s.Active() {
		t.Error("expected active after Select")
	}
	if s.Index() != 3 {
		t.Errorf("expected index 3, got %d", s.Index())
	}
}

func TestSelect_NegativeIndexClampsToZero(t *testing.T) {
	s := NewSelectedState().Select(-5)
	if s.Index() != 0 {
		t.Errorf("expected index 0 for negative input, got %d", s.Index())
	}
}

func TestDeselect_ClearsActive(t *testing.T) {
	s := NewSelectedState().Select(2).Deselect()
	if s.Active() {
		t.Error("expected inactive after Deselect")
	}
	if s.Index() != 2 {
		t.Errorf("expected index preserved after Deselect, got %d", s.Index())
	}
}

func TestMoveUp_DecrementsIndex(t *testing.T) {
	s := NewSelectedState().Select(3).MoveUp()
	if s.Index() != 2 {
		t.Errorf("expected index 2, got %d", s.Index())
	}
}

func TestMoveUp_ClampsAtZero(t *testing.T) {
	s := NewSelectedState().Select(0).MoveUp()
	if s.Index() != 0 {
		t.Errorf("expected index clamped at 0, got %d", s.Index())
	}
}

func TestMoveDown_IncrementsIndex(t *testing.T) {
	s := NewSelectedState().Select(1).MoveDown(5)
	if s.Index() != 2 {
		t.Errorf("expected index 2, got %d", s.Index())
	}
}

func TestMoveDown_ClampsAtMax(t *testing.T) {
	s := NewSelectedState().Select(4).MoveDown(4)
	if s.Index() != 4 {
		t.Errorf("expected index clamped at 4, got %d", s.Index())
	}
}

func TestMoveUp_InactiveIsNoop(t *testing.T) {
	s := NewSelectedState().MoveUp()
	if s.Active() {
		t.Error("expected move on inactive state to remain inactive")
	}
}

func TestMoveDown_InactiveIsNoop(t *testing.T) {
	s := NewSelectedState().MoveDown(10)
	if s.Active() {
		t.Error("expected move on inactive state to remain inactive")
	}
}

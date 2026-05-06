package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPinStore_Empty(t *testing.T) {
	p := NewPinStore()
	assert.Equal(t, 0, p.Count())
	assert.Empty(t, p.All())
}

func TestPinStore_Toggle_Add(t *testing.T) {
	p := NewPinStore()
	p.Toggle("svc-a")
	assert.True(t, p.IsPinned("svc-a"))
	assert.Equal(t, 1, p.Count())
}

func TestPinStore_Toggle_Remove(t *testing.T) {
	p := NewPinStore()
	p.Toggle("svc-a")
	p.Toggle("svc-a")
	assert.False(t, p.IsPinned("svc-a"))
	assert.Equal(t, 0, p.Count())
}

func TestPinStore_Toggle_EmptyKeyIgnored(t *testing.T) {
	p := NewPinStore()
	p.Toggle("")
	assert.Equal(t, 0, p.Count())
}

func TestPinStore_All_Sorted(t *testing.T) {
	p := NewPinStore()
	p.Toggle("svc-c")
	p.Toggle("svc-a")
	p.Toggle("svc-b")
	assert.Equal(t, []string{"svc-a", "svc-b", "svc-c"}, p.All())
}

func TestPinStore_Clear(t *testing.T) {
	p := NewPinStore()
	p.Toggle("svc-a")
	p.Toggle("svc-b")
	p.Clear()
	assert.Equal(t, 0, p.Count())
	assert.Empty(t, p.All())
}

func TestApplyPins_NilStore(t *testing.T) {
	targets := []string{"a", "b", "c"}
	result := ApplyPins(targets, nil)
	assert.Equal(t, targets, result)
}

func TestApplyPins_Nopins(t *testing.T) {
	p := NewPinStore()
	targets := []string{"a", "b", "c"}
	result := ApplyPins(targets, p)
	assert.Equal(t, targets, result)
}

func TestApplyPins_PinnedFirst(t *testing.T) {
	p := NewPinStore()
	p.Toggle("c")
	targets := []string{"a", "b", "c"}
	result := ApplyPins(targets, p)
	assert.Equal(t, []string{"c", "a", "b"}, result)
}

func TestApplyPins_PreservesRelativeOrder(t *testing.T) {
	p := NewPinStore()
	p.Toggle("b")
	p.Toggle("d")
	targets := []string{"a", "b", "c", "d", "e"}
	result := ApplyPins(targets, p)
	assert.Equal(t, []string{"b", "d", "a", "c", "e"}, result)
}

func TestApplyPins_DoesNotMutateOriginal(t *testing.T) {
	p := NewPinStore()
	p.Toggle("b")
	original := []string{"a", "b", "c"}
	copy := append([]string{}, original...)
	ApplyPins(original, p)
	assert.Equal(t, copy, original)
}

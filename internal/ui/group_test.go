package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGroupStore_Empty(t *testing.T) {
	g := NewGroupStore()
	assert.Empty(t, g.All())
}

func TestGroupStore_AddAndGet(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	g.Add("prod", "svc-b")
	assert.Equal(t, []string{"svc-a", "svc-b"}, g.Members("prod"))
}

func TestGroupStore_AddEmptyIgnored(t *testing.T) {
	g := NewGroupStore()
	g.Add("", "svc-a")
	g.Add("prod", "")
	assert.Empty(t, g.All())
}

func TestGroupStore_Remove(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	g.Remove("prod", "svc-a")
	assert.Empty(t, g.All())
}

func TestGroupStore_RemoveEmptiesGroup(t *testing.T) {
	g := NewGroupStore()
	g.Add("staging", "svc-x")
	g.Remove("staging", "svc-x")
	assert.NotContains(t, g.All(), "staging")
}

func TestGroupStore_GroupsFor(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	g.Add("eu", "svc-a")
	result := g.GroupsFor("svc-a")
	assert.Equal(t, []string{"eu", "prod"}, result)
}

func TestGroupStore_GroupsFor_Unknown(t *testing.T) {
	g := NewGroupStore()
	assert.Empty(t, g.GroupsFor("missing"))
}

func TestGroupStore_All_Sorted(t *testing.T) {
	g := NewGroupStore()
	g.Add("zzz", "k")
	g.Add("aaa", "k")
	g.Add("mmm", "k")
	assert.Equal(t, []string{"aaa", "mmm", "zzz"}, g.All())
}

func TestGroupStore_Clear(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	g.Clear()
	assert.Empty(t, g.All())
}

func TestGroupStore_Members_Unknown(t *testing.T) {
	g := NewGroupStore()
	assert.Nil(t, g.Members("nonexistent"))
}

package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBadge_NoGroups(t *testing.T) {
	g := NewGroupStore()
	assert.Equal(t, "", GroupBadge(g, "svc-a"))
}

func TestGroupBadge_Nil(t *testing.T) {
	assert.Equal(t, "", GroupBadge(nil, "svc-a"))
}

func TestGroupBadge_SingleGroup(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	badge := GroupBadge(g, "svc-a")
	assert.Equal(t, "[prod]", badge)
}

func TestGroupBadge_MultipleGroups(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	g.Add("eu", "svc-a")
	badge := GroupBadge(g, "svc-a")
	assert.Equal(t, "[eu,prod]", badge)
}

func TestGroupListView_Nil(t *testing.T) {
	assert.Equal(t, "(no groups)", GroupListView(nil))
}

func TestGroupListView_Empty(t *testing.T) {
	g := NewGroupStore()
	assert.Equal(t, "(no groups)", GroupListView(g))
}

func TestGroupListView_ShowsGroups(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	g.Add("prod", "svc-b")
	view := GroupListView(g)
	assert.Contains(t, view, "prod")
	assert.Contains(t, view, "svc-a")
	assert.Contains(t, view, "svc-b")
}

func TestGroupListView_MultipleGroups(t *testing.T) {
	g := NewGroupStore()
	g.Add("prod", "svc-a")
	g.Add("staging", "svc-b")
	view := GroupListView(g)
	assert.Contains(t, view, "prod")
	assert.Contains(t, view, "staging")
}

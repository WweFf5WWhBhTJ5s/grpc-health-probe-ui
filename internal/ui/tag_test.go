package ui

import (
	"testing"
)

func TestNewTagStore_Empty(t *testing.T) {
	s := NewTagStore()
	if len(s.AllTags()) != 0 {
		t.Fatal("expected empty tag store")
	}
}

func TestTagStore_AddAndGet(t *testing.T) {
	s := NewTagStore()
	s.Add("svc-a", "prod")
	s.Add("svc-a", "critical")
	tags := s.Get("svc-a")
	if len(tags) != 2 {
		t.Fatalf("expected 2 tags, got %d", len(tags))
	}
	// sorted: critical, prod
	if tags[0] != "critical" || tags[1] != "prod" {
		t.Errorf("unexpected order: %v", tags)
	}
}

func TestTagStore_AddEmptyTagIgnored(t *testing.T) {
	s := NewTagStore()
	s.Add("svc-a", "   ")
	if len(s.Get("svc-a")) != 0 {
		t.Fatal("whitespace-only tag should be ignored")
	}
}

func TestTagStore_Has(t *testing.T) {
	s := NewTagStore()
	s.Add("svc-a", "prod")
	if !s.Has("svc-a", "prod") {
		t.Error("expected Has to return true")
	}
	if s.Has("svc-a", "staging") {
		t.Error("expected Has to return false for absent tag")
	}
}

func TestTagStore_Remove(t *testing.T) {
	s := NewTagStore()
	s.Add("svc-a", "prod")
	s.Remove("svc-a", "prod")
	if s.Has("svc-a", "prod") {
		t.Error("tag should have been removed")
	}
	if len(s.Get("svc-a")) != 0 {
		t.Error("expected empty tag list after removal")
	}
}

func TestTagStore_RemoveAbsent(t *testing.T) {
	s := NewTagStore()
	// should not panic
	s.Remove("svc-unknown", "prod")
}

func TestTagStore_AllTags_Deduped(t *testing.T) {
	s := NewTagStore()
	s.Add("svc-a", "prod")
	s.Add("svc-b", "prod")
	s.Add("svc-b", "staging")
	all := s.AllTags()
	if len(all) != 2 {
		t.Fatalf("expected 2 unique tags, got %d: %v", len(all), all)
	}
}

func TestTagStore_GetUnknownTarget(t *testing.T) {
	s := NewTagStore()
	if tags := s.Get("nope"); tags != nil {
		t.Errorf("expected nil for unknown target, got %v", tags)
	}
}

func TestTagStore_ClearTarget(t *testing.T) {
	s := NewTagStore()
	s.Add("svc-a", "prod")
	s.Add("svc-a", "critical")
	s.ClearTarget("svc-a")
	if len(s.Get("svc-a")) != 0 {
		t.Error("expected all tags cleared")
	}
	if len(s.AllTags()) != 0 {
		t.Error("expected AllTags empty after clear")
	}
}

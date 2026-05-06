package ui

import (
	"strings"
	"testing"
)

func TestNewNotifyStore_Defaults(t *testing.T) {
	s := NewNotifyStore(0)
	if s.maxLen != 50 {
		t.Fatalf("expected default maxLen 50, got %d", s.maxLen)
	}
	if s.Len() != 0 {
		t.Fatalf("expected empty store, got %d entries", s.Len())
	}
}

func TestNotifyStore_Push(t *testing.T) {
	s := NewNotifyStore(10)
	s.Push("hello", NotifyInfo)
	if s.Len() != 1 {
		t.Fatalf("expected 1 entry, got %d", s.Len())
	}
	entries := s.All()
	if entries[0].Message != "hello" {
		t.Fatalf("unexpected message: %q", entries[0].Message)
	}
	if entries[0].Level != NotifyInfo {
		t.Fatalf("unexpected level: %d", entries[0].Level)
	}
}

func TestNotifyStore_EvictsWhenFull(t *testing.T) {
	s := NewNotifyStore(3)
	s.Push("first", NotifyInfo)
	s.Push("second", NotifyWarn)
	s.Push("third", NotifyError)
	s.Push("fourth", NotifyInfo)
	if s.Len() != 3 {
		t.Fatalf("expected 3 entries after eviction, got %d", s.Len())
	}
	if s.All()[0].Message != "second" {
		t.Fatalf("expected oldest to be evicted")
	}
}

func TestNotifyStore_Clear(t *testing.T) {
	s := NewNotifyStore(10)
	s.Push("a", NotifyInfo)
	s.Push("b", NotifyWarn)
	s.Clear()
	if s.Len() != 0 {
		t.Fatalf("expected empty store after clear, got %d", s.Len())
	}
}

func TestNotifyStore_AllReturnsCopy(t *testing.T) {
	s := NewNotifyStore(10)
	s.Push("msg", NotifyInfo)
	a := s.All()
	a[0].Message = "mutated"
	if s.All()[0].Message != "msg" {
		t.Fatal("All() should return a copy, not a reference")
	}
}

func TestLevelLabel(t *testing.T) {
	if LevelLabel(NotifyInfo) != "INFO" {
		t.Fatal("expected INFO")
	}
	if LevelLabel(NotifyWarn) != "WARN" {
		t.Fatal("expected WARN")
	}
	if LevelLabel(NotifyError) != "ERR " {
		t.Fatal("expected ERR ")
	}
}

func TestNotifyView_Empty(t *testing.T) {
	s := NewNotifyStore(10)
	out := NotifyView(s, 80)
	if out != "" {
		t.Fatalf("expected empty output for empty store, got %q", out)
	}
}

func TestNotifyView_Nil(t *testing.T) {
	out := NotifyView(nil, 80)
	if out != "" {
		t.Fatalf("expected empty output for nil store")
	}
}

func TestNotifyView_ContainsMessage(t *testing.T) {
	s := NewNotifyStore(10)
	s.Push("service down", NotifyError)
	out := NotifyView(s, 80)
	if !strings.Contains(out, "service down") {
		t.Fatalf("expected view to contain message, got:\n%s", out)
	}
}

func TestNotifyView_ContainsLevelLabel(t *testing.T) {
	s := NewNotifyStore(10)
	s.Push("latency spike", NotifyWarn)
	out := NotifyView(s, 80)
	if !strings.Contains(out, "WARN") {
		t.Fatalf("expected view to contain WARN label, got:\n%s", out)
	}
}

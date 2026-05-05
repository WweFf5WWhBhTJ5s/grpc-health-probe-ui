package ui

import (
	"strings"
	"testing"
)

func TestNewNoteStore_Empty(t *testing.T) {
	s := NewNoteStore()
	if s.Len() != 0 {
		t.Fatalf("expected 0 notes, got %d", s.Len())
	}
}

func TestNoteStore_SetAndGet(t *testing.T) {
	s := NewNoteStore()
	s.Set("svc-a", "check this first")
	n, ok := s.Get("svc-a")
	if !ok || n != "check this first" {
		t.Fatalf("unexpected note: %q %v", n, ok)
	}
}

func TestNoteStore_SetEmpty_Removes(t *testing.T) {
	s := NewNoteStore()
	s.Set("svc-a", "temp")
	s.Set("svc-a", "")
	if s.Has("svc-a") {
		t.Fatal("expected note to be removed")
	}
}

func TestNoteStore_SetTrimmed(t *testing.T) {
	s := NewNoteStore()
	s.Set("svc-b", "  spaces  ")
	n, _ := s.Get("svc-b")
	if n != "spaces" {
		t.Fatalf("expected trimmed note, got %q", n)
	}
}

func TestNoteStore_GetUnknown(t *testing.T) {
	s := NewNoteStore()
	_, ok := s.Get("missing")
	if ok {
		t.Fatal("expected false for unknown key")
	}
}

func TestNoteStore_All_Sorted(t *testing.T) {
	s := NewNoteStore()
	s.Set("z-svc", "last")
	s.Set("a-svc", "first")
	s.Set("m-svc", "middle")
	all := s.All()
	if all[0] != "a-svc" || all[1] != "m-svc" || all[2] != "z-svc" {
		t.Fatalf("unexpected order: %v", all)
	}
}

func TestNoteStore_Clear(t *testing.T) {
	s := NewNoteStore()
	s.Set("svc-a", "note")
	s.Clear()
	if s.Len() != 0 {
		t.Fatal("expected empty store after Clear")
	}
}

func TestNoteBadge_Present(t *testing.T) {
	s := NewNoteStore()
	s.Set("svc-a", "hello")
	if NoteBadge(s, "svc-a") != noteBadgeSymbol {
		t.Fatal("expected badge symbol")
	}
}

func TestNoteBadge_Absent(t *testing.T) {
	s := NewNoteStore()
	if NoteBadge(s, "svc-a") != "" {
		t.Fatal("expected empty badge")
	}
}

func TestNoteBadge_NilStore(t *testing.T) {
	if NoteBadge(nil, "svc-a") != "" {
		t.Fatal("expected empty badge for nil store")
	}
}

func TestNoteRow_Present(t *testing.T) {
	s := NewNoteStore()
	s.Set("svc-a", "important")
	row := NoteRow(s, "svc-a")
	if !strings.Contains(row, "important") {
		t.Fatalf("expected note text in row, got %q", row)
	}
}

func TestNoteRow_Absent(t *testing.T) {
	s := NewNoteStore()
	if NoteRow(s, "svc-a") != "" {
		t.Fatal("expected empty row when no note")
	}
}

func TestNoteListView_Empty(t *testing.T) {
	s := NewNoteStore()
	out := NoteListView(s)
	if out != "No notes." {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestNoteListView_ShowsEntries(t *testing.T) {
	s := NewNoteStore()
	s.Set("svc-a", "note one")
	s.Set("svc-b", "note two")
	out := NoteListView(s)
	if !strings.Contains(out, "svc-a") || !strings.Contains(out, "note one") {
		t.Fatalf("missing expected content: %q", out)
	}
}

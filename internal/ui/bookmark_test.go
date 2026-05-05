package ui

import (
	"testing"
)

func TestNewBookmarkStore_Empty(t *testing.T) {
	b := NewBookmarkStore()
	if b.Count() != 0 {
		t.Fatalf("expected 0 bookmarks, got %d", b.Count())
	}
}

func TestBookmark_Toggle_Add(t *testing.T) {
	b := NewBookmarkStore()
	added := b.Toggle("svc-a")
	if !added {
		t.Fatal("expected Toggle to return true on add")
	}
	if !b.IsBookmarked("svc-a") {
		t.Fatal("expected svc-a to be bookmarked")
	}
}

func TestBookmark_Toggle_Remove(t *testing.T) {
	b := NewBookmarkStore()
	b.Toggle("svc-a")
	removed := b.Toggle("svc-a")
	if removed {
		t.Fatal("expected Toggle to return false on remove")
	}
	if b.IsBookmarked("svc-a") {
		t.Fatal("expected svc-a to not be bookmarked after removal")
	}
}

func TestBookmark_All_Sorted(t *testing.T) {
	b := NewBookmarkStore()
	b.Toggle("zebra")
	b.Toggle("alpha")
	b.Toggle("mango")
	all := b.All()
	if len(all) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(all))
	}
	if all[0] != "alpha" || all[1] != "mango" || all[2] != "zebra" {
		t.Fatalf("unexpected order: %v", all)
	}
}

func TestBookmark_Clear(t *testing.T) {
	b := NewBookmarkStore()
	b.Toggle("svc-a")
	b.Toggle("svc-b")
	b.Clear()
	if b.Count() != 0 {
		t.Fatalf("expected 0 after clear, got %d", b.Count())
	}
}

func TestBookmark_IsBookmarked_Unknown(t *testing.T) {
	b := NewBookmarkStore()
	if b.IsBookmarked("ghost") {
		t.Fatal("expected ghost to not be bookmarked")
	}
}

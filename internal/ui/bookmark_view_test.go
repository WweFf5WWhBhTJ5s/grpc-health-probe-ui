package ui

import (
	"strings"
	"testing"
)

func TestBookmarkBadge_NotBookmarked(t *testing.T) {
	theme := DefaultTheme()
	badge := BookmarkBadge(false, theme)
	if badge != "" {
		t.Fatalf("expected empty badge for non-bookmarked, got %q", badge)
	}
}

func TestBookmarkBadge_Bookmarked(t *testing.T) {
	theme := DefaultTheme()
	badge := BookmarkBadge(true, theme)
	if !strings.Contains(badge, "★") {
		t.Fatalf("expected star in badge, got %q", badge)
	}
}

func TestBookmarkListView_Empty(t *testing.T) {
	theme := DefaultTheme()
	store := NewBookmarkStore()
	out := BookmarkListView(store, theme)
	if !strings.Contains(out, "No bookmarks") {
		t.Fatalf("expected placeholder, got %q", out)
	}
}

func TestBookmarkListView_Nil(t *testing.T) {
	theme := DefaultTheme()
	out := BookmarkListView(nil, theme)
	if !strings.Contains(out, "No bookmarks") {
		t.Fatalf("expected placeholder for nil store, got %q", out)
	}
}

func TestBookmarkListView_ShowsNames(t *testing.T) {
	theme := DefaultTheme()
	store := NewBookmarkStore()
	store.Toggle("payment-svc")
	store.Toggle("auth-svc")
	out := BookmarkListView(store, theme)
	if !strings.Contains(out, "payment-svc") {
		t.Fatalf("expected payment-svc in output, got %q", out)
	}
	if !strings.Contains(out, "auth-svc") {
		t.Fatalf("expected auth-svc in output, got %q", out)
	}
}

func TestBookmarkListView_ContainsTitle(t *testing.T) {
	theme := DefaultTheme()
	store := NewBookmarkStore()
	store.Toggle("svc-x")
	out := BookmarkListView(store, theme)
	if !strings.Contains(out, "Bookmarks") {
		t.Fatalf("expected title in output, got %q", out)
	}
}

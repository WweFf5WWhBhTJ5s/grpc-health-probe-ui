package ui

import (
	"testing"
)

func TestNewPagination_Defaults(t *testing.T) {
	p := NewPagination(5, 20)
	if p.PageSize != 5 {
		t.Errorf("expected PageSize 5, got %d", p.PageSize)
	}
	if p.Page != 0 {
		t.Errorf("expected Page 0, got %d", p.Page)
	}
	if p.Total != 20 {
		t.Errorf("expected Total 20, got %d", p.Total)
	}
}

func TestNewPagination_InvalidPageSize(t *testing.T) {
	p := NewPagination(0, 10)
	if p.PageSize != 10 {
		t.Errorf("expected default PageSize 10, got %d", p.PageSize)
	}
}

func TestTotalPages(t *testing.T) {
	cases := []struct {
		pageSize, total, want int
	}{
		{5, 20, 4},
		{5, 21, 5},
		{5, 0, 1},
		{10, 10, 1},
		{10, 11, 2},
	}
	for _, tc := range cases {
		p := NewPagination(tc.pageSize, tc.total)
		if got := p.TotalPages(); got != tc.want {
			t.Errorf("TotalPages(%d/%d) = %d, want %d", tc.total, tc.pageSize, got, tc.want)
		}
	}
}

func TestNext_ClampsAtLastPage(t *testing.T) {
	p := NewPagination(5, 10) // 2 pages
	p = p.Next()              // page 1
	p = p.Next()              // should clamp at 1
	if p.Page != 1 {
		t.Errorf("expected Page 1 after clamping, got %d", p.Page)
	}
}

func TestPrev_ClampsAtZero(t *testing.T) {
	p := NewPagination(5, 10)
	p = p.Prev() // already at 0
	if p.Page != 0 {
		t.Errorf("expected Page 0 after clamping, got %d", p.Page)
	}
}

func TestSlice(t *testing.T) {
	p := NewPagination(3, 7)
	s, e := p.Slice()
	if s != 0 || e != 3 {
		t.Errorf("page 0 slice: want 0-3, got %d-%d", s, e)
	}
	p = p.Next()
	s, e = p.Slice()
	if s != 3 || e != 6 {
		t.Errorf("page 1 slice: want 3-6, got %d-%d", s, e)
	}
	p = p.Next()
	s, e = p.Slice()
	if s != 6 || e != 7 {
		t.Errorf("page 2 slice: want 6-7, got %d-%d", s, e)
	}
}

func TestLabel(t *testing.T) {
	p := NewPagination(5, 5) // 1 page
	if label := p.Label(); label != "" {
		t.Errorf("expected empty label for single page, got %q", label)
	}

	p = NewPagination(5, 10) // 2 pages
	if label := p.Label(); label != "Page 1/2" {
		t.Errorf("expected 'Page 1/2', got %q", label)
	}
	p = p.Next()
	if label := p.Label(); label != "Page 2/2" {
		t.Errorf("expected 'Page 2/2', got %q", label)
	}
}

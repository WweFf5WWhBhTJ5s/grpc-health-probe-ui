package ui

import "fmt"

// Pagination manages paging through a list of targets in the terminal UI.
type Pagination struct {
	PageSize int
	Page     int
	Total    int
}

// NewPagination creates a Pagination with the given page size and total items.
func NewPagination(pageSize, total int) Pagination {
	if pageSize <= 0 {
		pageSize = 10
	}
	return Pagination{
		PageSize: pageSize,
		Page:     0,
		Total:    total,
	}
}

// TotalPages returns the number of pages needed to display all items.
func (p Pagination) TotalPages() int {
	if p.Total == 0 {
		return 1
	}
	pages := p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		pages++
	}
	return pages
}

// Next advances to the next page, clamping at the last page.
func (p Pagination) Next() Pagination {
	next := p.Page + 1
	if next >= p.TotalPages() {
		next = p.TotalPages() - 1
	}
	p.Page = next
	return p
}

// Prev moves to the previous page, clamping at zero.
func (p Pagination) Prev() Pagination {
	prev := p.Page - 1
	if prev < 0 {
		prev = 0
	}
	p.Page = prev
	return p
}

// Slice returns the start and end indices for the current page.
func (p Pagination) Slice() (start, end int) {
	start = p.Page * p.PageSize
	end = start + p.PageSize
	if end > p.Total {
		end = p.Total
	}
	return start, end
}

// Label returns a human-readable page indicator, e.g. "Page 1/3".
// Returns an empty string when there is only one page.
func (p Pagination) Label() string {
	if p.TotalPages() <= 1 {
		return ""
	}
	return fmt.Sprintf("Page %d/%d", p.Page+1, p.TotalPages())
}

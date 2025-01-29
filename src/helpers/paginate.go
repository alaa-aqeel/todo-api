package helpers

import (
	"math"
	"net/http"
	"strconv"
)

type Pagination struct {
	page       int
	limit      int
	total      int
	totalPages int
	data       any
}

func NewPagination(r *http.Request) *Pagination {
	p := &Pagination{}
	p.GetPaginationArgs(r)

	return p
}

func (p *Pagination) GetPaginationArgs(r *http.Request) {
	p.page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	p.limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))

	if p.page < 1 {
		p.page = 1
	}
	if p.limit < 1 || p.limit > 100 {
		p.limit = 10
	}
}

func (p *Pagination) SetTotal(total int64) {
	p.total = int(total)
	p.totalPages = int(math.Ceil(float64(p.total) / float64(p.limit)))
}

func (p *Pagination) SetData(data any) {
	p.data = data
}

func (p Pagination) Offset() int {
	return (p.page - 1) * p.limit
}

func (p Pagination) TotalPages() int {
	return p.totalPages
}

func (p Pagination) Total() int {
	return p.total
}

func (p Pagination) Limit() int {
	return p.limit
}

func (p Pagination) Page() int {
	return p.page
}

func (p Pagination) NextPage() int {
	if p.page < p.totalPages {
		return p.page + 1
	}
	return p.page
}

func (p Pagination) PreviousPage() int {
	if p.page > 1 {
		return p.page - 1
	}
	return p.page
}

func (p Pagination) HasNext() bool {
	return p.page < p.totalPages
}

func (p Pagination) HasPrevious() bool {
	return p.page > 1
}

func (p Pagination) IsLast() bool {
	return p.page == p.totalPages
}

func (p Pagination) IsFirst() bool {
	return p.page == 1
}

func (p Pagination) HasPagination() bool {
	return p.total > p.limit
}

func (p Pagination) IsEmpty() bool {
	return p.total == 0
}

func (p Pagination) ToMap() map[string]interface{} {
	return Map{
		"page":       p.page,
		"limit":      p.limit,
		"total":      p.total,
		"totalPages": p.totalPages,
		"data":       p.data,
	}
}

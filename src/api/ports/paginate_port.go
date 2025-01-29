package ports

import (
	"net/http"
)

type PaginatorPort interface {
	GetPaginationArgs(r *http.Request)
	Offset() int
	SetTotal(total int64)
	TotalPages() int
	Total() int
	Limit() int
	Page() int
	HasPagination() bool
	IsEmpty() bool
	ToMap() map[string]interface{}
	SetData(data any)
}

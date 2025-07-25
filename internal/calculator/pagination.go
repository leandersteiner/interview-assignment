package calculator

import (
	"github.com/leandersteiner/interview-assignment/internal/web"
	"net/http"
)

const DefaultPageSize = 5
const MaxPageSize = 20
const MinPageSize = 1

type PaginatedResult[T any] struct {
	Result T
	Metadata
}

type Metadata struct {
	CurrentPage  int
	PageSize     int
	TotalRecords int
	LastPage     int
	FirstPage    int
	NextPage     int
}

type Pagination struct {
	Page     int
	PageSize int
}

func (p *Pagination) Validate() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < MinPageSize {
		p.PageSize = DefaultPageSize
	}
	if p.PageSize > MaxPageSize {
		p.PageSize = MaxPageSize
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) Limit() int {
	return p.PageSize
}

func (p *Pagination) toMetadata(total int) Metadata {
	lastPage := total / p.PageSize
	nextPage := p.Page + 1
	if lastPage == 0 {
		lastPage = 1
	}
	if nextPage > lastPage {
		nextPage = lastPage
	}

	return Metadata{
		CurrentPage:  p.Page,
		PageSize:     p.PageSize,
		TotalRecords: total,
		LastPage:     lastPage,
		FirstPage:    1,
		NextPage:     nextPage,
	}
}

func parsePagination(r *http.Request) Pagination {
	query := r.URL.Query()

	return Pagination{
		Page:     web.GetIntParam(query, "page", 1),
		PageSize: web.GetIntParam(query, "page_size", DefaultPageSize),
	}
}

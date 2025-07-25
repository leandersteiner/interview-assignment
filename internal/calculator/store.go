package calculator

import (
	"sync"
)

type Store interface {
	storer
	getter
}

type ResultStore struct {
	results []Result
	mu      sync.RWMutex
}

func NewResultStore() *ResultStore {
	return &ResultStore{
		results: []Result{},
	}
}

func (s *ResultStore) Store(result Result) {
	s.mu.Lock()
	s.results = append([]Result{result}, s.results...)
	s.mu.Unlock()
}

func (s *ResultStore) Get(p Pagination) PaginatedResult[[]Result] {
	p.Validate()

	if len(s.results) == 0 || p.Offset() >= len(s.results) {
		return PaginatedResult[[]Result]{
			Result:   []Result{},
			Metadata: p.toMetadata(len(s.results)),
		}
	}

	startIndex := p.Offset()
	endIndex := startIndex + p.Limit()

	if endIndex > len(s.results) {
		endIndex = len(s.results)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	return PaginatedResult[[]Result]{
		Result:   s.results[startIndex:endIndex],
		Metadata: p.toMetadata(len(s.results)),
	}
}

package tools

import "sync"

type ConcurrencyMap[T any] struct {
	sync.Mutex

	slice []T
}

func NewConcurrencyMap[T any](length int) *ConcurrencyMap[T] {
	return &ConcurrencyMap[T]{
		slice: make([]T, 0, length),
	}
}

func (m *ConcurrencyMap[T]) Append(in ...T) {
	m.Lock()
	m.slice = append(m.slice, in...)
	m.Unlock()
}

func (m *ConcurrencyMap[T]) Get() []T {
	return m.slice
}

package mem_store

import "fmt"

type Store[K comparable, V any] interface {
	Find(id K) (*V, bool)
	Store(id K, value V) (*V, error)
}

type MemStore[K comparable, V any] struct {
	store map[K]V
}

func newMemStore[K comparable, V any]() *MemStore[K, V] {
	return &MemStore[K, V]{
		store: map[K]V{},
	}
}

func NewMemStore[K comparable, V any]() Store[K, V] {
	return newMemStore[K, V]()
}

func (s *MemStore[K, V]) Find(id K) (*V, bool) {
	val, found := s.store[id]
	return &val, found
}

func (s *MemStore[K, V]) Store(id K, value V) (*V, error) {
	s.store[id] = value

	return &value, nil
}

// Validation is still easy to implement
type MemStoreWithValidation[K comparable, V any] struct {
	*MemStore[K, V]
	validate func(V) error
}

func NewMemStoreWithValidation[K comparable, V any](validator func(val V) error) Store[K, V] {
	return &MemStoreWithValidation[K, V]{
		MemStore: newMemStore[K, V](),
		validate: validator,
	}
}

func (s *MemStoreWithValidation[K, V]) Find(id K) (*V, bool) {
	return s.MemStore.Find(id)
}

func (s *MemStoreWithValidation[K, V]) Store(id K, value V) (*V, error) {
	if err := s.validate(value); err != nil {
		return nil, fmt.Errorf("value is invalid: %w", err)
	}

	return s.MemStore.Store(id, value)
}

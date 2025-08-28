package set

import "iter"

type Set[T comparable] interface {
	Add(T)
	Del(T)
	Contains(T) bool
	AddMultiple(iter.Seq2[int, T])
	Intersection(Set[T]) Set[T]
}

func New[T comparable]() Set[T] {
	return &setM[T]{
		data: make(map[T]bool, 0),
	}
}

type setM[T comparable] struct {
	data map[T]bool
}

func (s *setM[T]) Add(item T) {
	s.data[item] = true
}

func (s *setM[T]) Del(item T) {
	delete(s.data, item)
}

func (s *setM[T]) Contains(item T) bool {
	_, contains := s.data[item]
	return contains
}

func (s *setM[T]) AddMultiple(items iter.Seq2[int, T]) {
	for _, item := range items {
		s.data[item] = true
	}
}

// returns new Set, that contains intersection of two
func (s *setM[T]) Intersection(other Set[T]) Set[T] {
	intersection := New[T]()
	for item := range s.data {
		if other.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

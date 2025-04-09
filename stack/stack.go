package stack

type Stack[T any] struct {
	data []T
	length int
}

func New[T any]() *Stack[T] {
	return new(Stack[T]).Init()
}

func (s *Stack[T]) Init() *Stack[T] {
	s.data = make([]T, 0)
	s.length = 0
	return s
}

func (s *Stack[T]) Push(elem T) {
	if s.length == len(s.data) {
		s.data = append(s.data, elem)
	} else {
		s.data[s.length] = elem
	}
	s.length++
}

func (s *Stack[T]) Pop() T {
	if s.length == 0 {
		panic("Cannot pop from empty stack")
	}
	s.length--
	return s.data[s.length]
}

func (s *Stack[T]) GetLast() T {
	return s.data[s.Len() - 1]
}

func (s *Stack[T]) Len() int {
	return s.length
}

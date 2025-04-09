package list

import (
	"iter"
)

type node[T any] struct {
	value T
	next  *node[T]
}

type List[T any] struct {
	head   *node[T]
	length int
}

func newNode[T any](value T) *node[T] {
	n := new(node[T])
	n.value = value
	return n
}

func New[T any]() *List[T] {
	l := new(List[T])
	return l
}

func (l *List[T]) AddFirst(value T) {
	tail := l.head
	l.head = newNode(value)
	l.head.next = tail
	l.length++
}

func (l *List[T]) AddLast(value T) {
	panic("TODO")
}

func (l *List[T]) PopFirst() T {
	panic("TODO")
}

func (l *List[T]) PopLast() T {
	panic("TODO")
}

func (l *List[T]) Remove(value T) T {
	panic("TODO")
}

func (l *List[T]) Find(value T) int {
	panic("TODO")
}

func (l *List[T]) Len() int {
	return l.length
}

func (l *List[T]) ElementsIter() iter.Seq[T] {
	headPtr := l.head
	return func(yield func(T) bool) {
		for headPtr != nil {
			if !yield(headPtr.value) {
				return
			}
			headPtr = headPtr.next
		}
	}
}

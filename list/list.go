// implementation of single linked list.
package list

import (
	"errors"
	"iter"
)

type node[T comparable] struct {
	value T
	next  *node[T]
}

type List[T comparable] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func newNode[T comparable](value T) *node[T] {
	n := new(node[T])
	n.value = value
	return n
}

func New[T comparable]() *List[T] {
	l := new(List[T])
	return l
}

func (l *List[T]) AddFirst(value T) {
	tail := l.head
	l.head = newNode(value)
	if tail == nil {
		l.tail = l.head
	} else {
		l.head.next = tail
	}
	l.length++
}

func (l *List[T]) AddLast(value T) {
	curTail := l.tail
	node := newNode(value)
	if curTail == nil {
		l.tail = node
		l.head = node
	} else {
		curTail.next = node
		l.tail = curTail.next
	}
	l.length++
}

func (l *List[T]) GetFirst() (T, error) {
	if l.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("Can't Get element from empty list")
	}
	return l.head.value, nil
}

func (l *List[T]) GetLast() (T, error) {
	if l.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("Can't Get element from empty list")
	}
	return l.tail.value, nil
}

func (l *List[T]) PopFirst() (T, error) {
	if l.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("Can't Pop from empty list")
	}
	poped := l.head
	l.head = poped.next
	l.length--
	return poped.value, nil
}

func (l *List[T]) PopLast() (T, error) {
	if l.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("Can't Pop from empty list")
	}
	if l.Len() == 1 {
		ret := l.head.value
		l.head = nil
		l.tail = nil
		l.length = 0
		return ret, nil
	}
	iterator := l.head
	l.length--
	for range l.Len() - 1 {
		iterator = iterator.next
	}
	ret := iterator.next.value
	iterator.next = nil
	return ret, nil
}

// Removes first element == value.
// return 0 if element was deleted,
// return -1 otherwise.
func (l *List[T]) Remove(value T) int {
	if l.Len() == 0 {
		return -1
	}
	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return 0
	}
	iterator := l.head
	for range l.length - 1 {
		if iterator.next.value == value {
			iterator.next = iterator.next.next
			l.length--
			return 0
		}
	}
	return -1
}

// return index of the first element == value,
// return -1 if no such element in list
func (l *List[T]) Find(value T) int {
	iterator := l.head
	for i := range l.length {
		if iterator.value == value {
			return i
		}
		iterator = iterator.next
	}
	return -1
}

func (l *List[T]) Len() int {
	return l.length
}

func (l *List[T]) IsEmpty() bool {
	return l.head == nil
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

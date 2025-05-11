package heap

import (
	"cmp"
	"errors"
)

func min2[T cmp.Ordered](p1, p2 T) T {
	if p1 < p2 {
		return p1
	}
	return p2
}

func max2[T cmp.Ordered](p1, p2 T) T {
	if p1 > p2 {
		return p1
	}
	return p2
}

func isSmaller[T cmp.Ordered](p1, p2 T) bool {
	return p1 < p2
}

func isGreater[T cmp.Ordered](p1, p2 T) bool {
	return p1 > p2
}

type Heap[T cmp.Ordered] struct {
	data       []T
	comparator func(T, T) bool
}

func NewHeap[T cmp.Ordered](comparator func(T, T) bool) *Heap[T] {
	h := new(Heap[T])
	h.comparator = comparator
	return h
}

func NewMinHeap[T cmp.Ordered]() *Heap[T] {
	h := new(Heap[T])
	h.data = make([]T, 0)
	h.comparator = isSmaller
	return h
}

func NewMaxHeap[T cmp.Ordered]() *Heap[T] {
	h := new(Heap[T])
	h.data = make([]T, 0)
	h.comparator = isGreater
	return h
}

func FromSliceMinHeap[T cmp.Ordered](data []T) *Heap[T] {
	h := new(Heap[T])
	h.data = make([]T, len(data))
	copy(h.data, data)
	h.comparator = isSmaller
	h.heapify()
	return h
}

func (h *Heap[T]) Push(elem T) {
	h.data = append(h.data, elem)
	h.siftup(len(h.data) - 1)
}

func (h *Heap[T]) Pop() (T, error) {
	if len(h.data) == 0 {
		var ret T
		return ret, errors.New("Pop from empty Heap")
	}
	if len(h.data) == 1 {
		poped := h.data[0]
		h.data = make([]T, 0)
		return poped, nil
	}
	temp := h.data[len(h.data)-1]
	h.data[len(h.data)-1] = h.data[0]
	h.data[0] = temp
	poped := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftdown(0)
	return poped, nil
}

func (h *Heap[T]) swap(i1, i2 int) {
	temp := h.data[i1]
	h.data[i1] = h.data[i2]
	h.data[i2] = temp
}

func (h *Heap[T]) siftup(index int) {
	for {
		parent := (index - 1) / 2
		if parent == index || !h.comparator(h.data[index], h.data[parent]) {
			break
		}
		h.swap(parent, index)
		index = parent
	}

}

func (h *Heap[T]) siftdown(index int) {
	n := len(h.data)
	for {
		j1 := 2*index + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		child := j1 // left child
		if j2 := j1 + 1; j2 < n && h.comparator(h.data[j2], h.data[j1]) {
			child = j2 // = 2*i + 2  // right child
		}
		if !h.comparator(h.data[child], h.data[index]) {
			break
		}
		h.swap(index, child)
		index = child
	}
}

func (h *Heap[T]) heapify() {
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.siftdown(i)
	}
}

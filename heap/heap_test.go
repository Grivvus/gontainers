package heap_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/Grivvus/gontainers/heap"
	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	h := heap.NewMinHeap[int]()
	_, err := h.Pop()
	assert.Error(t, err, "Expected error while poping from empty heap")
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for _, num := range nums {
		h.Push(num)
	}
	for i := range nums {
		elem, err := h.Pop()
		assert.NoErrorf(t, err, "Unexpected error %v on index %v", err, i)
		assert.Equalf(t, nums[len(nums)-1-i], elem, "Wrong number, expected %v got %v", nums[len(nums)-1-i], elem)
	}
}

func BenchmarkMinHeap(b *testing.B) {
	add1000Elems := func() {
		heap := heap.NewMinHeap[int]()
		for range 1000 {
			heap.Push(1366)
		}
	}

	for b.Loop() {
		add1000Elems()
	}
}

func TestMaxHeap(t *testing.T) {
	h := heap.NewMaxHeap[int]()
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, num := range nums {
		h.Push(num)
	}
	for i := range nums {
		elem, err := h.Pop()
		assert.NoErrorf(t, err, "Unexpected error %v on index %v", err, i)
		assert.Equalf(t, nums[len(nums)-1-i], elem, "Wrong number, expected %v got %v", nums[len(nums)-1-i], elem)
	}
}

func TestMinHeapOnRandom(t *testing.T) {
	nums := make([]int, 100)
	for i := range 100 {
		nums[i] = rand.Int() - 1_000_000_000
	}
	h := heap.MinHeapFromSlice(nums)
	slices.Sort(nums)
	for i := range 100 {
		num, err := h.Pop()
		assert.NoError(t, err, "Unexpected error %v on index %v", err, i)
		assert.Equalf(t, nums[i], num, "Wrong number, expected %v got %v", nums[i], num)
	}
}

func TestMaxHeapOnRandom(t *testing.T) {
	nums := make([]int, 100)
	for i := range 100 {
		nums[i] = rand.Int() - 1_000_000_000
	}
	h := heap.NewMaxHeap[int]()
	for _, num := range nums {
		h.Push(num)
	}
	slices.Sort(nums)
	for i := range 100 {
		num, err := h.Pop()
		assert.NoErrorf(t, err, "Unexpected error %v on index %v", err, i)
		assert.Equalf(t, nums[99-i], num, "Wrong number, expected %v got %v", nums[i], num)
	}
}

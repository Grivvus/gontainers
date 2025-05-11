package heap

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap[int]()
	_, err := h.Pop()
	if err == nil {
		t.Errorf("Expected error while poping from empty heap")
	}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for _, num := range nums {
		h.Push(num)
	}
	for i := range nums {
		elem, err := h.Pop()
		if err != nil {
			t.Errorf("Unexpected error %v on index %v", err.Error(), i)
		}
		if elem != nums[len(nums)-1-i] {
			t.Errorf("Wrong number, expected %v got %v", nums[len(nums)-1-i], elem)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap[int]()
	_, err := h.Pop()
	if err == nil {
		t.Errorf("Expected error while poping from empty heap")
	}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, num := range nums {
		h.Push(num)
	}
	for i := range nums {
		elem, err := h.Pop()
		if err != nil {
			t.Errorf("Unexpected error %v on index %v", err.Error(), i)
		}
		if elem != nums[len(nums)-1-i] {
			t.Errorf("Wrong number, expected %v got %v", nums[len(nums)-1-i], elem)
		}
	}
}

func TestMinHeapOnRandom(t *testing.T) {
	nums := make([]int, 100)
	for i := range 100 {
		nums[i] = rand.Int() - 1_000_000_000
	}
	h := FromSliceMinHeap(nums)
	slices.Sort(nums)
	for i := range 100 {
		num, err := h.Pop()
		if err != nil {
			t.Errorf("Unexpected error %v on index %v", err.Error(), i)
		}
		if num != nums[i] {
			t.Errorf("Wrong number, expected %v got %v", nums[i], num)
		}
	}
}

func TestMaxHeapOnRandom(t *testing.T) {
	nums := make([]int, 100)
	for i := range 100 {
		nums[i] = rand.Int() - 1_000_000_000
	}
	h := NewMaxHeap[int]()
	for _, num := range nums {
		h.Push(num)
	}
	slices.Sort(nums)
	for i := range 100 {
		num, err := h.Pop()
		if err != nil {
			t.Errorf("Unexpected error %v on index %v", err.Error(), i)
		}
		if num != nums[99-i] {
			t.Errorf("Wrong number, expected %v got %v", nums[i], num)
		}
	}
}

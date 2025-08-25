package list

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddFirst(t *testing.T) {
	l := New[int]()
	l.AddFirst(12)
	l.AddFirst(13)
	if l.Len() != 2 {
		t.Errorf("Unexpeted list Len, expeted %v, got %v", 2, l.Len())
	}
	fmt.Println("should be 12, 13")
	for elem := range l.ElementsIter() {
		fmt.Println(elem)
	}
}

func TestAddLast(t *testing.T) {
	l := New[int]()
	l.AddLast(3)
	if l.Len() != 1 && l.tail.value != 3 {
		t.Error("Error in AddLast operation in empty list")
	}
	l.AddLast(14)
	if l.Len() != 2 && l.tail.value != 14 && l.head.value != 3 {
		t.Error("Error in AddLast operation in existing list")
	}
}

func TestAddLastAndAddFirst(t *testing.T) {
	l := New[int]()
	l.AddLast(3)
	if l.Len() != 1 && l.tail.value == 3 {
		t.Errorf("Error in AddLast operation in empty list")
	}
	l.AddFirst(13)
	if l.Len() != 2 && l.head.value != 13 && l.tail.value != 3 && l.head.next != l.tail {
		t.Errorf("Error id AddFirst after AddLast")
	}
}

func TestAddFirstAndAddLast(t *testing.T) {
	l := New[int]()
	l.AddFirst(3)
	l.AddLast(14)
	if l.Len() != 2 && l.tail.value != 14 {
		t.Errorf("Error in AddLast operation, value of last element expected %d, got %d", 14, l.tail.value)
	}
	l.AddLast(15)
	if l.Len() != 3 && l.tail.value != 15 {
		t.Errorf("Error in AddLast operation, value of last element expected %d, got %d", 15, l.tail.value)
	}

	fmt.Println("should be 3, 14, 15")
	for elem := range l.ElementsIter() {
		fmt.Println(elem)
	}
}

func TestPopFirst(t *testing.T) {
	l := New[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	l.AddLast(5)
	for i := range 5 {
		elem, err := l.PopFirst()
		if elem != i+1 && err != nil {
			t.Errorf("Error while Poping from head, expected %d, got %d", i+1, elem)
		}
	}
	if l.Len() != 0 {
		t.Errorf("Error in PopFirst operation, wrong final length epxected 0, got %d", l.Len())
	}
}

func TestPopLast(t *testing.T) {
	l := New[int]()
	l.AddFirst(1)
	l.AddFirst(2)
	l.AddFirst(3)
	l.AddFirst(4)
	l.AddFirst(5)
	l.AddFirst(6)
	for i := range 6 {
		elem, err := l.PopLast()
		if elem != i+1 && err != nil {
			t.Errorf("Error while Poping from tail, expected %d, got %d", i+1, elem)
		}
	}
	if l.Len() != 0 {
		t.Errorf("Error in PopLast operation, wrong final length epxected 0, got %d", l.Len())
	}
}

func TestRemove(t *testing.T) {
	l := New[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	l.AddLast(5)
	for i := range 5 {
		if l.Remove(i+1) != nil && l.Len() != 5-i-1 {
			t.Errorf("Error in Remove operation")
		}
	}
	if l.Remove(12) == nil {
		t.Errorf("Error in Remove operation; trying remove from empty list")
	}
}

func BenchmarkIterateList(b *testing.B) {
	l := New[int]()
	for range 1000 {
		l.AddLast(1248)
	}

	for b.Loop() {
		for range l.ElementsIter() {
		}
	}
}

func BenchmarkAddStrLast(b *testing.B) {
	for b.Loop() {
		l := New[string]()
		for i := range 10000 {
			l.AddLast("iter" + strconv.Itoa(i))
		}
	}
}

func BenchmarkAddLast(b *testing.B) {
	for b.Loop() {
		l := New[int]()
		for i := range 10000 {
			l.AddLast(i)
		}
	}
}

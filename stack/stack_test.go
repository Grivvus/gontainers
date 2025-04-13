package stack

import "testing"

func TestIntStack(t *testing.T) {
	s := New[int]()
	s.Push(12)
	s.Push(13)
	s.Pop()
	s.Push(14)
}

func TestPopFromEmptyStack(t *testing.T) {


	s := New[int]()
	_, err := s.Pop()
	if err == nil || err.Error() != "Cannot Pop from empty stack" {
		t.Errorf("wrong result on pop from empty stack")
	}
}

func TestPushes(t *testing.T) {
	s := New[int]()
	for i := range 1000 {
		s.Push(i)
	}
	if s.Len() != len(s.data) {
		t.Error("Unexpected len of stack")
	}
}

func TestElementOrder(t *testing.T) {
	s := New[int]()
	for i := range 1000 {
		s.Push(i)
	}

	for i := 999; i >= 0; i-- {
		elem, err := s.Pop()
		if elem != i && err != nil {
			t.Errorf(`Unexpected order elements for Pop operation; expected %v, got %v`, i, elem)
		}
	}

	if s.Len() != 0 {
		t.Errorf(`Unexpected length of stack, expected %v, got %v`, 0, s.Len())
	}
}

func TestGetLastMethod(t *testing.T) {
	s := New[int]()
	s.Push(12)
	if s.GetLast() != 12 {
		t.Errorf(`Unexpected last element, expected %v, got %v`, 12, s.GetLast())
	}
	s.Push(13)
	if s.GetLast() != 13 {
		t.Errorf(`Unexpected last element, expected %v, got %v`, 13, s.GetLast())
	}
	s.Pop()
	s.Push(15)
	if s.GetLast() != 15 {
		t.Errorf(`Unexpected last element, expected %v, got %v`, 15, s.GetLast())
	}
}

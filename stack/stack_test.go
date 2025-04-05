package stack

import "testing"

func TestIntStack(t *testing.T) {
	s := New[int]()
	s.Push(12)
	s.Push(13)
	s.Pop()
	s.Push(14)
}

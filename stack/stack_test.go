package stack_test

import (
	"testing"

	"github.com/Grivvus/gontainers/stack"
	"github.com/stretchr/testify/assert"
)

func TestIntStack(t *testing.T) {
	s := stack.New[int]()
	s.Push(12)
	s.Push(13)
	elem, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 13, elem)
	s.Push(14)
}

func TestPopFromEmptyStack(t *testing.T) {

	s := stack.New[int]()
	_, err := s.Pop()
	assert.ErrorIs(t, err, stack.ErrPopedFromEmpty)
}

func TestElementOrder(t *testing.T) {
	s := stack.New[int]()
	for i := range 1000 {
		s.Push(i)
	}

	for i := 999; i >= 0; i-- {
		elem, err := s.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, elem)
	}

	assert.Equal(t, 0, s.Len())
}

func TestGetLastMethod(t *testing.T) {
	s := stack.New[int]()
	s.Push(12)
	assert.Equal(t, 12, s.GetLast())
	s.Push(13)
	assert.Equal(t, 13, s.GetLast())
	_, err := s.Pop()
	assert.NoError(t, err)
	s.Push(15)
	assert.Equal(t, 15, s.GetLast())
}

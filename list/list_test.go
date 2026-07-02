package list_test

import (
	"strconv"
	"testing"

	"github.com/Grivvus/gontainers/list"
	"github.com/stretchr/testify/assert"
)

func TestAddFirst(t *testing.T) {
	l := list.New[int]()
	l.AddFirst(12)
	l.AddFirst(13)
	assert.Equal(t, 2, l.Len())
	first, err := l.GetFirst()
	assert.NoError(t, err)
	last, err := l.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 13, first)
	assert.Equal(t, 12, last)
}

func TestAddLast(t *testing.T) {
	l := list.New[int]()
	l.AddLast(3)
	last, err := l.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, 3, last)
	l.AddLast(14)
	first, err := l.GetFirst()
	assert.NoError(t, err)
	last, err = l.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, 14, last)
	assert.Equal(t, 3, first)
}

func TestAddLastAndAddFirst(t *testing.T) {
	l := list.New[int]()
	l.AddLast(3)
	last, err := l.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 1, l.Len())
	assert.Equal(t, 3, last)
	l.AddFirst(13)
	first, err := l.GetFirst()
	assert.NoError(t, err)
	last, err = l.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, 13, first)
	assert.Equal(t, 3, last)
}

func TestAddFirstAndAddLast(t *testing.T) {
	l := list.New[int]()
	l.AddFirst(3)
	l.AddLast(14)
	last, err := l.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, 14, last)
	l.AddLast(15)
	last, err = l.GetLast()
	assert.NoError(t, err)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, 15, last)
}

func TestPopFirst(t *testing.T) {
	l := list.New[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	l.AddLast(5)
	for i := range 5 {
		elem, err := l.PopFirst()
		assert.NoError(t, err)
		assert.Equal(t, i+1, elem)
	}
	assert.Equal(t, 0, l.Len())
}

func TestPopLast(t *testing.T) {
	l := list.New[int]()
	const cnt = 6
	for i := 1; i <= cnt; i++ {
		l.AddFirst(i)
	}
	for i := range cnt {
		elem, err := l.PopLast()
		assert.NoError(t, err)
		assert.Equal(t, i+1, elem)
	}
	assert.Equal(t, 0, l.Len())
}

func TestRemove(t *testing.T) {
	l := list.New[int]()
	const cnt = 5
	for i := 1; i <= cnt; i++ {
		l.AddLast(i)
	}
	for i := range cnt {
		err := l.Remove(i + 1)
		assert.NoError(t, err)
		assert.Equal(t, cnt-i-1, l.Len())
	}
	err := l.Remove(12)
	assert.ErrorIs(t, err, list.ErrNoSuchElement)
}

func BenchmarkIterateList(b *testing.B) {
	l := list.New[int]()
	for range 1000 {
		l.AddLast(1248)
	}

	for b.Loop() {
		var elem int
		for iterElem := range l.ElementsIter() {
			elem = iterElem
		}
		_ = elem
	}
}

func BenchmarkAddStrLast(b *testing.B) {
	for b.Loop() {
		l := list.New[string]()
		for i := range 10000 {
			l.AddLast("iter" + strconv.Itoa(i))
		}
	}
}

func BenchmarkAddLast(b *testing.B) {
	for b.Loop() {
		l := list.New[int]()
		for i := range 10000 {
			l.AddLast(i)
		}
	}
}

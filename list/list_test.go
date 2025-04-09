package list

import (
	"fmt"
	"testing"
)

func TestAddFirst(t *testing.T) {
	l := New[int]()
	l.AddFirst(12)
	l.AddFirst(13)
	if l.Len() != 2 {
		t.Errorf(`Unexpeted list Len, expeted %v, got %v`, 2, l.Len())
	}
	for elem := range l.Elements() {
		fmt.Println(elem)
	}
}

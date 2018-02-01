package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i < 3; i++ {
		s.Push(i)
		t.Log(s)
	}
	for i := 0; i < 3; i++ {
		item := s.Pop()
		t.Log(item, s)
	}
}

package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i < 3; i++ {
		s.Push(i)
		fmt.Println(s)
	}
	for i := 0; i < 3; i++ {
		item := s.Pop()
		fmt.Println(item, s)
	}
}

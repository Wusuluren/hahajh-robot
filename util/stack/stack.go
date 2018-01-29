package stack

type Stack struct {
	stack []interface{}
	size  int
}

func (s *Stack) Push(item interface{}) {
	s.stack = append(s.stack, item)
	s.size += 1
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	item := s.stack[s.size-1]
	s.size -= 1
	return item
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]interface{}, 0, 8),
		size:  0,
	}
}

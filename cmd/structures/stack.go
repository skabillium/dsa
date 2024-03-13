package structures

type StackNode struct {
	value any
	prev  *StackNode
}

type Stack struct {
	Length int
	top    *StackNode
}

func NewStack() *Stack {
	return &Stack{Length: 0, top: nil}
}

func (s *Stack) Push(item any) {
	node := &StackNode{value: item, prev: nil}

	s.Length++
	if s.top == nil {
		s.top = node
		return
	}

	node.prev = s.top
	s.top = node
}

func (s *Stack) Pop() any {
	if s.Length == 0 {
		return nil
	}

	value := s.top.value
	s.top = s.top.prev
	s.Length--

	return value
}

func (s *Stack) Peek() any {
	if s.top == nil {
		return nil
	}

	return s.top.value
}

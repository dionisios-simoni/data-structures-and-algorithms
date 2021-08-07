package stack_array

import "log"

// Stack is a data structure that is implemented on top of an int slice.
type Stack struct {
	Items []int
	Len   int
}

// Pop removes the item on the top of the Stack
func (s *Stack) Pop() int {
	removed := s.peek()
	top := len(s.Items) - 1
	s.Items = s.Items[:top]
	s.Len--
	return removed
}

// Push adds an element on the top of the Stack
func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
	s.Len++
}

// peek returns the current item at the top of the Stack
func (s *Stack) peek() int {
	if len(s.Items) == 0 {
		log.Fatal("Stack is empty")
	}

	top := len(s.Items) - 1
	return s.Items[top]
}

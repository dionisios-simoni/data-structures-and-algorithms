package stack_array

import "log"

// stack is a data structure that is implemented on top of an int slice.
type stack struct {
	items []int
}

// pop removes the item on the top of the stack
func (s *stack) pop() {
	top := len(s.items) - 1
	s.items = s.items[:top]
}

// push adds an element on the top of the stack
func (s *stack) push(item int) {
	s.items = append(s.items, item)
}

// peek returns the current item at the top of the stack
func (s *stack) peek() int {
	if  len(s.items) == 0 {
		log.Fatal("stack is empty")
	}

	top := len(s.items) - 1
	return s.items[top]
}

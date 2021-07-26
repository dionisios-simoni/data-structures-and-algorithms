package stack_linked_list

import (
	"log"
)

// stack is a data structure that is implemented using linked list.
type stack struct {
	top, bottom *node
	length      int
}

// node represents a linked list entry
type node struct {
	value int
	next  *node
}

// pop removes the item on the top of the stack
func (s *stack) pop() {

	if s.isEmpty() {
		log.Fatal("stack is empty")
	}

	if s.top == s.bottom {
		s.bottom = nil
	}

	s.top = s.top.next
	s.length--
}

// push adds an element on the top of the stack
func (s *stack) push(v int) {
	n := &node{value: v}

	if s.length == 0 {
		s.top = n
		s.bottom = n
	} else {
		n.next = s.top
		s.top = n
	}
	s.length++
}

// peek returns the current item at the top of the stack
func (s *stack) peek() int {
	if s.isEmpty() {
		log.Fatal("stack is empty")
	}
	return s.top.value
}

func (s *stack) isEmpty() bool {
	return s.length == 0
}
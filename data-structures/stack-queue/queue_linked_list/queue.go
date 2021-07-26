package stack_linked_list

import (
	"log"
)

// queue is a data structure that is implemented using linked list.
type queue struct {
	start, end *node
	length     int
}

// node represents a linked list entry
type node struct {
	value int
	next  *node
}

// dequeue removes the first item in the queue
func (s *queue) dequeue() {

	if s.isEmpty() {
		log.Fatal("queue is empty")
	}

	if s.start == s.end {
		s.end = nil
	}

	s.start = s.start.next
	s.length--
}

// enqueue adds an element at the end of the queue
func (s *queue) enqueue(v int) {
	n := &node{value: v}

	if s.length == 0 {
		s.start = n
		s.end = n
	} else {
		s.end.next = n
		s.end = n
	}
	s.length++
}

// peek returns the current item at the start of the queue
func (s *queue) peek() int {
	if s.isEmpty() {
		log.Fatal("queue is empty")
	}
	return s.start.value
}

func (s *queue) isEmpty() bool {
	return s.length == 0
}

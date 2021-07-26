package stack_array

import "log"

// queue is a data structure that is implemented on top of an int slice.
type queue struct {
	items []int
}

// dequeue removes the first item on the queue
func (s *queue) dequeue() {
	s.items = s.items[1:]
}

// enqueue adds an element at the end of the queue
func (s *queue) enqueue(item int) {
	s.items = append(s.items, item)
}

// peek returns the current item at the start of the queue
func (s *queue) peek() int {
	if len(s.items) == 0 {
		log.Fatal("stack is empty")
	}
	return s.items[0]
}

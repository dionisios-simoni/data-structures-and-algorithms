package queue_array

import "log"

// Queue is a data structure that is implemented on top of an int slice.
type Queue struct {
	Items []int
	Len   int
}

// Dequeue removes the first item on the Queue
func (q *Queue) Dequeue() int {
	removed := q.Peek()
	q.Items = q.Items[1:]
	q.Len--
	return removed
}

// Enqueue adds an element at the end of the Queue
func (q *Queue) Enqueue(item int) {
	q.Items = append(q.Items, item)
	q.Len++
}

// Peek returns the current item at the start of the Queue
func (q *Queue) Peek() int {
	if len(q.Items) == 0 {
		log.Fatal("stack is empty")
	}
	return q.Items[0]
}

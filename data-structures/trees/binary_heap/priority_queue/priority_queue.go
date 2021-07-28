package priority_queue

import (
	"data-structures-and-algorithms/data-structures/trees/binary_heap/max_heap"
)

// PriorityQueue defines a priority queue where the element of highest
// value holds the priority.
type PriorityQueue struct {
	heap max_heap.MaxHeap
}

// Enqueue adds a value in the queue
func (pq *PriorityQueue) Enqueue(value int) {
	pq.heap.Insert(value)
}

// Dequeue removes and returns the element with the highest priority
func (pq *PriorityQueue) Dequeue() int {
	return pq.heap.ExtractMax()
}

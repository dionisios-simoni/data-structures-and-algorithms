package min_heap

import (
	"data-structures-and-algorithms/data-structures/trees/binary_heap/common"
	"log"
)

// minHeap represents a binary heap data structure
type minHeap struct {
	values []int
}

// insert adds an element on the binary heap
func (mh *minHeap) insert(value int) {
	mh.values = append(mh.values, value)
	mh.minHeapifyUp(len(mh.values) - 1)
}

// extractMin removes and returns the element with the lowest value / highest priority
func (mh *minHeap) extractMin() int {

	if len(mh.values) == 0 {
		log.Fatal("binary heap is currently empty")
	}

	extracted := mh.values[0]

	// replace last with first
	mh.values[0] = mh.values[len(mh.values)-1]

	// shrink array
	mh.values = mh.values[:len(mh.values)-1]

	// preserve correctness in order
	mh.minHeapifyDown(0)

	return extracted
}

// minHeapifyUp compares the added element with its parent; if they are not
// in the correct order then a swap of elements occurs until the correct
// order was preserved.
func (mh *minHeap) minHeapifyUp(index int) {
	for mh.values[index] < mh.values[common.Parent(index)] {
		mh.swap(common.Parent(index), index)
		index = common.Parent(index)
	}
}

// minHeapifyDown compares the parent element with its children; if they are not
// in the correct order then a swap of elements occurs until the correct
// order was preserved.
func (mh *minHeap) minHeapifyDown(index int) {
	lastIndex := len(mh.values) - 1
	l, r := common.Left(index), common.Right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || mh.values[l] < mh.values[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if mh.values[index] > childToCompare {
			mh.swap(index, childToCompare)
			index = childToCompare
			l, r = common.Left(index), common.Right(index)
		} else {
			return
		}

	}

}

func (mh *minHeap) swap(i1, i2 int) {
	mh.values[i1], mh.values[i2] = mh.values[i2], mh.values[i1]
}

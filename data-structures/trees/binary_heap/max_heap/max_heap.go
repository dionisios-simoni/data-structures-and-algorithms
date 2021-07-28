package max_heap

import (
	"data-structures-and-algorithms/data-structures/trees/binary_heap/common"
	"log"
)

// MaxHeap represents a binary heap data structure
type MaxHeap struct {
	Values []int
}

// Insert adds an element on the binary heap
func (mh *MaxHeap) Insert(value int) {
	mh.Values = append(mh.Values, value)
	mh.maxHeapifyUp(len(mh.Values) - 1)
}

// ExtractMax removes and returns the element with the highest value / priority
func (mh *MaxHeap) ExtractMax() int {

	if len(mh.Values) == 0 {
		log.Fatal("binary heap is currently empty")
	}

	extracted := mh.Values[0]

	// replace last with first
	mh.Values[0] = mh.Values[len(mh.Values)-1]

	// shrink array
	mh.Values = mh.Values[:len(mh.Values)-1]

	// preserve correctness in order
	mh.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp compares the added element with its parent; if they are not
// in the correct order then a swap of elements occurs until the correct
// order was preserved.
func (mh *MaxHeap) maxHeapifyUp(index int) {
	for mh.Values[common.Parent(index)] < mh.Values[index] {
		mh.swap(common.Parent(index), index)
		index = common.Parent(index)
	}
}

// maxHeapifyDown compares the parent element with its children; if they are not
// in the correct order then a swap of elements occurs until the correct
// order was preserved.
func (mh *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(mh.Values) - 1
	l, r := common.Left(index), common.Right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex || mh.Values[l] > mh.Values[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if mh.Values[index] < mh.Values[childToCompare] {
			mh.swap(index, childToCompare)
			index = childToCompare
			l, r = common.Left(index), common.Right(index)
		} else {
			return
		}
	}
}

func (mh *MaxHeap) swap(i1, i2 int) {
	mh.Values[i1], mh.Values[i2] = mh.Values[i2], mh.Values[i1]
}

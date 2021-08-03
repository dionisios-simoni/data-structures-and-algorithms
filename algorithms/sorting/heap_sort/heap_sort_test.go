package heap_sort

import (
	"data-structures-and-algorithms/data-structures/trees/binary_heap/max_heap"
	"testing"
)

func TestHeapSort(t *testing.T) {

	mh := &max_heap.MaxHeap{}
	mh.Insert(22)
	mh.Insert(17)
	mh.Insert(19)
	mh.Insert(12)
	mh.Insert(15)
	mh.Insert(11)
	mh.Insert(7)
	mh.Insert(6)
	mh.Insert(9)
	mh.Insert(10)
	mh.Insert(5)

	var want = []int{5, 6, 7, 9, 10, 11, 12, 15, 17, 19, 22}

	got := heapSort(mh)
	for i := range mh.Values {
		if got[i] != want[i] {
			t.Errorf("invalid sort, expected: %v got: %v", want, got)
		}
	}
}

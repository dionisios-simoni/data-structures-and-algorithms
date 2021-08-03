package heap_sort

import (
	"data-structures-and-algorithms/data-structures/trees/binary_heap/max_heap"
)

func heapSort(mh *max_heap.MaxHeap) []int {

	result := make([]int, len(mh.Values))
	for i := len(mh.Values) - 1; i >= 0; i-- {
		result[i] = mh.ExtractMax()
	}
	return result
}

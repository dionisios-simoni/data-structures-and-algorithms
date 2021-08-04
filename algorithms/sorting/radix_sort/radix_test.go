package radix_sort

import "testing"

var expected = []int{273, 383, 465, 722, 1465}

func TestRadixSortSort(t *testing.T) {
	input := []int{465, 273, 1465, 722, 383}
	got := radixSort(input, 1)
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("invalid sort, expected: %v got: %v", expected, got)
		}
	}
}

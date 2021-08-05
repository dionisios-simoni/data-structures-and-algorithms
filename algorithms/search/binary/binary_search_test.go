package binary

import "testing"

func TestBubbleSort(t *testing.T) {

	got := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)
	if got != true {
		t.Errorf("invalid search result, expected: %t got: %t", true, got)
	}

	got = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1)
	if got != true {
		t.Errorf("invalid search result, expected: %t got: %t", true, got)
	}

	got = binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11)
	if got != false {
		t.Errorf("invalid search result, expected: %t got: %t", false, got)
	}

}

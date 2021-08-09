package min_heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	mh := &minHeap{}

	t.Run("insert into min binary heap", func(t *testing.T) {
		for i := 100; i >= 10; i -= 10 {
			mh.insert(i)
		}

		want := []int{10, 20, 50, 40, 30, 90, 60, 100, 70, 80}
		got := mh.values

		for i := 0; i < len(want); i++ {
			if want[i] != got[i] {
				t.Errorf("order of binary heap elements is invalid got %v but want %v", got, want)
			}
		}

	})

	t.Run("extracting min value should maintain valid element order", func(t *testing.T) {

		got := mh.extractMin()
		want := 10
		if got != want {
			t.Errorf("invalid value extracted, want: %d but got %d", want, got)
		}

		gotValues := mh.values
		wantValues := []int{20, 30, 50, 40, 80, 90, 60, 100, 70}

		for i := 0; i < len(wantValues); i++ {
			if gotValues[i] != wantValues[i] {
				t.Errorf("order of binary heap elements is invalid got %v but want %v", gotValues, wantValues)
			}
		}

	})
}

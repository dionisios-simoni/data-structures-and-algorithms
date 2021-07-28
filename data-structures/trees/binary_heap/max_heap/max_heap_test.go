package max_heap

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	mh := &MaxHeap{}

	t.Run("Insert into max binary heap", func(t *testing.T) {
		for i := 10; i <= 100; i += 10 {
			mh.Insert(i)
		}

		want := []int{100, 90, 60, 70, 80, 20, 50, 10, 40, 30}
		got := mh.Values

		for i := 0; i < len(want); i++ {
			if want[i] != got[i] {
				t.Errorf("order of binary heap elements is invalid got %v but want %v", got, want)
			}
		}

	})

	t.Run("extracting max value should maintain valid element order", func(t *testing.T) {

		got := mh.ExtractMax()
		want := 100
		if got != want {
			t.Errorf("invalid value extracted, want: %d but got %d", want, got)
		}

		gotValues := mh.Values
		wantValues := []int{90, 80, 60, 70, 30, 20, 50, 10, 40}

		for i := 0; i < len(wantValues); i++ {
			if gotValues[i] != wantValues[i] {
				t.Errorf("order of binary heap elements is invalid got %v but want %v", gotValues, wantValues)
			}
		}

	})
}

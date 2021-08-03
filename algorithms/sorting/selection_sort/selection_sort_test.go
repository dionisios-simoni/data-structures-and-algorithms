package selection_sort

import (
	"testing"
)

var expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var cases = []struct {
	description string
	input       []int
	want        []int
}{
	{"worst", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected},
	{"average", []int{8, 4, 7, 2, 10, 5, 9, 1, 6, 3}, expected},
	{"best", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected},
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range cases {
		got := selectionSort(tt.input)
		for i := range tt.want {
			if got[i] != expected[i] {
				t.Errorf("invalid sort, expected: %v got: %v", tt.want, got)
			}
		}
	}
}

func BenchmarkBubbleWorst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectionSort(cases[0].input)
	}
}

func BenchmarkBubbleAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectionSort(cases[1].input)
	}
}

func BenchmarkBubbleBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectionSort(cases[2].input)
	}
}

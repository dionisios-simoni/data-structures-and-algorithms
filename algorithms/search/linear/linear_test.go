package linear

import "testing"

var cases = []struct {
	input []int
	value int
	want  bool
}{
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 4, true},
	{[]int{8, 4, 7, 2, 10, 5, 9, 1, 6, 3}, 12, false},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range cases {
		got := linearSearch(tt.input, tt.value)
		if got != tt.want {
			t.Errorf("invalid search result, expected: %v got: %v", tt.want, got)
		}
	}
}

package shell_sort

import (
	"fmt"
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

func TestShellSort(t *testing.T) {
	for _, tt := range cases {
		got := shellSort(tt.input)
		for i := range tt.want {
			if got[i] != expected[i] {
				t.Errorf("invalid sort, expected: %v got: %v", tt.want, got)
			}
		}
	}
}

func TestAnother(t *testing.T) {
	fmt.Println(shellSort([]int{10, 9, 8, 7, 6, 5, 4, 3}))
}

package shell_sort

import "data-structures-and-algorithms/algorithms/sorting/insertion_sort"

// Shell sort algorithm time complexity:
// best case is O(n)
// average case n*log(n)^2 or n^(3/2)
// worst case (depends on gap sequence)

// space complexity is O(1)

func shellSort(x []int) []int {
	gap := len(x) / 3
	distance := gap

	for gap > 1 {
		i := 0
		for distance < len(x) {
			if x[i] > x[distance] {
				x[i], x[distance] = x[distance], x[i]
			}
			i++
			distance = i + gap
		}
		gap--
		distance = gap
	}
	return insertion_sort.InsertionSort(x)
}

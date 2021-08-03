package selection_sort

// Selection sort algorithm time complexity:
// best case is Ω(n^2) - same or slower than O(n^2)
// average case is Θ(n^2) - the same as O(n^2)
// worst case is O(n^2) - the same or faster than O(n^2)

// space complexity is O(1)

func selectionSort(x []int) []int {
	for i := range x {
		smallest := i
		for j := i + 1; j < len(x); j++ {
			if x[j] < x[smallest] {
				smallest = j
			}
		}
		x[i], x[smallest] = x[smallest], x[i]
	}
	return x
}

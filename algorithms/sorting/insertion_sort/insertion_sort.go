package insertion_sort

// Insertion sort algorithm time complexity:
// best case is Ω(n) - same or slower than O(n)
// average case is Θ(n^2) - the same as O(n^2)
// worst case is O(n^2) - the same or faster than O(n^2)

// space complexity is O(1)

func InsertionSort(x []int) []int {
	for i := 1; i < len(x); i++ {
		j := i
		for j > 0 && x[j-1] > x[j] {
			x[j-1], x[j] = x[j], x[j-1]
			j--
		}
	}
	return x
}

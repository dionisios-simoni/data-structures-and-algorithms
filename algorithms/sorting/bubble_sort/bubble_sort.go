package bubble_sort

// Bubble sort algorithm time complexity:
// best case is Ω(n) - same or slower than O(n)
// average case is Θ(n^2) - the same as O(n^2)
// worst case is O(n^2) - the same or faster than O(n^2)

// space complexity is O(1)

func bubbleSort(x []int) []int {

	// is set to true if a swap occurs
	flag := true

	for flag {

		flag = false

		for i := 0; i < len(x)-1; i++ {
			if x[i] > x[i+1] {
				x[i], x[i+1] = x[i+1], x[i]
				flag = true
			}
		}
	}
	return x
}

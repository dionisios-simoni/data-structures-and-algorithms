package quick_sort

// Merge sort algorithm time complexity:
// best case is Ω(n log(n)) - same or slower than O(n log(n))
// average case is Θ(n log(n)) - the same as O(n log(n))
// worst case is O(n^2) - the same or faster than O(n^2)

// space complexity is O(log(n))

func quickSort(x []int) []int {

	if len(x) <= 1 {
		return x
	}

	pivot := len(x) / 2

	x[pivot], x[len(x)-1] = x[len(x)-1], x[pivot]
	pivot = len(x) - 1

	i, j := 0, 0

	for i < len(x)-1 {
		if x[i] > x[pivot] {
			i++
		} else {
			x[i], x[j] = x[j], x[i]
			j++
			i++
		}
	}

	x[pivot], x[j] = x[j], x[pivot]
	pivot = j

	quickSort(x[:pivot])
	quickSort(x[pivot+1:])

	return x
}

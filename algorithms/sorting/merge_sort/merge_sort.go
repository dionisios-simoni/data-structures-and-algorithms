package merge_sort

// Merge sort algorithm time complexity:
// best case is Ω(n log(n)) - same or slower than O(n log(n))
// average case is Θ(n log(n)) - the same as O(n log(n))
// worst case is O(n log(n)) - the same or faster than O(n log(n))

// space complexity is O(n)

func mergeSort(x []int) []int {

	length := len(x)

	if length == 1 {
		return x
	}

	left := x[:length/2]
	right := x[length/2:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(a []int, b []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// handle elements that are yet to be appended
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}

	for ; j < len(b); j++ {
		result = append(result, b[j])
	}

	return result
}

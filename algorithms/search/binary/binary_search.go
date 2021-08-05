package binary

func binarySearch(x []int, v int) bool {

	mid := len(x) / 2

	if x[mid] == v {
		return true
	}

	if len(x) == 1 {
		return false
	}

	if x[mid] < v {
		return binarySearch(x[mid:], v)
	} else {
		return binarySearch(x[:mid], v)
	}
}

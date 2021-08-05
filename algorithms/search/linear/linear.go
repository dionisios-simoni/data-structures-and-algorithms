package linear

func linearSearch(x []int, v int) bool {
	for i := range x {
		if x[i] == v {
			return true
		}
	}
	return false
}

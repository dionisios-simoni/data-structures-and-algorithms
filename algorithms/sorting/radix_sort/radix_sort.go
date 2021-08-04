package radix_sort

import (
	"math"
)

func radixSort(x []int, place int) []int {

	if place > 4 {
		return x
	}

	result := []int{}
	bin := make([][]int, 9)

	for i := range x {
		idx := digit(x[i], place)
		bin[idx] = append(bin[idx], x[i])
	}

	for j := 0; j < 9; j++ {
		if bin[j] != nil {
			result = append(result, bin[j]...)
		}
	}

	return radixSort(result, place+1)
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

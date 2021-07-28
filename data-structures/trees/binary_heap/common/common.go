package common

// Parent returns the parent index of a given child index
func Parent(i int) int {
	return (i - 1) / 2
}

// Left returns the left child index
func Left(i int) int {
	return 2*i + 1
}

// Right returns the left child index
func Right(i int) int {
	return 2*i + 2
}

package recursion

import "fmt"

// Direct recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// iterative approach
func factorialLoop(n int) int {
	var res = 1
	for n >= 1 {
		res *= n
		n--
	}
	return res
}

// ========================================================================

// Direct recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// iterative approach
func fibonacciLoop(n int) int {

	if n < 2 {
		return n
	}

	sequence := make([]int, n+1)
	sequence[0] = 0
	sequence[1] = 1

	for i := 2; i <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence[n]
}

// ========================================================================

//Direct recursion
func reverse(str string) string {
	if len(str) <= 1 {
		return str
	}
	return reverse(str[1:]) + string(str[0])
}

// iterative approach
func reverseLoop(word string) string {
	r := []rune(word)
	for i, j := 0, len(word)-1; i < len(word)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ========================================================================

// example tail recursive function
func printNum(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)

	// tail recursive call
	printNum(n - 1)
}

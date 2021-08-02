package recursion

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	for _, tt := range casesFactorial {
		if got := factorial(tt.input); got != tt.want {
			t.Errorf("unexpected result, want: %d but got:%d", tt.want, got)
		}
	}
}

func TestFactorialLoop(t *testing.T) {
	for _, tt := range casesFactorial {
		if got := factorialLoop(tt.input); got != tt.want {
			t.Errorf("unexpected result, want: %d but got:%d", tt.want, got)
		}
	}
}

func TestFibonacci(t *testing.T) {
	for _, tt := range casesFibonacci {
		if got := fibonacci(tt.input); got != tt.want {
			t.Errorf("unexpected result, want: %d but got:%d", tt.want, got)
		}
	}
}

func TestFibonacciLoop(t *testing.T) {
	for _, tt := range casesFibonacci {
		if got := fibonacciLoop(tt.input); got != tt.want {
			t.Errorf("unexpected result, want: %d but got:%d", tt.want, got)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, tt := range casesReverse {
		if got := reverse(tt.input); got != tt.want {
			t.Errorf("unexpected result, want: %s but got: %s", tt.want, got)
		}
	}
}

func TestReverseLoop(t *testing.T) {
	for _, tt := range casesReverse {
		if got := reverseLoop(tt.input); got != tt.want {
			t.Errorf("unexpected result, want: %s but got: %s", tt.want, got)
		}
	}
}

func TestTailRecursive(t *testing.T) {
	printNum(10)
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(5)
	}
}

func BenchmarkFactorialLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialLoop(5)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(5)
	}
}

func BenchmarkFibonacciLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciLoop(5)
	}
}

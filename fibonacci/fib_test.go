package fib

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13}
	if observed := Fibonacci(8); !testEq(observed, expected) {
		t.Fatalf("Fibonacci() = %v, want %v", observed, expected)
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

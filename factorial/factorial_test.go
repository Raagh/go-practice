package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	expected := 120
	if observed := Factorial(5); observed != expected {
		t.Fatalf("Fibonacci() = %v, want %v", observed, expected)
	}
}

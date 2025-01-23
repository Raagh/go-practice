package prime

import "math"

// ### Problem Statement
// Write a function that checks if a given number is prime.
//
// ### Explanation
// A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself. In other words, a prime number is a number that is only divisible by 1 and itself.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

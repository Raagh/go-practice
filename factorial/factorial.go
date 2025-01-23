package factorial

// ### Coding Assignment: Factorial Function
//
// **Objective:**
// Write a function in Go that calculates the factorial of a given non-negative integer.
//
// **Instructions:**
// 1. Define a function named `Factorial` that takes a single integer parameter `number` and returns an integer.
// 2. The function should calculate the factorial of the given number. The factorial of a non-negative integer `n` is the product of all positive integers less than or equal to `n`.
// 3. Ensure that the function handles edge cases, such as when the input is `0` or `1`.
//
// **Example:**
// - Input: `5`
// - Output: `120` (since 5! = 5 * 4 * 3 * 2 * 1 = 120)
//
// **Constraints:**
// - The input number will be a non-negative integer.
// - You should not use any built-in functions for calculating the factorial.
//
// **Hint:**
// - Consider using a loop or recursion to solve the problem.

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

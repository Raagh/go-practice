package fib

// ### Problem Statement
// Write a function that generates the first `n` numbers in the Fibonacci sequence.
//
// ### Explanation
// The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones. It starts with 0 and 1. So, the sequence looks like this:
// ```
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
// ```
//
// ### Steps to Solve
// 1. **Start with the first two numbers**: 0 and 1.
// 2. **Generate the next numbers**: Each subsequent number is the sum of the previous two numbers.
// 3. **Repeat until you have `n` numbers**.
//
// ### Example
// Let's say `n = 5`. We want the first 5 numbers in the Fibonacci sequence.
//
// 1. Start with 0 and 1.
// 2. Next number: 0 + 1 = 1.
// 3. Next number: 1 + 1 = 2.
// 4. Next number: 1 + 2 = 3.
// 5. Next number: 2 + 3 = 5.
//
// So, the first 5 numbers in the Fibonacci sequence are: `0, 1, 1, 2, 3`.

func Fibonacci(numbers int) []int {
	fibGen := fib()
	fibs := []int{}
	for i := 0; i < numbers; i++ {
		fibs = append(fibs, fibGen())
	}

	return fibs
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		next := a
		a, b = b, a+b
		return next
	}
}

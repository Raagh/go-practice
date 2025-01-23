package minmax

// ## Exercise: Find the Minimum and Maximum in a List
//
// ### Objective
// Write a function in Go that takes a list of integers and returns the minimum and maximum values in the list.
//
// ### Instructions
//  1. Open the `minmax.go` file.
//  2. Implement the `minmax` function to find the minimum and maximum values in the given list of integers.
//  3. The function signature is:
//     ```go
//     func minmax(list []int) (min, max int)
//     ```
//  4. The function should return the minimum and maximum values as a tuple.
//
// ### Example
// Given the list:
// ```go
// list := []int{3, 5, 7, 2, 8, -1, 4, 10, 12}
// ```
// The function should return:
// ```go
// min = -1
// max = 12
// ```
//
// ### Constraints
// - The list will contain at least one integer.
// - You are not allowed to use any built-in functions for finding the minimum or maximum values.
//
// Good luck!
func minmax(list []int) (min, max int) {
	min = list[0]
	max = list[0]
	for _, v := range list {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return min, max
}

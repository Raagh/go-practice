package chapter2

import "slices"

func SelectionSort(array []int) []int {
	new := make([]int, len(array))

	for i := range new {
		smallestIndex := findSmallestIndex(array)
		smallest := array[smallestIndex]

		new[i] = smallest
		array = slices.Delete(array, smallestIndex, smallestIndex+1)
	}

	return new
}

func findSmallestIndex(array []int) int {
	smallest := array[0]
	smallest_index := 0
	for i, v := range array {
		if v < smallest {
			smallest = v
			smallest_index = i
		}
	}

	return smallest_index
}

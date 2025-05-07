package chapter2

import "slices"

func SelectionSort(array []int) []int {
	new := make([]int, len(array))

	for i := range array {
		smallestIndex := findSmallestIndex(array)
		smallestElement := array[smallestIndex]

		array = slices.Delete(array, smallestIndex, smallestIndex+1)
		new[i] = smallestElement
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

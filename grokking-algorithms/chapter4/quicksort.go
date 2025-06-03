package chapter4

import (
	"math/rand/v2"
	"slices"
)

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivotIndex := rand.IntN(len(arr))
	pivot := arr[pivotIndex]

	withoutResults := slices.Delete(arr, pivotIndex, pivotIndex+1)
	smaller, larger := getSmallerAndLarger(withoutResults, pivot)

	return append(append(QuickSort(smaller), pivot), QuickSort(larger)...)
}

func getSmallerAndLarger(array []int, pivot int) (smaller []int, larger []int) {
	smaller = []int{}
	larger = []int{}

	for _, v := range array {
		if v < pivot {
			smaller = append(smaller, v)
		} else {
			larger = append(larger, v)
		}
	}

	return smaller, larger
}

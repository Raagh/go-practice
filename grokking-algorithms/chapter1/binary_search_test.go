package chapter1

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
	}

	for _, test := range tests {
		result := BinarySearch(test.arr, test.target)

		if result != test.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; want %d", test.arr, test.target, result, test.expected)
		}
	}
}

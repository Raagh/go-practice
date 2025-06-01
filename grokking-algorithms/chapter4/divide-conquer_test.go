package chapter4

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2}, 3},
		{[]int{5, 7}, 12},
		{[]int{-1, 1}, 0},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		result := Sum(tt.input)
		if result != tt.expected {
			t.Errorf("sumFirstTwo(%v) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{10, 20, 30, 40, 50}, 5},
	}

	for _, tt := range tests {
		result := Count(tt.input)
		if result != tt.expected {
			t.Errorf("Count(%v) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 5, 3, 9, 2}, 9},
		{[]int{10, 2, 8, 6}, 10},
		{[]int{7}, 7},
		{[]int{-1, -5, -3}, -1},
	}

	for _, test := range tests {
		result := Max(test.input)
		if result != test.expected {
			t.Errorf("Max(%v) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		array    []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 7, 9}, 3, 1},
		{[]int{1, 3, 5, 7, 9}, 9, 4},
		{[]int{1, 3, 5, 7, 9}, 1, 0},
		{[]int{1, 3, 5, 7, 9}, 5, 2},
		{[]int{1, 3, 5, 7, 9}, -1, -1},
		{[]int{1, 3, 5, 7, 9}, 10, -1},
		{[]int{}, 1, -1},
		{[]int{2}, 2, 0},
		{[]int{2}, 3, -1},
	}

	for _, tt := range tests {
		result := BinarySearch(tt.array, tt.target)
		if result != tt.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; want %d", tt.array, tt.target, result, tt.expected)
		}
	}
}

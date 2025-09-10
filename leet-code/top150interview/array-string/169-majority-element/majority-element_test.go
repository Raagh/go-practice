package majority_element

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Basic case with odd number of elements",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "Case with more elements",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "Edge case with a single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Case with all elements being the same",
			nums:     []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "Case with negative numbers",
			nums:     []int{-1, -1, -1, 2},
			expected: -1,
		},
		{
			name:     "Case with exactly n/2 + 1 occurrences",
			nums:     []int{1, 2, 1, 2, 1},
			expected: 1,
		},
		{
			name:     "Case with larger array",
			nums:     []int{8, 8, 7, 7, 7, 8, 8},
			expected: 8,
		},
		{
			name:     "Case where majority element is distributed",
			nums:     []int{3, 3, 4, 2, 4, 4, 3, 3, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.nums)
			if result != tt.expected {
				t.Errorf("majorityElement(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

// This test will also run with the implementation provided by the LeetCode solution
// The goal is to create enough test cases to verify the implementation works correctly
// without revealing the implementation itself


package jumpgame

import (
	"testing"
)

func TestCanJump_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true, // Jump 1 step from index 0 to 1, then 3 steps to the last index
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false, // You will always arrive at index 3 with value 0. From there, you cannot jump to the last index
		},
		{
			name:     "Simple path",
			nums:     []int{1},
			expected: true, // Can reach end by single steps
		},
		{
			name:     "Exact jumps",
			nums:     []int{2, 0, 0},
			expected: true, // Jump directly to the end from start
		},
		{
			name:     "Zero at start",
			nums:     []int{0, 2, 3},
			expected: false, // Can't move from starting position
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanJump(tt.nums)
			if result != tt.expected {
				t.Errorf("CanJump(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestCanJump_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true, // Already at the last index
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: true, // No jump needed
		},
		{
			name:     "All zeros except last",
			nums:     []int{0, 0, 0, 0},
			expected: false, // Can't move from starting position
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: false, // Can't reach last position
		},
		{
			name:     "Large first jump",
			nums:     []int{5, 0, 0, 0, 0, 0},
			expected: true, // Can jump over all zeros to the end
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanJump(tt.nums)
			if result != tt.expected {
				t.Errorf("CanJump(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestCanJump_ComplexCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Multiple paths",
			nums:     []int{2, 4, 2, 1, 0, 2, 0},
			expected: true, // Multiple possible successful paths
		},
		{
			name:     "Decreasing jump distances",
			nums:     []int{4, 3, 2, 1, 0},
			expected: true, // Each position has just enough jump distance
		},
		{
			name:     "Alternating zeros",
			nums:     []int{1, 0, 1, 0, 1},
			expected: true, // Must make exact jumps over zeros
		},
		{
			name:     "Long array with blockage",
			nums:     []int{3, 2, 1, 0, 4, 5, 6, 7, 8, 9},
			expected: false, // Zero at index 3 prevents reaching the end
		},
		{
			name:     "Need to make small jumps",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: true, // Need to make only single-step jumps
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanJump(tt.nums)
			if result != tt.expected {
				t.Errorf("CanJump(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

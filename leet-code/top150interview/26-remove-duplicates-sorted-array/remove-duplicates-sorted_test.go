package removeDuplicatesSorted

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		result   []int
	}{
		{
			name:     "Example 1: [1,1,2]",
			input:    []int{1, 1, 2},
			expected: 2,
			result:   []int{1, 2},
		},
		{
			name:     "Example 2: [0,0,1,1,1,2,2,3,3,4]",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
			result:   []int{0, 1, 2, 3, 4},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
			result:   []int{},
		},
		{
			name:     "Array with one element",
			input:    []int{1},
			expected: 1,
			result:   []int{1},
		},
		{
			name:     "Array with all duplicates",
			input:    []int{1, 1, 1, 1, 1},
			expected: 1,
			result:   []int{1},
		},
		{
			name:     "Array with no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
			result:   []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the original
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			// Call the function
			got := RemoveDuplicates(inputCopy)

			// Check if the returned count is correct
			if got != tt.expected {
				t.Errorf("RemoveDuplicates() returned %v, expected %v", got, tt.expected)
			}

			// Check if the first 'got' elements of the array are correct
			if !reflect.DeepEqual(inputCopy[:got], tt.result) {
				t.Errorf("After RemoveDuplicates(), first %v elements should be %v, got %v", got, tt.result, inputCopy[:got])
			}
		})
	}
}


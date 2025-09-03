package remove_element

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		val           int
		expectedCount int
		expectedNums  []int
	}{
		{
			name:          "Example 1: Remove all 3s",
			nums:          []int{3, 2, 2, 3},
			val:           3,
			expectedCount: 2,
			expectedNums:  []int{2, 2},
		},
		{
			name:          "Example 2: Remove all 2s",
			nums:          []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:           2,
			expectedCount: 5,
			expectedNums:  []int{0, 1, 3, 0, 4},
		},
		{
			name:          "Empty array",
			nums:          []int{},
			val:           1,
			expectedCount: 0,
			expectedNums:  []int{},
		},
		{
			name:          "All elements are the target value",
			nums:          []int{5, 5, 5, 5},
			val:           5,
			expectedCount: 0,
			expectedNums:  []int{},
		},
		{
			name:          "No elements are the target value",
			nums:          []int{1, 2, 3, 4},
			val:           5,
			expectedCount: 4,
			expectedNums:  []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalNums := make([]int, len(tt.nums))
			copy(originalNums, tt.nums)

			result := removeElement(tt.nums, tt.val)

			if result != tt.expectedCount {
				t.Errorf("removeElement(%v, %d) returned %d, expected %d", originalNums, tt.val, result, tt.expectedCount)
			}

			// Since the problem allows for the order to be changed, we need to check
			// that the first k elements contain the expected values, regardless of order
			if result > 0 {
				// Get the first k elements where k is the result
				remainingElements := tt.nums[:result]

				// Check that the remaining elements match the expected elements
				// (regardless of order)
				remainingMap := make(map[int]int)
				for _, num := range remainingElements {
					remainingMap[num]++
				}

				expectedMap := make(map[int]int)
				for _, num := range tt.expectedNums {
					expectedMap[num]++
				}

				if !reflect.DeepEqual(remainingMap, expectedMap) {
					t.Errorf("After removeElement(%v, %d), the first %d elements do not contain the expected values.\nGot: %v\nExpected to contain: %v",
						originalNums, tt.val, result, remainingElements, tt.expectedNums)
				}
			}
		})
	}
}

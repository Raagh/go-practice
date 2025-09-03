package mergesortedarray

import (
	"reflect"
	"testing"
)

// TestMerge tests the merge function with various test cases
func TestMerge(t *testing.T) {
	testCases := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			name:   "Example 1: Basic merge",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "Example 2: Second array is empty",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			expect: []int{1},
		},
		{
			name:   "Example 3: First array is empty",
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			expect: []int{1},
		},
		{
			name:   "Duplicate values",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{1, 2, 3},
			n:      3,
			expect: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:   "Negative values",
			nums1:  []int{-3, -1, 0, 0, 0, 0},
			m:      3,
			nums2:  []int{-4, -2, 5},
			n:      3,
			expect: []int{-4, -3, -2, -1, 0, 5},
		},
		{
			name:   "Large value differences",
			nums1:  []int{1, 100, 1000, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 200, 2000},
			n:      3,
			expect: []int{1, 2, 100, 200, 1000, 2000},
		},
		{
			name:   "Already sorted case",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{4, 5, 6},
			n:      3,
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of nums1 to avoid modifying the original
			nums1Copy := make([]int, len(tc.nums1))
			copy(nums1Copy, tc.nums1)

			// Call the merge function (which will be implemented)
			merge(nums1Copy, tc.m, tc.nums2, tc.n)

			// Check if the result matches the expected
			if !reflect.DeepEqual(nums1Copy, tc.expect) {
				t.Errorf("merge(%v, %d, %v, %d) = %v; want %v",
					tc.nums1, tc.m, tc.nums2, tc.n, nums1Copy, tc.expect)
			}
		})
	}
}

// TestMergeEdgeCases tests edge cases for the merge function
func TestMergeEdgeCases(t *testing.T) {
	testCases := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			name:   "Both arrays empty",
			nums1:  []int{},
			m:      0,
			nums2:  []int{},
			n:      0,
			expect: []int{},
		},
		{
			name:   "Single element in each array",
			nums1:  []int{1, 0},
			m:      1,
			nums2:  []int{2},
			n:      1,
			expect: []int{1, 2},
		},
		{
			name:   "First array has all smaller elements",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{4, 5, 6},
			n:      3,
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Second array has all smaller elements",
			nums1:  []int{4, 5, 6, 0, 0, 0},
			m:      3,
			nums2:  []int{1, 2, 3},
			n:      3,
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Arrays with same elements",
			nums1:  []int{1, 1, 1, 0, 0, 0},
			m:      3,
			nums2:  []int{1, 1, 1},
			n:      3,
			expect: []int{1, 1, 1, 1, 1, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of nums1 to avoid modifying the original
			nums1Copy := make([]int, len(tc.nums1))
			copy(nums1Copy, tc.nums1)

			// Call the merge function (which will be implemented)
			merge(nums1Copy, tc.m, tc.nums2, tc.n)

			// Check if the result matches the expected
			if !reflect.DeepEqual(nums1Copy, tc.expect) {
				t.Errorf("merge(%v, %d, %v, %d) = %v; want %v",
					tc.nums1, tc.m, tc.nums2, tc.n, nums1Copy, tc.expect)
			}
		})
	}
}

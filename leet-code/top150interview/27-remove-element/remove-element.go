package remove_element

import "slices"

// removeElement removes all occurrences of val from nums and returns the number of elements
// that remain. The first k elements of nums will contain the elements that were not removed.
// The order of elements may be changed.
//
// LeetCode Problem 27: Remove Element
// https://leetcode.com/problems/remove-element/
//
// Example 1:
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
//
// Example 2:
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,3,0,4,_,_,_]
//
// Constraints:
// - 0 <= nums.length <= 100
// - 0 <= nums[i] <= 50
// - 0 <= val <= 100
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = slices.Delete(nums, i, i+1)
			i--
		}
	}

	return len(nums)
}

package removeDuplicatesSorted

// RemoveDuplicates removes duplicates from a sorted array in-place
// and returns the number of unique elements.
//
// Given an integer array nums sorted in non-decreasing order, the function removes
// the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same.
//
// Then returns the number of unique elements in nums.
//
// Example:
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,...]
// (The first 2 elements of nums should be 1 and 2 respectively)
//
// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,...]
// (The first 5 elements of nums should be 0, 1, 2, 3, and 4 respectively)
//
// Constraints:
// - 1 <= nums.length <= 3 * 10^4
// - -100 <= nums[i] <= 100
// - nums is sorted in non-decreasing order
func RemoveDuplicates(nums []int) int {
	found := map[int]bool{}
	k := 0

	for i := 0; i < len(nums); i++ {
		if found[nums[i]] {
			continue
		}

		found[nums[i]] = true
		nums[k] = nums[i]
		k++
	}

	return k
}

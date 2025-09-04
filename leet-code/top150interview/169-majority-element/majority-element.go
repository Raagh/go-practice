package majority_element

// majorityElement finds the element that appears more than ⌊n/2⌋ times in the array.
// The majority element is guaranteed to exist in the array.
//
// Example 1:
// Input: nums = [3,2,3]
// Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
//
// Constraints:
// - n == nums.length
// - 1 <= n <= 5 * 10^4
// - -10^9 <= nums[i] <= 10^9
func majorityElement(nums []int) int {
	majorityThreashold := len(nums) / 2
	counts := map[int]int{}

	for i := 0; i < len(nums); i++ {
		v := nums[i]

		if count, ok := counts[v]; ok {
			counts[v] = count + 1
		} else {
			counts[v] = 1
		}

		if counts[v] > majorityThreashold {
			return v
		}
	}

	return -1
}

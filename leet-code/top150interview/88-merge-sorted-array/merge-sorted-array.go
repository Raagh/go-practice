package mergesortedarray

// merge merges two sorted arrays nums1 and nums2 into nums1
// nums1 has a size of m+n where the first m elements are filled and the rest are 0s
// nums2 has n elements
// The final sorted array should be stored inside nums1
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Start from the end of both arrays
	i := m - 1      // Index for the last element in nums1
	j := n - 1      // Index for the last element in nums2
	k := m + n - 1  // Index for the last position in the merged array (nums1)

	// While there are elements in both arrays to compare
	for i >= 0 && j >= 0 {
		// Compare elements and place the larger one at the end
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are remaining elements in nums2, copy them to nums1
	// Note: if there are remaining elements in nums1, they are already in place
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

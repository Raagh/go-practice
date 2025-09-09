package firstoccurrence

// strStr returns the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.
//
// Example 1:
// Input: haystack = "hello", needle = "ll"
// Output: 2
//
// Example 2:
// Input: haystack = "aaaaa", needle = "bba"
// Output: -1
//
// Example 3:
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
//
// Constraints:
// - 1 <= haystack.length, needle.length <= 10^4
// - haystack and needle consist of only lowercase English characters.
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

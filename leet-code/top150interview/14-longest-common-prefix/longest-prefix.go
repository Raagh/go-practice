package longestcommonprefix

import (
	"math"
	"strings"
)

// longestCommonPrefix finds the longest common prefix string amongst an array of strings.
// If there is no common prefix, it returns an empty string "".
//
// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
// Constraints:
// - 1 <= strs.length <= 200
// - 0 <= strs[i].length <= 200
// - strs[i] consists of only lowercase English letters.
func longestCommonPrefix(strs []string) string {
	// find shortest string
	shortest := ""
	shortestLen := math.MaxInt32
	for _, v := range strs {
		if len(v) < shortestLen {
			shortestLen = len(v)
			shortest = v
		}
	}

	// create all possible prefixes
	prefixes := map[string]bool{}
	for i := 0; i < len(shortest); i++ {
		prefix := []string{}
		for j := 0; j <= i; j++ {
			prefix = append(prefix, string(shortest[j]))
		}

		if len(prefix) > 0 {
			prefixes[strings.Join(prefix, "")] = true
		}
	}

	// compare prefixes with strings
	foundPre := ""
	for pre := range prefixes {
		all := true
		for _, str := range strs {
			if str == shortest {
				continue
			}

			if !strings.Contains(str[:len(pre)], pre) {
				all = false
			}
		}

		if all && len(pre) > len(foundPre) {
			foundPre = pre
		}
	}

	return foundPre
}

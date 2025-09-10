package longestcommonprefix

import (
	"testing"
)

func TestLongestCommonPrefix_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "Example 1",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Example 2",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "All Same",
			strs:     []string{"test", "test", "test"},
			expected: "test",
		},
		{
			name:     "One Letter Common",
			strs:     []string{"apple", "animal", "auto"},
			expected: "a",
		},
		{
			name:     "Longer Common Prefix",
			strs:     []string{"strawberry", "strawbale", "strawman"},
			expected: "straw",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.strs)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %s, expected %s", tt.strs, result, tt.expected)
			}
		})
	}
}

func TestLongestCommonPrefix_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "Empty Array",
			strs:     []string{},
			expected: "",
		},
		{
			name:     "Single String",
			strs:     []string{"alone"},
			expected: "alone",
		},
		{
			name:     "Single letter",
			strs:     []string{"a"},
			expected: "a",
		},
		{
			name:     "Empty String in Array",
			strs:     []string{"abc", "", "abcd"},
			expected: "",
		},
		{
			name:     "First String Shorter",
			strs:     []string{"a", "abc", "abcd"},
			expected: "a",
		},
		{
			name:     "No Common Prefix with Very Different Lengths",
			strs:     []string{"a", "b", "c"},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.strs)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %s, expected %s", tt.strs, result, tt.expected)
			}
		})
	}
}

func TestLongestCommonPrefix_AdditionalCases(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name: "Very Long Prefix",
			strs: []string{
				"thisisaverylongprefixtotest",
				"thisisaverylongprefixtotestmore",
				"thisisaverylongprefixtotesteverything",
			},
			expected: "thisisaverylongprefixtotest",
		},
		{
			name:     "Case Sensitive",
			strs:     []string{"Apple", "apple", "apPle"},
			expected: "",
		},
		{
			name:     "Special Characters",
			strs:     []string{"!@#$%", "!@#$%^", "!@#$%&"},
			expected: "!@#$%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.strs)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %s, expected %s", tt.strs, result, tt.expected)
			}
		})
	}
}

package firstoccurrence

import (
	"testing"
)

func TestStrStr_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Example 1",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "Example 2",
			haystack: "aaaaa",
			needle:   "bba",
			expected: -1,
		},
		{
			name:     "Example 3",
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			name:     "Needle at end",
			haystack: "testing",
			needle:   "ing",
			expected: 4,
		},
		{
			name:     "Repeated patterns",
			haystack: "mississippi",
			needle:   "issippi",
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strStr(tt.haystack, tt.needle)
			if result != tt.expected {
				t.Errorf("strStr(%s, %s) = %d, expected %d", tt.haystack, tt.needle, result, tt.expected)
			}
		})
	}
}

func TestStrStr_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Empty needle",
			haystack: "hello",
			needle:   "",
			expected: 0, // Empty needle is always found at index 0
		},
		{
			name:     "Needle longer than haystack",
			haystack: "abc",
			needle:   "abcdef",
			expected: -1,
		},
		{
			name:     "Identical strings",
			haystack: "test",
			needle:   "test",
			expected: 0,
		},
		{
			name:     "Single character haystack and needle",
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
		{
			name:     "Single character not found",
			haystack: "abc",
			needle:   "d",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strStr(tt.haystack, tt.needle)
			if result != tt.expected {
				t.Errorf("strStr(%s, %s) = %d, expected %d", tt.haystack, tt.needle, result, tt.expected)
			}
		})
	}
}

func TestStrStr_AdditionalCases(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Multiple occurrences",
			haystack: "ababababa",
			needle:   "aba",
			expected: 0, // Should return the first occurrence
		},
		{
			name:     "Partial matches",
			haystack: "ababac",
			needle:   "abac",
			expected: 2,
		},
		{
			name:     "Special characters",
			haystack: "hello!@#$",
			needle:   "!@#",
			expected: 5,
		},
		{
			name:     "Long strings",
			haystack: "thisisaverylongstringtotest",
			needle:   "longstring",
			expected: 11,
		},
		{
			name:     "Case sensitivity",
			haystack: "aBcDeF",
			needle:   "bcd",
			expected: -1, // Should be case sensitive
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strStr(tt.haystack, tt.needle)
			if result != tt.expected {
				t.Errorf("strStr(%s, %s) = %d, expected %d", tt.haystack, tt.needle, result, tt.expected)
			}
		})
	}
}

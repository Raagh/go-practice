package lastword

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Basic case with two words",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "String with leading, trailing, and multiple spaces",
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Multiple words without trailing spaces",
			input:    "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "Single word",
			input:    "word",
			expected: 4,
		},
		{
			name:     "Single word with leading spaces",
			input:    "    word",
			expected: 4,
		},
		{
			name:     "Single word with trailing spaces",
			input:    "word    ",
			expected: 4,
		},
		{
			name:     "Single character word",
			input:    "a",
			expected: 1,
		},
		{
			name:     "Multiple single character words",
			input:    "a b c d",
			expected: 1,
		},
		{
			name:     "Empty words between valid words",
			input:    "day   night",
			expected: 5,
		},
		{
			name:     "Edge case with minimum constraint",
			input:    "a",
			expected: 1,
		},
		{
			name:     "Edge case with longer last word",
			input:    "Today is a beautiful day in the neighborhood",
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLastWord(tt.input)
			if result != tt.expected {
				t.Errorf("lengthOfLastWord(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// This test focuses on edge cases and special situations
func TestLengthOfLastWord_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Many spaces between words",
			input:    "hello          world",
			expected: 5,
		},
		{
			name:     "Alternating words and spaces",
			input:    "a b c d e f g",
			expected: 1,
		},
		{
			name:     "Very long word",
			input:    "supercalifragilisticexpialidocious",
			expected: 34,
		},
		{
			name:     "Word with numbers and special characters (if allowed by constraints)",
			input:    "hello123",
			expected: 8,
		},
		{
			name:     "Many trailing spaces",
			input:    "last                   ",
			expected: 4,
		},
		{
			name:     "Many leading spaces",
			input:    "                   first",
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLastWord(tt.input)
			if result != tt.expected {
				t.Errorf("lengthOfLastWord(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

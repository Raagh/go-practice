package countchars

import (
	"reflect"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{
			input: "example string",
			expected: map[rune]int{
				'e': 2,
				'x': 1,
				'a': 1,
				'm': 1,
				'p': 1,
				'l': 1,
				' ': 1,
				's': 1,
				't': 1,
				'r': 1,
				'i': 1,
				'n': 1,
				'g': 1,
			},
		},
		{
			input:    "",
			expected: map[rune]int{},
		},
		{
			input: "aaa",
			expected: map[rune]int{
				'a': 3,
			},
		},
	}

	for _, test := range tests {
		result := countCharacters(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

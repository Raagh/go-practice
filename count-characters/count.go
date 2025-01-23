package countchars

// ### Exercise: Count Character Occurrences in a String
//
// In this exercise, you will write a function in Go to count the occurrences of each character in a given string. This is a common task that can help you understand how to work with strings and maps in Go.
//
// Given the string `"example string"`, your function should return the following counts:
// - 'e': 2
// - 'x': 1
// - 'a': 1
// - 'm': 1
// - 'p': 1
// - 'l': 1
// - ' ': 1
// - 's': 1
// - 't': 1
// - 'r': 1
// - 'i': 1
// - 'n': 1
// - 'g': 1
func countCharacters(input string) map[rune]int {
	counts := map[rune]int{}

	for _, v := range input {
		counts[v]++
	}

	return counts
}

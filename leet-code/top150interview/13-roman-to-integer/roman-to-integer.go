package romantointeger

// romanToInt converts a Roman numeral to an integer.
// Symbols: 'I' = 1, 'V' = 5, 'X' = 10, 'L' = 50, 'C' = 100, 'D' = 500, 'M' = 1000
// When a smaller value precedes a larger value, the smaller value is subtracted from the larger value.
// Otherwise, values are added.
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		currentLetter := romanMap[s[i]]
		if i+1 < len(s) && currentLetter < romanMap[s[i+1]] {
			sum -= romanMap[s[i]]
		} else {
			sum += romanMap[s[i]]
		}
	}
	return sum
}

package romantointeger

import (
	"testing"
)

func TestRomanToInt_BasicNumerals(t *testing.T) {
	tests := []struct {
		name     string
		roman    string
		expected int
	}{
		{"Single I", "I", 1},
		{"Single V", "V", 5},
		{"Single X", "X", 10},
		{"Single L", "L", 50},
		{"Single C", "C", 100},
		{"Single D", "D", 500},
		{"Single M", "M", 1000},
		{"Multiple Is", "III", 3},
		{"Basic Addition", "VI", 6},
		{"Basic Numerals", "VIII", 8},
		{"Mixed Basic", "MDCLXVI", 1666}, // 1000 + 500 + 100 + 50 + 10 + 5 + 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := romanToInt(tt.roman)
			if result != tt.expected {
				t.Errorf("romanToInt(%s) = %d, expected %d", tt.roman, result, tt.expected)
			}
		})
	}
}

func TestRomanToInt_SubtractionRules(t *testing.T) {
	tests := []struct {
		name     string
		roman    string
		expected int
	}{
		{"IV", "IV", 4},          // 5-1
		{"IX", "IX", 9},          // 10-1
		{"XL", "XL", 40},         // 50-10
		{"XC", "XC", 90},         // 100-10
		{"CD", "CD", 400},        // 500-100
		{"CM", "CM", 900},        // 1000-100
		{"MCMXC", "MCMXC", 1990}, // 1000 + (1000-100) + (100-10)
		{"MCMIV", "MCMIV", 1904}, // 1000 + (1000-100) + (5-1)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := romanToInt(tt.roman)
			if result != tt.expected {
				t.Errorf("romanToInt(%s) = %d, expected %d", tt.roman, result, tt.expected)
			}
		})
	}
}

func TestRomanToInt_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		roman    string
		expected int
	}{
		{"Empty String", "", 0},
		{"Largest Valid Roman Numeral", "MMMCMXCIX", 3999}, // 3000 + 900 + 90 + 9
		{"Complex Combination", "MMMDCCCLXXXVIII", 3888},   // 3000 + 800 + 80 + 8
		{"All Subtraction Rules", "CMXCIX", 999},           // 900 + 90 + 9
		{"Repeated Numerals", "MMMMM", 5000},               // 5 × 1000
		{"Long Simple Pattern", "IIIIIIIII", 9},            // 9 × 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := romanToInt(tt.roman)
			if result != tt.expected {
				t.Errorf("romanToInt(%s) = %d, expected %d", tt.roman, result, tt.expected)
			}
		})
	}
}

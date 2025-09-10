package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Simple increasing prices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5, // Buy at 1, sell at 6, profit = 5
		},
		{
			name:     "Decreasing prices",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0, // No profit can be made
		},
		{
			name:     "Flat prices",
			prices:   []int{3, 3, 3, 3, 3},
			expected: 0, // No profit can be made
		},
		{
			name:     "Single price",
			prices:   []int{5},
			expected: 0, // Can't buy and sell with only one price
		},
		{
			name:     "Empty slice",
			prices:   []int{},
			expected: 0, // No prices provided
		},
		{
			name:     "Complex case",
			prices:   []int{3, 2, 6, 5, 0, 3},
			expected: 4, // Buy at 2, sell at 6, profit = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfit(tt.prices)
			if result != tt.expected {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, result, tt.expected)
			}
		})
	}
}


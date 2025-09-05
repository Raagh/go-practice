package besttimetobuyandsellstock

import "math"

// MaxProfit calculates the maximum profit that can be made from buying and selling a stock.
// You can only buy once and sell once, and buying must come before selling.
// If no profit can be made, return 0.
//
// Parameters:
// - prices: A slice of integers where prices[i] is the price of the stock on day i.
//
// Returns:
// - The maximum profit that can be achieved, or 0 if no profit can be made.
func MaxProfit(prices []int) int {
	minBuy := math.MaxInt
	maxProfit := 0

	for _, price := range prices {
		if price < minBuy {
			minBuy = price
		} else if price-minBuy > maxProfit {
			maxProfit = price - minBuy
		}
	}

	return maxProfit
}

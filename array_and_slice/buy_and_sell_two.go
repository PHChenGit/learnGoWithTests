package array_and_slice

import "fmt"

func BuyAndSellTwo(prices []int) int  {
	totalProfit := 0
	minPrice := 0
	maxPrice := 0
	idx := 0
	priceLength := len(prices)-1
	fmt.Printf("len: %d", priceLength)
	for idx < priceLength {
		for idx < priceLength && prices[idx] >= prices[idx+1] {
			idx += 1
		}
		minPrice = prices[idx]

		for idx < priceLength && prices[idx] < prices[idx+1] {
			idx += 1
		}
		maxPrice = prices[idx]

		totalProfit += maxPrice - minPrice
	}
	return totalProfit
}
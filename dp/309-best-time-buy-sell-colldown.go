// Package dp provides ...
package dp

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// every day have three choice
	sell := make([]int, len(prices)+1)
	buy := make([]int, len(prices)+1)
	buy[1] = -prices[0]
	for idx := 2; idx < len(prices)+1; idx++ {
		buy[idx] = max(buy[idx-1], sell[idx-2]-prices[idx-1])
		sell[idx] = max(buy[idx-1]+prices[idx-1], sell[idx-1])
	}

	return sell[len(prices)]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

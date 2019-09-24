package dp

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	profit := make([]int, n)
	minLocal := prices[0]

	profit_r := make([]int, n)
	maxLocal_r := prices[n-1]

	for i := 1; i < len(prices); i++ {
		minLocal = min(prices[i], minLocal)
		maxLocal_r = max(prices[n-i-1], maxLocal_r)
		profit[i] = max(profit[i-1], prices[i]-minLocal)
		profit_r[n-i-1] = max(profit_r[n-i], maxLocal_r-prices[n-i-1])
	}

	res := 0

	res = max(res, profit[n-1])
	for i := 1; i < len(prices)-1; i++ {
		res = max(res, profit[i]+profit_r[i+1])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

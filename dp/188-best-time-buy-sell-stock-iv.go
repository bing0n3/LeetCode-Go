// Package dp provides ...
package dplocalMax

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	n := len(prices)

	if k >= n/2 {
		p := 0
		l_m := prices[0]

		for i := 1; i < n; i++ {
			if prices[i] > l_m {
				p += prices[i] - l_m
			}
			l_m = prices[i]
		}

		return p
	}

	local_max := make([][]int, n)
	profits := make([][]int, n)
	for i := range profits {
		profits[i] = make([]int, k+1)
		local_max[i] = make([]int, k+1)
	}

	for i := 1; i <= k; i++ {
		for j := 1; j < n; j++ {
			local_max[j][i] = max(profits[j-1][i-1], local_max[j-1][i]) + (prices[j] - prices[j-1])
			profits[j][i] = max(local_max[j][i], profits[j-1][i])
		}
	}

	return profits[n-1][min(k, n/2)]

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

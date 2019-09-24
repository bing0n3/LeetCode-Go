// Package dp provides ...
package dp

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	dp := make([][]int, len(costs))
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]

	for i := 1; i < len(costs); i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][1], dp[i-1][0]) + costs[i][2]

	}
	return min(dp[len(costs)-1][0], min(dp[len(costs)-1][1], dp[len(costs)-1][2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

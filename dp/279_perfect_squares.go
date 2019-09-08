package perfectSquare

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	var min func(int, int) int

	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}

	return dp[n]

}

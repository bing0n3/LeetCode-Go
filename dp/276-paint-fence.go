package dp

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return k
	}

	dp := make([]int, n)
	dp[0] = k
	dp[1] = k * k

	for i := 2; i < n; i++ {
		dp[i] = (k - 1) * (dp[i-1] + dp[i-2])
	}

	return dp[len(dp)-1]
}

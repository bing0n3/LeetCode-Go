package uniquepath

// method 1
// backtraning, exceed time limitation
func uniquePaths(m int, n int) int {
	cnt := 0
	generatePath(&cnt, m, n)
	return cnt
}

func generatePath(cnt *int, m int, n int) {
	if m == 1 && n == 1 {
		*cnt++
		return
	}

	if m == 1 {
		generatePath(cnt, m, n-1)
		return
	}

	if n == 1 {
		generatePath(cnt, m-1, n)
		return
	}

	generatePath(cnt, m-1, n)
	generatePath(cnt, m, n-1)
}

// method 2
// dynamic programmibf
func uniquePathsDP(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1]
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[n-1][m-1]
}

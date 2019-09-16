package dp

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// 每一层把一层投射到
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	res := 0
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
				if res < dp[i][j] {
					res = dp[i][j]
				}
			}
		}
	}

	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

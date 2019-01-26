package unique_path_ii

// 回溯法
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	rowNum := len(obstacleGrid)
	colNum := len(obstacleGrid[0])
	dp := make([][]int, rowNum)
	for i := 0; i < rowNum; i++ {
		dp[i] = make([]int, colNum)
	}
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
			} else if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[rowNum-1][colNum-1]
}


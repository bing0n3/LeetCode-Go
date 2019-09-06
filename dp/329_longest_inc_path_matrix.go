// Package longestIncPath provides ...
package longestIncPath

/*
Backtracking without memory works,
but will get Time Limit Exceeded
*/
func longestIncreasingPath(matrix [][]int) int {

	if len(matrix) == 0 {
		return 0
	}
	//backtracking
	maxCnt := 0
	var res int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res = dfs(matrix, i, j, -1)
			if res > maxCnt {
				maxCnt = res
			}
		}
	}

	return maxCnt

}

func dfs(matrix [][]int, i, j int, prv int) int {
	// return the largest one from four direction
	if matrix[i][j] == -1 {
		return 0
	}

	if matrix[i][j] <= prv {
		return 0
	}

	tmp := matrix[i][j]
	matrix[i][j] = -1

	// is boarder？
	up, down, right, left := 0, 0, 0, 0

	if i+1 < len(matrix) {
		up = dfs(matrix, i+1, j, tmp)
	}
	if i-1 >= 0 {
		down = dfs(matrix, i-1, j, tmp)
	}

	if j+1 < len(matrix[0]) {
		right = dfs(matrix, i, j+1, tmp)
	}

	if j-1 >= 0 {
		left = dfs(matrix, i, j-1, tmp)
	}

	matrix[i][j] = tmp

	return max(left, max(down, max(right, up))) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
Dynamic Programming method:
We can use memorization in here.
The reason is that it is a increase path, so that visited[i][i] is not necessary.
*/
func longestIncreasingPath_dp(matrix [][]int) int {

	if len(matrix) == 0 {
		return 0
	}
	//backtracking
	maxCnt := 0

	var res int

	// memorization
	dp := [][]int{}
	for i := 0; i < len(matrix); i++ {
		dp = append(dp, []int{})
		for j := 0; j < len(matrix[0]); j++ {
			dp[i] = append(dp[i], 0)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res = dfs2(matrix, dp, i, j)
			if res > maxCnt {
				maxCnt = res
			}
		}
	}

	return maxCnt

}

func dfs2(matrix [][]int, dp [][]int, i, j int) int {
	// return the largest one from four direction
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	// is boarder？
	up, down, right, left := 0, 0, 0, 0

	if i+1 < len(matrix) {
		if matrix[i+1][j] <= matrix[i][j] {
			up = 0
		} else {
			up = dfs2(matrix, dp, i+1, j)
		}
	}
	if i-1 >= 0 {
		if matrix[i-1][j] <= matrix[i][j] {
			down = 0
		} else {
			down = dfs2(matrix, dp, i-1, j)
		}
	}

	if j+1 < len(matrix[0]) {
		if matrix[i][j+1] <= matrix[i][j] {
			right = 0
		} else {
			right = dfs2(matrix, dp, i, j+1)
		}

	}

	if j-1 >= 0 {
		if matrix[i][j-1] <= matrix[i][j] {
			left = 0
		} else {
			left = dfs2(matrix, dp, i, j-1)
		}
	}

	dp[i][j] = max(left, max(down, max(right, up))) + 1

	return dp[i][j]
}

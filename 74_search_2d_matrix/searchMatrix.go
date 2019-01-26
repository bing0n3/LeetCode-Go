package search_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	whichRow := -1

	for i := 0; i < m; i++ {
		if target == matrix[i][n-1] || target == matrix[i][0] {
			return true
		}

		if target < matrix[i][n-1] && target > matrix[i][0] {
			whichRow = i
		}
	}

	if whichRow == -1 {
		return false
	}

	for j := 0; j < n; j++ {
		if matrix[whichRow][j] == target {
			return true
		}
	}
	return false
}

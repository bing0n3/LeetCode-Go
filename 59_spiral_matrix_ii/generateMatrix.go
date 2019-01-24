package spiral_matrix_ii

func generateMatrix(n int) [][]int {
	nums := make([][]int, n)
	for k := 0; k < n; k++ {
		nums[k] = make([]int, n)
	}
	row, col := n-1, n-1
	r, c := 0, 0

	count := 1

	for r <= row && c <= col {
		for i := c; i <= col; i++ {
			nums[r][i] = count
			count++
		}
		r++
		for i := r; i <= row; i++ {
			nums[i][col] = count
			count++
		}
		col--
		if r <= row {
			for i := col; i >= c; i-- {
				nums[row][i] = count
				count++
			}
		}
		row--
		if c <= col {
			for i := row; i >= r; i-- {
				nums[i][c] = count
				count++
			}
		}
		c++

	}
	return nums
}

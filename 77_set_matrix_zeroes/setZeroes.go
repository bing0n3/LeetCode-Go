package setZeroes

func setZeroes(matrix [][]int) {
	col0 := 1
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < cols; j++ {
			// when an element, set the first col and row element to zero to store the state.
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
}

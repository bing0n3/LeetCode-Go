package nqueens

func solveNQueens(n int) [][]string {
	var result [][]string
	t := make([][]rune, n)
	for i := range t {
		t[i] = make([]rune, n)
		for j := range t[i] {
			t[i][j] = '.'
		}
	}
	backtracking(&result, t, n, 0)
	return result
}

func backtracking(res *[][]string, matrix [][]rune, n, rowCount int) {
	if rowCount == n {
		var resultMatrix []string
		for _, runes := range matrix {
			resultMatrix = append(resultMatrix, string(runes))
		}
		*res = append(*res, resultMatrix)
		return
	}

	for j := 0; j < n; j++ {
		matrix[rowCount][j] = 'Q'
		if isValid(matrix, rowCount, j) {
			backtracking(res, matrix, n, rowCount+1)
		}
		matrix[rowCount][j] = '.'
	}

}

func isValid(list [][]rune, row, col int) bool {

	// get a string
	line := list[row]

	// compare current row and the previous row
	for i := 0; i < row; i++ {
		// two queen in the same colum
		if line[col] == list[i][col] {
			return false
		}

		diff := row - i
		// if two quuen in the diagonal line
		if col-diff >= 0 && list[i][col-diff] == 'Q' {
			return false
		}

		if col+diff < len(line) && list[i][col+diff] == 'Q' {
			return false
		}
	}

	return true
}

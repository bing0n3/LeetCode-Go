package rotateimage

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-1-i] = matrix[j][n-1-i], matrix[i][j]
			matrix[i][j], matrix[n-1-i][n-j-1] = matrix[n-1-i][n-j-1], matrix[i][j]
			matrix[i][j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j]
		}
	}

}

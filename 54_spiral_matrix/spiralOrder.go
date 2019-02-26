package spiralmatrixii

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	m := len(matrix)
	if m == 0 {
		return result
	}

	if len(matrix) == 1 {
		return matrix[0]
	}

	n := len(matrix[0])

	if n == 0 {
		return result
	}

	if n == 1 {
		for _, arr := range matrix {
			result = append(result, arr[0])
		}
		return result
	}

	// add the top line
	result = append(result, matrix[0]...)
	// add the right line
	for i := 1; i < m; i++ {
		result = append(result, matrix[i][n-1])
	}
	// add the bottom line
	for i := n - 2; i >= 0; i-- {
		result = append(result, matrix[m-1][i])
	}
	// add the left line
	for i := m - 2; i >= 1; i-- {
		result = append(result, matrix[i][0])
	}

	remain := [][]int{}
	// remove added element
	// remove added element

	for i := 1; i < m-1; i++ {
		remain = append(remain, matrix[i])
	}
	for i := range remain {
		remain[i] = remain[i][1 : n-1]
	}

	return append(result, spiralOrder(remain)...)
}

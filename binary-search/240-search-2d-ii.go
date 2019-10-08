package bsearch

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	x, y := 0, m-1

	for y >= 0 && x < n {
		cur := matrix[x][y]

		if cur == target {
			return true
		} else if cur < target {
			x++
		} else {
			y--
		}
	}

	return false
}

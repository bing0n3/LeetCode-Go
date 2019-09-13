package matrix01

func updateMatrix(matrix [][]int) [][]int {
	const maxUint = ^uint(0)
	const maxInt = int(maxUint >> 1)
	queue := [][]int{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				matrix[i][j] = maxInt
			}
		}
	}
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {

		// queue.pop(0)
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, dir := range direction {
			row := i + dir[0]
			col := j + dir[1]
			if row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) && matrix[row][col] > 1+matrix[i][j] {
				matrix[row][col] = 1 + matrix[i][j]
				queue = append(queue, []int{row, col})
			}
		}

	}

	return matrix
}

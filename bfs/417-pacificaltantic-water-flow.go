package dfs

func pacificAtlantic(matrix [][]int) [][]int {
	// bfs
	res := [][]int{}

	if len(matrix) == 0 || matrix == nil {
		return res
	}

	var N = len(matrix)
	var M = len(matrix[0])

	pacific := make([][]bool, N)
	atlantic := make([][]bool, N)

	for i := 0; i < N; i++ {
		pacific[i] = make([]bool, M)
		atlantic[i] = make([]bool, M)
	}

	// build two queue for pacific and atlantic to coordnaitio
	pQ := [][]int{}
	aQ := [][]int{}

	//initial queue
	for i := 0; i < M; i++ {
		pQ = append(pQ, []int{0, i})
		aQ = append(aQ, []int{N - 1, i})
		pacific[0][i] = true
		atlantic[N-1][i] = true
	}

	for i := 0; i < N; i++ {
		pQ = append(pQ, []int{i, 0})
		aQ = append(aQ, []int{i, M - 1})
		pacific[i][0] = true
		atlantic[i][M-1] = true
	}

	bfs(matrix, pQ, pacific)
	bfs(matrix, aQ, atlantic)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func bfs(matrix [][]int, queue [][]int, visited [][]bool) {
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var N = len(matrix)
	var M = len(matrix[0])
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			x := cur[0] + d[0]
			y := cur[1] + d[1]
			if x < 0 || x >= N || y >= M || y < 0 || matrix[x][y] < matrix[cur[0]][cur[1]] || visited[x][y] {
				continue
			}
			visited[x][y] = true
			queue = append(queue, []int{x, y})
		}
	}
}

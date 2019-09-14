package dfs

func maxAreaOfIsland(grid [][]int) int {

	res, cnt := 0, 0
	for i := range grid {
		for j := range grid[0] {
			cnt = dfs(&grid, i, j)
			if cnt > res {
				res = cnt
			}
		}
	}

	return res

}

func dfs(grid *[][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == -1 || (*grid)[i][j] == 0 {
		return 0
	}

	var left, right, down, up int

	(*grid)[i][j] = -1

	right = dfs(grid, i+1, j)
	up = dfs(grid, i, j+1)
	left = dfs(grid, i-1, j)
	down = dfs(grid, i, j-1)

	return right + down + left + up + 1
}

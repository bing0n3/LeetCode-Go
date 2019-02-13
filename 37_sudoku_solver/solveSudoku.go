package sudoku_solver

var (
	cubeMaps []map[byte]interface{}
	rowMaps  []map[byte]interface{}
	colMaps  []map[byte]interface{}
)

func solveSudoku(board [][]byte) {
	cubeMaps = make([]map[byte]interface{}, 9)
	rowMaps = make([]map[byte]interface{}, 9)
	colMaps = make([]map[byte]interface{}, 9)

	for i := 0; i < 9; i++ {
		cubeMaps[i] = make(map[byte]interface{})
		rowMaps[i] = make(map[byte]interface{})
		colMaps[i] = make(map[byte]interface{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				cubeMaps[i/3*3+j/3][board[i][j]] = struct{}{}
				rowMaps[i][board[i][j]] = struct{}{}
				colMaps[j][board[i][j]] = struct{}{}
			}
		}
	}
	backtracking(&board, 0)
}

func backtracking(board *[][]byte, count int) bool {
	if count == 81 {
		return true
	}
	i := count / 9
	j := count % 9

	if (*board)[i][j] != '.' {
		return backtracking(board, count+1)
	}

	for b := byte('1'); b <= '9'; b++ {
		_, okcube := cubeMaps[i/3*3+j/3][b]
		_, okrow := rowMaps[i][b]
		_, okcol := colMaps[j][b]
		if !okcube && !okrow && !okcol {
			(*board)[i][j] = b
			cubeMaps[i/3*3+j/3][b] = struct{}{}
			rowMaps[i][b] = struct{}{}
			colMaps[j][b] = struct{}{}
			if backtracking(board, count+1) {
				return true
			}
			(*board)[i][j] = '.'
			delete(cubeMaps[i/3*3+j/3], b)
			delete(rowMaps[i], b)
			delete(rowMaps[j], b)
		}
	}
	(*board)[i][j] = '.'
	return false
}

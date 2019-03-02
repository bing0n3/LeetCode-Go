package wordsearch

func exist(board [][]byte, word string) bool {

	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if findNeighbor(board, word, i, j) {
				return true
			}
		}
	}

	return false
}

func findNeighbor(board [][]byte, word string, i, j int) bool {

	if len(word) == 0 {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[0] {
		return false
	} else if len(word) == 1 {
		return true
	}

	board[i][j] = '.'

	exist := findNeighbor(board, word[1:], i-1, j) || findNeighbor(board, word[1:], i, j-1) || findNeighbor(board, word[1:], i+1, j) || findNeighbor(board, word[1:], i, j+1)

	board[i][j] = word[0]

	return exist
}

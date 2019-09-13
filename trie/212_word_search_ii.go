// Package wordSearch provides ...
package wordSearch

type Trie struct {
	Children map[byte]*Trie
	Word     string
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Children: make(map[byte]*Trie), Word: ""}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for i, _ := range word {
		if _, ok := parent.Children[word[i]]; !ok {
			newChild := &Trie{Children: make(map[byte]*Trie), Word: ""}
			parent.Children[word[i]] = newChild
		}
		parent = parent.Children[word[i]]
	}
	parent.Word = word
}

func findWords(board [][]byte, words []string) []string {
	// build a trie
	root := Constructor()
	for _, word := range words {
		root.Insert(word)
	}

	res := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, &root, &res)
		}
	}
	return res
}

func dfs(board [][]byte, i, j int, root *Trie, res *[]string) {
	c := board[i][j]
	node := root
	if c == '#' {
		return
	}
	if _, ok := node.Children[c]; !ok {
		return
	}

	node = node.Children[c]

	if node.Word != "" {
		*res = append(*res, node.Word)
		node.Word = ""
	}

	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, node, res)
	}
	if j > 0 {
		dfs(board, i, j-1, node, res)
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, node, res)
	}
	if j < len(board[0])-1 {
		dfs(board, i, j+1, node, res)
	}
	board[i][j] = c
}

// Package binaryTreePath provides ...
package binaryTreePath

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	collect := [][]string{}
	record := []string{}
	res := []string{}

	dfs(root, &collect, record)
	for _, list := range collect {
		res = append(res, strings.Join(list, "->"))
	}
	return res
}

func dfs(root *TreeNode, collect *[][]string, record []string) {
	tmp := make([]string, len(record)+1)
	copy(tmp, append(record, strconv.Itoa(root.Val)))
	if root.Left == nil && root.Right == nil {
		*collect = append(*collect, tmp)
		return
	}
	if root.Left != nil {
		dfs(root.Left, collect, tmp)
	}
	if root.Right != nil {
		dfs(root.Right, collect, tmp)
	}
}

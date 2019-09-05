package binary_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	cache := []int{}
	dfs(root, 0, &cache)
	return cache
}

func dfs(root *TreeNode, depth int, cache *[]int) {
	if root == nil {
		return
	}
	if len(*cache) < depth+1 {
		*cache = append(*cache, root.Val)
	} else {
		(*cache)[depth] = root.Val
	}

	dfs(root.Left, depth+1, cache)
	dfs(root.Right, depth+1, cache)

}

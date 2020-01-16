package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	var wrapList [][]int

	queue = append(queue, root)

	for len(queue) > 0 {
		var level []int
		cnt := len(queue)
		var tmpQ []*TreeNode
		for i := 0; i < cnt; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				tmpQ = append(tmpQ, node.Left)
			}
			if node.Right != nil {
				tmpQ = append(tmpQ, node.Right)
			}

		}
		queue = tmpQ
		wrapList = append(wrapList, level)
	}

	return wrapList
}

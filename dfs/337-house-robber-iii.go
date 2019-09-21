package dfs
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    if root == nil {
        return 0;
    }
    if root.Left == nil && root.Right == nil {
        return root.Val
    }
    
    var leftVal, subLeftLevel, rightVal, subRightLevel int
    if root.Left != nil {
        leftVal = rob(root.Left)
        subLeftLevel = rob(root.Left.Left) + rob(root.Left.Right)
    }
    if root.Right != nil {
        rightVal = rob(root.Right)
        subRightLevel = rob(root.Right.Left) + rob(root.Right.Right)
    }
    
    return max(subLeftLevel + subRightLevel + root.Val, leftVal + rightVal)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


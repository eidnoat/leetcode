/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    var (
		dfs func(node *TreeNode, left bool)
		sum int
	)

	dfs = func(node *TreeNode, left bool) {
		if node.Left == nil && node.Right == nil && left {
			sum += node.Val
		}

		if node.Left != nil {
			dfs(node.Left, true)
		}
		if node.Right != nil {
			dfs(node.Right, false)
		}
	}
	dfs(root, false)

	return sum
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.44MB，击败 67.71%
*/

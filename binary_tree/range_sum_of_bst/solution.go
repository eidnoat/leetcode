/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    var (
		dfs func(node *TreeNode)
		ret int
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val >= low && node.Val <= high {
			ret += node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：10.52MB，击败 89.36%
*/

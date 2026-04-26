/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var (
		dfs func(node *TreeNode, v int)
		ret int
	)

	dfs = func(node *TreeNode, v int) {
		v = v*10+node.Val
		if node.Left == nil && node.Right == nil {
			ret += v
			return
		}
		if node.Left != nil {
			dfs(node.Left, v)
		}
		if node.Right != nil {
			dfs(node.Right, v)
		}
	}
	dfs(root, 0)

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.03MB，击败 50.95%
*/

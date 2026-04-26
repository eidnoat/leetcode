/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    var (
		dfs func(node *TreeNode, v int)
		ret int
	)

	dfs = func(node *TreeNode, v int) {
		if node == nil {
			return
		}

		if node.Val >= v {
			v, ret = node.Val, ret+1
		}
		dfs(node.Left, v)
		dfs(node.Right, v)
	}
	dfs(root, math.MinInt)

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：14.20MB，击败 35.20%
*/

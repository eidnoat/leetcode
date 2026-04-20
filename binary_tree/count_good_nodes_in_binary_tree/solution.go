/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(tree *TreeNode) int {
	var (
		dfs func(root *TreeNode, maxV int)
		ret int
	)
	dfs = func(root *TreeNode, maxV int) {
		if root == nil {
			return
		}

		if maxV <= root.Val {
			ret++
		}

		dfs(root.Left, max(maxV, root.Val))
		dfs(root.Right, max(maxV, root.Val))
	} 
	dfs(tree, math.MinInt)

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：13.93MB，击败 46.82%
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(tree *TreeNode) *TreeNode {
	var dfs func(l, r *TreeNode, odd bool)
	dfs = func(l, r *TreeNode, odd bool) {
		if l == nil || r == nil {
			return
		}

		if odd {
			l.Val, r.Val = r.Val, l.Val
		}

		dfs(l.Left, r.Right, !odd)
		dfs(l.Right, r.Left, !odd)
	}
	dfs(tree.Left, tree.Right, true)

	return tree
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：10.59MB，击败 36.36%
*/

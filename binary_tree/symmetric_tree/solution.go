/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    var dfs func(l, r *TreeNode) bool

	dfs = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}

		return l.Val == r.Val && dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
	}

	return dfs(root.Left, root.Right)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.60MB，击败 53.89%
*/

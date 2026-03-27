/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(input, p, q *TreeNode) *TreeNode {
	var (
		ret *TreeNode
		dfs func(root *TreeNode) (bool, bool)
	)
	dfs = func(root *TreeNode) (bool, bool) {
		if root == nil {
			return false, false
		}

		lp, lq := dfs(root.Left)
		rp, rq := dfs(root.Right)

		pok, qok := lp || rp || root == p, lq || rq || root == q
		if pok && qok && ret == nil {
			ret = root
		}

		return pok, qok
	}
	dfs(input)

	return ret
}

/*
	执行用时分布：20ms，击败 5.50%
	消耗内存分布：9.23MB，击败 15.72%
*/

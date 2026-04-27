/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var dfs func(node *TreeNode) (int, int, bool)
	dfs = func(node *TreeNode) (int, int, bool) {
		if node == nil {
			return 0, 0, true
		}

		lmin, lmax, lok := dfs(node.Left)
		rmin, rmax, rok := dfs(node.Right)
		curMin, curMax := node.Val, node.Val
		if node.Left != nil {
			curMin = lmin
		}
		if node.Right != nil {
			curMax = rmax
		}

		return curMin, curMax, lok && rok && (node.Left == nil || node.Val > lmax) && (node.Right == nil || node.Val < rmin)
	}
	_, _, ok := dfs(root)

	return ok
}

/*
	执行用时分布：1ms，击败 7.43%
	消耗内存分布：7.00MB，击败 53.77%
*/

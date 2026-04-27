/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) int {
    var (
		dfs func(node *TreeNode) (bool, int, int, int)
		ret int
	)
	dfs = func(node *TreeNode) (bool, int, int, int) {
		if node == nil {
			return true, 0, 0, 0
		}

		lex, rex := node.Left != nil, node.Right != nil
		lok, lmin, lmax, lsum := dfs(node.Left)
		rok, rmin, rmax, rsum := dfs(node.Right)

		ok := lok && rok && (!lex || lmax < node.Val) && (!rex || rmin > node.Val)
		if !ok {
			return false, 0, 0, 0
		}
		ret = max(ret, lsum+rsum+node.Val)
        
		curSum := node.Val

        curMin := node.Val
        if lex {
            curMin, curSum = lmin, curSum+lsum
        }

		curMax := node.Val
		if rex {
			curMax, curSum = rmax, curSum+rsum
		}

		return ok, curMin, curMax, curSum
	}
	dfs(root)

	return ret
}

/*
	执行用时分布：15ms，击败 84.00%
	消耗内存分布：18.85MB，击败 36.00%
*/

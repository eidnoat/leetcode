/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(input *TreeNode) int {
    var (
		ret int
		dfs func(root *TreeNode) (int, int)
	)
	dfs = func(root *TreeNode) (int, int) {
		curMin, curMax := root.Val, root.Val
		if root.Left != nil {
			lmin, lmax := dfs(root.Left)
			curMin, curMax = min(curMin, lmin), max(curMax, lmax)
			ret = max(ret, max(abs(root.Val-lmax), abs(root.Val-lmin)))
		}
		if root.Right != nil {
			rmin, rmax := dfs(root.Right)
			curMin, curMax = min(curMin, rmin), max(curMax, rmax)
			ret = max(ret, max(abs(root.Val-rmax), abs(root.Val-rmin)))
		}

		return curMin, curMax
		
	}
	dfs(input)

	return ret
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：5.86MB，击败 94.20%
*/
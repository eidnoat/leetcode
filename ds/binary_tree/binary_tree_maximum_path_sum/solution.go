/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(input *TreeNode) int {
    var (
		ret int
		dfs func(root *TreeNode) (int, int)
	)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		lv, rv := root.Val+max(0, max(dfs(root.Left))), root.Val+max(0, max(dfs(root.Right)))
		ret = max(ret, lv+rv-root.Val)

		return lv, rv
	}
	dfs(input)

	return ret
}

/*
	执行用时分布：0ms，击败100.00%
	消耗内存分布：10.15MB，击败76.19%
*/

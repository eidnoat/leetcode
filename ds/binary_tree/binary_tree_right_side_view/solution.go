/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(input *TreeNode) []int {
    var (
		ret []int
		dfs func(root *TreeNode, level int)

		maxlevel = -1
	)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level > maxlevel {
			ret = append(ret, root.Val)
			maxlevel = level
		}
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(input)

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.05MB，击败 99.82%
*/

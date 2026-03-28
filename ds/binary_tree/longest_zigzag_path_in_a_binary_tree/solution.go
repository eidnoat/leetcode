/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(input *TreeNode) int {
    var (
		ret int
		dfs func(root *TreeNode) (int, int)
	)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		_, lr := dfs(root.Left)
		rl, _ := dfs(root.Right)

		curL, curR := 1+lr, 1+rl
		ret = max(ret, max(curL, curR))

		return curL, curR
	}
	dfs(input)

	return ret-1
}

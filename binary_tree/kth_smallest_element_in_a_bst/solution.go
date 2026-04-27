/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var (
		dfs func(node *TreeNode)
		cnt int
		ret = -1
	)
	dfs = func(node *TreeNode) {
		if node == nil || ret != -1 {
			return
		}

		dfs(node.Left)
		if cnt += 1; cnt == k {
			ret = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)

	return ret
}

/*	
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：9.02MB，击败 24.14%
*/

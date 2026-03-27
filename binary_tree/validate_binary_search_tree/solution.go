/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(input *TreeNode) bool {
    var dfs func(root *TreeNode) (bool, int int)
	dfs = func(root *TreeNode) (bool, int int) {
		if root == nil {
			return true, 0, 0
		}

		lok, lmin, lmax := dfs(root.Left)
		rok, rmin, rmax := dfs(root.Right)
		
		ok := (lok && rok) && (root.Left == nil || lmax < root.Val) && (root.Right == nil || rmin > root.Val)
		min, max := lmin, rmax
		if root.Left == nil {
			min = root.Val
		}
		if root.Right == nil {
			max = root.Val
		}

		return ok, min, max
	}
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：6.96MB，击败 71.90%
*/

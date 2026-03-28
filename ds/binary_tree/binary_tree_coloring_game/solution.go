/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(input *TreeNode, n int, x int) bool {
    var (
		r1   *TreeNode
		cnt  func(root *TreeNode) int
		find func(root *TreeNode)
	)
	find = func(root *TreeNode) {
		if root == nil || r1 != nil {
			return
		}
		if root.Val == x {
			r1 = root
			return
		}

		find(root.Left)
		find(root.Right)
	}
	find(input)

	cnt = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1+cnt(root.Left)+cnt(root.Right)
	}

	if lc, rc := cnt(r1.Left), cnt(r1.Right); lc>n/2 || rc>n/2 || (n-lc-rc-1)>n/2 {
		return true
	}
	
	return false
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.37MB，击败 100.00%
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    var dfs func(node *TreeNode, sum int) bool
	dfs = func(node *TreeNode, sum int) bool {
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			return sum >= limit
		}

		var lok, rok bool
		if node.Left != nil {
			lok = dfs(node.Left, sum)
			if !lok {
				node.Left = nil
			}
		}
		if node.Right != nil {
			rok = dfs(node.Right, sum)
			if !rok {
				node.Right = nil
			}
		}

		return lok || rok
	}
	if !dfs(root, 0) {
		return nil
	}

	return root
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：8.59MB，击败 7.69%
*/

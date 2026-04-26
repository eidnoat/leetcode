/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    var dfs func(node1, node2 *TreeNode)
	dfs = func(node1, node2 *TreeNode) {
		if node2 == nil {
			return
		}

		node1.Val += node2.Val
		if node2.Left != nil && node1.Left == nil {
			node1.Left = &TreeNode{}
		}
		if node2.Right != nil && node1.Right == nil {
			node1.Right = &TreeNode{}
		}
		dfs(node1.Left, node2.Left)
		dfs(node1.Right, node2.Right)
	}
	if root1 == nil {
		root1, root2 = root2, root1
	}

	dfs(root1, root2)
	return root1
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：8.38MB，击败 87.57%
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    var (
		dfs  func(node, parent *TreeNode)
		hash = make(map[int]bool)
		ret  []*TreeNode
	)
	for _, v := range to_delete {
		hash[v] = true
	}
	dfs = func(node, parent *TreeNode) {
		if node == nil {
			return
		}

		if hash[node.Val] {
			if node == parent.Left {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		} else {
			if hash[parent.Val] {
				ret = append(ret, node)
			}
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	tmp := &TreeNode{Val: -1, Left: root}
	hash[-1] = true
	dfs(root, tmp)

	return ret
}

/*
	执行用时分布：9ms，击败 24.53%
	消耗内存分布：8.14MB，击败 41.51%
*/

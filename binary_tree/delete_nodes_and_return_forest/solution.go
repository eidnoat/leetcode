/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(tree *TreeNode, to_delete []int) []*TreeNode {
	if len(to_delete) == 0 {
		return []*TreeNode{tree}
	}

    var (
		dfs func(root, parent *TreeNode)
		ret []*TreeNode

		hash = make(map[int]bool)
	)
	for _, v := range to_delete {
		hash[v] = true
	}

	dfs = func(root, parent *TreeNode) {
		if root == nil {
			return
		}

		if hash[root.Val] {
			if root == parent.Left {
				parent.Left = nil
			} else {
				parent.Right =nil
			}
		} else {
			if hash[parent.Val] {
				ret = append(ret, root)
			}
		}

		dfs(root.Left, root)
		dfs(root.Right, root)
	}

	tmp := &TreeNode{Val: to_delete[0], Left: tree}
	dfs(tree, tmp)

	return ret
}

/*
	执行用时分布：6ms，击败 64.81%
	消耗内存分布：8.13MB，击败 55.56%
*/

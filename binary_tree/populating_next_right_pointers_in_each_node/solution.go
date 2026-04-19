/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var dfs func(node1, node2 *Node)
	dfs = func(node1, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}

		node1.Next = node2
		dfs(node1.Left, node1.Right)
		dfs(node1.Right, node2.Left)
		dfs(node2.Left, node2.Right)
	}
	dfs(root.Left, root.Right)

	return root
}

/*
	执行用时分布：9ms，击败 9.70%
	消耗内存分布：7.79MB，击败 66.04%
*/

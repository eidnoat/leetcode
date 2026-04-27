/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
    var (
		dfs1, dfs2   func(node *TreeNode, query int)
		lv, rv       int
		ret          [][]int
	)
	dfs1 = func(node *TreeNode, query int) {
		if node == nil {
			return
		}
		if node.Val < query {
			lv = max(lv, node.Val)
			dfs1(node.Right, query)
		} else if node.Val == query {
			lv = node.Val
			return
		} else {
			dfs1(node.Left, query)
		}
	}
	dfs2 = func(node *TreeNode, query int) {
		if node == nil {
			return
		}
		if node.Val < query {
			dfs2(node.Right, query)
		} else if node.Val == query {
			rv = node.Val
			return
		} else {
			rv = min(rv, node.Val)
			dfs2(node.Left, query)
		}
	}

	for _, query := range queries {
		lv, rv = math.MinInt, math.MaxInt
		
		if dfs1(root, query); lv == math.MinInt {
			lv = -1
		}
		if dfs2(root, query); rv == math.MaxInt {
			rv = -1
		}
		ret = append(ret, []int{lv, rv})
	}
	return ret
}

/*
	可以通过基本测试用例，但是超时
*/

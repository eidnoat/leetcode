/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal(root *TreeNode) [][]int {
    var (
		dfs     func(node *TreeNode, h, l int)
		ret     [][]int
		entries []*Entry
	)
	dfs = func(node *TreeNode, h, l int) {
		if node == nil {
			return
		}
		entries = append(entries, &Entry{h, l, node.Val})

		dfs(node.Left, h+1, l-1)
		dfs(node.Right, h+1, l+1)
	}
	dfs(root, 0, 0)
		
	slices.SortFunc(entries, func(e1, e2 *Entry) int {
		if e1.l < e2.l || (e1.l == e2.l && e1.h < e2.h) || (e1.l == e2.l && e1.h == e2.h && e1.v < e2.v) {
			return -1
		}

		return 1
	})

	for i, e := range entries {
		if i == 0 {
			ret = append(ret, []int{e.v})
			continue
		}

		lastE := entries[i-1]
		if e.l == lastE.l {
			ret[len(ret)-1] = append(ret[len(ret)-1], e.v)
		} else {
			ret = append(ret, []int{e.v})
		}
	}

	return ret
}

type Entry struct {
	h, l, v int
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.57MB，击败 81.82%
*/

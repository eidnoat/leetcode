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
		dfs func(node *TreeNode)
		arr []int
		ret [][]int
	)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	for _, q := range queries {
		var l, r int
		p := sort.Search(len(arr), func(i int) bool {
			return arr[i] >= q
		})
		if p == 0 {
			if arr[0] == q {
				l, r = arr[0], arr[0]
			} else {
				l, r = -1, arr[0]
			}
		} else if p == len(arr) {
			l, r = arr[len(arr)-1], -1
		} else {
			if arr[p] == q {
				l, r = arr[p], arr[p]
			} else {
				l, r = arr[p-1], arr[p]
			}
		}

		ret = append(ret, []int{l, r})
	}

	return ret
}

/*
	执行用时分布：118ms，击败 6.06%
	消耗内存分布：48.98MB，击败 9.09%
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(tree *TreeNode) *TreeNode {
	list, root, levels, cnt := []*TreeNode{tree}, &TreeNode{Val: tree.Val}, []*TreeNode{}, 0
	if tree.Left == nil {
		return root
	}

	for len(list) > 0 {
		list2 := make([]*TreeNode, 0)
		for len(list) > 0 {
			head := list[0]
			list = list[1:]

			if head.Left != nil {
				list2 = append(list2, head.Left)
			}
			if head.Right != nil {
				list2 = append(list2, head.Right)
			}
		}

		list, cnt = list2, cnt+1
		if cnt%2 == 1 {
			tmp := slices.Clone(list)
			slices.Reverse(tmp)
			levels = append(levels, tmp...)
		} else {
			levels = append(levels, slices.Clone(list)...)
		}

	}

	arr := []*TreeNode{root}
	for len(levels) > 0 {
		head := levels[0]
		levels = levels[1:]

		if arr[0].Left == nil {
			ln := &TreeNode{Val: head.Val}
			arr[0].Left = ln

			arr = append(arr, ln)
		} else if arr[0].Right == nil {
			rn := &TreeNode{Val: head.Val}
			arr[0].Right = rn

			arr = append(arr, rn)

			arr = arr[1:]
		}
	}

	return root
}

/*
	执行用时分布：17ms，击败 6.06%
	消耗内存分布：10.30MB，击败 69.70%
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var preLeftNode, leftNode, rightNode, postRightNode *ListNode
    for i, tmp := 1, head; i <= right; i, tmp = i+1, tmp.Next {
		if i == left-1 {
			preLeftNode = tmp
		}
		if i == left {
			leftNode = tmp
		}
		if i == right {
			rightNode = tmp
			postRightNode = tmp.Next
		}
	}

	if preLeftNode != nil {
		preLeftNode.Next = nil
	}
	rightNode.Next = nil
	newLeftNode, newRightNode := reverse1(leftNode)
	if preLeftNode != nil {
		preLeftNode.Next = newLeftNode
	} else {
		head = newLeftNode
	}
	newRightNode.Next = postRightNode

	return head
}

func reverse2(head *ListNode) (*ListNode, *ListNode) {
	var dfs func(pre, cur *ListNode) *ListNode

	dfs = func(pre, cur *ListNode) *ListNode {
		if cur == nil {
			return pre
		}

		next := cur.Next
		cur.Next = pre
		return dfs(cur, next)
	}

	return dfs(nil, head), head
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.96MB，击败 20.99%
*/

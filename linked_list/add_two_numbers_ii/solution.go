/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret, l1, l2, sign := &ListNode{}, reverse(l1), reverse(l2), 0
	for tmp := ret; l1 != nil || l2 != nil || sign > 0; tmp = tmp.Next {
		v1, v2 := 0, 0
		if l1 != nil {
			l1, v1 = l1.Next, l1.Val
		}
		if l2 != nil {
			l2, v2 = l2.Next, l2.Val
		}
		
		tmp.Next, sign = &ListNode{Val: (v1+v2+sign)%10}, (v1+v2+sign)/10
	}

	return reverse(ret.Next)
}

func reverse(head *ListNode) *ListNode {
	var dfs func(pre, cur *ListNode) *ListNode
	dfs = func(pre, cur *ListNode) *ListNode {
		if cur == nil {
			return pre
		}

		next := cur.Next
		cur.Next = pre
		return dfs(cur, next)
	}

	return dfs(nil, head)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：6.08MB，击败 46.25%
*/

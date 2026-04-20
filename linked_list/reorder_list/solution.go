/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}

	cnt, head2 := cnt(head), (*ListNode)(nil)
	for i, tmp := 0, head; ; i, tmp = i+1, tmp.Next {
		if (cnt % 2 == 0 && i == cnt/2-1) || (cnt % 2 == 1 && i == cnt/2) {
			head2 = tmp.Next
			tmp.Next = nil
			break
		}
	}
	
	head2 = reverse(head2)

	for tmp := head; tmp != nil && head2 != nil; {
		newTmp, newHead2 := tmp.Next, head2.Next
		tmp.Next, head2.Next = head2, tmp.Next
		tmp, head2 = newTmp, newHead2
	}

	return
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, cur, next := (*ListNode)(nil), head, head.Next
	for cur != nil {
		cur.Next = prev
		prev, cur = cur, next
		if next != nil {
			next = next.Next
		}
	}

	return prev
}

func cnt(head *ListNode) int {
	i := 0
	for tmp := head; tmp != nil; i, tmp = i+1, tmp.Next {
	}

	return i
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：7.25MB，击败 62.33%
*/

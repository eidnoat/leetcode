/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    prev, cur, ret := (*ListNode)(nil), head, (*ListNode)(nil)
	for cur != nil {
		tail := cur
		for i := 1; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		next := tail.Next

		newHead, newTail := reverse(cur, tail)

		if prev != nil {
			prev.Next = newHead
		}
		newTail.Next = next

		prev, cur = newTail, next

		if ret == nil {
			ret = newHead
		}
	}
	if ret == nil {
		ret = head
	}

	return ret
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	if head == nil || head == tail {
		return head, head
	}

	prev, cur, next := (*ListNode)(nil), head, head.Next
	for prev != tail {
		cur.Next = prev

		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		} else {
			next = nil
		}
	}

	return tail, head
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：5.16MB，击败 97.75%
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
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

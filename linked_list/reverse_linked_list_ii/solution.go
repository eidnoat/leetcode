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

func reverse1(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	pre, cur, next := head, head.Next, head.Next.Next
	pre.Next = nil
	for cur != nil {
		pre, cur.Next, cur = cur, pre, next
		if cur != nil {
			next = cur.Next
		}
	}

	return pre, head
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.90MB，击败 39.60%
*/
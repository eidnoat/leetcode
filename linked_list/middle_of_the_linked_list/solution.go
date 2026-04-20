/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	cnt := 0
	for tmp := head; tmp != nil; cnt, tmp = cnt+1, tmp.Next {
	}
	cnt--

	var ret *ListNode
	for i, tmp := 0, head; tmp != nil; i, tmp = i+1, tmp.Next {
		if cnt % 2 == 0 && i == cnt/2 {
			ret = tmp
		} else if cnt % 2 == 1 && i == cnt/2+1 {
			ret = tmp
		}
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.80MB，击败 74.36%
*/

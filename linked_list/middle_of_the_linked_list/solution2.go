/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var cnt int
    for tmp := head; tmp != nil; tmp = tmp.Next {
		cnt++
	}

	var ret *ListNode
	for i, tmp := 0, head; i < cnt; i, tmp = i+1, tmp.Next {
		if (cnt % 2 == 0 && i == cnt/2) || (cnt % 2 == 1 && i == cnt/2+1) {
			ret = tmp
			break
		}
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.80MB，击败 66.92%
*/

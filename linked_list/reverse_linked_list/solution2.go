/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var dfs func(prev, cur *ListNode) *ListNode
	dfs = func(prev, cur *ListNode) *ListNode {
		if cur == nil {
			return prev
		}

		cur.Next, prev, cur = prev, cur, cur.Next
		return dfs(prev, cur)
	}
	return dfs(nil, head)
}

/*
	执行用时分布：3ms，击败 0.56%
	消耗内存分布：4.77MB，击败 5.18%
*/

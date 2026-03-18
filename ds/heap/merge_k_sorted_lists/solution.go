/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    h := &NodeHeap{}
	for _, node := range lists {
		heap.Push(h, node)
	}

	ret := &ListNode{}
	cur := ret
	for ; h.Len() > 0; cur = cur.Next {
		top := heap.Pop(h).(*ListNode)
		cur.Next = &ListNode{Val: top.Val}
		if top.Next != nil {
			heap.Push(h, top.Next)
		}
	}

	return ret.Next
}

type NodeHeap struct {
	Nodes []*ListNode
}

func (h *NodeHeap) Push(v any) {
	h.Nodes = append(h.Nodes, v.(*ListNode))
}

func (h *NodeHeap) Pop() any {
	v := h.Nodes[len(h.Nodes)-1]
	h.Nodes = h.Nodes[:len(h.Nodes)-1]
	return v
}

func (h *NodeHeap) Len() int {
	return len(h.Nodes)
}

func (h *NodeHeap) Less(i, j int) bool {
	return h.Nodes[i].Val < h.Nodes[j].Val
}

func (h *NodeHeap) Swap(i, j int) {
	h.Nodes[i], h.Nodes[j] = h.Nodes[j], h.Nodes[i]
}

/*
	执行用时分布：4ms，击败 38.49%
	消耗内存分布：6.90MB，击败 11.84%
*/

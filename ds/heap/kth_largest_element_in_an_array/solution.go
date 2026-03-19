func findKthLargest(nums []int, k int) int {
    h := &IntHeap{}
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return h.IntSlice[0]
}

type IntHeap struct {
	sort.IntSlice
}

func (h *IntHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *IntHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

/*
	执行用时分布：50ms，击败 29.17%
	消耗内存分布：10.61MB，击败 13.25%
*/
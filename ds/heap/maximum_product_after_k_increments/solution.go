const MAX_VALUE = 1000000000 + 7

func maximumProduct(nums []int, k int) int {
    h := &IntHeap{IntSlice: nums}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, heap.Pop(h).(int)+1)
	}

	ret := 1
	for _, v := range nums {
		ret = (ret * v) % MAX_VALUE
	}

	return ret
}

type IntHeap struct{ sort.IntSlice }

func (h *IntHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *IntHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
func (h *IntHeap) Top() any {
	return h.IntSlice[0]
}

/*
	执行用时分布：176ms，击败 76.19%
	消耗内存分布：10.45MB，击败 28.57%
*/
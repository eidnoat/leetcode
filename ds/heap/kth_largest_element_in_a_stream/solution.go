type KthLargest struct {
	h *IntHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
	ret := KthLargest{h: &IntHeap{}, k: k}
	for _, num := range nums {
		ret.Add(num)
	}
	return ret
}


func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	return this.h.Top().(int)
}

type IntHeap struct{ sort.IntSlice }

func (h *IntHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *IntHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
func (h *IntHeap) Top() any { return h.IntSlice[0] }


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

/*
	执行用时分布：13ms，击败 51.46%
	消耗内存分布：10.92MB，击败 77.67%

*/
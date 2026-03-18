func lastStoneWeight(stones []int) int {
	h := &IntHeap{IntSlice: stones}
	heap.Init(h)
	for h.Len() > 1 {
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x == y {
			continue
		}
		heap.Push(h, abs(x-y))
	}
	if h.Len() == 0 {
		return 0
	}

	return h.IntSlice[0]
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

type IntHeap struct{ sort.IntSlice }

func (h *IntHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *IntHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
func (h *IntHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.78MB，击败 58.59%
*/


func resultsArray(queries [][]int, k int) []int {
	h, ret := &IntHeap{}, make([]int, 0)
	for _, query := range queries {
		heap.Push(h, dis(query[0], query[1]))
		if h.Len() < k {
			ret = append(ret, -1)
			continue
		} else if h.Len() > k {
			heap.Pop(h)
		}

		ret = append(ret, h.Top().(int))
	}

	return ret
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

func (h *IntHeap) Top() any {
	return h.IntSlice[0]
}

func (h *IntHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func dis(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	return x + y
}

/*
	执行用时分布：97ms，击败 63.16%
	消耗内存分布：44.85MB，击败 10.53%
*/ 
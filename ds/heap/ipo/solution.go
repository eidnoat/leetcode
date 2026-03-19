func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// capitalHeap: 小根堆，每次获取最小 capital 的项目 id
	// profitHeap: 大根堆，每次获取净利润最大的项目 id

	capitalHeap := &EntryHeap{lessFunc: func(entry1, entry2 *Entry) bool { return entry1.val < entry2.val }}
	profitHeap := &EntryHeap{lessFunc: func(entry1, entry2 *Entry) bool { return entry1.val > entry2.val }}
	for i, c := range capital {
		heap.Push(capitalHeap, &Entry{id: i, val: c})
	}

	for i := 0; i < k; i++ {
		for capitalHeap.Len() > 0 && capitalHeap.Top().(*Entry).val <= w {
			id := heap.Pop(capitalHeap).(*Entry).id
			heap.Push(profitHeap, &Entry{id: id, val: profits[id]})

		}
		if profitHeap.Len() == 0 {
			break
		}
		w += profits[heap.Pop(profitHeap).(*Entry).id]
	}

	return w
}

type Entry struct {
	id  int
	val int
}

type EntryHeap struct {
	entries  []*Entry
	lessFunc func(entry1, entry2 *Entry) bool
}

func (h *EntryHeap) Len() int {
	return len(h.entries)
}

func (h *EntryHeap) Less(i, j int) bool {
	return h.lessFunc(h.entries[i], h.entries[j])
}

func (h *EntryHeap) Swap(i, j int) {
	h.entries[i], h.entries[j] = h.entries[j], h.entries[i]
}

func (h *EntryHeap) Push(v any) {
	h.entries = append(h.entries, v.(*Entry))
}

func (h *EntryHeap) Pop() any {
	v := h.entries[len(h.entries)-1]
	h.entries = h.entries[:len(h.entries)-1]
	return v
}

func (h *EntryHeap) Top() any {
	return h.entries[0]
}

/*
	执行用时分布：218ms，击败 5.84%
	消耗内存分布：16.12MB，击败 18.25%
*/

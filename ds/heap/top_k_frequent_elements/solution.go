func topKFrequent(nums []int, k int) []int {
    hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}

	entries := make([]*Entry, 0)
	for num, freq := range hash {
		entries = append(entries, &Entry{num: num, freq: freq})
	}
	
	h := &EntryHeap{}
	for _, entry := range entries {
		heap.Push(h, entry)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret := make([]int, 0)
	for h.Len() > 0 {
		ret = append(ret, heap.Pop(h).(*Entry).num)
	}

	return ret
}

type Entry struct {
	num, freq int
}

type EntryHeap struct {
	entries []*Entry
}

func (h *EntryHeap) Len() int {
	return len(h.entries)
}

func (h *EntryHeap) Less(i, j int) bool {
	return h.entries[i].freq < h.entries[j].freq
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

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：7.70MB，击败 25.92%
*/
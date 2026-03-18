var arr1, arr2 []int

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	arr1, arr2 = nums1, nums2

	ret, h := make([][]int, 0), &EntryHeap{}
	for i := 0; i < len(arr1); i++ {
		heap.Push(h, &Entry{i, 0})
	}

	for cnt := 0; cnt < k; cnt++ {
		top := heap.Pop(h).(*Entry)
		ret = append(ret, []int{arr1[top.i], arr2[top.j]})
		if top.j+1 < len(arr2) {
			heap.Push(h, &Entry{top.i, top.j + 1})
		}
	}

	return ret
}

type Entry struct {
	i, j int
}

type EntryHeap struct {
	entries []*Entry
}

func (h *EntryHeap) Push(v any) { h.entries = append(h.entries, v.(*Entry)) }
func (h *EntryHeap) Pop() any {
	v := h.entries[len(h.entries)-1]
	h.entries = h.entries[:len(h.entries)-1]
	return v
}
func (h *EntryHeap) Less(i, j int) bool {
	return arr1[h.entries[i].i]+arr2[h.entries[i].j] < arr1[h.entries[j].i]+arr2[h.entries[j].j]
}
func (h *EntryHeap) Len() int      { return len(h.entries) }
func (h *EntryHeap) Swap(i, j int) { h.entries[i], h.entries[j] = h.entries[j], h.entries[i] }
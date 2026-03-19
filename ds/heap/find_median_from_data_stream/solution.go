type MedianFinder struct {
    minHeap, maxHeap *IntHeap
}


func Constructor() MedianFinder {
    return MedianFinder{
		minHeap: &IntHeap{LessFunc: func(v1, v2 int) bool { return v1 < v2 }},
		maxHeap: &IntHeap{LessFunc: func(v1, v2 int) bool { return v1 > v2 }},
	}
}


func (f *MedianFinder) AddNum(num int)  {
	if  f.minHeap.Len() == 0 && f.maxHeap.Len() == 0 {
		heap.Push(f.maxHeap, num)
		return
	}
 
	if num <= f.maxHeap.Top() {
		heap.Push(f.maxHeap, num)
	} else {
		heap.Push(f.minHeap, num)
	}
	
	for f.maxHeap.Len() - f.minHeap.Len() > 1 {
		heap.Push(f.minHeap, heap.Pop(f.maxHeap).(int))
	}

	for f.maxHeap.Len() < f.minHeap.Len() {
		heap.Push(f.maxHeap, heap.Pop(f.minHeap).(int))
	}
}


func (f *MedianFinder) FindMedian() float64 {
	var ret float64

    if f.minHeap.Len() == f.maxHeap.Len() {
		ret = float64(f.maxHeap.Top()+f.minHeap.Top()) / 2.0
	} else {
		ret = float64(f.maxHeap.Top())
	}

	return ret
}

type IntHeap struct {
	sort.IntSlice
	LessFunc func(v1, v2 int ) bool
}

func (h *IntHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *IntHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func (h *IntHeap) Less(i, j int) bool {
	return h.LessFunc(h.IntSlice[i], h.IntSlice[j])
}

func (h *IntHeap) Top() int {
	return h.IntSlice[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

 /*
	执行用时分布：108ms，击败 75.66%
	消耗内存分布：22.66MB，击败 69.67%
 */
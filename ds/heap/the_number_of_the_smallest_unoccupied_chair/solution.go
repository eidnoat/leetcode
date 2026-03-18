func smallestChair(times [][]int, targetFriend int) int {
	targetFriendStartTime := times[targetFriend][0]

	friendArrivedHeap := &FriendHeap{LessFunc: func(i, j int) bool { return times[i][0] < times[j][0] }}
	for i := 0; i < len(times); i++ {
		heap.Push(friendArrivedHeap, i)
	}

	friendLeavedHeap := &FriendHeap{LessFunc: func(i, j int) bool { return times[i][1] < times[j][1] }}

	chairHeap := &ChairHeap{}
	for i := 0; i < len(times); i++ {
		heap.Push(chairHeap, i)
	}

	ret, friendToChair := 0, make(map[int]int)
	for curTime := 1; curTime <= targetFriendStartTime; curTime++ {
		for friendLeavedHeap.Len() > 0 && times[friendLeavedHeap.Top().(int)][1] == curTime { // 同一时间可能有多个人离开
			leavedFriendID := friendLeavedHeap.Top().(int)
			heap.Push(chairHeap, friendToChair[heap.Pop(friendLeavedHeap).(int)])
			delete(friendToChair, leavedFriendID)
		}

		if friendArrivedHeap.Len() > 0 && times[friendArrivedHeap.Top().(int)][0] == curTime {
			if curTime == targetFriendStartTime {
				ret = chairHeap.Top().(int)
			}

			arrivedFriendID := friendArrivedHeap.Top().(int)
			friendToChair[heap.Pop(friendArrivedHeap).(int)] = heap.Pop(chairHeap).(int)
			heap.Push(friendLeavedHeap, arrivedFriendID)
		}
	}

	return ret
}

type IntHeap struct{ sort.IntSlice }
type ChairHeap struct{ IntHeap }
type FriendHeap struct {
	IntHeap
	LessFunc func(i, j int) bool
}

func (h *IntHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *IntHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
func (h *IntHeap) Top() any {
	return h.IntSlice[0]
}

func (h *FriendHeap) Less(i, j int) bool {
	return h.LessFunc(h.IntSlice[i], h.IntSlice[j])
}

/*
	执行用时分布：61ms，击败 50.00%
	消耗内存分布：10.18MB，击败71.43%
*/
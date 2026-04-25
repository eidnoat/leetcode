func maxSlidingWindow(nums []int, k int) []int {
	q, ret := &MonoQueue{}, make([]int, 0)
    for i, v := range nums {
		q.Push(v)
		if i >= k {
			q.Pop()
		}
		if i >= k-1 {
			ret = append(ret, q.Extreme())
		}
	}

	return ret
}

type MonoQueue struct {
	q1, q2 []int
}

func (q *MonoQueue) Push(v int) {
	q1, q2 := q.q1, q.q2
	defer func() { q.q1, q.q2 = q1, q2 }()

	q1 = append(q1, v)
	for len(q2) > 0 && v > q2[len(q2)-1] {
		q2 = q2[:len(q2)-1]
	}
	q2 = append(q2, v)
}

func (q *MonoQueue) Pop() {
	q1, q2 := q.q1, q.q2
	defer func() { q.q1, q.q2 = q1, q2 }()

	v := q1[0]
	if q2[0] == v {
		q2 = q2[1:]
	}
	q1 = q1[1:]
}

func(q *MonoQueue) Extreme() int {
	return q.q2[0]
}

/*
	执行用时分布：34ms，击败 13.14%
	消耗内存分布：10.88MB，击败 10.44%
*/

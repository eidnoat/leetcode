func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
    l, r, sum, ret, q := 0, 0, 0, 0, &MonoQueue{}
	cost := func() int64 { return int64(q.Extreme())+int64(r-l+1)*int64(sum) }

	for ; r < len(chargeTimes); r++ {
		sum += runningCosts[r]
		q.Push(chargeTimes[r])

		for l <= r && cost() > budget {
			l, sum = l+1, sum-runningCosts[l]
			q.Pop()
		}
		ret = max(ret, r-l+1)
	}

	return ret
}

type MonoQueue struct {
	q1, q2 []int
}

func (q *MonoQueue) Push(v int) {
	q1, q2 := q.q1, q.q2
	defer func() { q.q1, q.q2 = q1, q2 }()

	for len(q2) > 0 && v > q2[len(q2)-1] {
		q2 = q2[:len(q2)-1]
	}
	q1, q2 = append(q1, v), append(q2, v)
}

func (q *MonoQueue) Pop() {
	q1, q2 := q.q1, q.q2
	defer func() { q.q1, q.q2 = q1, q2 }()
	
	if q2[0] == q1[0] {
		q2 = q2[1:]
	}
    q1 = q1[1:]
}

func (q *MonoQueue) Extreme() int {
	return q.q2[0]
}

/*
	执行用时分布：13ms，击败 35.00%
	消耗内存分布：9.20MB，击败 70.00%
*/

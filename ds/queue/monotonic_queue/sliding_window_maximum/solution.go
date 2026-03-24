func maxSlidingWindow(nums []int, k int) []int {
    q, ret := &MonoQueue{}, make([]int, 0)
	for i, v := range nums {
		q.Push(v)
		if i < k-1 {
			continue
		}
		if i >= k {
			q.Pop()
		}

		ret = append(ret, q.Extremum())
	}

	return ret
}

type MonoQueue struct {
	q1, q2 []int
}

func (q *MonoQueue) Push(v int) {
	tmp1, tmp2 := q.q1, q.q2
	defer func(){ q.q1, q.q2 = tmp1, tmp2}()

	for ; len(tmp2) > 0 && v > tmp2[len(tmp2)-1]; tmp2 = tmp2[:len(tmp2)-1] {
	}
	tmp1, tmp2 = append(tmp1, v), append(tmp2, v)
}

func (q *MonoQueue) Pop() {
	tmp1, tmp2 := q.q1, q.q2
	defer func(){ q.q1, q.q2 = tmp1, tmp2}()

	if tmp2[0] == tmp1[0] {
		tmp2 = tmp2[1:]
	}
	tmp1 = tmp1[1:]
}

func (q *MonoQueue) Extremum() int {
	return q.q2[0]
}

/*
	执行用时分布：38ms，击败 13.31%
	消耗内存分布：11.05MB，击败 8.40%
*/

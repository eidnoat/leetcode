func minimumSum(nums []int) int {
    q := &MonoQueue{}
	for i := 2; i < len(nums); i++ {
		q.Push(nums[i])
	}
	
	ret := math.MaxInt

	lMin := nums[0]
	for j := 1; j < len(nums)-1; j++ {
		rMin := q.Min()

		if lMin < nums[j] && rMin < nums[j] {
			ret = min(ret, lMin+nums[j]+rMin)
		}
		
		lMin = min(lMin, nums[j])
		q.Pop()
	}
	
	if ret == math.MaxInt {
		return -1
	}
	return ret
}

type MonoQueue struct {
	q1, q2 []int
}

func (q *MonoQueue) Push(v int) {
	q.q1 = append(q.q1, v)
	for len(q.q2) > 0 && v < q.q2[len(q.q2)-1] {
		q.q2 = q.q2[:len(q.q2)-1]
	}
	q.q2 = append(q.q2, v)
}

func (q *MonoQueue) Pop() int {
	v := q.q1[0]
	q.q1 = q.q1[1:]
	if v == q.q2[0] {
		q.q2 = q.q2[1:]
	}

	return v
}

func (q *MonoQueue) Min() int {
	return q.q2[0]
}

/*
	执行用时分布：1ms，击败 67.12%
	消耗内存分布：10.63MB，击败 36.98%
*/

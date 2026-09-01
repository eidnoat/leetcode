func carPooling(trips [][]int, capacity int) bool {
	m := 0
	for _, trip := range trips {
		m = max(m, trip[2])
	}

	diff := make([]int, m+1)
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}

	cur := 0
	for _, v := range diff {
		cur += v
		if cur > capacity {
			return false
		}
	}

	return true
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：5.18MB，击败 31.82%
*/
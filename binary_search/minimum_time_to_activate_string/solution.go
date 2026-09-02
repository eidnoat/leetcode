func minTime(s string, order []int, k int) int {
	check := func(t int) bool {
		dst := make([]int, t+1)
		copy(dst, order)
		sort.Slice(dst, func(i, j int) bool { return dst[i] < dst[j] })

		cnt := 0
		for i := range dst {
			l, r := 0, len(s)-1
			if i > 0 {
				l = dst[i-1] + 1
			}
			cnt += (dst[i] - l + 1) * (r - dst[i] + 1)
		}

		return cnt >= k
	}

	l, r, ret := 0, len(s)-1, -1
	for l <= r {
		m := (l + r) / 2
		if check(m) {
			if m-1 >= l && check(m-1) {
				r = m - 1
			} else {
				ret = m
				break
			}
		} else {
			l = m + 1
		}
	}

	return ret
}

/*
	执行用时分布：191ms，击败 -%
	消耗内存分布：10.99MB，击败 20.00%
*/

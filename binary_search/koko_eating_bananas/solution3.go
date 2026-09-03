func minEatingSpeed(piles []int, h int) int {
	check := func(k int) bool {
		cnt := 0
		for _, pile := range piles {
			cnt += pile / k
			if pile % k > 0 {
				cnt++
			}
		}

		return cnt <= h
	}

	l, r, ret := 1, slices.Max(piles), 0
	for l <= r {
		m := (l+r)/2
		if check(m) {
			if m-1 >= l && check(m-1) {
				r = m-1
			} else {
				ret = m
				break
			}
		} else {
			l = m+1
		}
	}

	return ret
}

/*
	执行用时分布：15ms，击败 15.90%
	消耗内存分布：8.07MB，击败 23.08%
*/

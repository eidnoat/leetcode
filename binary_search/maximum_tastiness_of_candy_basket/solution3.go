func maximumTastiness(price []int, k int) int {
	sort.Slice(price, func(i, j int) bool { return price[i] < price[j] })

	check := func(diff int) bool {
		p, cnt := 0, 1
		for cnt < k {
			newP := sort.Search(len(price), func(i int) bool { return price[i]-price[p] >= diff})
			if newP == len(price) {
				break
			}

			p, cnt = newP, cnt+1
		}

		return cnt == k
	}

	l, r, ret := 0, price[len(price)-1]-price[0], 0
	for l <= r {
		m := (l+r)/2
		if check(m) {
			if m+1 <= r && check(m+1) {
				l = m+1
			} else {
				ret = m
				break
			}		
		} else {
			r = m-1
		}
	}

	return ret
}

/*
	执行用时分布：479ms，击败 5.26%
	消耗内存分布：9.29MB，击败 94.74%
*/

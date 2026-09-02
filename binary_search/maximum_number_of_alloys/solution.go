func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	check := func(cnt, no int) bool {
		need := 0
		for i := 0; i < n; i++ {
			need += max(0, composition[no][i]*cnt-stock[i]) * cost[i]
		}

		return need <= budget
	}

	ret := 0
	for i := 0; i < k; i++ {
		l, r, cur := 0, int(1e10), 0
		for l <= r {
			m := (l + r) / 2
			if check(m, i) {
				if m+1 <= r && check(m+1, i) {
					l = m + 1
				} else {
					cur = m
					break
				}
			} else {
				r = m - 1
			}
		}
		ret = max(ret, cur)
	}

	return ret
}

/*
	执行用时分布：17ms，击败 7.41%
	消耗内存分布：8.25MB，击败 11.11%
*/
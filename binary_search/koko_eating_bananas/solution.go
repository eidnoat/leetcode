func minEatingSpeed(piles []int, h int) int {
	check := func(v int) bool {
		h := h
		for i, pile := range piles {
			if (pile % v) == 0 {
				h -= pile / v
			} else {
				h -= pile/v + 1
			}

			if h < 0 {
				return false
			}
			if h == 0 {
				if i == len(piles)-1 {
					break
				} else {
					return false
				}
			}
		}

		return true
	}

	l, r := 1, slices.Max(piles)
	for l <= r {
		m := (l + r) / 2
		if check(m) {
			if m-1 >= l && check(m-1) {
				r = m - 1
			} else {
				return m
			}
		} else {
			l = m + 1
		}
	}

	return 0
}

/*
	执行用时分布：19ms，击败 5.39%
	消耗内存分布：8.07MB，击败 19.92%
*/
func repairCars(ranks []int, cars int) int64 {
	// 给定能力值 rank 和时间 t，设着期间修复的车辆数为 n
	// t = rank * n * n => n = (t/rank)^(1/2)

	cnt := func(t int) int {
		ret := 0
		for _, rank := range ranks {
			ret += int(math.Sqrt(float64(t)/float64(rank)))
		}
		return ret
	}

    l, r := 1, slices.Min(ranks)*cars*cars
	for l <= r {
		m := (l+r)/2
		if cnt(m) >= cars {
			if m-1 >= l && cnt(m-1) >= cars {
				r = m-1
			} else {
				return int64(m)
			}
		} else {
			l = m+1
		}
	}

	return 0
}

/*
	执行用时分布：51ms，击败 5.71%
	消耗内存分布：9.30MB，击败 51.43%
*/
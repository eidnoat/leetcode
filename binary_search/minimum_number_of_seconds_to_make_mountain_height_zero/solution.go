func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
    // 设 workerTimes[i] 移除 n 米的山需要 x 秒
	// workerTimes[i] + workerTimes[i] * 2 + ... + workerTimes[i] * n = x
	// 令 workerTimes[i] = t
	// t * (1+2+...+n) = x => t*n*(n+1)/2 = x => n*(n+1) = (2*x)/t => n*n+n+1/4 = (2*x)/t+1/4
	// => (n+1/2)^2 = (2*x)/t+1/4 => n+1/2 = (2x/t+0.25)^0.5 => n = (2x/t+0.25)^0.5 - 0.5
	// => 给定时间 x，workerTimes[i] 可以移除 (2x/t+0.25)^0.5 - 0.5 米的山

	moveHeight := func(t, workerTime int) int {
		tf, wf := float64(t), float64(workerTime)
		return int(math.Floor(math.Sqrt(2*tf/wf+0.25) - 0.5))
	}
	moveHeightTotal := func(t int) int {
		var moveHeightTotal int
		for _, workerTime := range workerTimes {
			moveHeightTotal += moveHeight(t, workerTime)
		}

		return moveHeightTotal
	}

	costTime := func(workerTime int) int {
		return workerTime * mountainHeight * (mountainHeight + 1) / 2
	}

	l, r := 1, costTime(slices.Min(workerTimes))
	for l <= r {
		m := (l + r) / 2
		if ht := moveHeightTotal(m); ht < mountainHeight {
			l = m + 1
		} else {
			if m-1 >= l && moveHeightTotal(m-1) >= mountainHeight {
				r = m - 1
			} else {
				return int64(m)
			}
		}
	}

	return 0
}

/*
	执行用时分布：30ms，击败47.96%
	消耗内存分布：7.54MB，击败55.86%
*/
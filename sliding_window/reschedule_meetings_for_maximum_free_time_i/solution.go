func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
    w, ret, tmp := make([]int, len(startTime)), 0, 0
	for i := 0; i < len(startTime); i++ {
		w[i] = endTime[i]-startTime[i]
		if i < k {
			tmp += w[i]
		}
		if i == k-1 {
			if i+1 < len(startTime) {
				ret = startTime[i+1]-tmp
			} else {
				ret = eventTime-tmp
			}
		}
	}

	for l, r := 0, k; r < len(startTime); l, r = l+1, r+1 {
		tmp += w[r]-w[l]
		le, re := endTime[l], eventTime
		if r+1 < len(startTime) {
			re = startTime[r+1]
		}
		ret = max(ret, re-le-tmp)
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：12.20MB，击败 15.56%
*/

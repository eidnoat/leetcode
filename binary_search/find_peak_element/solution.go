func findPeakElement(nums []int) int {
    l, r := 0, len(nums)-1
	for l <= r {
		m := (l+r)/2
		v, lv, rv := nums[m], math.MinInt, math.MinInt
		if m-1 >= 0 {
			lv = nums[m-1]
		}
		if m+1 <= len(nums)-1 {
			rv = nums[m+1]
		}

		if lv < v && v < rv {
			l = m+1
		} else if lv < v && v > rv {
			return m
		} else if lv > v && v > rv {
			r = m-1
		} else {
			r = m-1
		}
	}

	return -1
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.38MB，击败 74.83%
*/

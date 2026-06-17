func maxArea(height []int) int {
    l, r, ret := 0, len(height)-1, 0
	for l < r {
		ret = max(ret, (r-l) * min(height[l], height[r]))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return ret
}

/*
执行用时分布：0ms，击败 100.00%
消耗内存分布：9.04MB，击败 99.74%
*/

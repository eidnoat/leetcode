func minSubArrayLen(target int, nums []int) int {
    l, r, sum, ret := 0, 0, 0, math.MaxInt
	for ; r < len(nums); r++ {
		sum += nums[r]
		for ; sum >= target; sum, l = sum-nums[l], l+1 {
			ret = min(ret, r-l+1)
		}
	}
	if ret == math.MaxInt {
		ret = 0
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：8.97MB，击败 34.61%
*/

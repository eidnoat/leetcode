func countSubarrays(nums []int, k int64) int64 {
	l, r, sum, ret := 0, 0, 0, int64(0)
	getScore := func() int64 {
		return int64(r-l+1) * int64(sum)
	}

	for ; r < len(nums); r++ {
		for sum += nums[r]; getScore() >= k; l, sum = l+1, sum-nums[l] {
		}
		ret += int64(r-l+1)
	}

	return ret
}

/*
	执行用时分布：4ms，击败 26.76%
	消耗内存分布：9.58MB，击败 87.32%
*/

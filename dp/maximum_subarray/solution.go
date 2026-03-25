func maxSubArray(nums []int) int {
	// dp[i] = max(nums[i], dp[i-1]+nums[i])
	last, ret := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		last, ret = max(nums[i], last+nums[i]), max(ret, max(nums[i], last+nums[i]))
	}
	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：10.14MB，击败 13.43%
*/

func rob(nums []int) int {
    // dp[i] = nums[i]+max(dp[i+2]...dp[len(nums)-1])
	// dp[len(nums)-1] = nums[len(nums)-1]
	// dp[len(nums)-2] = nums[len(nums)-2]
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	v1, v2, maxV := nums[len(nums)-2], nums[len(nums)-1], nums[len(nums)-1]
	for i := len(nums)-3; i >= 0; i-- {
		dp[i] = nums[i]+maxV
        v1, v2, maxV = v, v1, max(maxV, v1)
	}
	return max(v1, v2)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.79MB，击败 87.21%
*/

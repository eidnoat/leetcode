func maxSubarraySum(nums []int, k int) int64 {
	// if i < k-1:
	// 	dp[i] = 0
	// else
    // 	dp[i] = sum[i-k+1...i]+max(0, dp[i-k])

	prefixSum, dp, ret := make([]int64, len(nums)+1), make([]int64, len(nums)), int64(math.MinInt64)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i]+int64(nums[i])
		if i < k-1 {
			continue
		} else if i == k-1 {
			dp[i] = prefixSum[i+1]
		} else if i > k-1 {
			dp[i] = prefixSum[i+1]-prefixSum[i-k+1]+max(0, dp[i-k])
		}
		ret = max(ret, dp[i])
	}

	return ret
}

/*
	执行用时分布：8ms，击败 40.63%
	消耗内存分布：16.18MB，击败 25.00%
*/
func maxSubArray(nums []int) int {
	prefixSum, lastMin, ret := make([]int, len(nums)+1), 0, math.MinInt
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i]+nums[i]
		lastMin, ret = min(lastMin, prefixSum[i+1]), max(ret, prefixSum[i+1]-lastMin)
	}

	return ret
}

/*
	执行用时分布：8ms，击败 5.63%
	消耗内存分布：10.02MB，击败 32.32%
*/

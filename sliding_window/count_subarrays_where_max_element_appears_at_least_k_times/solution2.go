func countSubarrays(nums []int, k int64) int64 {
	l, r, ret := int64(0), int64(0), int64(0)
	for sum, score := int64(0), int64(0); r < int64(len(nums)); r++ {
		sum += int64(nums[r])
		score = sum * (r-l+1)
		for ; score >= k; l++ {
			sum -= int64(nums[l])
			score = sum * (r-l)
		}
		ret += r-l+1
	}

	return ret
}

/*
	执行用时分布：1ms，击败 36.96%
	消耗内存分布：9.46MB，击败 82.61%
*/


func numberOfSubarrays(nums []int, k int) int {
	sum, hash, ret := 0, map[int]int{0:1}, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]%2
		ret += hash[sum-k]

		hash[sum]++
	} 

	return ret
}

/*
	执行用时分布：31ms，击败 13.33%
	消耗内存分布：9.27MB，击败 8.57%
*/

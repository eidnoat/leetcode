func countSubarrays(nums []int, k int) int {
	prefixSum, hash, ret, found := make([]int, len(nums)+1), map[int]int{0: 1}, 0, false
	for i := 0; i < len(nums); i++ {
		v := -1
		if nums[i] == k {
			v = 0
			found = true
		} else if nums[i] > k {
			v = 1
		}
		prefixSum[i+1] = prefixSum[i] + v

		if found {
			ret += hash[prefixSum[i+1]] + hash[prefixSum[i+1]-1]
		} else {
			hash[prefixSum[i+1]]++
		}
	}

	return ret
}

/*
	执行用时分布：10ms，击败 81.82%
	消耗内存分布：10.14MB，击败 18.18%
*/

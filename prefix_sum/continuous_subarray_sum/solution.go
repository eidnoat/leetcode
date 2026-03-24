func checkSubarraySum(nums []int, k int) bool {
    prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i]+nums[i]
	}

	hash := map[int]int{0:-1}
	for j := 0; j < len(nums); j++ {
		if i, ok := hash[prefixSum[j+1]%k]; ok && j-i >= 2 {
			return true
		}
		if _, ok := hash[prefixSum[j+1]%k]; !ok {
			hash[prefixSum[j+1]%k] = j
		}
	}

	return false
}

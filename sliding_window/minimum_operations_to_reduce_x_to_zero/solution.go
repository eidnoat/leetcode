func minOperations(nums []int, x int) int {
	prefixSums, sum := []int{0}, 0
	for _, num := range nums { // sumOf[i...j] = prefixSums[j+1]-prefixSums[i]
		sum += num
		prefixSums = append(prefixSums, sum)
	}

	target, ret, hash := sum-x, -1, make(map[int]int, 0)
	for j, prefixSum := range prefixSums {
		hash[prefixSum] = j
		if i, ok := hash[prefixSum-target]; ok {
			ret = max(ret, j-i)
		}
	}
	if ret == -1 {
		return -1
	}

	return len(nums) - ret
}

/*
	执行用时分布：87ms，击败 5.33%
	消耗内存分布：15.52MB，击败 6.00%
*/

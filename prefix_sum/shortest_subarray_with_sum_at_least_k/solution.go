func shortestSubarray(nums []int, k int) int {
    prefixSums, sum := []int{0}, 0 // sumof[i...j] = prefixSums[j+1]-prefixSum[i]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixSums = append(prefixSums, sum)
	}

	monoQ, ret := make([]int, 0), math.MaxInt
	for j, prefixSum := range prefixSums {
		p := sort.Search(len(monoQ), func(i int) bool {
			return prefixSum-prefixSums[monoQ[i]] < k
		})
		if p-1 >= 0 && prefixSum-prefixSums[monoQ[p-1]] >= k {
			ret = min(ret, j-monoQ[p-1])
		}

		for len(monoQ) > 0 && prefixSum <= prefixSums[monoQ[len(monoQ)-1]] {
			monoQ = monoQ[:len(monoQ)-1]
		}
		monoQ = append(monoQ, j)
	}
	if ret == math.MaxInt {
		ret = -1
	}

	return ret
}

/*
	执行用时分布：48ms，击败 12.66%
	消耗内存分布：9.77MB，击败 89.87%
*/
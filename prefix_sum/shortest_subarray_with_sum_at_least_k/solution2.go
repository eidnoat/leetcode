func shortestSubarray(nums []int, k int) int {
    prefixSums, sum := []int{0}, 0 // sumof[i...j] = prefixSums[j+1]-prefixSum[i]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixSums = append(prefixSums, sum)
	}

	monoQ, ret := make([]int, 0), math.MaxInt
	for j, prefixSum := range prefixSums {
		for ; len(monoQ) > 0 && prefixSum-prefixSums[monoQ[0]] >= k; ret, monoQ = min(ret, j-monoQ[0]), monoQ[1:] {
		}

		for ; len(monoQ) > 0 && prefixSum <= prefixSums[monoQ[len(monoQ)-1]]; monoQ = monoQ[:len(monoQ)-1] {
		}
		monoQ = append(monoQ, j)
	}
	if ret == math.MaxInt {
		ret = -1
	}

	return ret
}

/*
	执行用时分布：14ms，击败 35.44%
	消耗内存分布：10.07MB，击败 63.29%
*/

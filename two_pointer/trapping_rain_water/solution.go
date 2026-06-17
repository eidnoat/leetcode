func trap(height []int) int {
    records, ret := make([]int, len(height)), 0

	tempMax := 0
	for i := 0; i < len(height); i++ {
		records[i] = tempMax
		tempMax = max(tempMax, height[i])
	}

	tempMax = 0
	for i := len(nums)-1; i >= 0; i-- {
		records[i] = min(records[i], tempMax)
		tempMax = max(tempMax, height[i])

		ret += max(records[i]-height[i], 0)
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：7.58MB，击败 98.04%
*/
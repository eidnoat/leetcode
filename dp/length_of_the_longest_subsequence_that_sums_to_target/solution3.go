func lengthOfLongestSubsequence(nums []int, target int) int {
	var (
		dfs func(i, target int) int
		mem = make([][]int, len(nums)+1)
	)
	for i := 0; i < len(mem); i++ {
		mem[i] = make([]int, 1001)
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -2
		}
	}
	mem[len(nums)][0] = 0

	prefixSums := make([]int, len(nums)+1)
	for i, num := range nums {
		prefixSums[i+1] = prefixSums[i]+num
	}

	minVals := make([]int, len(nums))
	for i := len(nums)-1; i >= 0; i-- {
		if i == len(nums)-1 {
			minVals[i] = nums[i]
		} else {
			minVals[i] = min(minVals[i+1], nums[i])
		}
	}

	dfs = func(i, target int) int {
		if target == 0 {
			return 0
		}
		if i == len(nums) || target < 0 {
			return -1
		}

		if mem[i][target] != -2 {
			return mem[i][target]
		}

		var ret int
		if target < minVals[i] || target > prefixSums[len(nums)]-prefixSums[i] {
			ret = -1
		} else {
			path1, path2 := dfs(i+1, target-nums[i]), dfs(i+1, target)
			if path1 != -1 {
				path1 += 1
			}
			ret = max(path1, path2)
		}
		mem[i][target] = ret

		return mem[i][target]
	}

	return dfs(0, target)
}

/*
	执行用时分布：285ms，击败 19.23%
	消耗内存分布：19.96MB，击败 52.56%
*/

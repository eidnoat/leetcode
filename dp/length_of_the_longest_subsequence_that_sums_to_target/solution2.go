func lengthOfLongestSubsequence(nums []int, target int) int {
	var (
		dfs func(i, target int) int
		mem = map[int]map[int]int{len(nums): {0: 0}}
	)

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
		if i == len(nums) {
			return -1
		}

		if v, ok := mem[i][target]; ok {
			return v
		}
		if mem[i] == nil {
			mem[i] = make(map[int]int)
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
	可以通过基本测试用例，但整体会超时。
*/

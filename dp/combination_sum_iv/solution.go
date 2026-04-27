func combinationSum4(nums []int, target int) int {
    /**
		dp[0] = 1
		for _, v := range nums {
			dp[i] += dp[i-v]
		}
	**/

	var (
		dfs func(target int) int
		mem = map[int]int{0:1}
	)
	dfs = func(target int) int {
		if v, ok := mem[target]; ok {
			return v
		}

		ret := 0
		for _, num := range nums {
			if target-num < 0 {
				continue
			}
			ret += dfs(target-num)
		}
		mem[target] = ret

		return ret
	}

	return dfs(target)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.94MB，击败 5.15%
*/

var MAX = 1000000000+7

func countGoodStrings(low int, high int, zero int, one int) int {
    /*
		dp[0] = 1
		dp[i] = dp[i-zero]+dp[i-one]
	*/
	var (
		dfs func(target int) int
		mem = map[int]int{0:1}
		ret int
	)
	dfs = func(target int) int {
		if target < 0 {
			return 0
		}
		if _, ok := mem[target]; !ok {
			mem[target] = (dfs(target-zero)+dfs(target-one))%MAX
		}
		return mem[target]
	}
	for i := high; i >= low; i-- {
		ret += dfs(i)
		ret %= MAX
	}

	return ret
}

/*
	执行用时分布：116ms，击败 5.06%
	消耗内存分布：33.20MB，击败 5.06%
*/

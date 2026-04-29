func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
	dp[0] = 1
	
	for _, coin := range coins {
		for j := 0; j < len(dp); j++ {
			if j-coin >= 0 {
				dp[j] += dp[j-coin]
			}
		}
	}

	return dp[amount]
}

/*
	执行用时分布：1ms，击败 78.40%
	消耗内存分布：4.18MB，击败 79.00%
*/

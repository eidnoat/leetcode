func minCostClimbingStairs(cost []int) int {
    // dp[len(cost)-1] = cost[len(cost)-1]
	// dp[i] = cost[i]+min(dp[i+1], dp[i+2])
	cost = append(cost, 0)
	v1, v2 := cost[len(cost)-2], cost[len(cost)-1]
	for i := len(cost)-3; i >= 0; i-- {
		v := cost[i]+min(v1, v2)
		v1, v2 = v, v1
	}

	return min(v1, v2)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.49MB，击败 85.74%
*/

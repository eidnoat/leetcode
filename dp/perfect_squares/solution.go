func numSquares(n int) int {
	/*
		dp[i] = min(dp[j]+dp[k]) // j+k = i
	*/
    dp := make([]int, n+1)
	for i := 1; i*i < len(dp); i++ {
		dp[i*i] = 1
	}
	for i := 2; i < len(dp); i++ {
		if dp[i] != 0 {
			continue
		}
		for j, k := 1, i-1; j <= k; j, k = j+1, k-1 {
			if dp[i] == 0 {
				dp[i] = dp[j]+dp[k]
			} else {
				dp[i] = min(dp[i], dp[j]+dp[k])
			}
		}
	}

	return dp[n]
}

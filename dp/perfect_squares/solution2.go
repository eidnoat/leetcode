func numSquares(n int) int {
	/*
		dp[i] = min(dp[i-k^2]+1) // k^2 < i
	*/
    dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		for j := 1; j*j <= i; j++ {
			if dp[i] == 0 {
				dp[i] = dp[i-j*j]+1
			} else {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
	}

	return dp[n]
}

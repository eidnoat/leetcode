func maximumTastiness(price []int, k int) int {
    var (
		dfs  func(i, diff int, path []int)
		ret  int
	)

	slices.Sort(price)
	dfs = func(i, diff int, path []int) {
		if len(path) == k {
			ret = max(ret, diff)
			return
		}
		if i == len(price) {
			return
		}
		
		// not select
		dfs(i+1, diff, path)

		// select
		if len(path) >= 1 {
			diff = min(diff, price[i]-path[len(path)-1])
		}
		dfs(i+1, diff, append(path, price[i]))
	}
	dfs(0, math.MaxInt, []int{})

	return ret
}

/*
	回溯，可以通过基本的测试用例，但当 len(price) 和  k 较大时，会超时。
*/

func change(amount int, coins []int) int {
	slices.Sort(coins)

	var (
		dfs func(amount, start int) int
		mem = make([][]int, 5001)
	)
	for i := 0; i < len(mem); i++ {
		arr1 := make([]int, 300)
		for j := 0; j < len(arr1); j++ {
			arr1[i] = -1
		}
		mem = append(mem, arr1)
	}

	dfs = func(amount, start int) int {
		if start == len(coins) || amount < 0 {
			return 0
		}
		if amount == 0 {
			return 1
		}
		if mem[amount][start] != -1 {
			return mem[amount][start]
		}

		ret := 0
		for i, coin := range coins[start:] {
			ret += dfs(amount-coin, i)
		}
		mem[amount][start] = ret

		return ret
	}

	return dfs(amount, 0)
}

/*
	可以通过基本测试用例，但是会超时。
*/

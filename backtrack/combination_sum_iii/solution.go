func combinationSum3(k int, n int) [][]int {
    var (
		ret [][]int
		dfs func(i int, nums []int)
	)
	dfs = func(i int, nums []int) {
		if i > 9 {
			return
		}

		sum := 0
		for _, num := range nums {
			sum += num
		}
		if sum > n {
			return
		}

		if len(nums) == k {
			if sum == n {
				ret = append(ret, slices.Clone(nums))
			}
			return
		}

		dfs(i+1, nums) // not select
		dfs(i+1, append(nums, i)) // select
	}
	dfs(1, []int{})

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.82MB，击败 25.24%
*/

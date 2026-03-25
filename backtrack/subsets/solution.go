func subsets(nums []int) [][]int {
    var (
		ret [][]int
		dfs func(i int, set []int)
	)
	dfs = func(i int, set []int) {
		if i == len(nums) {
			ret = append(ret, slices.Clone(set))
			return
		}

		dfs(i+1, append(set, nums[i])) // select
		dfs(i+1, set) // not select
	}
	dfs(0, []int{})

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.99MB，击败 99.83%
*/

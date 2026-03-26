func combinationSum2(candidates []int, target int) [][]int {
    var (
		ret [][]int
		dfs func(i, sum int, nums []int, last bool)
	)
	slices.Sort(candidates)

	dfs = func(i, sum int, nums []int, last bool) {
		if sum > target {
			return
		}
		if i == len(candidates) {
			if sum == target {
				ret = append(ret, slices.Clone(nums))
			}
			return
		}

		dfs(i+1, sum, nums, false) // not select

		if i-1 >= 0 && candidates[i-1] == candidates[i] && !last {
			return
		}
		dfs(i+1, sum+candidates[i], append(nums, candidates[i]), true) // select
	}
	dfs(0, 0, []int{}, false)

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：6.88MB，击败 5.00%
*/

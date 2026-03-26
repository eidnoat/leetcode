func subsetsWithDup(input []int) [][]int {
    var (
        ret [][]int
		dfs func(i int, nums []int, last bool)
    )
    slices.Sort(input)

	dfs = func(i int, nums []int, last bool) {
		if i == len(input) {
			ret = append(ret, slices.Clone(nums))
			return
		}

		dfs(i+1, nums, false) // not select

		if i-1 >= 0 && input[i-1] == input[i] && !last {
			return
		}
		dfs(i+1, append(nums, input[i]), true) // select
	}
	dfs(0, []int{}, false)

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.35MB，击败 21.34%
*/

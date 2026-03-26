func combine(n int, k int) [][]int {
    var (
        ret [][]int
        dfs func(i int, nums []int)
    )
    dfs = func(i int, nums []int){
        if len(nums) == k {
            ret = append(ret, slices.Clone(nums))
            return
        }
		if i > n {
			return
		}
		
		dfs(i+1, nums) // not select
		dfs(i+1, append(nums, i)) // select
    }
	dfs(1, []int{})

	return ret
}

/*
	执行用时分布：53ms，击败 10.09%
	消耗内存分布：65.51MB，击败 25.11%
*/

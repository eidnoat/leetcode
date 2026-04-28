func minPathSum(grid [][]int) int {
	var (
		mem = map[int]map[int]int{len(grid) - 1: {len(grid[0]) - 1: grid[len(grid)-1][len(grid[0])-1]}}
		dfs func(i, j int) int
	)
	dfs = func(i, j int) int {
		if v, ok := mem[i][j]; ok {
			return v
		}
		if mem[i] == nil {
			mem[i] = make(map[int]int)
		}

		dv, rv := math.MaxInt, math.MaxInt
		if i < len(grid)-1 {
			dv = dfs(i+1, j)
		}
		if j < len(grid[0])-1 {
			rv = dfs(i, j+1)
		}
		mem[i][j] = grid[i][j] + min(dv, rv)

		return mem[i][j]
	}

	return dfs(0, 0)
}

/*
	执行用时分布：16ms，击败 2.51%
	消耗内存分布：8.86MB，击败 5.02%
*/

func partition(input string) [][]string {
    var (
		ret [][]string
		dfs func(i int, strs []string)
	)
	check := func (s string) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}

	dfs = func(i int, strs []string) {
		if i == len(input) {
			ret = append(ret, slices.Clone(strs))
			return
		}
		for j := i; j < len(input); j++ {
			if !check(input[i:j+1]) {
				continue
			}
			dfs(j+1, append(strs, input[i:j+1]))
		}
	}
	dfs(0, []string{})

	return ret
}

/*
	执行用时分布：19ms，击败 37.00%
	消耗内存分布：23.88MB，击败 69.00%
*/

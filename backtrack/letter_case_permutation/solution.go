func letterCasePermutation(s string) []string {
    var (
		ret []string
		dfs func(i int, raw []byte)
	)
	dfs = func(i int, raw []byte) {
		if i == len(s) {
			ret = append(ret, string(raw))
			return
		}

		if isAlpha(s[i]) {
			dfs(i+1, append(raw, tran(s[i])))
		}
		dfs(i+1, append(raw, s[i]))
	}
	dfs(0, []byte{})

	return ret
}

func tran(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return 'A' + (c-'a')
	} else {
		return 'a' + (c-'A')
	}
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：7.84MB，击败 97.67%
*/

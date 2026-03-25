var hash = map[int][]byte{2: []byte{'a', 'b', 'c'}, 3: []byte{'d', 'e', 'f'}, 4: []byte{'g', 'h', 'i'}, 5: []byte{'j', 'k', 'l'}, 6: []byte{'m', 'n', 'o'}, 7: []byte{'p', 'q', 'r', 's'}, 8: []byte{'t', 'u', 'v'}, 9: []byte{'w', 'x', 'y', 'z'}}

func letterCombinations(digits string) []string {
    var ret []string
	var dfs func(i int, raw []byte)

	dfs = func(i int, raw []byte) {
		if i >= len(digits) {
			ret = append(ret, string(raw))
			return
		}

		for _, c := range hash[int(digits[i]-'0')] {
			dfs(i+1, append(raw, c))
		}
	}

	dfs(0, []byte{})

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.79MB，击败 99.95%
*/

func removeInvalidParentheses(input string) []string {
    var (
		ret []string
		dfs func(i int, raw []byte, lc, rc int)
	)
    hash := make(map[string]bool)
	dfs = func(i int, raw []byte, lc, rc int) {
		if i == len(input) {
			if lc != rc {
				return
			}
            if hash[string(raw)] {
                return
            }
			if len(ret) > 0 {
				if len(raw) > len(ret[0]) {
					ret = []string{}
				} else if len(raw) < len(ret[0]) {
					return
				}
			}

			ret = append(ret, string(raw))
            hash[string(raw)] = true
			return
		}

		if input[i] == '(' {
			dfs(i+1, raw, lc, rc) // not select
			dfs(i+1, append(raw, '('), lc+1, rc) // select
		} else if input[i] == ')' {
			dfs(i+1, raw, lc, rc) // not select
			if lc > rc {
				dfs(i+1, append(raw, ')'), lc, rc+1) // select
			}
		} else {
			dfs(i+1, append(raw, input[i]), lc, rc) // select
		}
	}
	dfs(0, []byte{}, 0, 0)

	return ret
}

/*
	执行用时分布：97ms，击败 32.17%
	消耗内存分布：8.60MB，击败 33.91%
*/

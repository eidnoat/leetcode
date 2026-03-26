func punishmentNumber(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		if check(i) {
			ret += i * i
		}
	}
	return ret
}

func check(v int) bool {
	var (
		ret bool
		dfs func(i, sum int)
		raw = []byte(fmt.Sprintf("%d", v*v))
	)
	dfs = func(i, sum int) {
		if ret {
			return
		}
		if i == len(raw) {
			if sum == v {
				ret = true
			}
			return
		}
		if sum > v {
			return
		}
		if raw[i] == '0' {
			dfs(i+1, sum)
			return
		}

		for j := i; j < len(raw); j++ {
			dfs(j+1, sum+raw2Num(raw[i:j+1]))
		}
	}
	dfs(0, 0)
	
	return ret
}

func raw2Num(raw []byte) int {
	ret := 0
	for _, c := range raw {
		ret = (ret * 10) + int(c-'0')
	}
	return ret
}

/*
	执行用时分布：27ms，击败 56.76%
	消耗内存分布：5.18MB，击败 8.11%
*/

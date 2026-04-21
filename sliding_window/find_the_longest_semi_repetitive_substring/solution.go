func longestSemiRepetitiveSubstring(s string) int {
    l, r, cnt, ret := 0, 0, 0, 0
	for ; r < len(s); r++ {
		if r-1>0 && s[r-1]==s[r] {
			cnt++
		}
		for ; cnt > 1; l++ {
			if s[l]==s[l+1] {
				cnt--
			}
		}
		ret = max(ret, r-l+1)
	}

	return ret
}

/*
	执行用时分布：7ms，击败 15.91%
	消耗内存分布：6.25MB，击败 13.64%
*/

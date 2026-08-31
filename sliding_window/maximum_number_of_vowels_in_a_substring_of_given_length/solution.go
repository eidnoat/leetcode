func maxVowels(s string, k int) int {
	l, r, ret := 0, 0, 0 
	for ; r < k; r++ {
		if !is(s[r]) {
			continue
		}
		ret++
	}

	cur := ret
	for ; r < len(s); l, r = l+1, r+1 {
		if is(s[l]) {
			cur--
		}
		if is(s[r]) {
			cur++
		}
		ret = max(ret, cur)
	}

	return ret
}

func is(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：6.40MB，击败 66.37%
*/
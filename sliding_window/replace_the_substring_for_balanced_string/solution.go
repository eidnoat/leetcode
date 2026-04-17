func balancedString(s string) int {
    hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}
	for _, k := range []byte{'Q', 'W', 'E', 'R'} {
		if hash[k] <= len(s)/4 {
			delete(hash, k)
		} else {
			hash[k] -= len(s)/4
		}
	}
	if len(hash) == 0 {
		return 0
	}

	records := make(map[byte]int)
	contain := func() bool {
		for k, v :=  range hash {
			if records[k] < v {
				return false
			}
		}
		return true
	}

	l, r, ret := 0, 0, math.MaxInt
	for ; r < len(s); r++ {
		records[s[r]]++
		for ; l<=r && contain(); l++ {
			ret = min(ret, r-l+1)
			records[s[l]]--
		}
	}
	if ret == math.MaxInt {
		ret = 0
	}

	return ret
}

/*
	执行用时分布：43ms，击败 6.58%
	消耗内存分布：4.71MB，击败 55.26%
*/

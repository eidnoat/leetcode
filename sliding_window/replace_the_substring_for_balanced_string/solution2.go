func balancedString(s string) int {
    hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	target := make(map[byte]int)
	for k, v := range hash {
		if v > len(s)/4 {
			target[k] = v-len(s)/4
		}
	}
	if len(target) == 0 {
		return 0
	}

	l, r, ret, records := 0, 0, math.MaxInt, make(map[byte]int)
	contain := func() bool {
		for k, v := range target {
			if records[k] < v {
				return false
			}
		}

		return true
	}

	for ; r < len(s); r++ {
		records[s[r]]++
		for ; contain(); l, records[s[l]] = l+1, records[s[l]]-1 {
			ret = min(ret, r-l+1)
		}
	}

	return ret
}
func maxPossibleScore(start []int, d int) int {
    slices.Sort(start)
    check := func(v int) bool {
        num := start[0]
        for i := 0; i < len(start)-1; i++ {
            if num = max(num+v, start[i+1]); !(num >= start[i+1] && num <= start[i+1]+d) {
                return false
            }
        }

        return true
    }

    l, r := 0, start[len(start)-1]+d-start[0]
    for l <= r {
        m := (l+r)/2
        if check(m) {
            if m+1 <= r && check(m+1) {
                l = m+1
            } else {
                return m
            }
        } else {
            r = m-1
        }
    }

    return 0
}

/*
	执行用时分布：65ms，击败 10.53%
	消耗内存分布：11.42MB，击败 47.37%
*/
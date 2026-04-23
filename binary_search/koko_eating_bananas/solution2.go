func minEatingSpeed(piles []int, h int) int {
    l, r, cnt := 1, slices.Max(piles), func(v int) int {
		ret := 0
		for _, pile := range piles {
			if pile%v == 0 {
				ret += pile/v
			} else {
				ret += pile/v+1
			}
		}
		return ret
	}

	return l+sort.Search(r-l+1, func(i int) bool { 
		return cnt(i+l)<=h
	 })
}

/*
	执行用时分布：16ms，击败 10.67%
	消耗内存分布：8.01MB，击败 52.17%
*/

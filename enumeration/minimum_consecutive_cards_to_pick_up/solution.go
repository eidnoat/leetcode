func minimumCardPickup(cards []int) int {
    hash, ret := make(map[int]int), math.MaxInt
	for i, card := range cards {
		if lp, ok := hash[card]; ok {
			ret = min(ret, i-lp+1)
		}
		hash[card] = i
	}
	if ret == math.MaxInt {
		return -1
	}

	return ret
}

/*
	执行用时分布：52ms，击败 46.05%
	消耗内存分布：14.06MB，击败 10.53%
*/

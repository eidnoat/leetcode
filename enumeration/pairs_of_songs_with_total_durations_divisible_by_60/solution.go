func numPairsDivisibleBy60(time []int) int {
	hash, ret := make(map[int]int), 0
	for _, t := range time {
		ret += hash[(60-(t%60))%60]
		hash[t%60]++
	}
	return ret
}

/*
	执行用时分布：5ms，击败 41.18%
	消耗内存分布：8.59MB，击败 50.98%
*/

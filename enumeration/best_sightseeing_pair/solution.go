func maxScoreSightseeingPair(values []int) int {
	last, ret := values[0], math.MinInt
	for j := 1; j < len(values); j++ {
		last, ret = max(last, values[j]+j), max(ret, last+values[j]-j)
	}
	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：8.89MB，击败 16.33%
*/

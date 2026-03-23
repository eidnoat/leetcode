func dailyTemperatures(temperatures []int) []int {
    sk, ret := make([]int, 0), make([]int, len(temperatures))
	for i, t := range temperatures {
		for ; len(sk) > 0 && temperatures[sk[len(sk)-1]] < t; sk = sk[:len(sk)-1]  {
			ret[sk[len(sk)-1]] = i-sk[len(sk)-1]
		}
		sk = append(sk, i)
	}

	return ret
}

/*daily_temperatures
	执行用时分布：14ms，击败 14.22%
	消耗内存分布：10.56MB，击败 18.26%
*/
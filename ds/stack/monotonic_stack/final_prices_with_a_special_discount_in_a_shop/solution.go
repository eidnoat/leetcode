func finalPrices(prices []int) []int {
    sk, ret := make([]int, 0), slices.Clone(prices)
	for i, p := range prices {
		for ; len(sk) > 0 && p <= prices[sk[len(sk)-1]]; sk = sk[:len(sk)-1] {
			ret[sk[len(sk)-1]] -= p
		}
		sk = append(sk, i)
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.68MB，击败 6.97%
*/

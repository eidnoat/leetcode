func maximumSum(nums []int) int {
	tran := func(v int) int {
		ret := 0
		for ; v > 0; v /= 10 {
			ret += v%10
		}
		return ret
	}

    hash, ret := make(map[int]int), -1
	for _, num := range nums {
		if v, ok := hash[tran(num)]; ok {
			ret = max(ret, v+num)
			hash[tran(num)] = max(v, num)
		} else {
			hash[tran(num)] = num
		}
	}

	return ret
}

/*
	执行用时分布：28ms，击败 32.94%
	消耗内存分布：10.37MB，击败 28.23%
*/

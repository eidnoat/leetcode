func numSubarrayBoundedMax(nums []int, left int, right int) int {
	cnt := func(v int) int {
		c, ret := 0, 0
		for _, num := range nums {
			if num < v {
				c++
				ret += c
			} else {
				c = 0
			}
		}
		return ret
	}

	return cnt(right+1)-cnt(left)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：10.06MB，击败 36.11%
*/
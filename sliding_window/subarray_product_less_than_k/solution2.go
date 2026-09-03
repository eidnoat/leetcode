func numSubarrayProductLessThanK(nums []int, k int) int {
    l, r, cur, ret := 0, 0, 1, 0
	for ; r < len(nums); r++ {
		cur *= nums[r]
		for ; cur > k; l, cur = l+1, cur/nums[l] {
		}
		ret += r-l+1
	}

	return ret
}

/*
	执行用时分布：4ms，击败 28.09%
	消耗内存分布：8.58MB，击败 87.08%
*/

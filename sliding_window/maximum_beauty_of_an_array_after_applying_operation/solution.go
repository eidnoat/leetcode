func maximumBeauty(nums []int, k int) int {
    slices.Sort(nums)
	l, r, ret := 0, 0, 0
	for ; r < len(nums); r++ {
		for ; nums[r]-nums[l] > 2*k; l++ {
		}
		ret = max(ret, r-l+1)
	}

	return ret
}

/*
	执行用时分布：44ms，击败 65.26%
	消耗内存分布：9.93MB，击败 31.58
*/

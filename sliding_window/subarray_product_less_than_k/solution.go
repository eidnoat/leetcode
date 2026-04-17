func numSubarrayProductLessThanK(nums []int, k int) int {
    l, r, product, ret := 0, 0, 1, 0
	for ; r < len(nums); r++ {
		product *= nums[r]
		for ; l <= r && product >= k; product, l = product/nums[l], l+1 {
		}
        ret += r-l+1
	}

	return ret
}

/*
	执行用时分布：4ms，击败 43.52%
	消耗内存分布：9.16MB，击败 7.78%
*/

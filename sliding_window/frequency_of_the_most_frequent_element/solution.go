func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	l, r, ret := 0, 1, 1
	for ; r < len(nums); r++ {
		k -= (r-l)*(nums[r]-nums[r-1])
		for ; k < 0; l++ {
			k += nums[r]-nums[l]
		}
		ret = max(ret, r-l+1)
	}

	return ret
}

/*
	执行用时分布：46ms，击败 8.70%
	消耗内存分布：10.06MB，击败 8.70%
*/
func maximumCount(nums []int) int {
	l, r := sort.Search(len(nums), func(i int) bool { return nums[i] >= 0 })-1, sort.Search(len(nums), func(i int) bool { return nums[i] > 0 })
	return max(l+1, len(nums)-r)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：7.41MB，击败 42.37%
*/

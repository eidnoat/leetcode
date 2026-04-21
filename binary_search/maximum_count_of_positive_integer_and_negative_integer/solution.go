func maximumCount(nums []int) int {
	return max(len(nums)-sort.Search(len(nums), func(i int) bool { return nums[i] > 0 } ), sort.Search(len(nums), func(i int) bool { return nums[i] >= 0 } ))
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：7.36MB，击败 80.22%
*/

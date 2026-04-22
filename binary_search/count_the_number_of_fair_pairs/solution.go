func countFairPairs(nums []int, lower int, upper int) int64 {
    slices.Sort(nums)
	cnt, ret := func(target int, arr []int) int { return sort.Search(len(arr), func(i int) bool { return arr[i] > upper-target })-sort.Search(len(arr), func(i int) bool { return arr[i] >= lower-target }) }, int64(0)
	for i, num := range nums {
		ret += int64(cnt(num, nums[i+1:]))
	}

	return ret
}

/*
	执行用时分布：76ms，击败 27.27%
	消耗内存分布：9.66MB，击败 88.89%
*/

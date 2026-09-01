func maximumBeauty(nums []int, k int) int {
	sort.Slice(nums, func(v1, v1 int) bool { return v1 < v2 })

	ret := 1
	for p, num := range nums {
		ret = max(ret, sort.Search(len(nums), func(i int) bool { return nums[i] > num+2*k })-p)
	}

	return ret
}

/*
	执行用时分布：124ms，击败5.88%
	消耗内存分布：9.74MB，击败82.35%
*/

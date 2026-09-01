func maximumBeauty(nums []int, k int) int {
	diff := make([]int, slices.Max(nums)+1)
	for _, num := range nums {
		l, r := max(0, num-k), min(len(diff)-1, num+k)
		diff[l] += 1
		if r+1 < len(diff) {
			diff[r+1] -= 1
		}
	}

	ret, cur := 1, 0
	for _, v := range diff {
		cur += v
		ret = max(ret, cur)
	}

	return ret
}

/*
	执行用时分布：4ms，击败 100.00%
	消耗内存分布：9.95MB，击败 29.41%
*/

func triangleNumber(nums []int) int {
	slices.Sort(nums)
	ret := 0

	for i := len(nums)-1; i >= 0; i-- {
		v := nums[i]
		l, r := 0, i-1

		for l < r {
			if s := nums[l]+nums[r]; s <= v {
				l++
			} else {
				ret += (r-l)
				r--
			}
		}
	}

	return ret
}

/*
	执行用时分布：21ms，击败 85.61%
	消耗内存分布：4.66MB，击败 51.51%
*/

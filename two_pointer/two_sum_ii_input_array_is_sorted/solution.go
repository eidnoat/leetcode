func twoSum(nums []int, target int) []int {
    l, r, ret := 0, len(nums)-1, make([]int, 2)
	for l < r {
		sum := nums[l]+nums[r]
		if sum < target {
			l++
		} else if sum == target {
			ret[0], ret[1] = l+1, r+1
			break
		} else {
			r--
		}
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：7.72MB，击败 14.74%
*/

func rob(nums []int) int {
    if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(nums[0]+dp(nums[2:len(nums)-1]), dp(nums[1:]))
}

func dp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	v1, v2 := max(nums[len(nums)-2], nums[len(nums)-1]), nums[len(nums)-1]
	for i := len(nums)-3; i >= 0; i-- {
		v1, v2 = max(nums[i]+v2, v1), v1
	}
	return max(v1, v2)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.78MB，击败 84.38%
*/

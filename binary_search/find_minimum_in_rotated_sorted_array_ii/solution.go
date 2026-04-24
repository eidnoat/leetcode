func findMin(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}

    l, m, r := 0, (len(nums)-1)/2, len(nums)-1
	
	var lmin int
	if nums[l] < nums[m] {
		lmin = nums[l]
	} else if nums[l] == nums[m] {
		lmin = min(nums[l], findMin(nums[l+1:m]))
	} else {
		lmin = min(nums[m], findMin(nums[l+1:m]))
	}

	var rmin int
	if nums[m] < nums[r] {
		rmin = nums[m]
	} else if nums[m] == nums[r] {
		rmin = min(nums[m], findMin(nums[m+1:r]))
	} else {
		rmin = min(nums[r], findMin(nums[m+1:r]))
	}

	return min(lmin, rmin)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.71MB，击败 53.66%
*/

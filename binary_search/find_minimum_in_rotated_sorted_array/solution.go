func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return min(nums[0], nums[1])
	}

	l, m, r := 0, (len(nums)-1)/2, len(nums)-1
	if nums[l] < nums[m] && nums[m] < nums[r] {
		return nums[l]
	} else if nums[l] < nums[m] && nums[m] > nums[r] {
		return min(nums[l], findMin(nums[m+1:]))
	} else if nums[l] > nums[m] && nums[m] < nums[r] {
		return min(nums[m], findMin(nums[:m]))
	} else {
		return nums[len(nums)-1]
	}
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.15MB，击败 77.14%
*/

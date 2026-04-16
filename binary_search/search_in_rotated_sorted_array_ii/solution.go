func search(nums []int, target int) bool {
    l, r, m := 0, len(nums)-1, (len(nums)-1)/2
	if nums[m] == target {
		return true
	}

	if m-1 >= l {
		if nums[l] < nums[m-1] && find(nums[l:m], target) {
			return true
		}
		if nums[l] >= nums[m-1] && search(nums[l:m], target) {
			return true
		}
	}

	if m+1 <= r {
		if nums[m+1] < nums[r] && find(nums[m+1:r+1], target) {
			return true
		}
		if nums[m+1] >= nums[r] && search(nums[m+1:r+1], target) {
			return true
		}
	}

	return false
}

func find(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l+r)/2
		if nums[m] < target {
			l = m+1
		} else if nums[m] == target {
			return true
		} else {
			r = m-1
		}
	}

	return false
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.79MB，击败 68.52%
*/

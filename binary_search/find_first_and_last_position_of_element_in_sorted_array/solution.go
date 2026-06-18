func searchRange(nums []int, target int) []int {
    return []int{findL(nums, target), findR(nums, target)}
}

func findR(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l+r)/2
		if nums[m] < target {
			l = m+1
		} else if nums[m] == target {
			if m+1 <= r && nums[m+1] == target {
				l = m+1
			} else {
				return m
			}
		} else {
			r = m-1
		}
	}

	return -1
}

func findL(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l+r)/2
		if nums[m] < target {
			l = m+1
		} else if nums[m] == target {
			if m-1 >= l && nums[m-1] == target {
				r = m-1
			} else {
				return m
			}
		} else {
			r = m-1
		}
	}

	return -1
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：6.09MB，击败 97.78%
*/

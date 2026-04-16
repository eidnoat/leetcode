func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
	for l <= r {
		m := (l+r)/2
		if nums[l] <= nums[m] { // left side is ordered
			p := find(nums, l, m, target)
			if p != -1 {
				return p
			}
			l = m+1
		} else { // right side is ordered
			p := find(nums, m, r, target)
			if p != -1 {
				return p
			}
			r = m-1
		}
	}

	return -1
}

func find(nums []int, start, end, target int) int {
	l, r := start, end
	for l <= r {
		m := (l+r)/2
		if nums[m] < target {
			l = m+1
		} else if nums[m] == target {
			return m
		} else {
			r = m-1
		}
	}

	return -1
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.22MB，击败 70.08%
*/

var max_v = 10^5

func threeSumClosest(nums []int, target int) int {
    slices.Sort(nums)
	ret := max_v

	for i := 0; i < len(nums)-2; i++ {
		tmp := nums[i]+towSumClosest(nums[i+1:], target-nums[i])
		if abs(tmp-target) < abs(ret-target) {
			ret = tmp
		}
	}

	return ret
}

func towSumClosest(nums []int, target int) int {
	l, r, ret := 0, len(nums)-1, max_v

	for l < r {
		sum := nums[l]+nums[r]
		if abs(sum-target) < abs(ret-target) {
			ret = sum
			if ret == target {
				return ret
			}
		}

		if sum < target {
			l++
		} else {
			r--
		}
	}

	return ret
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

/*
	执行用时分布：7ms，击败 52.48%
	消耗内存分布：4.54MB，击败 50.32%
*/	

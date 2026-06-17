func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	ret = make([][]int, 0)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}

		for _, tmp range twoSum(nums[i+1:], -nums[i]) {
			ret = append(ret, append(tmp, nums[i]))
		}
	}

	return ret
}

func twoSum(nums []int, target int) [][]int {
	l, r, ret := 0, len(nums)-1, make([][]int, 0)
	for l < r {
		sum, sign := nums[l]+nums[r], 0
		if sum < target {
		} else if sum == target {
			ret = append(ret, []int{nums[l], nums[r]})
		} else {
			sign = 1
		}

		if sign == 0 {
			for {
				l++
				if l >= len(nums) || nums[l] != nums[l-1] {
					break
				}
			}
		} else {
			for {
				r--
				if r < 0 || nums[r] != nums[r+1] {
					break
				}
			}
		}
	}

	return ret
}

/*
	执行用时分布：33ms，击败 16.75%
	消耗内存分布：9.71MB，击败 72.82%
*/
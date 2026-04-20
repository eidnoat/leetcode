func threeSum(nums []int) [][]int {
    slices.Sort(nums)
	towSum := func(arr []int, target int) [][]int {
		ret := make([][]int, 0)
		l, r := 0, len(arr)-1
		for l < r {
			sum, moveL, moveR := arr[l]+arr[r], false, false
			if sum < target {
				l, moveL = l+1, true
			} else if sum == target {
				ret, l, moveL, r, moveR = append(ret, []int{arr[l], arr[r]}), l+1, true, r-1, true
			} else {
				r, moveR = r-1, true
			}

			for ; moveL && l < r && arr[l] == arr[l-1]; l++ {
			}
			for ; moveR && l < r && arr[r] == arr[r+1]; r-- {
			}
		}
		return ret
	}

	ret := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}

		for _, ts := range towSum(nums[i+1:], -num) {
			ret = append(ret, append(tmp, num))
		}
	}

	return ret
}

/*
	执行用时分布：39ms，击败 10.68%
	消耗内存分布：9.73MB，击败 71.53%
*/

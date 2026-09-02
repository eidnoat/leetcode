func minOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := sum - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}

	l, r, cur, ret := 0, 0, 0, math.MaxInt
	for ; r < len(nums); r++ {
		cur += nums[r]
		for ; cur > target; l++ {
			cur -= nums[l]
		}

		if cur == target {
			ret = min(ret, len(nums)-(r-l+1))
		}
	}

	if ret == math.MaxInt {
		ret = -1
	}

	return ret
}
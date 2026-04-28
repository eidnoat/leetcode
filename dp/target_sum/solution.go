func findTargetSumWays(nums []int, target int) int {
    if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}

	return findTargetSumWays(nums[1:], target-nums[0])+findTargetSumWays(nums[1:], target+nums[0])
}

/*
	执行用时分布：471ms，击败 20.14%
	消耗内存分布：3.95MB，击败 90.42%
*/

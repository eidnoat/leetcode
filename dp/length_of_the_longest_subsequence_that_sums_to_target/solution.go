func lengthOfLongestSubsequence(nums []int, target int) int {
	var (
		dfs func(nums []int, target int) int
	)
	dfs = func(nums []int, target int) int {
		if len(nums) == 0 {
			if target == 0 {
				return 0
			} else {
				return -1
			}
		}

		path1, path2 := dfs(nums[1:], target-nums[0]), dfs(nums[1:], target)
		if path1 != -1 {
			path1 += 1
		}

		return max(path1, path2)
	}

	return dfs(nums, target)
}

/*
	可以通过基本测试用例，但整体会超时。
*/

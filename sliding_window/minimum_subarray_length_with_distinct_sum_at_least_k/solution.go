func minLength(nums []int, k int) int {
    l, r, cur, ret, visited := 0, 0, 0, math.MaxInt, make(map[int]int)
	for ; r < len(nums); r++ {
		if visited[nums[r]] += 1; visited[nums[r]] == 1 {
			cur += nums[r]
		}
		for ; cur >= k; l++ {
			if visited[nums[l]] == 1 {
				cur -= nums[l]
			}
			visited[nums[l]], ret = visited[nums[l]]-1, min(ret, r-l+1)
		}
	}
	if ret == math.MaxInt {
		ret = -1
	}

	return ret
}

/*
	执行用时分布：122ms，击败 57.53%
	消耗内存分布：14.09MB，击败 67.12%
*/

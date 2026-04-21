func countSubarrays(nums []int, k int) int64 {
    l, r, cnt, ret, target := 0, 0, 0, int64(0), slices.Max(nums)
	for ; r < len(nums); r++ {
		if nums[r] == target {
			cnt++
		}
		for ; cnt >= k; ret, l = ret+int64(len(nums)-r), l+1 {
			if nums[l] != target {
				continue
			}
			cnt--
		}
	}

	return ret
}

/*
	执行用时分布：3ms，击败 70.18%
	消耗内存分布：10.59MB，击败 21.93%
*/

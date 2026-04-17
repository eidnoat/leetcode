func numberOfSubarrays(nums []int, k int) int {
    arr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 == 1 {
			arr = append(arr, i)
		}
	}

	ret := 0
	for i := 0; i+k-1 < len(arr); i++ {
		l, r := -1, len(nums)
		if i > 0 {
			l = arr[i-1]
		}
		if i+k < len(arr) {
			r = arr[i+k]
		}
		lCnt, rCnt := max(0, arr[i]-l), max(0, r-arr[i+k-1])
		ret += lCnt*rCnt
	}

	return ret
}

/*
	执行用时分布：3ms，击败 73.33%
	消耗内存分布：9.01MB，击败 25.71%
*/

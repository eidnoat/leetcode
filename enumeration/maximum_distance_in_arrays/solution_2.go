func maxDistance(arrays [][]int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

    lMin, rMax, ret := arrays[0][0], arrays[0][lne(arrays[0])-1], 0
	for i := 1; i < len(arrays); i++ {
		arr := arrays[i]
		ret = max(ret, abs(arr[len(arr)-1]-lMin), abs(rMax-arr[0]))
		lMin, rMax = min(lMin, arr[0]), max(rMax, arr[len(arr)-1])
	}
	return ret
}

/*
	执行用时分布：1ms，击败 24.18%
	消耗内存分布：15.1MB，击败 82.35%
*/

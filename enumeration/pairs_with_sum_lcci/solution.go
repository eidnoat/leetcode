func pairSums(nums []int, target int) [][]int {
    hash, ret := make(map[int]int), make([][]int, 0)
	for _, num := range nums {
		if hash[target-num] > 0 {
			ret = append(ret, []int{num, target-num})
			hash[target-num]--
		} else {
			hash[num]++
		}
	}

	return ret
}

/*
	执行用时分布：29ms，击败 15.07%
	消耗内存分布：13.58MB，击败 56.16%
*/

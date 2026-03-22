const MAX_VALUE = 1000000000 + 7

func specialTriplets(nums []int) int {
	lHash, rHash := map[int]int{nums[0]: 1}, make(map[int]int)
	for i := 2; i < len(nums); i++ {
		rHash[nums[i]]++
	}

	ret := 0
	for j := 1; j < len(nums)-1; j++ {
		ret += (lHash[nums[j]*2] * rHash[nums[j]*2]) % MAX_VALUE
		ret %= MAX_VALUE

		lHash[nums[j]]++
		rHash[nums[j+1]]--
	}

	return ret
}

/*
	执行用时分布：263ms，击败 25.58%
	消耗内存分布：19.39MB，击败 34.88%
*/

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	nextPrefixSum, previousPrefixSum := make([]int, 27), make([]int, 27)
	for i := 0; i < 26; i++ {
		nextPrefixSum[i+1], previousPrefixSum[i+1] = nextPrefixSum[i]+nextCost[i], previousPrefixSum[i]+previousCost[i]
	}
	nextCostCalculate := func(sb, tb byte) int {
		if sb <= tb {
			i := sb-'a'
			j := tb-'a'-1
			return nextPrefixSum[j+1]-nextPrefixSum[i]
		} else {
			i1 := sb-'a'
			j1 := 'z'-'a'
			i2 := 'a'-'a'
			j2 := tb-'a'-1
			return nextPrefixSum[j1+1]-nextPrefixSum[i1]+nextPrefixSum[j2+1]-nextPrefixSum[i2]
		}
	}
	previousCostCalculate := func(sb, tb byte) int {
		if sb >= tb {
			i := tb-'a'+1
			j := sb-'a'
			return previousPrefixSum[j+1]-previousPrefixSum[i]
		} else {
			i1 := 'a'-'a'
			j1 := sb-'a'
			i2 := tb-'a'+1
			j2 := 'z'-'a'
			return previousPrefixSum[j1+1]-previousPrefixSum[i1]+previousPrefixSum[j2+1]-previousPrefixSum[i2]
		}
	}

	ret := 0
	for i := 0; i < len(s); i++ {
		ret += min(nextCostCalculate(s[i], t[i]), previousCostCalculate(s[i], t[i]))
	}

	return int64(ret)
}

/*
	执行用时分布：26ms，击败 58.62%
	消耗内存分布：8.35MB，击败 86.21%
*/

func findLongestSubarray(array []string) []string {
	prefixSum := make([]int, len(array)+1) // alphaCnt - numCnt
	for i := 0; i < len(array); i++ {
		if isAlpha(array[i][0]) {
			prefixSum[i+1] = prefixSum[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i] - 1
		}
	}

	ret, hash := []string{}, map[int]int{0: -1}
	for j, _ := range array {
		if i, ok := hash[prefixSum[j+1]]; ok && j-i > len(ret) {
			ret = array[i+1 : j+1]
		}
		if _, ok := hash[prefixSum[j+1]]; !ok {
			hash[prefixSum[j+1]] = j
		}
	}

	return ret
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

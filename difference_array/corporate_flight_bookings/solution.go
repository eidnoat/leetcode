func corpFlightBookings(bookings [][]int, n int) []int {
    diff := make([]int, n)
	for _, booking := range bookings {
		f, l, s := booking[0]-1, booking[1]-1, booking[2]
		diff[f] += s
		if l+1 < len(diff) {
			diff[l+1] -= s
		}
	}

	for i := 1; i < len(diff); i++ {
		diff[i] += diff[i-1]
	}

	return diff
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：10.35MB，击败 94.38%
*/

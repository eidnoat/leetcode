func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	total, anger, record := 0, 0, 0
	for i, c := range customers {
		total, anger = total+c, anger+grumpy[i]*c
		if i == minutes-1 {
			record = anger
		}
	}
	if minutes >= len(customers) {
		return total
	}

	cur := record
	for l, r := 0, minutes; r < len(customers); l, r = l+1, r+1 {
		cur += grumpy[r]*customers[r] - grumpy[l]*customers[l]
		record = max(record, cur)
	}

	return total - anger + record
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：8.00MB，击败 65.91%
*/


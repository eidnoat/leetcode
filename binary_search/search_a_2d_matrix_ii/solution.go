func searchMatrix(matrix [][]int, target int) bool {
    h, l := 0, len(matrix[0])-1
	for h < len(matrix) && l >= 0 {
		if target < matrix[h][l] {
			l--
		} else if target == matrix[h][l] {
			return true
		} else {
			h++
		}
	}

	return false
}

/*
	执行用时分布：15ms，击败 82.51%
	消耗内存分布：7.87MB，击败 77.82%
*/
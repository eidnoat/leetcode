func searchMatrix(matrix [][]int, target int) bool {
	l, r, h := 0, len(matrix)-1, -1
	for l <= r {
		m := (l+r)/2
		if target < matrix[m][0] {
			r = m-1
		} else if target > matrix[m][len(matrix[m])-1] {
			l = m+1
		} else {
			h = m
			break
		}
	}
	if h == -1 {
		return false
	}

	l, r = 0, len(matrix[h])-1
	for l <= r {
		m := (l+r)/2
		if target < matrix[h][m] {
			r = m-1
		} else if target == matrix[h][m] {
			return true
		} else {
			l = m+1
		}
	}

	return false
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.30MB，击败 25.95%
*/

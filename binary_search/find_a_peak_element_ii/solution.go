func findPeakGrid(mat [][]int) []int {
    h, l := 0, 0
	for {
		v, lv, rv, tv, bv := mat[h][l], -1, -1, -1, -1
		if l-1 >= 0 {
			lv = mat[h][l-1]
		}
		if l+1 < len(mat[0]) {
			rv = mat[h][l+1]
		}
		if h-1 >= 0 {
			tv = mat[h-1][l]
		}
		if h+1 < len(mat) {
			bv = mat[h+1][l]
		}

		newH, newL := -1, -1
		if lv > v {
			v, newH, newL = lv, h, l-1
		}
		if rv > v {
			v, newH, newL = rv, h, l+1
		}
		if tv > v {
			v, newH, newL = tv, h-1, l
		}
		if bv > v {
			v, newH, newL = bv, h+1, l
		}

		if newH == -1 && newL == -1 {
			return []int{h, l}
		}

		h, l = newH, newL
	}

	return []int{-1, -1}
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：11.81MB，击败 40.00%
*/

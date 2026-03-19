const (
	MIN_VALUE, MAX_VALUE = -100000, 100000
)

func findClosestElements(arr []int, k int, x int) []int {
    search := func() int { // 找到最后一个小于等于 x 的元素的下标
		l, r := 0, len(arr)-1
		for l <= r {
			m := (l+r)/2
			if arr[m] <= x {
				if m+1 <= r && arr[m+1] <= x {
					l = m+1
				} else {
					return m
				}
			} else {
				r = m-1
			}
		}
		return -1
	}

	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	
	p := search()
	l, r := p, p+1
	for i := 0; i < k; i++ {
		lv, rv := MIN_VALUE, MAX_VALUE
		if l >= 0 {
			lv = arr[l]
		}
		if r < len(arr) {
			rv = arr[r]
		}

		if abs(lv-x) <= abs(rv-x) {
			l--
		} else {
			r++
		}
	}

	ret := make([]int, 0)
	for i := l+1; i < r; i++ {
		ret = append(ret, arr[i])
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：8.63MB，击败 21.95%
*/
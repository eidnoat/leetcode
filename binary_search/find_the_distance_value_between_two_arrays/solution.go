func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	slices.Sort(arr2)
    check := func(v int) bool {
		l := sort.Search(len(arr2), func(i int) bool { return arr2[i] >= v-d } )
		r := sort.Search(len(arr2), func(i int) bool { return arr2[i] > v+d } )-1
		return !(l<=r)
	}

	ret := 0
	for _, v := range arr1 {
		if !check(v) {
			continue
		}
		ret++
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：4.79MB，击败 34.00%
*/

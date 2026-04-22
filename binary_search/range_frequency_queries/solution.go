type RangeFreqQuery struct {
    arr  []int
	hash map[int][]int
}


func Constructor(arr []int) RangeFreqQuery {
    rfq := RangeFreqQuery{arr: arr, hash: make(map[int][]int)}
	for i, v := range arr {
		rfq.hash[v] = append(rfq.hash[v], i)
	}

	return rfq
}


func (q *RangeFreqQuery) Query(left int, right int, value int) int {
    arr := q.hash[value]
	l, r := sort.Search(len(arr), func(i int) bool { return arr[i] >= left }), sort.Search(len(arr), func(i int) bool { return arr[i] > right })-1
	return r-l+1
}


/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */

/*
	执行用时分布：83ms，击败 60.00%
	消耗内存分布：64.95MB，击败 88.75%
*/

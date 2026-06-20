type entry struct {
	v, p int
}

type RangeFreqQuery struct {
	ens []*entry
}


func Constructor(arr []int) RangeFreqQuery {
	ens := make([]*entry, len(arr))
    for i, v := range arr {
		ens[i] = &entry{v, i}
	}

	slices.SortFunc(ens, func(v1, v2 *entry) int {
		if v1.v != v2.v {
			return v1.v - v2.v
		} else {
			return v1.p - v2.p
		}
	})

	return RangeFreqQuery{ens}
}


func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	ens := this.ens

    l, r := slices.Search(len(ens), func(i int) bool { return ens[i].v >= value }), slices.Search(len(ens), func(i int) bool { return ens[i].v > value })
	if l == r {
		return 0
	}

	ens = ens[l:r]
	l, r = slices.search(len(ens), func(i int) bool { return ens[i].p >= left }), slices.search(len(ens), func(i int) bool { return ens[i].p > right })

	return r-l
}


/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
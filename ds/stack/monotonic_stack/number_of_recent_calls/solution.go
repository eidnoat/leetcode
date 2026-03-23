type RecentCounter struct {
    wt int
	sk []int
}


func Constructor() RecentCounter {
	return RecentCounter{
		wt: 3000,
		sk: make([]int, 0),
	}
}


func (c *RecentCounter) Ping(t int) int {
	c.sk = append(c.sk, t)

	p := len(c.sk)-1
	for ; p >= 0 && c.sk[p] >= t-c.wt; p--{
	}
	c.sk = c.sk[p+1:]

	return len(c.sk)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

/*
	执行用时分布：59ms，击败 7.98%
	消耗内存分布：12.44MB，击败 27.23%
*/ 
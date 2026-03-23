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

func (this *RecentCounter) Ping(t int) int {
	this.sk = append(this.sk, t)
	this.sk = this.sk[sort.Search(len(this.sk), func(i int) bool { return this.sk[i] >= t-this.wt }):]
	return len(this.sk)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

 /*
	执行用时分布：39ms，击败 19.25%
	消耗内存分布：11.31MB，击败 74.65%
 */
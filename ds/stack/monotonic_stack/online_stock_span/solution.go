type StockSpanner struct {
	prices []int
	sk     []int
}


func Constructor() StockSpanner {
    return StockSpanner{}
}


func (s *StockSpanner) Next(price int) int {
	prices, sk, ret := s.prices, s.sk, 1

	prices = append(prices, price)
	for ; len(sk) > 0 && price >= prices[sk[len(sk)-1]]; sk = sk[:len(sk)-1] {
	}
	if len(sk) == 0 {
		ret = len(prices)
	} else {
		ret = len(prices)-1-sk[len(sk)-1]
	}
	sk = append(sk, len(prices)-1)

	s.prices, s.sk = prices, sk
	return ret
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

 /*
	执行用时分布：41ms，击败 46.56%
	消耗内存分布：14.21MB，击败 20.61%
 */
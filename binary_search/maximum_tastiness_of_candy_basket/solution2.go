func maximumTastiness(price []int, k int) int {
	slices.Sort(price)

	check := func(v int) bool {
		base, cnt := price[0], 1
		for i := 1; i < len(price); i++ {
			if price[i]-base < v {
				continue
			}

			base, cnt = price[i], cnt+1
			if cnt == k {
				break
			}
		}
		return cnt == k
	}

	l, r := 0, price[len(price)-1]-price[0]
	return sort.Search(r-l+1, func(i int) bool {
		return !check(i)
	})-1
}

/*
	枚举答案+二分搜索
*/

/*
	执行用时分布：45ms，击败 37.04%
	消耗内存分布：10.16MB，击败 37.04%
*/

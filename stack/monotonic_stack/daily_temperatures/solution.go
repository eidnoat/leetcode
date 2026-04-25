func dailyTemperatures(temperatures []int) []int {
	stack, ret := make([]int, 0), make([]int, len(temperatures))
    for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			ret[stack[len(stack)-1]] = i-stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ret
}

/*
	执行用时分布：8ms，击败 49.63%
	消耗内存分布：9.78MB，击败 94.42%
*/

func carFleet(target int, position []int, speed []int) int {
    // no2Pos, no2Spd, arr := make(map[int]int), make(map[int]int), make([]int, len(position))
	// for i := 0; i < len(position); i++ {
	// 	no2Pos[i], no2Spd[i], arr[i] = position[i], speed[i], i
	// }
    arr := make([]int, len(position))
	for i := 0; i < len(position); i++ {
		arr[i] = i
	}

	slices.SortFunc(arr, func(a, b int) int { return position[a]-position[b] })
	// sk, nextLessSpeed := make([]int, 0), make([]int, len(position))
	// for i := 0; i < len(nextLessSpeed); i++ {
	// 	nextLessSpeed[i] = -1
	// }
	// for _, no := range arr {
	// 	for ; len(sk) > 0 && speed[no] < speed[sk[len(sk)-1]]; sk = sk[:len(sk)-1] {
	// 		nextLessSpeed[sk[len(sk)-1]] = no
	// 	}
	// 	sk := append(sk, no)
	// }

	ret, lastCost := 1, cost(target-position[arr[len(arr)-1]], speed[arr[len(arr)-1]])
	for i := len(arr)-2; i >= 0; i-- {
		curCost := cost(target-position[arr[i]], speed[arr[i]])
		if curCost > lastCost {
			ret++
			lastCost = curCost
		}
	}

	return ret
}

func cost(d, s int) float64 {
	return float64(d)/float64(s)
}

/*
	执行用时分布：38ms，击败 95.08%
	消耗内存分布：10.65MB，击败 98.36%
*/

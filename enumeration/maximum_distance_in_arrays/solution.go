func maxDistance(arrays [][]int) int {
    lMinP, rMaxP := 0, 0
    for i, arr := range arrays {
        if arr[0] < arrays[lMinP][0] {
            lMinP = i
        }
        if arr[len(arr)-1] > arrays[rMaxP][len(arrays[rMaxP])-1] {
            rMaxP = i
        }
    }

    ret1 := 0
    for i := 0; i < len(arrays); i++ {
        if i == lMinP {
            continue
        }
        arr := arrays[i]
        ret1 = max(ret1, arr[len(arr)-1]-arrays[lMinP][0])
    }

    ret2 := 0
    for i := 0; i < len(arrays); i++ {
        if i == rMaxP {
            continue
        }
        arr := arrays[i]
        ret2 = max(ret2, arrays[rMaxP][len(arrays[rMaxP])-1]-arr[0])
    }

    return max(ret1, ret2)
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：15.94MB，击败 68.63%
*/
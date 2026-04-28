const MAX = 1000000000+7

var (
	alphaCntOfDig = map[int]int{2:3,3:3,4:3,5:3,6:3,7:4,8:3,9:4}
	mem           = map[int]map[int]int{2:{0:1},3:{0:1},4:{0:1},5:{0:1},6:{0:1},7:{0:1},8:{0:1},9:{0:1}}
)

func countTexts(pressedKeys string) int {
	ret := 1
    for i := 0; i < len(pressedKeys); {
		j := i+1
		for ; j < len(pressedKeys) && pressedKeys[j] == pressedKeys[i]; j++ {
		}
		ret *= cnt(int(pressedKeys[i]-'0'), j-i)
		ret %= MAX
		i = j
	}

	return ret
}

func cnt(dig, repeatCnt int) int {
	if v, ok := mem[dig][repeatCnt]; ok {
		return v
	}

	ret := 0
	for i := 1; i <= min(alphaCntOfDig[dig], repeatCnt); i++ {
		ret += cnt(dig, repeatCnt-i)
		ret %= MAX
	}
	mem[dig][repeatCnt] = ret

	return ret
}

/*
	执行用时分布：71ms，击败 15.25%
	消耗内存分布：23.33MB，击败 6.78%
*/

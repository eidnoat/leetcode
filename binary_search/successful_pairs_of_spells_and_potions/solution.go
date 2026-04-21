func successfulPairs(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)

	ret := make([]int, len(spells))
	for i, spell := range spells {
		ret[i] = len(potions)-sort.Search(len(potions), func(i int) bool { return int64(potions[i])*int64(spell) >= success })
	}

	return ret
}

/*
	执行用时分布：32ms，击败 71.51%
	消耗内存分布：11.66MB，击败 16.76%
*/

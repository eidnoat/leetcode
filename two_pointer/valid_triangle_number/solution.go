func triangleNumber(nums []int) int {
    slices.Sort(nums)
    ret := 0

    for i := 0; i < len(nums)-2; i++ {
        v := nums[i]
        for l := i+1; l < len(nums)-1; l++ {
            tmp := nums[l+1:]
            ret += sort.Search(len(tmp), func(r int) bool {
                return v + nums[l] <= tmp[r]
            })
        }
    }

    return ret
}

/*
    执行用时分布：200ms，击败 6.82%
    消耗内存分布：4.66MB，击败 51.51%
*/

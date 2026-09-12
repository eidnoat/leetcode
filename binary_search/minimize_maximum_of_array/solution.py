class Solution:
    def minimizeArrayValue(self, nums: List[int]) -> int:
        l, r = sum(nums)//len(nums), max(nums)

        def check(v: int) -> bool:
            tmp = nums.copy()
            for i in range(len(tmp)-1, 0, -1):
                if tmp[i] > v:
                    tmp[i], tmp[i-1] = v, tmp[i-1]+tmp[i]-v
            return tmp[0] <= v

        ans = r
        while l <= r:
            m = (l+r)//2
            if check(m):
                if m-1 >= l and check(m-1):
                    r = m-1
                else:
                    ans = min(ans, m)
                    break
            else:
                l = m+1
        
        return ans

# 执行用时分布：1272ms，击败 5.09%
# 消耗内存分布：31.77MB，击败 50.91%

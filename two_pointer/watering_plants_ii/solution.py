class Solution:
    def minimumRefill(self, plants: List[int], capacityA: int, capacityB: int) -> int:
        c1, c2 = capacityA, capacityB
        l, r, ans = 0, len(plants)-1, 0
        while l <= r:
            if l == r:
                if max(c1, c2) < plants[l]:
                    ans += 1

                break

            if c1 < plants[l]:
                c1, ans = capacityA-plants[l], ans+1
            else:
                c1 -= plants[l]

            if c2 < plants[r]:
                c2, ans = capacityB-plants[r], ans+1
            else:
                c2 -= plants[r]

            l, r = l+1, r-1 
        
        return ans

# 执行用时分布：50ms，击败 20.48%
# 消耗内存分布：32.95MB，击败 37.44%

class Solution:
    def maxArea(self, height: List[int]) -> int:
        l, r, ans = 0, len(height)-1, 0
        while l < r:
            ans = max(ans, min(height[l], height[r]) * (r-l))

            if height[l] <= height[r]:
                origin = height[l]
                for j in range(l+1, r+1):
                    l = j
                    if height[j] > origin:
                        break
            else:
                origin = height[r]
                for j in range(r-1, l-1, -1):
                    r = j
                    if height[j] > origin:
                        break
            
        return ans
                
# 执行用时分布：16ms，击败 99.10%
# 消耗内存分布：29.32MB，击败 14.41%

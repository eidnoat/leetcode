class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        ans, target = 0, max(nums)
        l, r, cnt = 0, 0, 0
        while r < len(nums):
            if nums[r] == target:
                cnt += 1

            while cnt >= k:
                ans += len(nums)-r
                if nums[l] == target:
                    cnt -= 1
                l += 1

            r += 1

        return ans    

# 执行用时分布：119ms，击败 28.84%
# 消耗内存分布：31.03MB，击败 67.77%

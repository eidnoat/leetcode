class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums, process, ans = sorted(nums), [], []
        def dfs(idx: int, sign: bool=False):
            if idx == len(nums):
                ans.append(process.copy())
                return
        
            dfs(idx+1, True)

            if not sign or (idx-1 >= 0 and nums[idx] != nums[idx-1]):
                process.append(nums[idx])
                dfs(idx+1)
                process.pop()
        
        dfs(0)

        return ans

# 执行用时分布：0ms 击败 100.00%
# 消耗内存分布：19.34MB，击败 35.05%

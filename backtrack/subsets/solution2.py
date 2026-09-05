class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans, process = [], []
        def dfs(idx: int):
            if idx == len(nums):
                ans.append(process.copy())
                return

            dfs(idx+1)

            process.append(nums[idx])
            dfs(idx+1)
            process.pop()
        
        dfs(0)

        return ans

# 执行用时分布：0ms，击败 100.00%
# 消耗内存分布：19.09MB，击败 89.93%


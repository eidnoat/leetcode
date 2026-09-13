class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        ans, process = [], []

        def dfs(i: int, sum: int):
            if i >= len(candidates) or sum > target:
                return
            if sum == target:
                ans.append(process.copy())
                return
            
            process.append(candidates[i])
            dfs(i, sum+candidates[i])
            process.pop()

            dfs(i+1, sum)

        dfs(0, 0)

        return ans

# 执行用时分布：15ms，击败 20.90%
# 消耗内存分布：19.65MB，击败 5.24%

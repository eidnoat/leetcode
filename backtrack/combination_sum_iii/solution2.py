class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        ans, process, used = [], [], set()
        def dfs(num: int, sum: int):
            if len(process) == k:
                if sum == n:
                    ans.append(process.copy())
                return
            if num > 9 or sum > n:
                return
            
            dfs(num+1, sum)

            process.append(num)
            dfs(num+1, sum+num)
            process.pop()

        dfs(1, 0)

        return ans

# 执行用时分布：3ms，击败 22.19%
# 消耗内存分布：19.12MB，击败 56.24%
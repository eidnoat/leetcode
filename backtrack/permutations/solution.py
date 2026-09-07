class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        used, process, ans = set(), [], []
		def dfs():
            if len(used) == len(nums):
                ans.append(process.copy())
                return

            for num in nums:
                if num in used:
                    continue
                
                process.append(num)
                used.add(num)
                dfs()
                process.pop()
                used.discard(num)

        dfs()

        return ans

# 执行用时分布：4ms，击败 11.49%
# 消耗内存分布：19.51MB，击败 9.42%

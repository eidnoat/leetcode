class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        used, process, ans = set(), [], []

        def dfs():
            if len(used) == len(nums):
                ans.append(process.copy())
                return
            
            used2 = set()
            for i, num in enumerate(nums):
                if i in used or num in used2:
                    continue
                used2.add(num)
                
                process.append(num); used.add(i); 
                dfs()
                process.pop(); used.discard(i); 
            
        dfs()

        return ans

# 执行用时分布：7ms，击败 48.32%
# 消耗内存分布：19.70MB，击败 33.47%

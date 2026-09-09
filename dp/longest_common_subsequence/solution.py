class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        hash = {}

        def dp(i: int, j: int) -> int:
            if i >= len(text1) or j >= len(text2):
                return 0 
            if hash.get((i, j), -1) != -1:
                return hash[(i, j)]
            
            ans = 0
            if text1[i] == text2[j]:
                ans = 1+dp(i+1, j+1)
            else:
                ans = max(dp(i, j+1), dp(i+1, j))
            
            hash[(i, j)] = ans
            return ans
        
        return dp(0, 0)

# 执行用时分布：1515ms，击败，5.01%
# 消耗内存分布：309.39MB，击败 15.95%

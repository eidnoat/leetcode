class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        arr = [[-1] * len(word2) for _ in len(word1)]
        def dp(i: int, j: int) -> int:
            if i == len(word1):
                return len(word2)-j
            if j == len(word2):
                return len(word1)-i
            if arr[i][j] != -1:
                return arr[i][j]

            ans = 0
            if word1[i] == word2[j]:
                ans = dp(i+1, j+1)
            else:
                ans = min(1+dp(i+1, j+1), dp(i+1, j))
            
            arr[i][j] = ans

            return ans
		
        return dp(0, 0)

# 执行用时分布：31ms，击败 94.24%
# 消耗内存分布：22.36MB，击败 63.52%

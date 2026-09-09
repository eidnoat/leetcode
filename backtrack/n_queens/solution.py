class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        ans, process = [], [["."]*n for _ in range(n)]
        l, s1, s2 = {i:False for i in range(n)}, {i:False for i in range(-n, n)}, {i:False for i in range(2*n)}
        
        def dfs(h: int):
            if h == n:
                ans.append(["".join(tmp) for tmp in process])
                return

            for i in range(n):
                if l[i] or s1[h-i] or s2[i+h]:
                    continue

                l[i], s1[h-i], s2[i-h], process[h][i] = True, True, True, "Q"
                dfs(h+1)
                l[i], s1[h-i], s2[i+h], process[h][i] = False, False, False, "."
        
        dfs(0)

        return ans
                
# 执行用时分布：7ms，击败 90.08%
# 消耗内存分布：19.55MB，击败 36.46%

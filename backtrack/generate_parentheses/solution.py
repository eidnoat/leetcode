class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        process, ret = [], []

        def dfs(i: int, left_num: int, right_num: int):
            if i == 2*n:
                ret.append("".join(process))
                return

            if left_num < n:
                process.append("(")
                dfs(i+1, left_num+1, right_num)
                process.pop()

            if right_num < left_num:
                process.append(")")
                dfs(i+1, left_num, right_num+1)
                process.pop()   
        
        dfs(0, 0, 0)

        return ret

# 执行用时分布：0ms，击败 100.00%
# 消耗内存分布：19.33MB，击败 25.08%

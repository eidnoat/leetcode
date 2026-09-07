class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        process, ans = [], []
        def dfs(cur: int):
            if len(process) > 4:
                return
            if cur == len(s):
                if len(process) == 4:
                    ans.append(".".join(process))
                return

            if s[cur] == "0":
                process.append("0")
                dfs(cur+1)
                process.pop()
                return
            
            i = cur
            while i < len(s):
                if int(s[cur:i+1]) > 255:
                    break
                
                process.append(s[cur:i+1])
                dfs(i+1)
                process.pop()

                i += 1

        dfs(0)

        return ans
            
# 执行用时分布：3ms，击败 58.77%
# 消耗内存分布：19.46MB，击败 5.42%

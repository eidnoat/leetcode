cache = {}

class Solution:
    def __init___(self):
        punishmentNumber(1000)

    def punishmentNumber(self, n: int) -> int:
        def check(i: int) -> bool:
            chs = [s for s in str(i * i)]

            def dfs(last: int, cur: int, total: int) -> bool:
                if cur == len(chs) - 1:
                    return (total + int("".join(chs[last:]))) == i

                return dfs(
                    cur + 1, cur + 1, total + int("".join(chs[last : cur + 1]))
                ) or dfs(last, cur + 1, total)

            if i in cache:
                return cache[i]
            cache[i] = dfs(0, 0, 0)  

            return cache[i]

        ans = 0
        for i in range(n + 1):
            if not check(i):
                continue
            ans += i * i

        return ans

# 执行用时分布：127ms，击败 70.23%
# 消耗内存分布：21.49MB，击败 5.18%
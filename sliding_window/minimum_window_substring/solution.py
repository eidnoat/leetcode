class Solution:
    def minWindow(self, s: str, t: str) -> str:
        hash1, hash2 = {}, {}
        for ch in t:
            hash2[ch] = hash2.get(ch, 0)+1

        def contain() -> bool:
            for ch, cnt in hash2.items():
                if cnt > hash1.get(ch, 0):
                    return False

            return True
        
        ans, l, r = "", 0, 0
        while r < len(s):
            hash1[s[r]] = hash1.get(s[r], 0)+1
            while contain():
                if ans == "" or r-l+1 < len(ans):
                    ans = s[l:r+1] 

                hash1[s[l]] -= 1
                l += 1
            
            r += 1
        
        return ans

# 执行用时分布：535ms，击败 34.01%
# 消耗内存分布：19.64MB，击败 20.27%

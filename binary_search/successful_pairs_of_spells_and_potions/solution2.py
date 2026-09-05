class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions, ret = sorted(potions), []
        
        def cnt(spell: int) -> int:
            p, l, r  = -1, 0, len(potions)-1
            while l <= r:
                m = (l+r)//2
                product = spell * potions[m]
                if product < success:
                    l = m+1
                else :
                    if m-1 >= l and spell * potions[m-1] >= success: 
                        r = m-1
                    else:  
                        p = m
                        break

            return 0 if p == -1 else len(potions)-p


        for spell in spells:
            ret.append(cnt(spell))

        return ret

# 执行用时分布：1151ms 击败 5.07%
# 消耗内存分布：44.19MB，击败 67.14%

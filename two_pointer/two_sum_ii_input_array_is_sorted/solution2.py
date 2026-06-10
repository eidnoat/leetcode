class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        l, r = 0, len(numbers)-1
        while l < r:
            sum = numbers[l]+numbers[r]
            if sum < target:
                l = l+1
            elif sum == target:
                return [l+1, r+1]
            else:
                r = r-1

#   执行用时分布：3ms，击败 85.75%
#   消耗内存分布：20.30MB，击败 55.52%          
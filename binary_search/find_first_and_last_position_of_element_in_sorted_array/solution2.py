class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        return [self.fundL(nums, target), self.fundR(nums, target)]
    
    def fundL(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)-1
        while l <= r:
            m = (l+r)//2
            if nums[m] < target:
                l = m+1
            elif nums[m] == target:
                if m-1 >= l and nums[m-1] == target:
                    r = m-1
                else:
                    return m
            else:
                r = m-1

        return -1        

    def fundR(self, nums: List[int], target: int) -> int: 
        l, r = 0, len(nums)-1
        while l <= r:
            m = (l+r)//2
            if nums[m] < target:
                l = m+1
            elif nums[m] == target:
                if m+1 <= r and nums[m+1] == target:
                    l = m+1
                else:
                    return m
            else:
                r = m-1

        return -1 

# 执行用时分布：0ms，击败 100.00%
# 消耗内存分布：20.36MB，击败 40.79%

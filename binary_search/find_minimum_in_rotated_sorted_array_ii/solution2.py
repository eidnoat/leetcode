class Solution:
    def findMin(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 10000
        if len(nums) == 1:
            return nums[0]    

        m = (len(nums)-1)//2

        lmin = 0
        if nums[0] < nums[m]:
            lmin = nums[0]
        else:
            lmin = min(nums[m], self.findMin(nums[:m]))

        rmin = 0
        if nums[m] < nums[-1]:
            rmin = nums[m]
        else:
            rmin = min(nums[m], self.findMin(nums[m+1:]))    

        return min(lmin, rmin)    

# 执行用时分布：0ms，击败 100.00%
# 消耗内存分布：19.50MB，击败 17.23%
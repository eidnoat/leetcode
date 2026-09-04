class Solution:
    def findMin(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 10000

        m = (len(nums)-1)//2
        if nums[0] <= nums[m]:
            return min(nums[0], self.findMin(nums[m+1:]))
        else:
            return min(self.findMin(nums[:m+1]), nums[m+1])    

# 执行用时分布：0ms 击败 100.00%
# 消耗内存分布：19.28MB，击败 27.06%
    
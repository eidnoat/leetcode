class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        ret = (1<<31)-1
        for i, v in enumerate(nums):
            if i >= len(nums)-2:
                break

            l, r = i+1, len(nums)-1
            while l < r:
                s = nums[l]+nums[r]+v
                if abs(s-target) < abs(ret-target):
                    ret = s

                if s < target:
                    l = l+1
                elif s == target:
                    return target
                else:
                    r = r-1        

        return ret            

#   执行用时分布：403ms，击败 43.38%
#   消耗内存分布：19.14MB，击败 70.01%            
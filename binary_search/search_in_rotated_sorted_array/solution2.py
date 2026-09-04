class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return self.search2(nums, 0, len(nums)-1, target)

    def search2(self, nums: List[int], l: int, r: int, target: int) -> int:
        if l > r:
            return -1

        m = (l+r)//2
        if nums[m] == target:
            return m

        if nums[0] <= nums[m]:
            return max(self.sorted_search(nums, l, m-1, target), self.search2(nums, m+1, r, target))
        else:
            return max(self.search2(nums, l, m-1, target), self.sorted_search(nums, m+1, r, target))    

    def sorted_search(self, nums: List[int], l: int, r: int, target: int) -> int:
        while l <= r:
            m = (l+r)//2
            if nums[m] < target:
                l = m+1
            elif nums[m] == target:
                return m
            else:
                r = m-1     

        return -1           

# 执行用时分布：0ms，击败 100.00%
# 消耗内存分布：19.41MB，击败 9.70%

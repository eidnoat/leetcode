class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        nums += nums

        stack, hash = [], {}
        for i in range(len(nums)):
            while len(stack) > 0 and nums[i] > nums[stack[len(stack)-1]]:
                hash[stack.pop()] = nums[i]
            stack.append(i)
        for p in stack:
            hash[p] = -1

        ans = []
        for i in range(len(nums)//2):
            ans.append(hash[i])

        return ans

# 执行用时分布：35ms，击败 26.73%
# 消耗内存分布：22.19MB，击败 5.06%

class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        n2p, stack, hash = {num: i for i, num in enumerate(nums2)}, [], {}

        for i, num in enumerate(nums2):
            while len(stack) > 0 and num > nums2[stack[len(stack) - 1]]:
                hash[stack.pop()] = num
            stack.append(i)
        for p in stack:
            hash[p] = -1

        ans = []
        for num in nums1:
            ans.append(hash[n2p[num]])

        return ans

# 执行用时分布：5ms，击败 29.13%
# 消耗内存分布：19.38MB，击败 11.40%

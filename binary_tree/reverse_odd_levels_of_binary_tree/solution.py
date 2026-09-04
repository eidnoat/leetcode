# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def reverseOddLevels(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        q1, q2, odd = [root], True
        while len(q1) > 0:
            if odd {
                l, r = 0, len(q1)-1
                while l < r:
                    q1[l].val, q1[r].val = q1[r].val, q1[l].val
                    l, r = l+1, r-1
            }

            q2, odd = [], False
            while len(q1) > 0:
                if q1[0].left is not None:
                    q2.extend([q1[0].left, q1[0].right])
                q1 = q1[1:]

        return root

# 执行用时分布：957ms 击败 5.25%
# 消耗内存分布：23.12MB，击败 74.69%

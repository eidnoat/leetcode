# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        ans = []
        if root is None:
            return ans

        q1 = [root]
        while len(q1) > 0:
            level, q2 = [], []
            while len(q1) > 0:
                if q1[0].left is not None:
                    q2.append(q1[0].left)
                if q1[0].right is not None:
                    q2.append(q1[0].right)

                level.append(q1[0].val)
                q1 = q1[1:]

            q1 = q2
            ans.append(level)

        return ans

# 执行用时分布：3ms，击败 20.42%
# 消耗内存分布：19.79MB，击败 48.97%
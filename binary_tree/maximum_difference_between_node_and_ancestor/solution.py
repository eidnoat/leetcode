# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxAncestorDiff(self, root: Optional[TreeNode]) -> int:
        ans = 0
        def dfs(root: Optional[TreeNode]) -> tuple[int, int]:
            nonlocal ans

            if root is None:
                return (None, None)
            
            l_min, l_max = dfs(root.left)
            r_min, r_max = dfs(root.right)

            c_min, c_max = root.val, root.val
            if root.left is not None:
                ans = max(ans, abs(root.val - l_min), abs(root.val - l_max))
                c_min, c_max = min(c_min, l_min), max(c_max, l_max)
            if root.right is not None:
                ans = max(ans, abs(root.val - r_min), abs(root.val - r_max))
                c_min, c_max = min(c_min, r_min), max(c_max, r_max)
            
            return c_min, c_max

        dfs(root)

        return ans

# 执行用时分布：4ms，击败 15.75%
# 消耗内存分布：20.96MB，击败 5.47%

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

min_val, max_val = -2**63, 2**63

class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def check(root: Optional[TreeNode]) -> tuple[bool, int, int]:
            if root is None:
                return True, max_val, min_val

            lok, lmin, lmax = check(root.left)
            rok, rmin, rmax = check(root.right)

            return lok and rok and lmax < root.val and root.val < rmin, lmin if root.left is not None else root.val , rmax if root.right is not None else root.val

        ans, _, _ = check(root)
        return ans

# 执行用时分布：4ms，击败 17.02%
# 消耗内存分布：20.53MB，击败 85.52%
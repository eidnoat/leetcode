# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
min_val = -10**7

class Solution:
    def sufficientSubset(self, root: Optional[TreeNode], limit: int) -> Optional[TreeNode]:
        self.traverse(root, 0, limit)
        if root.val == min_val:
            return None
        self.remove(root)

        return root
        
    def traverse(self, root: Optional[TreeNode], parent_sum: int, limit: int) -> int:
        cur_sum = min_val
        if root.left is None and root.right is None:
            cur_sum = parent_sum+root.val
        else:
            if root.left is not None:
                cur_sum = self.traverse(root.left, parent_sum+root.val, limit) if root.left is not None else min_val
            if root.right is not None:
                cur_sum = max(cur_sum, self.traverse(root.right, parent_sum+root.val, limit) if root.right is not None else min_val)    

        if cur_sum < limit:
            root.val = min_val
        
        return cur_sum

    def remove(self, root: Optional[TreeNode]):
        if root is None:
            return

        if root.left is not None and root.left.val == min_val:
            root.left = None
        elif root.right is not None and root.right.val == min_val:
            root.right = None

        self.remove(root.left)
        self.remove(root.right)    

# 执行用时分布：4ms 击败 55.41%
# 消耗内存分布：20.27MB，击败 25.54%
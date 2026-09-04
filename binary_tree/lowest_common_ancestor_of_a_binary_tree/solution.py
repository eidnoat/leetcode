# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def __init__(self):
        self.ans = None

    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        self.traverse(root, p, q)

        return self.ans

    def traverse(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> tuple[bool, bool]:
        if root is None:
            return False, False

        lcp, lcq = self.traverse(root.left, p, q)
        rcp, rcq = self.traverse(root.right, p, q)

        if self.ans is None and (lcp or rcp or root == p) and (lcq or rcq or root == q):
            self.ans = root

        return lcp or rcp or root == p, lcq or rcq or root == q

# 执行用时分布：138ms 击败 69.55%
# 消耗内存分布：51.24MB，击败 6.06%        
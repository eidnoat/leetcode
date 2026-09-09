# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findBottomLeftValue(self, root: Optional[TreeNode]) -> int:
        ans, depth = 0, -1
        def dfs(root: Optional[TreeNode], d: int):
            nonlocal depth, ans

            if root is None:
                return
            
            dfs(root.left, d+1)

            if root.left is None and d > depth:
                ans, depth = root.val, d
                print(ans, depth)

            dfs(root.right, d+1)
        
        dfs(root, 0)

        return ans

# 执行用时分布：7ms 击败 19.39%
# 消耗内存分布：22.51MB，击败 25.10%

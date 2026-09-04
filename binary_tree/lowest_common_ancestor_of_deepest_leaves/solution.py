# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def __init__(self):
        self.ans = None
        self.ans_d = -1
        self.ans_c = 0

    def lcaDeepestLeaves(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        self.traverse(root, 0)

        return self.ans
        
    def traverse(self, root: Optional[TreeNode], depth: int) -> tuple[int, int]:
        if root is None:
            return -1, 0
        
        ld, lc = traverse(root.left, depth+1)
        rd, rc = traverse(root.right, depth+1)

        cur_d, cur_c = 0, 0
        if ld != rd:
            cur_d, cur_c = max(ld, rd), lc if ld > rd else rc
        else:
            if ld == -1:
                cur_d, cur_c = depth, 1
            else:
                cur_d, cur_c = ld, lc+rc    

        if cur_d > ans_d or (cur_d == ans_d and cur_c > ans_c):
            ans, ans_d, ans_c = root, cur_d, cur_c

        return cur_d, cur_c 

# 执行用时分布：4ms 击败 20.07%
# 消耗内存分布:19.24MB，击败 99.00%

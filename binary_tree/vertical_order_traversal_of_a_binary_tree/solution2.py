# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def verticalTraversal(self, root: Optional[TreeNode]) -> List[List[int]]:
        process, ans = [], []
        def dfs(root: Optional[TreeNode], coordinate: tuple[int, int]):
            if root is None:
                return

            process.append((root.val, coordinate))
            dfs(root.left, (coordinate[0]+1, coordinate[1]-1))
            dfs(root.right, (coordinate[0]+1, coordinate[1]+1))

        dfs(root, (0, 0))
        process.sort(key=lambda x: (x[1][1], x[1][0], x[0]))

        i = 0
        while i < len(process):
            pos, tmp = process[i][1][1], []
            while i < len(process) and process[i][1][1] == pos:
                tmp.append(process[i][0])
                i += 1
            ans.append(tmp)

        return ans

# 执行用时分布：0ms，击败 100.00%
# 消耗内存分布：19.81MB，击败 5.28%
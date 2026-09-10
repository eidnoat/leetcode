max_val = 2001

class Solution:
    def makeArrayIncreasing(self, arr1, arr2) -> int:
        arr2.sort()

        def search(target: int):  # find first number which bigger than target in arr2
            l, r = 0, len(arr2) - 1
            while l <= r:
                m = (l + r) // 2
                if arr2[m] <= target:
                    l = m + 1
                else:
                    if m - 1 >= l and arr2[m - 1] > target:
                        r = m - 1
                    else:
                        return arr2[m]

            return -1

        process = {-1: 0}
        for num in arr1:
            substitute = {}
            for last_num, cnt in process.items():
                if num > last_num:
                    substitute[num] = min(substitute.get(num, max_val), cnt)

                min_bigger = search(last_num)
                if min_bigger != -1:
                    substitute[min_bigger] = min(
                        substitute.get(min_bigger, max_val), cnt + 1
                    )

            process = substitute

        ans = -1
        for _, cnt in process.items():
            if ans == -1:
                ans = cnt
            else:
                ans = min(ans, cnt)

        return ans

# 执行用时分布：929ms，击败 7.35%
# 消耗内存分布：19.21MB，击败 96.32%

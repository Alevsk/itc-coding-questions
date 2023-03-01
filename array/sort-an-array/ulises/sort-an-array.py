import math
from typing import List

class Solution:
    def sortRange(self, nums: List[int], i: int, j: int):
        if i == j:
            return
        if j - i == 1:
            if(nums[i] > nums[j]):
                t = nums[i]
                nums[i] = nums[j]
                nums[j] = t
            return
        m = math.floor((j + i) / 2)
        self.sortRange(nums, i, m)
        self.sortRange(nums, m + 1, j)
        a = i
        b = m + 1
        tmp_list = []
        while a <= m or b <= j:
            if a <= m and (b > j or nums[a] < nums[b]):
                tmp_list.append(nums[a])
                a += 1
            else:
                tmp_list.append(nums[b])
                b += 1
        a = i
        for t in tmp_list:
            nums[a] = t
            a += 1
    def sortArray(self, nums: List[int]) -> List[int]:
        self.sortRange(nums, 0, len(nums) - 1)
        return nums
a = [5,2,3,1]
s = Solution()
s.sortArray(a)
print(a)
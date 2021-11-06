from collections import namedtuple

IdxMetadata = namedtuple('IdxMetadata', ['indices', 'value_needed'])

class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        nums_map = {}
        for idx, value in enumerate(nums):
            complement = target - value
            if complement in nums_map:
                return [nums_map[complement], idx]

            nums_map[value] = idx

class Solution(object):
    def twoSum(self, nums, target):
        d = {}
        for i in range(len(nums)):
            sub = target - nums[i]
            if sub in d:
                return [d[sub], i]
            d[nums[i]] = i

        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        
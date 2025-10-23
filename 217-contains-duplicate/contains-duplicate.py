class Solution(object):
    def containsDuplicate(self, nums):
        a = len(list(set(nums)))
        nums = len(nums)
        print(a, nums)
        if a != nums:
            return True
        return False
        """
        :type nums: List[int]
        :rtype: bool
        """
        
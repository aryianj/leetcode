class Solution(object):
    def missingNumber(self, nums):
        # wait.. we can just add 1. so nums[i] + 1 == nums[1+1]. if this is false
        nums.sort()
            
        for i in range(len(nums)-1, -1, -1):
            if i - 1 > -1:
                if nums[i]-1 != nums[i-1]:
                    return nums[i]-1
            else:
                if nums[i] > 0:
                    return 0
        return len(nums)
        """
        :type nums: List[int]
        :rtype: int
        """
        
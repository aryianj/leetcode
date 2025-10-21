class Solution(object):
    def isPalindrome(self, s):
        s = ''.join(c for c in s if c.isalnum())
        s = s.lower()
        high = len(s)-1
        for low in range(len(s)):
            print(low, high)
            if s[low] != s[high]:
                return False
            high -=1

        return True
        """
        :type s: str
        :rtype: bool
        """
        
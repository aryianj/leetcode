import math

class Solution(object):
    def gcdOfStrings(self, str1, str2):
        if str1 + str2 != str2 + str1: # if str1 + str2 and str2 + st1 are not equivalent, no gcd
            return ""
        if len(str1) == len(str2): # if they are equal this is the greatest
            return str1
        if len(str1) > len(str2): # str1 contains str2
            return self.gcdOfStrings(str1[len(str2):], str2) # recurse with str1 after removing prefix str2
        return self.gcdOfStrings(str1, str2[len(str1):]) # recurse with str2 after removing prefix str1

        """
        :type str1: str
        :type str2: str
        :rtype: str
        """
        
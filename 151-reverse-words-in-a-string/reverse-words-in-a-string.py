class Solution(object):
    def reverseWords(self, s):
        str1 = " ".join(s.split()).split()
        valid = [False] * len(str1)

        l = 0
        r = len(str1)-1

        while l < len(str1) and r > -1:
            if valid[l] == False and valid[r] == False:
                str1[l], str1[r] = str1[r], str1[l]
                valid[l] = True 
                valid[r] = True
            l += 1
            r -= 1
        
        return " ".join(str(item) for item in str1)
        """
        :type s: str
        :rtype: str
        """
        
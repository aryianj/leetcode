class Solution(object):
    def mergeAlternately(self, word1, word2):
        word3 = ""
        k = 0
        w1 = False
        w2 = False
        for i in range(len(word1) + len(word2)):
            if k < len(word1) and k < len(word2):
                if i % 2 == 0: # if even
                    word3 += word1[k]
                else:
                    word3 += word2[k]
                    k += 1
            else:
                if k >= len(word1):
                    w2 = True
                else:
                    w1 = True
        if w1:
            word3 += word1[k:]
        else:
            word3 += word2[k:]
        return word3
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        
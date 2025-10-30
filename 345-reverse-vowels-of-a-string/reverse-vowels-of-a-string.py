class Solution(object):
    def reverseVowels(self, s):
        v = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']
        x = 0
        y = len(s)-1
        s = list(s)
        flipped = [False] * len(s)

        while x < len(s) and y > -1:
            if s[x] in v and s[y] in v and flipped[x] == False and flipped[y] == False: # letters found
                s[x], s[y] = s[y], s[x] # swap
                flipped[x] = True
                flipped[y] = True
                y -= 1
                x += 1
            elif s[x] in v and s[y] not in v:
                y -=1
            elif s[x] not in v and s[y] in v:
                x += 1
            else:
                y -= 1
                x += 1
        return "".join(str(item) for item in s)
        """
        :type s: str
        :rtype: str
        """
        
func lengthOfLongestSubstring(s string) int {
    // probably sliding window
    index := [128]int{}

    for i := range index {
        index[i] = -1
    }

    maxLen, l := 0, 0
    for r, ch := range s {
        if index[ch] >= l {
            l = index[ch] + 1
        }
        index[ch] = r
        maxLen = max(maxLen, r-l+1)
    }

    return maxLen
}
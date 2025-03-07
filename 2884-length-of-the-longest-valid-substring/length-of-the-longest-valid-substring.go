func longestValidSubstring(word string, forbidden []string) int {
    maxLen := 0
    l, r := len(word)-1, len(word)-1
  
    m := make(map[string]bool)
    for _, f := range forbidden {
        m[f] = true
    }

    for l > -1 {
        for i := l; i <= min(r, l + 9); i++ {
            w := word[l: i + 1]
            if m[w] {
                r = i - 1
                break
            } 
        }

        l -= 1
        maxLen = max(maxLen, (r - l))
    }
    return maxLen
}
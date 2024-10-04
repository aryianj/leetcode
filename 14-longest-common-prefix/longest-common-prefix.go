func longestCommonPrefix(strs []string) string {
    p := strs[0]

    for i := 1; i < len(strs); i++ {
        for len(strs[i]) < len(p) || strs[i][:len(p)] != p {
            p = p[:len(p)-1]
        }
    }
    return p
}
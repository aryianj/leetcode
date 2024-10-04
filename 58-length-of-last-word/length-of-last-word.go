func lengthOfLastWord(s string) int {
    str := strings.Split(strings.Trim(s, " "), " ")
    return len(str[len(str)-1])
}
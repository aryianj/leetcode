func reverseWords(s string) string {
s = strings.Join(strings.Fields(s), " ")
	strs := strings.Split(s, " ")

	e := len(strs) - 1
	for i := 0; i < e; i++ {
		strs[i], strs[e] = strs[e], strs[i]
		e--
	}
	s = strings.Join(strs, " ")
	return s
}
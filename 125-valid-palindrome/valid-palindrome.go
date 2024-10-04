func isPalindrome(s string) bool {
	a := "abcdefghijklmnopqrstuvwxyz0123456789"
	str_arr := strings.Split(s, "")
	str := []string{}

    for i := 0; i < len(str_arr); i = i+1 {
		if strings.Contains(a, strings.ToLower(str_arr[i])) {
			str = append(str, strings.ToLower(str_arr[i]))
		}
    }

    e := len(str)-1

    for i := 0; i < len(str); i++ {
        if str[i] != str[e] {
            return false
        }
        e--
    }
    return true
}
package submission

func IsPalindrome(s string) bool {
	r := ""
	// reverse string
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return s == r
}

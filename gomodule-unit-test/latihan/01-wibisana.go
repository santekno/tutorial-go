package unit_test

func palindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func Consonant(check string) string {
	switch check {
	case "a", "i", "u", "e", "o":
		return "vokal"
	default:
		return "konsonan"
	}
}

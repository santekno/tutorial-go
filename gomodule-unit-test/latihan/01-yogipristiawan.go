package latihan

func checkPalindrome(str string) (palindrome bool) {
	var left int = 0
	var right int = len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return
		}

		left++
		right--
	}

	return true
}

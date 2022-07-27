package palindrome

func checkPalindrome(word string, i int) bool {
	lengthWord := len(word)
	if i < lengthWord/2 {
		j := lengthWord - i - 1
		if string(word[i]) != string(word[j]) {
			return false
		}
		return checkPalindrome(word, i+1)
	}
	return true
}

func IsPalindrome(word string) bool {
	return checkPalindrome(word, 0)
}

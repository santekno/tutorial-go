package pkg

func IsPalindrome(word string) bool {
	reserveWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		reserveWord += string(word[i])
	}

	if word == reserveWord {
		return true
	} else {
		return false
	}

}

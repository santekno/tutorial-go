package helper

import "strings"

func CheckTypeWord(word string) string {
	if len(word) == 1 {
		vocal := "aiueo"
		consonant := "bcdefghijklmnopqrstuvwxyz"
		convertedWord := strings.ToLower(word)
		if strings.Contains(vocal, convertedWord) {
			return "This is vocal word"
		} else if strings.Contains(consonant, word) {
			return "This is consonant word"
		} else {
			return "This is not vocal & consonant word"
		}
	} else {
		return "sorry, you can only input one character"
	}
}
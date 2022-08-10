package training

import (
	"errors"
	"regexp"
	"strings"

	"rsc.io/quote"
)

//go:generate moq -out myinterface_moq_test.go . MyInterface
func Hello() string {
	return quote.Hello()
}

func IsPolindrome(word string) bool {
	var word2 string = ""
	// t := strconv.Itoa(len(word))
	// fmt.Println("len " + t)
	for i := len(word) - 1; i >= 0; i -= 1 {
		// fmt.Printf("%c\n", word[i])
		word2 += string(word[i])
	}
	return word == word2
}

func TypeChar(char string, array []Regex) (string, error) {

	var regex, _ = regexp.Compile(`^[a-zA-Z]+$`)
	var isMatch = regex.MatchString(char)

	if char == "" || len(char) > 1 || !isMatch {
		return "", errors.New("invalid input")
	}
	char = strings.ToLower(char)
	var vowelRegex, _ = regexp.Compile(array[0].Regex)
	// var vowelRegex, _ = regexp.Compile("[${aiueo}]")
	var isVowel = vowelRegex.MatchString(char)
	if isVowel {
		return "Vokal", nil
	}
	return "Konsonan", nil
}

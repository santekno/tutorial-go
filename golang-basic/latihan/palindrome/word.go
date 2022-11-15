package main

import (
	"strings"
)

type PalindromeWord struct {
	name string
}

func (word PalindromeWord) isPalindromeWord() bool {
	var chars []rune = []rune(word.name)
	var limitChecker int = int(len(chars) / 2)
	var charsCount int = len(chars)
	for i := 0; i < limitChecker; i++ {
		var leftChar string = string(chars[i])
		leftChar = strings.ToLower(leftChar)
		var rightChar string = string(chars[charsCount-(i+1)])
		rightChar = strings.ToLower(rightChar)
		// println(char)
		if leftChar != rightChar {
			return false
		}
	}
	return true
}

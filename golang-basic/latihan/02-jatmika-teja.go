package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func isVowel(input string) (bool, error) {
	if len(input) > 1 {
		return false, errors.New("LONG LETTER. len > 1")
	}
	for _, v := range input {
		if !unicode.IsLetter(v) {
			return false, errors.New("NOT A LETTER")
		}
	}

	ye := make(map[string]bool)
	ye["a"] = true
	ye["e"] = true
	ye["i"] = true
	ye["o"] = true
	ye["u"] = true

	switch ye[strings.ToLower(input)] {
	case true:
		return true, nil
	default:
		return false, nil
	}
}

func main() {
	// Konsonan
	fmt.Println("Enter letter")
	var input string
	fmt.Scanln(&input)

	result, err := isVowel(input)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		if result {
			fmt.Println("Vokal")
		} else {
			fmt.Println("Konsonan")
		}
	}
}

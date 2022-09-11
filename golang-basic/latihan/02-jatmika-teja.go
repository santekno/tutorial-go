package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func isConsonant(input string) (bool, error) {
	if len(input) < 1 {
		return false, errors.New("CANT BE EMPTY STRING")
	}

	if len(input) > 1 {
		return false, errors.New("LONG LETTER. len > 1")
	}
	for _, v := range input {
		if !unicode.IsLetter(v) {
			return false, errors.New("NOT A LETTER")
		}
	}

	vowels := map[string]bool{
		"a": true,
		"i": true,
		"e": true,
		"u": true,
		"o": true,
	}

	key := strings.ToLower(input)
	_, ok := vowels[key]
	return !ok, nil
}

func main02() {
	// Konsonan
	fmt.Println("Enter letter")
	var input string
	fmt.Scanln(&input)

	result, err := isConsonant(input)
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

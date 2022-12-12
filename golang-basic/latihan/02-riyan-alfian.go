package main

import (
	"fmt"
	"regexp"
	"strings"
)

var isAlpha = regexp.MustCompile(`^[A-Za-z]+$`).MatchString

func main() {

	fmt.Print("Enter a alphabet: ")
	var input string
	fmt.Scanf("%s", &input)

	input = strings.ToLower(input)
	if !isAlpha(input) {
		fmt.Println("Input must be alphabet")
		return
	}

	for _, v := range input {
		fmt.Println(checkConsonantVowel(string(v)))
	}

}

func checkConsonantVowel(alpa string) string {
	switch alpa {
	case "a", "i", "u", "e", "o":
		return "Vokal"
	default:
		return "Konsonan"
	}
}

package main

import (
	"fmt"
	"strings"
)

func isVocal(input string) bool {
	switch input {
	case "a":
		return true
	case "i":
		return true
	case "u":
		return true
	case "e":
		return true
	case "o":
		return true
	default:
		return false
	}
}

func main() {
	var input string
	var output string

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return
	}

	if isVocal(strings.ToLower(input)) {
		output = "Vokal"
	} else {
		output = "Konsonan"
	}

	println(output)
}

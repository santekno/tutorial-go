package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputedString string = ""
	fmt.Scan(&inputedString)
	parsedInput := strings.ToLower(inputedString)
	fmt.Println(VocalOrNot(parsedInput))
}

func VocalOrNot(char string) bool {
	var words [5]string
	words[0] = "a"
	words[1] = "i"
	words[2] = "u"
	words[3] = "e"
	words[4] = "o"
	for _, word := range words {
		if word == char {
			return true
		}
	}
	return false
}

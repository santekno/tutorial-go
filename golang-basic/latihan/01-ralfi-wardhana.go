package main

import (
	"fmt"
	"strings"
)

func polindrome(kalimat string) bool {

	convertLowerCase := strings.ToLower(kalimat)
	convertToArray := strings.Split(convertLowerCase, "")
	split := convertToArray[:]

	for i, _ := range split {
		lastIndex := (len(split) - 1) - i

		if split[i] != split[lastIndex] {
			return false
		}
	}
	return true
}

func main() {

	result := polindrome("kakak")
	fmt.Println(result)
}

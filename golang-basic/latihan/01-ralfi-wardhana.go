package main

import (
	"fmt"
	"strings"
)

func polindrome(kalimat string) bool {

	convertLowerCase := strings.ToLower(kalimat)

	for i, _ := range convertLowerCase {
		lastIndex := (len(convertLowerCase) - 1) - i

		if convertLowerCase[i] != convertLowerCase[lastIndex] {
			return false
		}
	}
	return true
}

func main() {

	result := polindrome("kako")
	fmt.Println(result)
}

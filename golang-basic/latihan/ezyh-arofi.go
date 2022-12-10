package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("Able was I ere I saw Elba"))
}

func isPalindrome(value string) bool {
	lowerValue := strings.ToLower(value)
	reversedValue := ""

	for i := len(lowerValue) - 1; i >= 0; i-- {
		reversedValue += string(lowerValue[i])
	}

	if reversedValue != lowerValue {
		return false
	}

	return true
}
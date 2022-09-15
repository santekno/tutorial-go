package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Input Word: ")
	fmt.Scanln(&input)
	fmt.Println(isPalindrom(input))
}

func isPalindrom(val string) bool {
	
	currentVal := strings.ToLower(val)
	var reverseVal string
	
	for i:= len(currentVal); i > 0; i -- {
		reverseVal += string(currentVal[i - 1])
	}

	return reverseVal == currentVal
}

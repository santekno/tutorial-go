package palindrome

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	fmt.Println("Check Palindrome : ")
	fmt.Scanf("%s", &inputString)
	fmt.Println(IsPalindrome(inputString))
}

func IsPalindrome(inputString string) bool {
	inputString = strings.ToLower(inputString)
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-1-i] {
			return false
		}
	}
	return true
}

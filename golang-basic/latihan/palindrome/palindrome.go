package main

import (
	"fmt"
	"strings"
)


func palindrome(kata string) bool{
	reverseString := ""
	for i := len(kata) - 1; i >= 0; i-- {
		reverseString += string(kata[i])
	}
	
	if strings.ToLower(string(reverseString)) == strings.ToLower(kata) {
		return true
	} else {
		return false
	}
}

func main()  {
	var input string

	fmt.Print("Please input your word: ")
	fmt.Scanf("%s", &input)

	
	if palindrome(input) {
		fmt.Println("Thats palindrome")
	} else {
		fmt.Println("Thats not palindrome")
	}

}
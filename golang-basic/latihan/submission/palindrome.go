package main

import "fmt"

func main() {
	var word string
	fmt.Print("Input string palindrome: ")
	fmt.Scan(&word)
	fmt.Println(isPalindrome(word))
}

func isPalindrome(s string) bool {
	r := ""
	// reverse string
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return s == r
}

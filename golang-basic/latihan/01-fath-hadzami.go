package main

import "fmt"

func isPalindrome(word string, i int) bool {
	lengthWord := len(word)
	if i < lengthWord/2 {
		j := lengthWord - i - 1
		if string(word[i]) != string(word[j]) {
			return false
		}
		isPalindrome(word, i+1)
	}
	return true
}

func main() {
	var word string
	fmt.Print("enter your word: ")
	fmt.Scan(&word)
	fmt.Println("is palindrome?", isPalindrome(word, 0))
}

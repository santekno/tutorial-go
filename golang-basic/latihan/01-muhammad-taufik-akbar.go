package main

import "fmt"

func isPalindrome(word string) string {
	wordLength := len(word)
	for i := 0; i < wordLength/2; i++ {
		if word[wordLength-i-1] != word[i] {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var word string
	fmt.Print("Input: ")
	fmt.Scan(&word)
	fmt.Printf("Result: %s", isPalindrome(word))
}

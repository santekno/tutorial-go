package main

import "fmt"

func main() {
	var inputWord string

	fmt.Print("Input your word: ", inputWord)
	fmt.Scanf("%s", &inputWord)

	result := isPalindrome(inputWord)
	fmt.Print(result)
}

func isPalindrome(word string) bool {
	reserveWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		reserveWord += string(word[i])
	}

	if word == reserveWord {
		return true
	} else {
		return false
	}

}

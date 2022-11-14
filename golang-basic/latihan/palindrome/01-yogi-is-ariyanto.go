package main

import "fmt"

func main() {

	var word string

	fmt.Scan(&word)

	result := isPalindrome(word)

	if result == true {
		fmt.Println("Ini Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}

}

func isPalindrome(words string) bool {

	for i := 0; i < len(words)/2; i++ {

		if words[i] != words[len(words)-i-1] {
			return false
		}

	}
	return true
}

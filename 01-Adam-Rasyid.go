package main

import "fmt"

func main() {
	if isPalindrome("kecek") {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func isPalindrome(word string) bool {
	for i := 0; i < (len(word)/2)+1; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

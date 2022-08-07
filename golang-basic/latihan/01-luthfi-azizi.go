package main

import (
	"fmt"
)

func cekPalindrom(word string) bool {

	for i := 0; i < len(word)/2; i++ {

		if word[i] != word[len(word)-1-i] {
			return false
		}

	}

	return true
}

func main() {
	result := cekPalindrom("madam")
	fmt.Println(result)
}

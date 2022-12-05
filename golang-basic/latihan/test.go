package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Print("Input string palindrome: ")
	fmt.Scan(&word)
	fmt.Println(isPalindrome(word))

Check:
	fmt.Print("Input string vocal/konsonan : ")
	fmt.Scan(&word)
	if len(word) != 1 {
		fmt.Println("!! Inputan tidak valid harus satu hurup !!")
		goto Check
	}

	fmt.Println(isVocal(word))
}

func reverse(s string) string {
	r := ""

	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

func isPalindrome(s string) bool {
	return s == reverse(s)
}

func isVocal(alpha string) string {
	vocals := []string{"a", "i", "u", "e", "o"}

	for _, r := range vocals {
		if strings.ToLower(r) == strings.ToLower(alpha) {
			return "Vocal"
		}
	}

	return "Konsonan"
}

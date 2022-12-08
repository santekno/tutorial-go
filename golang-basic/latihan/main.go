package main

import (
	"fmt"
	"latihan/submission"
)

func main() {
	var word string
	fmt.Print("Input string palindrome: ")
	fmt.Scan(&word)
	fmt.Println(submission.IsPalindrome(word))

Check:
	fmt.Print("Input string vocal/konsonan : ")
	fmt.Scan(&word)
	if len(word) != 1 {
		fmt.Println("!! Inputan tidak valid harus satu hurup !!")
		goto Check
	}

	fmt.Println(submission.IsVocal(word))
}

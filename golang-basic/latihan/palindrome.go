package main

import "fmt"

func main() {
	var input string
	fmt.Print("Masukkan kata yang ingin dicek:")
	fmt.Scan(&input)
	if IsPalindrome(input) {
		fmt.Print("Palindrome")
	} else {
		fmt.Print("Bukan Palindrome")
	}
}

func IsPalindrome(input string) bool {
	var inputLenDiv2 = len(input) / 2
	var inputLenFull = len(input)

	for i := 0; i <= inputLenDiv2; i++ {
		if input[i] == input[inputLenFull-i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}

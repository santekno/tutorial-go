/*
Soal 1
Buat program `palindrome` menggunakan golang

$ go main.go
$ madam

output
true

*/

package main

import "fmt"

func isPalindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	var str string

	fmt.Print("Cek Palindrome : ")
	fmt.Scan(&str)

	result := isPalindrome(str)

	if result == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

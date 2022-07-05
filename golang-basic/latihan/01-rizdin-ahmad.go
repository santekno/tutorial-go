package main

import "fmt"

func isPalindrome(quiz string) bool {
	for i := 0; i < len(quiz)/2; i++ {
		if quiz[i] != quiz[len(quiz)-i-1] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("Silahkan Input Kata Polindrome")

	inputData := ("")

	fmt.Scan(&inputData)
	result := isPalindrome(inputData)
	if result == true {
		fmt.Println("Anda Benar")
	} else {
		fmt.Println("Salah")
	}
}

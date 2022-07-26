/*
Edspert Bootcamp - Golang Backend
Tugas 1 : Golang Basic
Peserta : Rizdin Ahmad Farizi
*/

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
	inputData := ("")

	fmt.Scan(&inputData)
	result := isPalindrome(inputData)
	if result == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

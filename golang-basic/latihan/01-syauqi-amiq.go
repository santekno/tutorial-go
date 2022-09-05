package main

import "fmt"

func main() {
	input := "madam"
	fmt.Println(checkPalindrome(input))

	fmt.Println(checkConsonant("a"))
}

// * Time Complexity O(1) karena proses checking hanya 1 kali looping
// ! Koreksi jika salah :)
func checkPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		// Setiap looping akan membuat index reverse
		reverseIndex := len(str) - 1 - i

		//compare character index maju dengan character index mundur
		if str[i] != str[reverseIndex] {
			// return false jika tidak sama, jika sama lanjutkan looping sampai selesai
			return false
		}

	}
	//return true jika setelah looping selesai dan palindrome
	return true
}

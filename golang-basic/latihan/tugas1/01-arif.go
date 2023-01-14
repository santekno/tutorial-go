package main

import "fmt"

func main() {
	var kata string

	fmt.Scan(&kata)

	result := isPalindrome(kata)

	fmt.Println(result)
}

func isPalindrome(kata string) (isPalindrome bool) {
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] == kata[len(kata)-i-1] {
			isPalindrome = true
		}
	}
	return
}

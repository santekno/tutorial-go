package main

import "fmt"

func main()  {
	// menerima input dari user
	fmt.Print("Masukkan Kata : ")
	var kata string
	fmt.Scanf("%s", &kata)

	fmt.Println(isPalindrome(kata))
	
}

func isPalindrome(kat string) bool {
	// pengolahan palindrome
	for i := 0; i < len(kat)/2; i++ {
		if kat[i] != kat[len(kat)-i-1] {
			return false
		}
	}

	// hasil jika kata adalah palindrome
	return true
}
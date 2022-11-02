package main

import "fmt"

func main()  {
	// menerima input dari user
	fmt.Print("Masukkan Kata : ")
	var kata string
	fmt.Scanf("%s", &kata)

	// pengolahan palindrome
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] != kata[len(kata)-i-1] {
			fmt.Println(false)
			return
		}
	}

	// hasil jika kata adalah palindrome
	fmt.Println(true)
}
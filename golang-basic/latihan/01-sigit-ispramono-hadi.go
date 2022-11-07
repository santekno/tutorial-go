package main

import "fmt"

//Soal 1 : Pengecekan Kata Palindrome
func cekPalindrome(kata string) bool {
	kataDibalik := ""
	j := 0
	for i := len(kata) - 1; i >= 0; i-- {
		kataDibalik += string(kata[i])
		if kata[j] != kataDibalik[j] {
			return false
		}
		j += 1
	}
	return true
}

func main() {
	var masukan string
	fmt.Println("Masukkan kata : ")
	fmt.Scanln(&masukan)
	hasil1 := cekPalindrome(masukan)
	if hasil1 == true {
		fmt.Println("Kata termasuk palindrome")
	} else if hasil1 == false {
		fmt.Println("Kata tidak termasuk palindrome")
	}
}

package main

import "fmt"

func main() {

	fmt.Println(isPalindrom("a"))
	fmt.Println(isPalindrom("ab"))
	fmt.Println(isPalindrom("aba"))
	fmt.Println(isPalindrom("katak"))
}

func isPalindrom(kata string) bool {
	for i := 0; i < len(kata)/2; i++ {
		indexAwal := i
		indexAkhir := len(kata) - i - 1

		if kata[indexAwal] != kata[indexAkhir] {
			return false
		}

	}
	return true
}

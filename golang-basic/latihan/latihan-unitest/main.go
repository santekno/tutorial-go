package main

import (
	"fmt"
	"quis/isPalindrom"
	"quis/vocal"
)

func main() {
	isPalindromTrue := isPalindrom.IsPalindrom("kodok")
	isPalindromFalse := isPalindrom.IsPalindrom("batman")

	isvocalTrue := vocal.Vocal("a")
	isvocalFalse := vocal.Vocal("B")

	fmt.Println("Soal nomor satu")
	fmt.Println("kodok = ", isPalindromTrue)
	fmt.Println("batman = ", isPalindromFalse)

	fmt.Println("-----------------------------")
	fmt.Println("Soal nomor Dua")
	fmt.Println("Huruf A adalah ", isvocalTrue)
	fmt.Println("Huruf B adalah ", isvocalFalse)
}

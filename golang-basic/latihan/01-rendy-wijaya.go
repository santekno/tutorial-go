package main

import (
	"fmt"
	"latihan1/helper"
)

func main() {
	//	ambil dan simpan inputan user ke variable "input"
	var input string

	fmt.Print("Masukkan kata/kalimat: ")
	_, err := fmt.Scan(&input)

	// panic jika inputan error
	if err != nil {
		panic(err)
	}

	fmt.Println(helper.IsPalindrome(input))
}

package main

import (
	"fmt"
	"latihan1/helper"
)

func main() {

	//	ambil inputan user dan buat variable untuk menampung inputan
	var input string
	fmt.Print("Masukkan huruf: ")
	_, err := fmt.Scan(&input)

	//	panic jika inputan error
	if err != nil {
		panic(err)
	}

	fmt.Println(helper.IsVokal(input))
}

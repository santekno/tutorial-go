package main

import (
	"fmt"
	"strings"
)

func main() {
	// masukan semua huruf vokal ke variable, dan lakukan concat huruf vokal yang sudah di ubah menjadi Upper Case
	vokal := "aiueo"
	vokal += strings.ToUpper(vokal)

	// masukan semua huruf konsonan ke variable, dan lakukan concat huruf konsonan yang sudah di ubah menjadi Upper Case
	konsonan := "bcdfghjklmnpqrstvwxyz"
	konsonan += strings.ToUpper(konsonan)

	//	ambil inputan user dan buat variable untuk menampung inputan
	var input string
	fmt.Print("Masukkan huruf: ")
	_, err := fmt.Scan(&input)

	//	panic jika inputan error
	if err != nil {
		panic(err)
	}

	//	user hanya boleh input 1 huruf, jika kurang atau lebih return dan berikan pesan
	if len(input) < 1 || len(input) > 1 {
		fmt.Println("Inputan tidak boleh kurang dari atau lebih dari 1")
		return
	}

	// jika yang di input user valid, cek apakah inputan vokal atau konsonan
	// jika user menginput huruf atau simbol maka return dan beritahu jika yang di input bukan vokal atau konsonan
	if strings.Contains(vokal, input) {
		fmt.Printf("%s adalah vokal", input)
	} else if strings.Contains(konsonan, input) {
		fmt.Printf("%s adalah konsonan", input)
	} else {
		fmt.Printf("%s bukan vokal ataupun konsonan", input)
	}
}

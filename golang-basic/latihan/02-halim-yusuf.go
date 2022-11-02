package main

import (
	"fmt"
	"strings"
)

func main()  {

	// menentukan jenis huruf
	var (
		vokal = "aiueo"
		konsonan = "bcdfghjklmnpqrstvwxyz"
	)
	
	// mengubah huruf menjadi kapital
	vokal += strings.ToUpper(vokal)
	konsonan += strings.ToUpper(konsonan)

	// menerima input dari user
	fmt.Print("Masukkan 1 huruf : ")
	var huruf string
	fmt.Scanf("%s", &huruf)

	// validasi user input huruf tidak boleh lebih dari 1
	if len(huruf) > 1 {
		fmt.Println("Anda hanya boleh memasukkan 1 huruf")
		return
	}

	// pengkondisian sesuai inputan user
	if strings.Contains(vokal, huruf) {
		fmt.Println("huruf Vokal")
	} else if strings.Contains(konsonan, huruf) {
		fmt.Println("huruf Konsonan")
	} else {
		fmt.Println("yang Anda masukkan bukan karakter huruf")
	}
	
}
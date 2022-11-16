package main

import (
	"fmt"
	"strings"
)

func main()  {
	// menerima input dari user
	fmt.Print("Masukkan 1 huruf : ")
	var huruf string
	fmt.Scanf("%s", &huruf)

	// validasi user input huruf tidak boleh lebih dari 1
	if len(huruf) > 1 {
		fmt.Println("Anda hanya boleh memasukkan 1 huruf")
		return
	}

	fmt.Println(isVokalOrKonsonan(huruf))
	
}

func isVokalOrKonsonan(hur string) string {
	// menentukan jenis huruf
	var (
		vokal = "aiueo"
		konsonan = "bcdfghjklmnpqrstvwxyz"
	)

	// mengubah huruf menjadi kapital
	vokal += strings.ToUpper(vokal)
	konsonan += strings.ToUpper(konsonan)

	// pengkondisian sesuai inputan user
	if strings.Contains(vokal, hur) {
		return "huruf Vokal"
	} else if strings.Contains(konsonan, hur) {
		return "huruf Konsonan"
	} else {
		return "yang Anda masukkan bukan karakter huruf"
	}
}
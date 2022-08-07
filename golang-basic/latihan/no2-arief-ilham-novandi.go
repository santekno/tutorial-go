package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Masukkan Huruf : ")
	var checkhuruf string
	fmt.Scan(&checkhuruf)

	var validasi = strings.ToLower(checkhuruf)
	switch validasi {
	case "a", "i", "u", "e", "o":
		fmt.Println("huruf Vokal")

	default:
		fmt.Println("huruf Konsonan")
	}
}

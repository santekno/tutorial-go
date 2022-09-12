package main

import (
	"fmt"
	"strings"

	"4d63.com/strrev"
)

func main() {
	fmt.Print("masukan string ")
	var polidrome string
	fmt.Scanln(&polidrome)
	hasil := Cek_polidrome(polidrome)
	fmt.Println(hasil)
}

func Cek_polidrome(polidrome string) string {

	hasil := strrev.Reverse(polidrome)

	if len(polidrome) < 2 {
		return "String harus lebih dari 2 huruf"
	} else {
		if strings.ToLower(polidrome) == strings.ToLower(hasil) {
			return "true"
		} else {
			return "false"
		}
	}

}

package main

import (
	"fmt"
)

func main() {
	vocalKonsonan("a")
	vocalKonsonan("b")
}

func vocalKonsonan(letter string) {
	if letter == "a" || letter == "i" || letter == "u" || letter == "e" || letter == "o" {
		fmt.Println("Vocal")
	} else {
		fmt.Println("Konsonan")
	}
}

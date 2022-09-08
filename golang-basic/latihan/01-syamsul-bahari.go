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

	hasil := strrev.Reverse(polidrome)

	if len(polidrome) < 2 {
		fmt.Println("String harus lebih dari 2 huruf")
	} else {
		if strings.ToLower(polidrome) == strings.ToLower(hasil) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}

}

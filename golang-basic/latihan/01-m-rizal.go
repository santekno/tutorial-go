package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	kata := "madam"

	fmt.Println(isPalindrom(kata))

	VocalOrConsonant()
}

func isPalindrom(kata string) bool {
	var loop int
	if loop = len(kata) / 2; len(kata)%2 != 0 {
		loop = (len(kata) - 1) / 2
	}

	for i := 0; i < loop; i++ {
		if kata[i] != kata[len(kata)-i-1] {
			return false
		}
	}
	return true
}

func VocalOrConsonant() {
	var name string
	fmt.Print("masukan huruf : ")
	fmt.Scanf("%s", &name)

	name = strings.ToLower(name)

	if len(name) > 1 {
		fmt.Println("masukan huruf")
		os.Exit(3)
	}

	if name == "a" || name == "i" || name == "u" || name == "e" || name == "o" {
		fmt.Println("vocal")
	} else {
		fmt.Println("konsonan")
	}
}

package main

import (
	"fmt"
)

func cekpalindrom(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Masukkan Kalimat :")
	var str string
	fmt.Scan(&str)
	result := cekpalindrom(str)
	if result {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}

package main

import "fmt"

func main() {
	var inputan string
	fmt.Print("Masukan huruf: ")
	fmt.Scanln(&inputan)
	vokalkosnan(inputan)

}

func vokalkosnan(input string) {
	if input == "a" || input == "i" || input == "u" || input == "e" || input == "o" {
		println("huruf vokal")
	} else {
		fmt.Println("kosonan")
	}
}

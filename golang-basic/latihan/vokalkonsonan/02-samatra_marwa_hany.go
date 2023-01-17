package main

import "fmt"

func main() {
	var inputan string
	fmt.Print("Masukan huruf: ")
	fmt.Scanln(&inputan)
	fmt.Println(vokalkosnan(inputan))

}

func vokalkosnan(input string) string {
	var kembalian string = ""
	if input == "a" || input == "i" || input == "u" || input == "e" || input == "o" || input == "A" || input == "I" || input == "U" || input == "E" || input == "O" {
		kembalian = "vokal"
		return kembalian
	} else {
		kembalian = "konsonan"
		return kembalian
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func ConsonantOrVocal(input byte) bool {
	if input == 'a' || input == 'i' || input == 'u' || input == 'e' || input == 'o' ||
		input == 'A' || input == 'I' || input == 'U' || input == 'E' || input == 'O' {
		fmt.Println("vowel")
		return true
	} else {
		fmt.Println("consonant")
		return false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan huruf = ")
	input, _ := reader.ReadByte()
	ConsonantOrVocal(input)
}

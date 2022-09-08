package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Karakter = ")
	chr, _ := reader.ReadByte()

	isVowel(chr)
}

func isVowel(chr byte) {
	if chr == 'a' || chr == 'e' || chr == 'i' || chr == 'o' || chr == 'u' ||
		chr == 'A' || chr == 'E' || chr == 'I' || chr == 'O' || chr == 'U' {
		fmt.Printf("%c termasuk karakter Vokal \n", chr)
	} else {
		fmt.Printf("%c termsuk karakter Konsonan \n", chr)
	}

}

package main

import "fmt"

func isLetter(char string) bool {
	for _, r := range char {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func checkChar(char string) string {
	if len(char) != 1 {
		return "Hanya boleh 1 karakter"
	}

	if !isLetter(char) {
		return "Hanya boleh huruf"
	}

	switch char {
	case "a", "i", "u", "e", "o", "A", "I", "U", "E", "O":
		return "Vokal"
	}

	return "Konsonan"
}

func main() {
	var char string
	fmt.Print("Masukkan 1 huruf: ")
	fmt.Scan(&char)
	fmt.Println(checkChar(char))
}

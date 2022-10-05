package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isVocal("A"))      // Vokal
	fmt.Println(isVocal("Z"))      // Konsonan
	fmt.Println(isVocal("1"))      // error
	fmt.Println(isVocal("Achmad")) // error
}

func isVocal(str string) string {
	if len(str) > 1 {
		panic("it's looks like more than 1 letter")
	}

	if _, err := strconv.Atoi(str); err == nil {
		panic("looks like number")
	}

	vocalLetters := []string{"a", "i", "u", "e", "o"}

	for _, char := range vocalLetters {
		if strings.ToLower(str) == char {
			return "Vokal"
		}
	}
	return "Konsonan"
}

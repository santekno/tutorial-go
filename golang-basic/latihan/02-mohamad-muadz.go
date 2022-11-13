package main

import (
	"strings"
)

// func main() {
// 	fmt.Println("Enter 1 character")
// 	inputReader := bufio.NewReader(os.Stdin)
// 	character, _ := inputReader.ReadString('\n')
// 	character = strings.Trim(character, "\n")
// 	fmt.Println(consonantVocal(character))
// }

func consonantVocal(character string) string {
	if len(character) > 1 {
		return "Only one character allowed"
	}

	for _, r := range character {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return "Character must be alphabet"
		}
	}

	character = strings.ToLower(character)

	if character == "a" || character == "i" || character == "u" || character == "e" || character == "o" {
		return "vocal"
	}

	return "consonant"
}

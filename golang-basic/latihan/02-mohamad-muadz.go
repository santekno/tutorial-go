package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter 1 character")
	inputReader := bufio.NewReader(os.Stdin)
	character, _ := inputReader.ReadString('\n')
	character = strings.Trim(character, "\n")

	if len(character) > 1 {
		fmt.Println("Only one character allowed")
		return
	}

	for _, r := range character {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			fmt.Println("Character must be alphabet")
			return
		}
	}

	character = strings.ToLower(character)

	if character == "a" || character == "i" || character == "u" || character == "e" || character == "o" {
		fmt.Println("vocal")
		return
	}

	fmt.Println("consonant")

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func VowelCheck(value string) string {
	normalizedInput := strings.ToLower(value)

	switch normalizedInput {
	case "a", "i", "u", "e", "o":
		return "Vocal"
	}
	return "Konsonan"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan 1 huruf alfabet: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println(VowelCheck(input))
}

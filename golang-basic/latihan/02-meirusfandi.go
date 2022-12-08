package main

import (
	"fmt"
	"strings"
)

func typeSentences(ch string) string {
	if strings.EqualFold(ch, "a") || strings.EqualFold(ch, "i") || strings.EqualFold(ch, "u") || strings.EqualFold(ch, "e") || strings.EqualFold(ch, "o") {
		return "Vokal"
	} else {
		return "Konsonan"
	}
}

func main() {
	var chr string
	fmt.Print("Input sentences : ")
	fmt.Scan(&chr)
	fmt.Println(typeSentences(chr))
}

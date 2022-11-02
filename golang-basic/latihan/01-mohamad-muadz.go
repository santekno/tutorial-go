package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// check if entered word is polindrome

func main() {
	fmt.Println("Enter word(s):")
	inputReader := bufio.NewReader(os.Stdin)
	word, _ := inputReader.ReadString('\n')
	word = strings.Trim(strings.ToLower(word), "\n")
	var arrWord []string
	var reverse string
	for _, v := range word {
		arrWord = append(arrWord, string(v))
	}
	for i := len(arrWord) - 1; i >= 0; i-- {
		reverse += arrWord[i]
	}
	reverse = strings.Trim(reverse, "\n")

	fmt.Println(word == reverse)
}

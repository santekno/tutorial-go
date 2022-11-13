package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	fmt.Println("Enter word(s):")
// 	inputReader := bufio.NewReader(os.Stdin)
// 	word, _ := inputReader.ReadString('\n')
// 	word = strings.Trim(strings.ToLower(word), "\n")
// 	polindrome(word)
// }

func polindrome(word string) bool {
	var arrWord []string
	var reverse string
	for _, v := range word {
		arrWord = append(arrWord, string(v))
	}
	for i := len(arrWord) - 1; i >= 0; i-- {
		reverse += arrWord[i]
	}
	reverse = strings.Trim(reverse, "\n")

	isPolindrome := word == reverse
	fmt.Println(isPolindrome)
	return isPolindrome
}

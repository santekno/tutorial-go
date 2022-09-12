// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var word string

	isVokalOrNot := "adalah konsonan"

	fmt.Printf("Silahkan input huruf : ")
	fmt.Scanf("%s", &word)

	vokalOrKonsonan := isVokal(word)

	if vokalOrKonsonan {
		isVokalOrNot = "adalah huruf vokal"
	}

	fmt.Println(word, isVokalOrNot)

}

func isVokal(val string) bool {
	if val == "a" || val == "i" || val == "u" || val == "e" || val == "o" {
		return true
	}

	return false
}

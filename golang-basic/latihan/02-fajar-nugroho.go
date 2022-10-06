package main

import (
	"fmt"
	"strings"
)

func main() {

	output, status := checkIsVocal("E")
	if status == 0 {
		fmt.Println("Data Max 1 Word")
	} else {
		fmt.Println(output)
	}

}

func checkIsVocal(param string) (string, int) {
	var isVocal string
	var status int

	if len(param) > 1 {
		status = 0
	} else {
		status = 1
	}

	switch strings.ToUpper(param) {
	case "A", "I", "U", "E", "O":
		isVocal = "vocal"
	default:
		isVocal = "konsonan"
	}
	return isVocal, status
}

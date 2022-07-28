package main

import (
	"fmt"
	"strings"
)

func ExerciseOne(userInput string) bool {

	tempArr := strings.Split(userInput, "")
	for i, j := 0, len(tempArr)-1; i < j; i, j = i+1, j-1 {
		tempArr[i], tempArr[j] = tempArr[j], tempArr[i]
	}
	reservedString := strings.Join(tempArr, "")

	var result = reservedString == userInput
	return result
}

func main() {
	var inputedString string = ""
	fmt.Scan(&inputedString)
	fmt.Println(ExerciseOne(inputedString))
}

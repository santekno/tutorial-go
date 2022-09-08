package main

import (
	"errors"
	"fmt"
)

func isPalindrome(input string) (bool, error) {
	length := len(input)
	if length < 2 {
		return true, nil
	}

	var leftSlice string
	var rightSlice string
	var rightSliceFlip string = ""

	halfLength := int(length / 2)
	leftSlice = input[:halfLength]
	rightSlice = input[halfLength+(length%2):]

	// flip right
	strRune := []rune(fmt.Sprintf("%s", rightSlice))
	for i := range strRune {
		rightSliceFlip = fmt.Sprintf("%c%s", strRune[i], rightSliceFlip)
	}

	// fmt.Printf("%d, %s, %s, %s\n", halfLength, leftSlice, rightSlice, rightSliceFlip)

	if len(leftSlice) != len(rightSlice) {
		return false, errors.New("should be sliced evenly")
	}
	return leftSlice == rightSliceFlip, nil
}

func main_01() {
	fmt.Println("Enter word")
	var input string
	fmt.Scanln(&input)
	result, err := isPalindrome(input)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(result)
	}
}

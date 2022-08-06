package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(isPalindrome("MakAm"))
}

func isPalindrome(str string) bool {
    stringLength := len(str)
    if stringLength == 1 {
        return true
    }

    reversedString := []byte{}
    for i := len(str) - 1; i >= 0; i-- {
        reversedString = append(reversedString, str[i])
    }

    return strings.ToLower(str) == strings.ToLower(string(reversedString))
}

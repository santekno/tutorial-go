package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(isVocal("W"))
}

func isVocal(str string) string {
    if len(str) > 1 {
        return "Please input only 1 letter"
    }
    vocalLetters := []string{"a", "i", "u", "e", "o"}

    for _, char := range vocalLetters {
        if strings.ToLower(str) == char {
            return "Vokal"
        }
    }

    return "Konsonan"
}

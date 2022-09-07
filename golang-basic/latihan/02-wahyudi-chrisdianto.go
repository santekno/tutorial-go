package main

import (
	"fmt"
	"strings"
)  
  
func isVowel(input string ) bool { 
		switch true {  
			case input == "a":  
				return false
			case input == "i":  
				return false
			case input == "u":  
				return false
			case input == "e":  
				return false
			case input == "o":  
				return false
			default:
				return true
		} 
}  
func main() {  
	fmt.Print("input to check Vokal atau Konsonan: ")
	var input string
	fmt.Scanln(&input)
	 if isVowel(strings.ToLower(input)) == true {
      fmt.Println("Konsonan")
   } else {
      fmt.Println("Vokal")
	 }
} 
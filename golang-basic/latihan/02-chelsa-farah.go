package main

import "fmt"

func main() {
	isVokal('l')
}

func isVokal(character rune) {
	if(character == 'a' || character == 'b' || character =='c' || character =='d' || character =='e'){
		fmt.Println("Vokal")
	}else{
		fmt.Println("Konsonan")
	}
}
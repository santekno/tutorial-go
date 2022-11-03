package main

import "fmt"

var inputdata string

func main() {

	fmt.Scan(&inputdata)

	result := palindrome(inputdata)
	fmt.Println(result)

}

//Palindrome Function
func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ { // jika i dibawah length input di bagi 2 , i++ looping
		if input[i] != input[len(input)-i-1] { //perbandingan array input dari depan dan belakang
			return false
		}

	}
	return true
}

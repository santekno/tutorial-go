// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var words string
	var number int

	isPalindromeStr := "bukan palindrome"
	isPalindromeInt := "bukan palindrome"

	fmt.Printf("Silahkan input kata : ")
	fmt.Scanf("%s", &words)

	palindromeStr := isPalindromeString(words)

	if palindromeStr {
		isPalindromeStr = "adalah palindrome"
	}

	fmt.Printf("Silahkan input angka : ")
	fmt.Scanf("%d", &number)

	palindromeInt := isPalindromeInteger(number)

	if palindromeInt {
		isPalindromeInt = "adalah palindrome"
	}

	fmt.Println(words, isPalindromeStr)
	fmt.Println(number, isPalindromeInt)

}

func isPalindromeString(val string) bool {
	for i := 0; i < len(val); i++ {
		j := len(val) - 1 - i
		if val[i] != val[j] {
			return false
		}
	}
	return true
}

func isPalindromeInteger(val int) bool {
	var remainder, temp int
	var reverse int = 0

	temp = val

	// For Loop used in format of While Loop
	for {
		remainder = val % 10
		reverse = reverse*10 + remainder
		val /= 10

		if val == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == reverse {
		return true
	} else {
		return false
	}
}

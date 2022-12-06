package main

import "fmt"

func main() {
	cekPalindrome := ""
	fmt.Print("Masukan Kata yang ingin di cek palindrome : ")
	fmt.Scan(&cekPalindrome)
	isPalindrome(cekPalindrome)

}
func isPalindrome(str string) {
	revString := ""

	for i := len(str) - 1; i >= 0; i-- {
		revString = revString + string(str[i])
	}
	if str == revString {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
/* Soal 2
	Buat program untuk menentukan Huruf yang diinput Vokal atau Konsonan
 >>> input
	$ go main .go
	h
 >>> output
    Konsonan
*/

package main

import "fmt"

func main() {
	var vocal string
	fmt.Print("Masukkan huruf : ")
	fmt.Scan(&vocal)

	if vocal == "a" || vocal == "e" || vocal == "i" || vocal == "o" || vocal == "u" {
		fmt.Println("Vokal.")
	} else {
		fmt.Println("Konsonan")
	}
}

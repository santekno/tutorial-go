package main

import (
	"fmt"
	"strings"
)


func jenisHuruf(huruf string){
	
	if len(huruf) == 1 {
		vokal := "aiueo"
		konsonan := "bcdefghijklmnopqrstuvwxyz"
		
		if strings.Contains(vokal, strings.ToLower(huruf)) {
			fmt.Println("Vokal!")
		} else if strings.Contains(konsonan, huruf) {
			fmt.Println("Konsonan!")
		} else {
			fmt.Println("Masukkan input alphabet!")
		}
	} else {
		fmt.Println("anda hanya boleh memasukan 1 huruf")
	}
}


func main()  {
	var input string
	fmt.Print("Masukan Huruf : ")
	fmt.Scanf("%s", &input)
	jenisHuruf(input)

}
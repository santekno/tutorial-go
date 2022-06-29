package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println(i, " Satu")
		case 2:
			fmt.Println(i, " Dua")
		case 3:
			fmt.Println(i, " Tiga")
		case 4:
			fmt.Println(i, " Empat")
		case 5:
			fmt.Println(i, " Lima")
		case 6:
			fmt.Println(i, " Enam")
		case 7:
			fmt.Println(i, " Tujuh")
		case 8:
			fmt.Println(i, " Delapan")
		case 9:
			fmt.Println(i, " Sembilan")
		case 10:
			fmt.Println(i, " Sepuluh")
		default:
			fmt.Println(i, " Tidak terdeteksi")
		}
	}
}

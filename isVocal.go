package main

import "fmt"

func main() {
	cekHuruf := ""
	fmt.Print("Masukan huruf yang akan di cek vocal/consonant : ")
	fmt.Scan(&cekHuruf)
	isVocal(cekHuruf)

}
func isVocal(str string) {

	if len(str) > 1 {
		fmt.Println("hanya bisa cek satu huruf")
	} else {
		if str == "a" || str == "i" || str == "u" || str == "o" || str == "e" {
			fmt.Println("Huruf ini vocal")
		} else {
			fmt.Println("Huruf ini Consonant")
		}
	}

	// switch str{
	// case "a":
	// 		fmt.Println("Huruf ini vocal")
	// case "i":
	// 		fmt.Println("Huruf ini vocal")
	// case "u":
	// 		fmt.Println("Huruf ini vocal")
	// case "o":
	// 		fmt.Println("Huruf ini vocal")
	// case "e":
	// 		fmt.Println("Huruf ini vocal")
	// default:
	// 	fmt.Println("Huruf Consonant")
	// }

	// if str == "a" || str == "i" || str == "u" || str == "o" || str == "e"{
	// 	fmt.Println("Huruf ini vocal")
	// }if else{
	// 	fmt.Println("Huruf ini Consonant")
	// }
}
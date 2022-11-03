package main

import "fmt"

var inputdata string

func main() {

	fmt.Scan(&inputdata)

	result := cekFonologi(inputdata)
	if result == true { //true = vokal
		fmt.Println("Vokal")
	} else { //false = konsonan
		fmt.Println("Konsonan")
	}

}

//Fonologi Function
func cekFonologi(input string) bool {
	var x [5]string
	x[0] = "a"
	x[1] = "i"
	x[2] = "u"
	x[3] = "e"
	x[4] = "o"

	for i := 0; i < 5; i++ { //Looping array
		if input == x[i] { //Perbandingan input dgn array
			return true
		}

	}
	return false
}

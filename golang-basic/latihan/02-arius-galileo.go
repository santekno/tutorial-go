package main

import "fmt"

//konsonan

func main() {
	var inputHuruf string
	var hurufVokal = [5]string{"a", "i", "u", "e", "o"}
	var jawaban string
	jawaban = "Konsonan"
	fmt.Scan(&inputHuruf)
	for i := 0; i < len(hurufVokal); i++ {
		if hurufVokal[i] == inputHuruf {
			jawaban = "Vokal"
		}
	}
	fmt.Println(jawaban)
}

package tasks

import(
	"fmt"
)

func IsVocalOrConsonant(str string) {

	if str == "a" || str == "e" || str == "i" || str == "o" || str == "u" {  
		fmt.Println("Vokal")  
	  } else {  
		fmt.Println("Konsonan")  
	  }
}
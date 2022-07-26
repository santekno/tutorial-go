package main

// import "fmt"

// func main() {
// 	// Soal 2
// 	fmt.Print("Enter a letter : ")
// 	var input string
// 	fmt.Scanf("%s", &input)

// 	if (input >= "a" && input <= "z") || (input >= "A" && input <= "Z") {
// 		switch letter {
// 		case "a", "i", "u", "e", "o":
// 			fmt.Println("Vokal")
// 		default:
// 			fmt.Println("Konsonan")
// 		}
// 	} else {
// 		fmt.Println("Bukan huruf")
// 	}
// }

func checkLetter(letter string) string {
	result := "Bukan huruf"
	if (letter >= "a" && letter <= "z") || (letter >= "A" && letter <= "Z") {
		switch letter {
		case "a", "i", "u", "e", "o":
			result = "Vokal"
		default:
			result = "Konsonan"
		}
	}

	return result
}

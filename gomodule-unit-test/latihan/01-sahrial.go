package latihan

import "fmt"

func isVocalOrConsonant(value string) string {
	switch value {
	case "a":
		return "Vokal"
	case "i":
		return "Vokal"
	case "u":
		return "Vokal"
	case "e":
		return "Vokal"
	case "o":
		return "Vokal"
	case "A":
		return "Vokal"
	case "I":
		return "Vokal"
	case "U":
		return "Vokal"
	case "E":
		return "Vokal"
	case "O":
		return "Vokal"
	default:
		return "Konsonan"
	}
}

func main() {
	fmt.Print("Enter your character: ")
	var input string
	fmt.Scanf("%s", &input)

	result := isVocalOrConsonant(input)
	fmt.Println(result)
}

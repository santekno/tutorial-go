package main

import "fmt"

func vokalChecking(huruf string) string {
	switch huruf {
	case "a":
		huruf = "Vokal"
		break
	case "i":
		huruf = "Vokal"
		break
	case "u":
		huruf = "Vokal"
		break
	case "e":
		huruf = "Vokal"
		break
	case "o":
		huruf = "Vokal"
		break
	case "A":
		huruf = "Vokal"
		break
	case "I":
		huruf = "Vokal"
		break
	case "U":
		huruf = "Vokal"
		break
	case "E":
		huruf = "Vokal"
		break
	case "0":
		huruf = "Vokal"
		break
	default:
		huruf = "Konsonan"
	}
	return huruf
}

func main() {
	huruf := ""
	fmt.Scan(&huruf)
	isVokal := vokalChecking(huruf)
	fmt.Println(isVokal)
}

package main

// ! Koreksi jika salah
func checkConsonant(char string) bool {
	switch true {
	case char == "A" || char == "a":
		return false
	case char == "I" || char == "i":
		return false
	case char == "U" || char == "u":
		return false
	case char == "E" || char == "e":
		return false
	case char == "O" || char == "o":
		return false
	default:
		return true
	}
}

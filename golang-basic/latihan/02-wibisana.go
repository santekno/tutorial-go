package main

func Consonant(check string) string {
	switch check {
	case "a", "i", "u", "e", "o":
		return "vokal"
	default:
		return "konsonan"
	}
}

func main() {
	check := "i"
	println(Consonant(check))
}

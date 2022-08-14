package latihan

import "fmt"

func Reverse(s string) bool {
	var temp string
	for _, x := range s {
		temp = string(x) + temp
	}
	if s != temp {
		return false
	}
	return true
}

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &input)

	fmt.Println(Reverse(input))
}

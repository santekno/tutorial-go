package main

import "fmt"

// Program A
// func f() {
// 	fmt.Println(x)
// }
// func main() {
// 	x := 5
// 	f()
// }

// Program B
// func f(x int) {
// 	fmt.Println(x)
// }
// func main() {
// 	x := 5
// 	f(x)
// }

// Program C
// var x int = 5

// func f() {
// 	fmt.Println(x)
// }
// func main() {
// 	f()
// }

// Program D
func main() {
	fmt.Println(f1())
}
func f1() int {
	return f2()
}
func f2() int {
	return 1
}
